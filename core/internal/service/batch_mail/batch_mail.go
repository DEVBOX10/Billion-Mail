package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
	"time"
)

// ============= 任务相关操作 =============

// GetTasksWithPage 获取任务列表（分页）
func GetTasksWithPage(ctx context.Context, page, pageSize int, keyword string, status int) (total int, list []*v1.EmailTask, err error) {
	// 默认分页参数
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	model := g.DB().Model("email_tasks").Safe()

	// 添加查询条件
	if keyword != "" {
		model = model.WhereLike("task_name", "%"+keyword+"%").
			WhereOrLike("subject", "%"+keyword+"%")
	}
	if status != -1 {
		model = model.Where("task_process", status)
	}

	// 获取总数
	total, err = model.Count()
	if err != nil {
		return 0, nil, err
	}

	// 分页查询
	list = make([]*v1.EmailTask, 0)
	err = model.Page(page, pageSize).
		Order("create_time DESC").
		Scan(&list)

	return total, list, err
}

// DeleteTask
func DeleteTask(ctx context.Context, id int) error {
	// 删除任务前先移除任务执行器
	RemoveTaskExecutor(id)

	_, err := g.DB().Model("email_tasks").
		Where("id", id).
		Delete()
	return err
}

// CreateTask 创建任务
func CreateTask(ctx context.Context, addresser, subject, fullName string, templateId int,
	isRecord, unsubscribe, threads, trackOpen, trackClick int, etypes string, remark string, startTime int, add_type int) (int, error) {

	now := time.Now().Unix()
	// 生成任务名称
	taskName := "task_" + gconv.String(now)
	result, err := g.DB().Model("email_tasks").Insert(g.Map{
		"task_name":       taskName,
		"addresser":       addresser,
		"subject":         subject,
		"full_name":       fullName,
		"recipient_count": 0, // 初始为0，后续更新
		"task_process":    0,
		"pause":           0,
		"template_id":     templateId,
		"is_record":       isRecord,
		"unsubscribe":     unsubscribe,
		"threads":         threads,
		"etypes":          etypes,
		"track_open":      trackOpen,
		"track_click":     trackClick,
		"start_time":      startTime,
		"create_time":     now,
		"update_time":     now,
		"active":          1,
		"remark":          remark,
		"add_type":        add_type,
	})
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

// UpdateRecipientCount 更新收件人数量
func UpdateRecipientCount(ctx context.Context, taskId, count int) error {
	_, err := g.DB().Model("email_tasks").
		Where("id", taskId).
		Data(g.Map{"recipient_count": count}).
		Update()
	return err
}

// ============= 收件人相关操作 =============

// GetSentCount 获取已发送数量
func GetSentCount(ctx context.Context, taskId int) (int, error) {
	return g.DB().Model("recipient_info").
		Where("task_id", taskId).
		Where("is_sent", 1).
		Count()
}

// ImportRecipients 导入收件人信息
func ImportRecipients(ctx context.Context, taskId int, contacts []*entity.Contact) error {
	if len(contacts) == 0 {
		return nil
	}

	now := time.Now().Unix()
	// 构建批量插入的数据
	values := make([]g.Map, len(contacts))
	for i, contact := range contacts {
		values[i] = g.Map{
			"task_id":     taskId,
			"recipient":   contact.Email,
			"is_sent":     0,
			"sent_time":   0,
			"message_id":  "",
			"create_time": now,
		}
	}

	// 批量插入数据
	_, err := g.DB().Model("recipient_info").InsertIgnore(values)
	return err
}

// ============= 联系人组相关操作 =============

// GetGroupInfo 获取组信息
func GetGroupInfo(ctx context.Context, groupId int) (*v1.GroupInfo, error) {
	// 使用结构体map代替gdb.Record
	var result struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Count int    `json:"count"`
	}

	err := g.DB().Model("contact_groups cg").
		LeftJoin("contacts c", "cg.id = c.group_id").
		Fields("cg.id, cg.name, COUNT(CASE WHEN c.active = 1 THEN 1 END) as count").
		Where("cg.id", groupId).
		Group("cg.id, cg.name").
		Scan(&result)

	if err != nil {
		return nil, err
	}

	return &v1.GroupInfo{
		Id:    result.Id,
		Name:  result.Name,
		Count: result.Count,
	}, nil
}

// GetActiveContacts 获取组内活跃联系人
func GetActiveContacts(ctx context.Context, groupId int) ([]*entity.Contact, error) {
	var contacts []*entity.Contact
	err := g.DB().Model("contacts").
		Where("group_id", groupId).
		Where("active", 1).
		Scan(&contacts)
	return contacts, err
}

// ============= 业务逻辑组合 =============
// 增加类型参数 add_type 默认0
func CreateTaskWithRecipients(ctx context.Context, req *v1.CreateTaskReq, addType int) (int, error) {
	//  如果没传开始时间 则默认当前时间
	if req.StartTime == 0 {
		req.StartTime = int(time.Now().Unix())
	}
	// 开启事务  创建任务 导入收件人 触发器 更新收件人数量
	var taskId int
	var err error
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		etype := strings.Join(gconv.SliceStr(req.GroupIds), ",")
		// 创建任务
		taskId, err = CreateTask(
			ctx,
			req.Addresser,
			req.Subject,
			req.FullName,
			req.TemplateId,
			req.IsRecord,
			req.Unsubscribe,
			req.Threads,
			req.TrackOpen,
			req.TrackClick,
			etype,
			req.Remark,
			int(time.Now().Unix()),
			addType,
		)
		if err != nil {
			return gerror.New(public.LangCtx(ctx, "Failed to create task {}", err.Error()))
		}

		// 导入收件人信息
		totalRecipients := 0
		for _, groupId := range req.GroupIds {
			contacts, err := GetActiveContacts(ctx, groupId)
			if err != nil {
				return gerror.New(public.LangCtx(ctx, "Failed to get contacts for group {}: {}", groupId, err.Error()))
			}
			if len(contacts) == 0 {
				continue
			}
			err = ImportRecipients(ctx, taskId, contacts)
			if err != nil {
				return gerror.New(public.LangCtx(ctx, "Failed to import recipients for group {}: {}", groupId, err.Error()))
			}

			totalRecipients += len(contacts)
		}

		if totalRecipients == 0 {
			return gerror.New("No recipients found in the selected groups")
		}

		// 注册任务执行器 todo 创建任务后要触发
		executor := NewTaskExecutor(ctx)
		RegisterTaskExecutor(taskId, executor)

		// 更新收件人数量
		return UpdateRecipientCount(ctx, taskId, totalRecipients)
	})

	if err != nil {
		return 0, err
	}

	return taskId, nil
}

// GetTaskInfo 获取任务信息
func GetTaskInfo(ctx context.Context, taskId int) (*entity.EmailTask, error) {
	var task entity.EmailTask
	err := g.DB().Model("email_tasks").
		Where("id", taskId).
		Scan(&task)

	if err != nil {
		return nil, fmt.Errorf("failed to get task info: %w", err)
	}

	return &task, nil
}

// UpdateTaskPauseStatus 更新任务暂停状态
func UpdateTaskPauseStatus(ctx context.Context, taskId int, isPaused bool) error {
	pauseValue := 0
	if isPaused {
		pauseValue = 1
	}

	_, err := g.DB().Model("email_tasks").
		Where("id", taskId).
		Data(g.Map{"pause": pauseValue}).
		Update()

	if err != nil {
		return fmt.Errorf("failed to update task pause status: %w", err)
	}

	return nil
}

// UpdateTaskProcessStatus 更新任务处理状态
func UpdateTaskProcessStatus(ctx context.Context, taskId int, status int) error {
	_, err := g.DB().Model("email_tasks").
		Where("id", taskId).
		Data(g.Map{"task_process": status}).
		Update()

	if err != nil {
		return fmt.Errorf("failed to update task process status: %w", err)
	}

	return nil
}

// InitScheduler 初始化调度器
func InitScheduler(ctx context.Context) {
	// 定时清理空闲执行器
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				CleanupIdleExecutors()
			case <-ctx.Done():
				return
			}
		}
	}()

	g.Log().Debug(ctx, "Email task scheduler initialized")
}
