package timers

import (
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/batch_mail"
	"billionmail-core/internal/service/domains"
	"billionmail-core/internal/service/maillog_stat"
	"context"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"sync"
	"time"
)

var (
	// 全局任务执行器
	executor     *batch_mail.TaskExecutor
	executorOnce sync.Once
)

// getExecutor 获取单例任务执行器
func getExecutor(ctx context.Context) *batch_mail.TaskExecutor {
	executorOnce.Do(func() {
		executor = batch_mail.NewTaskExecutor(ctx)
	})
	return executor
}

// Start 启动所有定时任务 (主入口函数)
func Start(ctx context.Context) (err error) {
	g.Log().Debug(ctx, "Starting timers...")

	// Start DNS Records checking
	// Ensure the DNS records are fresh on startup
	gtimer.AddOnce(5*time.Second, func() {
		domains.FreshRecords(ctx)
	})
	// Check DNS Records every 5 minutes
	gtimer.Add(5*time.Minute, func() {
		domains.FreshRecords(ctx)
	})

	// Start maillog analysis
	gtimer.AddOnce(5*time.Second, func() {
		me := maillog_stat.NewMallogEventHandler("", 1*time.Second)
		me.Start()
	})

	// Start maillog aggregation
	gtimer.AddOnce(5*time.Second, func() {
		maillog_stat.AggregateMaillogsTask(1 * time.Minute)
	})

	g.Log().Debug(ctx, "Start timers complete")

	// 初始化批处理邮件任务调度器
	batch_mail.InitScheduler(ctx)

	// ========== 邮件日志分析任务 ==========
	gtimer.AddOnce(5*time.Second, func() {
		me := maillog_stat.NewMallogEventHandler("", 1*time.Second)
		me.Start()
	})

	//// ========== 邮件任务处理方式1: 单一执行器 ==========
	//// 获取任务执行器
	//taskExecutor := getExecutor(ctx)
	//// 每分钟检查并执行任务
	//gtimer.Add(1*time.Minute, func() {
	//	// 创建新的上下文，设置超时时间
	//	taskCtx, cancel := context.WithTimeout(ctx, 55*time.Second)
	//	defer cancel()
	//
	//	// 如果任务执行器已经在运行，则跳过本次执行
	//	if taskExecutor.IsRunning() {
	//		g.Log().Debug(taskCtx, "Task executor is busy, skipping this run")
	//		return
	//	}
	//
	//	// 执行任务处理
	//	if err := taskExecutor.ProcessTask(taskCtx); err != nil {
	//		if err == context.DeadlineExceeded {
	//			g.Log().Warning(taskCtx, "Task processing timed out")
	//		} else {
	//			g.Log().Error(taskCtx, "Failed to process tasks: %v", err)
	//		}
	//	}
	//})

	// ========== 邮件任务处理方式2: 每任务一个执行器 ==========
	// 定时处理邮件发送任务
	_, err = gcron.Add(ctx, "0 */2 * * * *", func(ctx context.Context) {
		processEmailTasks(ctx)
	})
	if err != nil {
		g.Log().Error(ctx, "Failed to add email task cron: %v", err)
	}

	// 每10分钟清理一次空闲执行器
	_, err = gcron.Add(ctx, "0 */10 * * * *", func(ctx context.Context) {
		batch_mail.CleanupIdleExecutors()
		g.Log().Debug(ctx, "Idle task executors cleanup completed")
	})
	if err != nil {
		g.Log().Error(ctx, "Failed to add executor cleanup cron: %v", err)
	}

	g.Log().Debug(ctx, "All timers started successfully")
	return nil
}

// processEmailTasks 处理邮件发送任务 (多执行器模式)
func processEmailTasks(ctx context.Context) {
	// 获取待处理的任务
	var tasks []*entity.EmailTask
	err := g.DB().Model("email_tasks").
		Where("task_process IN (0,1)").              // 未开始或进行中
		Where("pause", 0).                           // 未暂停
		Where("start_time <= ?", time.Now().Unix()). // 开始时间已到
		Order("id ASC").                             // 按ID排序
		Scan(&tasks)

	if err != nil {
		g.Log().Error(ctx, "Failed to get pending email tasks: %v", err)
		return
	}

	if len(tasks) == 0 {
		return
	}

	g.Log().Debug(ctx, "Found %d pending email tasks", len(tasks))

	// 处理每个任务
	for _, task := range tasks {
		// 检查任务是否已经有执行器并正在运行
		executor := batch_mail.GetTaskExecutor(task.Id)
		if executor != nil && executor.IsRunning() {
			continue // 跳过正在运行的任务
		}

		// 创建新的执行器
		newCtx := gctx.New()
		executor = batch_mail.NewTaskExecutor(newCtx)
		batch_mail.RegisterTaskExecutor(task.Id, executor)

		// 启动任务处理
		go func(taskId int) {
			if err := executor.ProcessTask(newCtx); err != nil {
				g.Log().Error(newCtx, "Error processing task %d: %v", taskId, err)
			}
		}(task.Id)
	}
}
