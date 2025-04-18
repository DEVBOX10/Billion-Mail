package batch_mail

import (
	"billionmail-core/internal/model/entity"

	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"github.com/panjf2000/ants/v2"
	"golang.org/x/time/rate"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// 全局任务执行器映射
	taskExecutors      = make(map[int]*TaskExecutor)
	taskExecutorsMutex sync.RWMutex

	// 全局速率限制器
	globalLimiter = rate.NewLimiter(rate.Limit(5000), 100) // 默认每秒5000个请求，最多突发100个
)

// GetTaskExecutor 获取任务执行器
func GetTaskExecutor(taskId int) *TaskExecutor {
	taskExecutorsMutex.RLock()
	defer taskExecutorsMutex.RUnlock()
	return taskExecutors[taskId]
}

// GetOrCreateTaskExecutor 获取或创建任务执行器
func GetOrCreateTaskExecutor(ctx context.Context, taskId int) *TaskExecutor {
	taskExecutorsMutex.RLock()
	executor, exists := taskExecutors[taskId]
	taskExecutorsMutex.RUnlock()

	if !exists {
		taskExecutorsMutex.Lock()
		defer taskExecutorsMutex.Unlock()
		// 双重检查，避免竞态条件
		if executor, exists = taskExecutors[taskId]; !exists {
			executor = NewTaskExecutor(ctx)
			taskExecutors[taskId] = executor
		}
	}
	return executor
}

// RegisterTaskExecutor 注册任务执行器
func RegisterTaskExecutor(taskId int, executor *TaskExecutor) {
	taskExecutorsMutex.Lock()
	defer taskExecutorsMutex.Unlock()
	taskExecutors[taskId] = executor
}

// RemoveTaskExecutor 移除任务执行器
func RemoveTaskExecutor(taskId int) {
	taskExecutorsMutex.Lock()
	defer taskExecutorsMutex.Unlock()

	if executor, exists := taskExecutors[taskId]; exists {
		// 停止执行器的所有操作
		executor.Stop()
		delete(taskExecutors, taskId)
	}
}

// CleanupIdleExecutors 清理空闲的执行器
func CleanupIdleExecutors() {
	taskExecutorsMutex.Lock()
	defer taskExecutorsMutex.Unlock()

	now := time.Now()
	for id, executor := range taskExecutors {
		if !executor.IsRunning() && now.Sub(executor.lastActivity) > 30*time.Minute {
			executor.Stop()
			delete(taskExecutors, id)
		}
	}
}

// TaskExecutor 简化版任务执行器
type TaskExecutor struct {
	// 上下文和取消函数
	ctx    context.Context
	cancel context.CancelFunc

	// 运行状态
	isRunning    atomic.Bool
	isPaused     atomic.Bool
	lastActivity time.Time

	// 速率控制
	rateController *SimpleRateController

	// 工作池
	pool *ants.Pool
	wg   sync.WaitGroup

	// 指标统计
	sentCount   atomic.Int64
	failedCount atomic.Int64
	startTime   time.Time

	// 暂停/恢复控制
	pauseChan  chan struct{}
	resumeChan chan struct{}
}

// SendResult 发送结果
type SendResult struct {
	RecipientID int
	Success     bool
	MessageID   string
	Error       error
}

// NewTaskExecutor 创建任务执行器
func NewTaskExecutor(ctx context.Context) *TaskExecutor {
	taskCtx, cancel := context.WithCancel(ctx)

	executor := &TaskExecutor{
		ctx:            taskCtx,
		cancel:         cancel,
		lastActivity:   time.Now(),
		startTime:      time.Now(),
		pauseChan:      make(chan struct{}, 1),
		resumeChan:     make(chan struct{}, 1),
		rateController: NewSimpleRateController(1000), // 默认每分钟1000封
	}

	return executor
}

// IsRunning 检查任务是否正在运行
func (e *TaskExecutor) IsRunning() bool {
	return e.isRunning.Load()
}

// IsPaused 检查任务是否已暂停
func (e *TaskExecutor) IsPaused() bool {
	return e.isPaused.Load()
}

// ProcessTask 处理任务 todo 取到单个任务ID去执行
func (e *TaskExecutor) ProcessTask(ctx context.Context) error {
	// 防止重复运行
	if !e.isRunning.CompareAndSwap(false, true) {
		fmt.Println("邮件任务执行器已经在运行中，请稍后再试")
		return errors.New("task executor is already running")
	}

	defer e.isRunning.Store(false)

	// 更新活动时间
	e.lastActivity = time.Now()

	// 获取任务ID
	taskId, err := e.getTaskIdFromContext(ctx)
	if err != nil {
		fmt.Println("获取任务ID失败:", err.Error())
		return err
	}

	// 获取任务信息
	task, err := GetTaskInfo(ctx, taskId)
	if err != nil {
		fmt.Println("获取任务信息失败:", err.Error())
		return fmt.Errorf("failed to get task info: %w", err)
	}

	if task == nil || task.Id == 0 {
		fmt.Println("任务不存在 111")
		return fmt.Errorf("task %d not found", taskId)
	}

	// 检查任务是否应该运行
	if task.TaskProcess == 2 { // 已完成
		fmt.Println("当前取到的任务 已经完成")
		return nil
	}

	// 设置是否暂停状态
	if task.Pause == 1 {
		fmt.Println("当前取到的任务 已被暂停 无法执行")
		e.isPaused.Store(true)
	}

	// 配置速率控制器 todo
	e.configureRateController(task)

	// 获取模板信息
	template, err := e.getTemplateInfo(ctx, task.TemplateId)
	if err != nil {
		fmt.Println("获取模板信息失败:", err.Error())
		return fmt.Errorf("failed to get template: %w", err)
	}

	// 处理邮件内容
	emailContent := e.processEmailContent(ctx, template.Content, task)

	// 创建工作池
	poolSize := task.Threads
	if poolSize <= 0 {
		poolSize = 3 // 默认线程数
	}

	e.pool, err = ants.NewPool(poolSize, ants.WithPreAlloc(true), ants.WithPanicHandler(func(p interface{}) {
		fmt.Println("工作池 问题", p)
		g.Log().Error(ctx, "Worker panic: %v", p)
	}))

	if err != nil {
		fmt.Println("创建工作池失败:", err.Error())
		return fmt.Errorf("failed to create worker pool: %w", err)
	}

	defer e.pool.Release()

	// 更新任务状态为进行中
	if task.TaskProcess == 0 {
		if err := UpdateTaskProcessStatus(ctx, task.Id, 1); err != nil {
			fmt.Println("更新任务状态失败:", err.Error())
			return fmt.Errorf("failed to update task status: %w", err)
		}
	}

	// 处理任务  替换jwt  发件  更新发件状态
	if err := e.processTaskRecipients(ctx, task, emailContent); err != nil {
		fmt.Println("处理任务失败:", err.Error())
		if errors.Is(err, context.Canceled) {
			fmt.Println("任务被取消:", err.Error())
			g.Log().Info(ctx, "Task %d was canceled", task.Id)
			return nil
		}
		return err
	}

	// 检查任务是否完成
	completed, err := e.isTaskComplete(ctx, task.Id)
	if err != nil {
		fmt.Println("检查任务完成状态失败:", err.Error())
		return fmt.Errorf("failed to check completion: %w", err)
	}
	// 状态改2 完成
	if completed {
		if err := UpdateTaskProcessStatus(ctx, task.Id, 2); err != nil {
			return fmt.Errorf("failed to update task status: %w", err)
		}
		g.Log().Info(ctx, "Task %d completed successfully", task.Id)
	}

	return nil
}

// Stop 停止任务执行器
func (e *TaskExecutor) Stop() {
	if e.cancel != nil {
		e.cancel()
	}

	// 等待所有工作完成
	done := make(chan struct{})
	go func() {
		e.wg.Wait()
		close(done)
	}()

	// 最多等待3秒
	select {
	case <-done:
		// 工作已完成
	case <-time.After(3 * time.Second):
		// 超时，强制结束
	}

	// 释放工作池
	if e.pool != nil {
		e.pool.Release()
	}

	e.isRunning.Store(false)
}

// PauseTask 暂停任务
func (e *TaskExecutor) PauseTask(taskId int) error {
	// 如果已经是暂停状态，直接返回
	if e.isPaused.Load() {
		return nil
	}

	// 设置暂停状态
	e.isPaused.Store(true)

	// 更新数据库状态
	if err := UpdateTaskPauseStatus(context.Background(), taskId, true); err != nil {
		e.isPaused.Store(false) // 恢复状态
		return fmt.Errorf("failed to update task pause status: %w", err)
	}

	g.Log().Info(context.Background(), "Task %d paused successfully", taskId)
	return nil
}

// ResumeTask 恢复任务
func (e *TaskExecutor) ResumeTask(taskId int) error {
	// 如果不是暂停状态，直接返回
	if !e.isPaused.Load() {
		return nil
	}

	// 恢复运行状态
	e.isPaused.Store(false)

	// 发送恢复信号
	select {
	case e.resumeChan <- struct{}{}:
		// 成功发送恢复信号
	default:
		// 通道可能已满，重新创建
		e.resumeChan = make(chan struct{}, 1)
		e.resumeChan <- struct{}{}
	}

	// 更新数据库状态
	if err := UpdateTaskPauseStatus(context.Background(), taskId, false); err != nil {
		e.isPaused.Store(true) // 恢复状态
		return fmt.Errorf("failed to update task resume status: %w", err)
	}

	g.Log().Info(context.Background(), "Task %d resumed successfully", taskId)
	return nil
}

// getTaskIdFromContext 从上下文获取任务ID
func (e *TaskExecutor) getTaskIdFromContext(ctx context.Context) (int, error) {
	for id, executor := range taskExecutors {
		if executor == e {
			return id, nil
		}
	}
	return 0, errors.New("task id not found in context")
}

// configureRateController 配置速率控制器
func (e *TaskExecutor) configureRateController(task *entity.EmailTask) {
	// 设置合理的发送速率
	maxPerMinute := task.Threads * 60 // 默认每线程每秒1封
	if maxPerMinute <= 0 {
		maxPerMinute = 600 // 默认每分钟600封
	}

	e.rateController = NewSimpleRateController(maxPerMinute)
}

// processTaskRecipients 处理任务收件人  具体发送
func (e *TaskExecutor) processTaskRecipients(ctx context.Context, task *entity.EmailTask, emailContent string) error {
	const batchSize = 100 // 每次获取100个收件人
	var lastId = 0

	for {
		// 检查上下文是否取消
		select {
		case <-ctx.Done():
			fmt.Println("上下文取消，停止任务执行1:", ctx.Err())
			return ctx.Err()
		default:
		}

		// 检查暂停状态
		if e.isPaused.Load() {
			g.Log().Debug(ctx, "Task %d is paused, waiting for resume signal", task.Id)

			// 等待恢复信号
			select {
			case <-e.resumeChan:
				g.Log().Debug(ctx, "Task %d received resume signal", task.Id)
			case <-ctx.Done():
				fmt.Println("上下文取消，停止任务执行2:", ctx.Err())
				return ctx.Err()
			}
		}

		// 获取一批未发送的收件人
		recipients, err := e.getNextRecipientBatch(ctx, task.Id, lastId, batchSize)
		if err != nil {
			return fmt.Errorf("failed to get recipients: %w", err)
		}

		// 没有更多收件人，退出循环
		if len(recipients) == 0 {
			break
		}

		// 更新最后ID
		lastId = recipients[len(recipients)-1].Id

		// 处理这批收件人 todo 取要发送的人
		if err := e.processRecipientBatch(ctx, task, recipients, emailContent); err != nil {
			return err
		}

		// 调整发送速率
		e.rateController.AdjustRate()
	}

	// 等待所有任务完成
	e.wg.Wait()
	return nil
}

// getNextRecipientBatch 获取下一批收件人
func (e *TaskExecutor) getNextRecipientBatch(ctx context.Context, taskId, lastId, batchSize int) ([]*entity.RecipientInfo, error) {
	var recipients []*entity.RecipientInfo

	// 使用ID分页，效率更高
	err := g.DB().Model("recipient_info").
		Where("task_id", taskId).
		Where("is_sent", 0).
		Where("id > ?", lastId).
		Order("id ASC").
		Limit(batchSize).
		Scan(&recipients)

	return recipients, err
}

// processRecipientBatch 处理一批收件人
func (e *TaskExecutor) processRecipientBatch(ctx context.Context, task *entity.EmailTask, recipients []*entity.RecipientInfo, emailContent string) error {
	// 创建结果通道，缓冲区大小与收件人数量相同
	resultChan := make(chan *SendResult, len(recipients))

	// 为每个收件人提交发送任务
	submittedCount := 0
	for _, recipient := range recipients {
		// 再次检查是否暂停或取消
		if e.isPaused.Load() {
			select {
			case <-e.resumeChan:
				// 已恢复
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// 等待速率控制
		if err := e.rateController.Wait(ctx); err != nil {
			if errors.Is(err, context.Canceled) {
				return err
			}
			// 记录错误但继续
			g.Log().Warning(ctx, "Rate limit wait error: %v", err)
		}

		// 创建收件人副本避免闭包问题
		recipient := recipient

		// 增加等待计数
		e.wg.Add(1)
		submittedCount++

		// 提交到工作池
		err := e.pool.Submit(func() {
			defer e.wg.Done()

			// 个性化内容
			personalized := e.personalizeEmail(ctx, emailContent, task, recipient)

			// 发送邮件
			result := e.sendEmail(ctx, task, recipient, personalized)

			// 记录发送
			e.rateController.RecordSend()

			// 更新统计
			if result.Success {
				e.sentCount.Add(1)
			} else {
				e.failedCount.Add(1)
			}

			// 发送结果
			select {
			case resultChan <- result:
				// 成功发送结果
			case <-ctx.Done():
				// 上下文取消
			}
		})

		if err != nil {
			e.wg.Done() // 减少等待计数
			submittedCount--

			// 创建失败结果
			failResult := &SendResult{
				RecipientID: recipient.Id,
				Success:     false,
				Error:       fmt.Errorf("failed to submit to worker pool: %w", err),
			}

			// 发送结果
			select {
			case resultChan <- failResult:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	// 所有任务已提交，现在启动结果处理
	resultsDone := make(chan struct{})
	go func() {
		e.processSendResults(ctx, resultChan)
		close(resultsDone)
	}()

	// 单独关闭通道的goroutine
	go func() {
		// 等待所有已提交的任务完成
		for i := 0; i < submittedCount; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				// 用空结构体模拟"join"操作
				var wg sync.WaitGroup
				wg.Add(1)
				e.pool.Submit(func() {
					defer wg.Done()
				})
				wg.Wait()
			}
		}
		// 所有任务完成后关闭结果通道
		close(resultChan)
	}()

	// 等待结果处理完成或上下文取消
	select {
	case <-resultsDone:
		// 结果处理完成
	case <-ctx.Done():
		return ctx.Err()
	}

	return nil
}

// processSendResults 处理发送结果  分批处理 更新已发送的状态和message_id
func (e *TaskExecutor) processSendResults(ctx context.Context, resultChan <-chan *SendResult) {
	const batchSize = 100                        // 批量处理大小
	const flushInterval = 500 * time.Millisecond // 定时刷新间隔

	successResults := make([]*SendResult, 0, batchSize) // 成功的发送结果
	failedIDs := make([]int, 0, batchSize)              // 失败的ID列表

	// 创建定时刷新器
	ticker := time.NewTicker(flushInterval)
	defer ticker.Stop()

	// 刷新函数
	flushUpdates := func() {
		if len(successResults) == 0 && len(failedIDs) == 0 {
			return
		}

		// 处理成功记录
		if len(successResults) > 0 {
			// 准备批量插入/更新的数据
			batchData := make([]g.Map, 0, len(successResults))
			for _, result := range successResults {
				batchData = append(batchData, g.Map{
					"id":         result.RecipientID, // 主键
					"is_sent":    1,
					"sent_time":  time.Now().Unix(),
					"message_id": result.MessageID,
				})
			}

			// 使用Save方法进行批量插入/更新
			// OnDuplicate指定在遇到重复键时需要更新的字段
			_, err := g.DB().Model("recipient_info").
				Data(batchData).
				OnDuplicate("is_sent, sent_time, message_id").
				Save()

			if err != nil {
				g.Log().Error(ctx, "批量更新收件人状态失败: %v", err)

				// 如果批量更新失败，可以尝试单个更新作为备选方案
				for _, result := range successResults {
					_, err := g.DB().Model("recipient_info").
						Where("id", result.RecipientID).
						Data(g.Map{
							"is_sent":    1,
							"sent_time":  time.Now().Unix(),
							"message_id": result.MessageID,
						}).
						Update()

					if err != nil {
						g.Log().Error(ctx, "单个更新收件人(ID:%d)状态失败: %v",
							result.RecipientID, err)
					}
				}
			} else {
				g.Log().Debug(ctx, "成功批量更新 %d 个收件人状态", len(successResults))
			}

			// 清空成功结果
			successResults = successResults[:0]
		}

		// 清空失败记录
		if len(failedIDs) > 0 {
			failedIDs = failedIDs[:0]
		}
	}

	// 主循环
	for {
		select {
		case result, ok := <-resultChan:
			if !ok {
				// 通道已关闭，处理剩余结果
				flushUpdates()
				return
			}

			if result.Success {
				successResults = append(successResults, result)
			} else {
				failedIDs = append(failedIDs, result.RecipientID)
				g.Log().Warning(ctx, "发送邮件到收件人 %d 失败: %v",
					result.RecipientID, result.Error)
			}

			// 达到批处理大小则刷新
			if len(successResults)+len(failedIDs) >= batchSize {
				flushUpdates()
			}

		case <-ticker.C:
			// 定时刷新
			flushUpdates()

		case <-ctx.Done():
			// 上下文取消，处理剩余结果
			flushUpdates()
			return
		}
	}
}

// getTemplateInfo 获取模板信息
func (e *TaskExecutor) getTemplateInfo(ctx context.Context, templateId int) (*entity.EmailTemplate, error) {
	var template entity.EmailTemplate

	err := g.DB().Model("email_templates").
		Where("id", templateId).
		Scan(&template)

	if err != nil {
		return nil, err
	}

	if template.Id == 0 {
		return nil, fmt.Errorf("template %d not found", templateId)
	}

	return &template, nil
}

// processEmailContent 处理邮件内容
func (e *TaskExecutor) processEmailContent(ctx context.Context, content string, task *entity.EmailTask) string {
	// 处理退订链接
	if task.Unsubscribe == 1 {
		// 获取当前域名
		domain := g.Cfg().MustGet(ctx, "server.domain").String()
		if domain == "" {
			domain = g.Cfg().MustGet(ctx, "server.address").String()
		}

		// 生成退订URL占位符
		unsubscribeURL := fmt.Sprintf("http://%s/unsubscribe", domain)
		groupURL := fmt.Sprintf("http://%s/unsubscribe/user_group", domain)

		unsubscribeJumpURL := fmt.Sprintf("%s/unsubscribe.html?jwt=__JWT__&email=__EMAIL__&url_type=%s&url_unsubscribe=%s",
			domain, groupURL, unsubscribeURL)
		fmt.Printf("处理邮件内容 unsubscribeJumpURL: %s\n", unsubscribeJumpURL)
		// 替换退订链接占位符
		content = strings.ReplaceAll(content, "__UNSUBSCRIBE_URL__", unsubscribeJumpURL)
	}

	return content
}

// personalizeEmail 个性化邮件内容
func (e *TaskExecutor) personalizeEmail(ctx context.Context, content string, task *entity.EmailTask, recipient *entity.RecipientInfo) string {
	// 处理退订JWT
	if task.Unsubscribe == 1 {
		// 生成JWT
		jwtToken, err := GenerateUnsubscribeJWT(
			recipient.Recipient,
			recipient.TaskId,
			task.TemplateId,
			task.Id,
		)

		if err != nil {
			g.Log().Error(ctx, "Failed to generate unsubscribe JWT: %v", err)
		} else {
			// 替换JWT和邮箱
			content = strings.ReplaceAll(content, "__JWT__", jwtToken)
			content = strings.ReplaceAll(content, "__EMAIL__", recipient.Recipient)
		}
	}

	return content
}

// sendEmail 发送邮件
func (e *TaskExecutor) sendEmail(ctx context.Context, task *entity.EmailTask, recipient *entity.RecipientInfo, content string) *SendResult {
	// 生成消息ID
	domain := strings.Split(task.Addresser, "@")
	messageDomain := "localhost"
	if len(domain) == 2 {
		messageDomain = domain[1]
	}

	messageID := fmt.Sprintf("<%s@%s>", uuid.New().String(), messageDomain)

	// TODO: 实现实际的邮件发送逻辑
	// 这里仅模拟发送操作
	time.Sleep(10 * time.Millisecond)

	// 创建并返回结果
	return &SendResult{
		RecipientID: recipient.Id,
		MessageID:   messageID,
		Success:     true,
		Error:       nil,
	}
}

// isTaskComplete 检查任务是否完成
func (e *TaskExecutor) isTaskComplete(ctx context.Context, taskId int) (bool, error) {
	type CountResult struct {
		TotalCount int `json:"total_count"`
		SentCount  int `json:"sent_count"`
	}

	var result CountResult

	// 使用一次查询获取总数和已发送数
	err := g.DB().Model("recipient_info").
		Fields("COUNT(1) as total_count, SUM(CASE WHEN is_sent = 1 THEN 1 ELSE 0 END) as sent_count").
		Where("task_id", taskId).
		Scan(&result)

	if err != nil {
		return false, err
	}

	// 如果没有收件人，任务不能算完成
	if result.TotalCount == 0 {
		return false, nil
	}

	// 如果已发送数等于或超过总数，任务完成
	return result.SentCount >= result.TotalCount, nil
}

// GetMetrics 获取执行指标
func (e *TaskExecutor) GetMetrics() map[string]interface{} {
	duration := time.Since(e.startTime).Seconds()
	sent := e.sentCount.Load()
	failed := e.failedCount.Load()
	total := sent + failed

	var successRate float64

	if total > 0 {
		successRate = float64(sent) / float64(total)
	}

	return map[string]interface{}{
		"sent_count":    sent,
		"failed_count":  failed,
		"total_count":   total,
		"success_rate":  successRate,
		"current_speed": e.rateController.GetCurrentRate(),
		"max_rate":      e.rateController.GetMaxRate(),
		"duration_sec":  duration,
	}
}

// UpdateTaskSpeed 更新任务发送速度
func (e *TaskExecutor) UpdateTaskSpeed(taskId int, speedFactor float64) error {
	// 参数验证
	if speedFactor <= 0 {
		return fmt.Errorf("speed factor must be greater than zero")
	}

	// 获取任务信息
	task, err := GetTaskInfo(context.Background(), taskId)
	if err != nil {
		return fmt.Errorf("failed to get task info: %w", err)
	}

	if task == nil || task.Id == 0 {
		return fmt.Errorf("task %d not found", taskId)
	}

	// 计算新的每分钟最大发送数
	// 基础速率 = 线程数 * 60
	baseRate := task.Threads * 60
	if baseRate <= 0 {
		baseRate = 600 // 默认值
	}

	// 新速率 = 基础速率 * 速度因子
	newRate := int(float64(baseRate) * speedFactor)

	// 设置最小和最大限制
	minRate := 60   // 最低每分钟60封
	maxRate := 6000 // 最高每分钟6000封

	if newRate < minRate {
		newRate = minRate
	} else if newRate > maxRate {
		newRate = maxRate
	}

	// 创建新的速率控制器
	e.rateController = NewSimpleRateController(newRate)

	g.Log().Info(context.Background(), "Task %d speed updated to %d emails per minute (factor: %.2f)",
		taskId, newRate, speedFactor)

	return nil
}
