package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/batch_mail"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) UpdateTaskSpeed(ctx context.Context, req *v1.UpdateTaskSpeedReq) (res *v1.UpdateTaskSpeedRes, err error) {
	res = &v1.UpdateTaskSpeedRes{}

	// 参数验证
	if req.Speed <= 0 {
		res.Code = 400
		res.SetError(gerror.New(public.LangCtx(ctx, "Speed factor must be greater than zero")))
		return
	}

	// 获取任务信息
	var taskInfo *entity.EmailTask
	taskInfo, err = batch_mail.GetTaskInfo(ctx, req.TaskId)
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

	// 获取任务执行器
	executor := batch_mail.GetTaskExecutor(req.TaskId)
	if executor == nil {
		res.Code = 404
		res.SetError(gerror.New(public.LangCtx(ctx, "Task executor not found")))
		return
	}

	// 更新任务速度
	if err = executor.UpdateTaskSpeed(req.TaskId, req.Speed); err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to update task speed: {}", err.Error())))
		return
	}

	// 计算实际每分钟发送速率
	baseRate := taskInfo.Threads * 60
	if baseRate <= 0 {
		baseRate = 600
	}
	actualRate := int(float64(baseRate) * req.Speed)

	// 添加实际速率到响应
	res.Data = map[string]interface{}{
		"task_id":           req.TaskId,
		"speed_factor":      req.Speed,
		"emails_per_minute": actualRate,
	}

	res.SetSuccess(public.LangCtx(ctx, "Task speed updated successfully"))
	return
}
