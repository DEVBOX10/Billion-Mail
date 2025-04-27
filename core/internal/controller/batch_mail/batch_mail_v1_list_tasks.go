package batch_mail

import (
	"billionmail-core/internal/service/batch_mail"
	"billionmail-core/internal/service/contact"
	"billionmail-core/internal/service/email_template"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"

	"billionmail-core/api/batch_mail/v1"
)

func (c *ControllerV1) ListTasks(ctx context.Context, req *v1.ListTasksReq) (res *v1.ListTasksRes, err error) {

	res = &v1.ListTasksRes{}

	total, tasks, err := batch_mail.GetTasksWithPage(ctx, req.Page, req.PageSize, req.Keyword, req.Status)
	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get task list {}", err.Error())))
		return
	}

	task_list := make([]*v1.TaskDetail, len(tasks))
	for i, task := range tasks {
		detail := &v1.TaskDetail{
			EmailTask: *task,
		}

		template, err := email_template.GetTemplate(ctx, task.TemplateId)
		if err == nil {
			detail.TemplateName = template.TempName
		}

		groupIds := strings.Split(task.Etypes, ",")

		groups := make([]*v1.GroupInfo, 0)
		for _, groupId := range groupIds {
			groupIdInt := gconv.Int(groupId)

			groupInfo, err := batch_mail.GetGroupInfo(ctx, groupIdInt)
			if err != nil {
				continue
			}
			groups = append(groups, groupInfo)
		}
		detail.Groups = groups

		// Get actual sent count
		sentCount, _ := batch_mail.GetSentCount(ctx, task.Id)
		detail.SentCount = sentCount

		// Fix: Check if recipient_count is valid
		if task.RecipientCount <= 0 {
			groupIdsInt := make([]int, 0, len(groupIds))
			for _, groupId := range groupIds {
				groupIdsInt = append(groupIdsInt, gconv.Int(groupId))
			}

			actualCount, err := contact.GetGroupContactCount(ctx, groupIdsInt)
			if err == nil && actualCount > 0 {
				// Update the recipient count in the database
				_ = batch_mail.UpdateRecipientCount(ctx, task.Id, actualCount)
				detail.RecipientCount = actualCount
				task.RecipientCount = actualCount
			}
		}

		// Now calculate unsent_count with the corrected recipient_count
		if task.RecipientCount >= sentCount {
			detail.UnsentCount = task.RecipientCount - sentCount
		} else {
			// If still inconsistent, set unsent count to 0
			detail.UnsentCount = 0
			g.Log().Warning(ctx, "Task %d has inconsistent data: recipient_count(%d) < sent_count(%d)",
				task.Id, task.RecipientCount, sentCount)
		}

		// Calculate progress with safeguards against division by zero
		if task.RecipientCount > 0 {
			detail.Progress = int(float64(sentCount) / float64(task.RecipientCount) * 100)
			// Ensure progress is between 0 and 100
			if detail.Progress > 100 {
				detail.Progress = 100
			}
		} else {
			// Default progress for tasks with no recipients
			detail.Progress = 0
		}

		// Get success and error counts
		successCount, errorCount, err := batch_mail.GetTaskSendingStats(ctx, task.Id)
		if err == nil {
			detail.SuccessCount = successCount
			detail.ErrorCount = errorCount
		} else {
			detail.SuccessCount = 0
			detail.ErrorCount = 0
			g.Log().Error(ctx, "Failed to get sending statistics: ", err)
		}

		task_list[i] = detail
	}

	res.Data.Total = total
	res.Data.List = task_list
	res.SetSuccess(public.LangCtx(ctx, "Task list retrieved successfully"))
	return
}
