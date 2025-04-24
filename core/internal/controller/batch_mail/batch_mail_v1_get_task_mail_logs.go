package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/service/batch_mail"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func (c *ControllerV1) GetTaskMailLogs(ctx context.Context, req *v1.GetTaskMailLogsReq) (res *v1.GetTaskMailLogsRes, err error) {
	g.Log().Printf(ctx, "Getting task log details, TaskId: %d, Status: %d",
		req.TaskId, req.Status)

	res = &v1.GetTaskMailLogsRes{}

	// Verify task exists
	taskInfo, err := batch_mail.GetTaskInfo(ctx, req.TaskId)
	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get task information: {}", err.Error())))
		return
	}

	if taskInfo == nil || taskInfo.Id == 0 {
		res.Code = 404
		res.SetError(gerror.New(public.LangCtx(ctx, "Task not found: {}", req.TaskId)))
		return
	}

	// Build query conditions
	condition := g.Map{
		"ri.task_id": req.TaskId,
	}

	// Filter by status
	if req.Status == 1 {
		condition["sm.status"] = "sent"
		condition["sm.dsn LIKE"] = "2.%"
	} else if req.Status == 0 {
		// Keep consistent with GetTaskSendingStats
		condition["sm.status"] = "bounced"
	}

	startTime := int64(0)

	if taskInfo.StartTime > 0 {
		startTime = int64(taskInfo.StartTime)
	} else {
		startTime = time.Now().Add(-24 * time.Hour).Unix() // Default to last 24 hours
	}

	endTime := time.Now().Unix()

	// Set pagination
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	// Build base query
	baseQuery := g.DB().Model("mailstat_send_mails sm")
	baseQuery = baseQuery.LeftJoin("mailstat_message_ids mi", "sm.postfix_message_id=mi.postfix_message_id")
	baseQuery = baseQuery.LeftJoin("recipient_info ri", "mi.message_id=ri.message_id")

	// Ensure valid joins
	baseQuery.Where("mi.postfix_message_id IS NOT NULL")
	baseQuery.Where("ri.task_id IS NOT NULL")

	// Add time range filters
	if startTime > 0 {
		baseQuery.Where("sm.log_time_millis > ?", startTime*1000-1) // Convert to milliseconds
	}

	if endTime > 0 {
		baseQuery.Where("sm.log_time_millis < ?", endTime*1000+1) // Convert to milliseconds
	}

	// Add domain filter
	if req.Domain != "" {
		baseQuery.Where("sm.recipient LIKE ?", "%@"+req.Domain)
	}

	// Apply conditions
	baseQuery = baseQuery.Where(condition)

	// Get total count
	count, err := baseQuery.Count()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to get log count: %v", err)
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get log count: {}", err.Error())))
		return
	}

	var logs []*v1.TaskLogItem
	err = baseQuery.Fields(
		"sm.postfix_message_id",
		"sm.status",
		"sm.recipient", // Using sm.recipient directly instead of ri.recipient
		"sm.mail_provider",
		"sm.delay",
		"sm.delays",
		"sm.dsn",
		"sm.relay",
		"sm.description",
		"sm.log_time",
	).OrderDesc("sm.log_time_millis").
		Limit((req.Page-1)*req.PageSize, req.PageSize).
		Scan(&logs)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to query logs: %v", err)
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to query logs: {}", err.Error())))
		return
	}

	// Build response
	res.Data.Total = count
	res.Data.List = make([]*v1.TaskLogItem, 0)

	// Convert to API response format
	for _, log := range logs {
		//logTimeStr := time.Unix(log.LogTime, 0).Format("2006-01-02 15:04:05")

		res.Data.List = append(res.Data.List, &v1.TaskLogItem{
			PostfixMessageId: log.PostfixMessageId,
			Status:           log.Status,
			Recipient:        log.Recipient,
			MailProvider:     log.MailProvider,
			Delay:            log.Delay,
			Delays:           log.Delays,
			Dsn:              log.Dsn,
			Relay:            log.Relay,
			Description:      log.Description,
			LogTime:          log.LogTime,
		})
	}

	res.SetSuccess(public.LangCtx(ctx, "Task mail logs retrieved successfully"))
	return
}
