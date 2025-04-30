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

// ============= task related operations =============

// GetTasksWithPage get task list (pagination)
func GetTasksWithPage(ctx context.Context, page, pageSize int, keyword string, status int) (total int, list []*v1.EmailTask, err error) {
	// default pagination parameters
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	model := g.DB().Model("email_tasks").Safe()

	// add query conditions
	if keyword != "" {
		model = model.WhereLike("task_name", "%"+keyword+"%").
			WhereOrLike("subject", "%"+keyword+"%")
	}
	if status != -1 {
		model = model.Where("task_process", status)
	}

	// get total
	total, err = model.Count()
	if err != nil {
		return 0, nil, err
	}

	// pagination query
	list = make([]*v1.EmailTask, 0)
	err = model.Page(page, pageSize).
		Order("create_time DESC").
		Scan(&list)

	return total, list, err
}

// DeleteTask
func DeleteTask(ctx context.Context, id int) error {
	// delete task before removing task executor
	RemoveTaskExecutor(id)

	_, err := g.DB().Model("email_tasks").
		Where("id", id).
		Delete()
	return err
}

// CreateTask create task
func CreateTask(ctx context.Context, addresser, subject, fullName string, templateId int,
	isRecord, unsubscribe, threads, trackOpen, trackClick int, etypes string, remark string, startTime int, add_type int) (int, error) {

	now := time.Now().Unix()
	// generate task name
	taskName := "task_" + gconv.String(now)
	result, err := g.DB().Model("email_tasks").Insert(g.Map{
		"task_name":       taskName,
		"addresser":       addresser,
		"subject":         subject,
		"full_name":       fullName,
		"recipient_count": 0, // initial 0, update later
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

// UpdateRecipientCount update recipient count
func UpdateRecipientCount(ctx context.Context, taskId, count int) error {
	_, err := g.DB().Model("email_tasks").
		Where("id", taskId).
		Data(g.Map{"recipient_count": count}).
		Update()
	return err
}

// ============= recipient related operations =============

// GetSentCount get sent count
func GetSentCount(ctx context.Context, taskId int) (int, error) {
	return g.DB().Model("recipient_info").
		Where("task_id", taskId).
		Where("is_sent", 1).
		Count()
}

// ImportRecipients import recipient information with batch processing
func ImportRecipients(ctx context.Context, taskId int, contacts []*entity.Contact) error {
	if len(contacts) == 0 {
		return nil
	}

	// Set batch size to avoid PostgreSQL limitations
	const batchSize = 1000

	totalBatches := (len(contacts) + batchSize - 1) / batchSize

	totalImported := 0
	now := time.Now().Unix()

	for i := 0; i < totalBatches; i++ {
		// Calculate start and end indices for current batch
		startIdx := i * batchSize
		endIdx := (i + 1) * batchSize
		if endIdx > len(contacts) {
			endIdx = len(contacts)
		}

		// Get current batch
		currentBatch := contacts[startIdx:endIdx]

		// Prepare data for current batch
		values := make([]g.Map, len(currentBatch))
		for j, contact := range currentBatch {
			values[j] = g.Map{
				"task_id":     taskId,
				"recipient":   contact.Email,
				"is_sent":     0,
				"sent_time":   0,
				"message_id":  "",
				"create_time": now,
			}
		}

		// Insert current batch
		result, err := g.DB().Model("recipient_info").InsertIgnore(values)
		if err != nil {
			g.Log().Error(ctx, "Failed to import recipient batch %d/%d for task %d: %v",
				i+1, totalBatches, taskId, err)
			return fmt.Errorf("failed to import recipients (batch %d/%d): %w", i+1, totalBatches, err)
		}

		// Count affected rows
		affected, err := result.RowsAffected()
		if err != nil {
			g.Log().Warning(ctx, "Could not get affected rows for batch %d/%d: %v", i+1, totalBatches, err)
		} else {
			totalImported += int(affected)
			g.Log().Debug(ctx, "Task %d: Batch %d/%d imported %d recipients successfully",
				taskId, i+1, totalBatches, affected)
		}
	}

	g.Log().Info(ctx, "Task %d: Total %d recipients imported successfully", taskId, totalImported)
	return nil
}

// ============= contact group related operations =============

// GetGroupInfo get group info
func GetGroupInfo(ctx context.Context, groupId int) (*v1.GroupInfo, error) {
	// use struct map instead of gdb.Record
	var result struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Count int    `json:"count"`
	}

	err := g.DB().Model("bm_contact_groups cg").
		LeftJoin("bm_contacts c", "cg.id = c.group_id").
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

// GetActiveContacts get active contacts in group
func GetActiveContacts(ctx context.Context, groupId int) ([]*entity.Contact, error) {
	var contacts []*entity.Contact
	err := g.DB().Model("bm_contacts").
		Where("group_id", groupId).
		Where("active", 1).
		Scan(&contacts)
	return contacts, err
}

// ============= business logic combination =============
// add type parameter add_type default 0
func CreateTaskWithRecipients(ctx context.Context, req *v1.CreateTaskReq, addType int) (int, error) {
	var taskId int
	var err error
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		etype := strings.Join(gconv.SliceStr(req.GroupIds), ",")

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
			req.StartTime,
			addType,
		)
		if err != nil {
			return gerror.New(public.LangCtx(ctx, "Failed to create task {}", err.Error()))
		}

		var abnormalRecipients []struct {
			Recipient string `json:"recipient"`
			Count     int    `json:"count"`
		}

		err = tx.Model("abnormal_recipient").
			Where("count >= ?", 3).
			Fields("recipient, count").
			Scan(&abnormalRecipients)

		if err != nil {
			g.Log().Warning(ctx, "Failed to get the exception recipient list: %v", err)
		}

		abnormalMap := make(map[string]int)
		if len(abnormalRecipients) > 0 {
			for _, ar := range abnormalRecipients {
				abnormalMap[ar.Recipient] = ar.Count
			}
		}

		// import recipient information
		totalRecipients := 0
		totalSkipped := 0

		for _, groupId := range req.GroupIds {
			contacts, err := GetActiveContacts(ctx, groupId)
			if err != nil {
				return gerror.New(public.LangCtx(ctx, "Failed to get contacts for group {}: {}", groupId, err.Error()))
			}
			if len(contacts) == 0 {
				continue
			}

			// abnormal skip
			filteredContacts := make([]*entity.Contact, 0, len(contacts))
			skippedInGroup := 0

			for _, contact := range contacts {
				if _, exists := abnormalMap[contact.Email]; exists {
					skippedInGroup++
					continue
				}
				filteredContacts = append(filteredContacts, contact)
			}

			if skippedInGroup > 0 {
				totalSkipped += skippedInGroup

			}

			if len(filteredContacts) == 0 {
				continue
			}

			err = ImportRecipients(ctx, taskId, filteredContacts)
			if err != nil {
				return gerror.New(public.LangCtx(ctx, "Failed to import recipients for group {}: {}", groupId, err.Error()))
			}

			totalRecipients += len(filteredContacts)
		}

		// update recipient count
		return UpdateRecipientCount(ctx, taskId, totalRecipients)
	})

	if err != nil {
		return 0, err
	}

	return taskId, nil
}

// GetTaskInfo get task info
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

// UpdateTaskPauseStatus update task pause status
func UpdateTaskPauseStatus(ctx context.Context, taskId int, isPaused bool) error {
	pauseValue := 0
	processValue := 1

	if isPaused {
		pauseValue = 1
		processValue = 3
	}

	_, err := g.DB().Model("email_tasks").
		Where("id", taskId).
		Data(g.Map{
			"pause":        pauseValue,
			"task_process": processValue,
		}).
		Update()

	if err != nil {
		return fmt.Errorf("failed to update task pause status: %w", err)
	}

	g.Log().Info(ctx, "Updated task %d pause status: isPaused=%v, task_process=%d", taskId, isPaused, processValue)
	return nil
}

// UpdateTaskProcessStatus update task process status
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

// GetTaskSendingStats get task sending stats (success count and failed count)
func GetTaskSendingStats(ctx context.Context, taskID int) (int, int, error) {
	if taskID <= 0 {
		return 0, 0, nil
	}

	// query success count
	successQuery := g.DB().Model("mailstat_send_mails sm")
	successQuery = successQuery.LeftJoin("mailstat_message_ids mid", "sm.postfix_message_id=mid.postfix_message_id")
	successQuery = successQuery.LeftJoin("recipient_info ri", "mid.message_id=ri.message_id")
	successQuery = successQuery.Where("ri.task_id", taskID)
	successQuery = successQuery.Where("sm.status", "sent")
	successQuery = successQuery.Where("sm.dsn LIKE '2.%'")
	successCount, err := successQuery.Count()
	if err != nil {
		// fmt.Println("query success count failed:", err)
		return 0, 0, err
	}

	// query failed count bounced
	failedQuery := g.DB().Model("mailstat_send_mails sm")
	failedQuery = failedQuery.LeftJoin("mailstat_message_ids mid", "sm.postfix_message_id=mid.postfix_message_id")
	failedQuery = failedQuery.LeftJoin("recipient_info ri", "mid.message_id=ri.message_id")
	failedQuery = failedQuery.Where("ri.task_id", taskID)
	failedQuery = failedQuery.Where("sm.status = 'bounced'")
	//failedQuery = failedQuery.Where("sm.status = 'deferred'")
	failedCount, err := failedQuery.Count()
	if err != nil {
		// fmt.Println("query failed count failed:", err)
		return successCount, 0, err
	}
	// fmt.Println("success count:", successCount, "failed count:", failedCount)
	return successCount, failedCount, nil
}

func GetActualRecipientCount(ctx context.Context, taskId int) (int, error) {

	count, err := g.DB().Model("recipient_info").
		Where("task_id", taskId).
		Count()

	if err != nil {
		return 0, fmt.Errorf("failed to get recipient count: %w", err)
	}

	return count, nil
}
