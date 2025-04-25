package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/batch_mail"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

// UpdateTaskSpeed update task sending speed
func (c *ControllerV1) UpdateTaskSpeed(ctx context.Context, req *v1.UpdateTaskSpeedReq) (res *v1.UpdateTaskSpeedRes, err error) {
	res = &v1.UpdateTaskSpeedRes{}

	// parameter validation
	if req.Threads <= 0 {
		res.Code = 400
		res.SetError(gerror.New(public.LangCtx(ctx, "Threads must be greater than zero")))
		return
	}

	if req.Threads > 100 {
		res.Code = 400
		res.SetError(gerror.New(public.LangCtx(ctx, "Threads cannot exceed 100")))
		return
	}

	// get task info
	var taskInfo *entity.EmailTask
	taskInfo, err = batch_mail.GetTaskInfo(ctx, req.TaskId)
	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get task info: {}", err.Error())))
		return
	}

	if taskInfo == nil || taskInfo.Id == 0 {
		res.Code = 404
		res.SetError(gerror.New(public.LangCtx(ctx, "Task not found: {}", req.TaskId)))
		return
	}

	// get task executor
	executor := batch_mail.GetTaskExecutor(req.TaskId)
	if executor == nil {
		res.Code = 404
		res.SetError(gerror.New(public.LangCtx(ctx, "Task executor not found")))
		return
	}

	// call update task threads function
	if err = executor.UpdateTaskThreads(req.TaskId, req.Threads); err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to update task sending speed: {}", err.Error())))
		return
	}

	// calculate actual sending rate per minute
	actualRate := req.Threads * 20 * 60 // 20 emails per second per thread

	// add actual rate to response
	res.Data = map[string]interface{}{
		"task_id":           req.TaskId,
		"threads":           req.Threads,
		"emails_per_minute": actualRate,
	}

	res.SetSuccess(public.LangCtx(ctx, "Task sending speed updated successfully"))
	return
}
