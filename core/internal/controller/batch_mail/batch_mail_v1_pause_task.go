package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/service/batch_mail"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) PauseTask(ctx context.Context, req *v1.PauseTaskReq) (res *v1.PauseTaskRes, err error) {
	res = &v1.PauseTaskRes{}

	// 验证任务是否存在
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

	// 检查任务状态，只有进行中的任务才能暂停
	if taskInfo.TaskProcess != 1 { // 1表示进行中
		res.Code = 400

		var statusText string
		switch taskInfo.TaskProcess {
		case 0:
			statusText = "not started"
		case 2:
			statusText = "already completed"
		default:
			statusText = "in unknown state"
		}

		res.SetError(gerror.New(public.LangCtx(ctx, "Cannot pause task because it is {}", statusText)))
		return
	}

	// 如果任务已经暂停，直接返回成功
	if taskInfo.Pause == 1 {
		res.SetSuccess(public.LangCtx(ctx, "Task is already paused"))
		return
	}

	// 获取任务执行器
	executor := batch_mail.GetTaskExecutor(req.TaskId)
	if executor == nil {
		// 没有活动执行器，直接更新数据库状态
		if err := batch_mail.UpdateTaskPauseStatus(ctx, req.TaskId, true); err != nil {
			res.Code = 500
			res.SetError(gerror.New(public.LangCtx(ctx, "Failed to pause task: {}", err.Error())))
			return nil, err
		}
	} else {
		// 通过执行器暂停任务
		if err := executor.PauseTask(req.TaskId); err != nil {
			res.Code = 500
			res.SetError(gerror.New(public.LangCtx(ctx, "Failed to pause task: {}", err.Error())))
			return nil, err
		}
	}

	res.SetSuccess(public.LangCtx(ctx, "Task paused successfully"))
	return
}
