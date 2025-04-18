package batch_mail

import (
	"billionmail-core/internal/service/batch_mail"
	"billionmail-core/internal/service/email_template"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
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

		// 获取发送进度 progress 满分100

		sentCount, _ := batch_mail.GetSentCount(ctx, task.Id)
		detail.SentCount = sentCount
		detail.UnsentCount = task.RecipientCount - sentCount
		detail.Progress = int(float64(sentCount) / float64(task.RecipientCount) * 100)

		// 通过 recipient_info  查询当前任务收件人的 message-id
		detail.SuccessCount = 0
		detail.ErrorCount = 0
		// 根据 message-id  查日志表的日志记录
		// 记录中统计 成功数和失败数

		task_list[i] = detail
	}

	res.Data.Total = total
	res.Data.List = task_list
	res.SetSuccess(public.LangCtx(ctx, "Task list retrieved successfully"))
	return

}
