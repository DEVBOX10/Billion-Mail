package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/service/batch_mail"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) ResumeTask(ctx context.Context, req *v1.ResumeTaskReq) (res *v1.ResumeTaskRes, err error) {
	res = &v1.ResumeTaskRes{}

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

	// 检查任务状态，只有暂停的进行中任务才能恢复
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

		res.SetError(gerror.New(public.LangCtx(ctx, "Cannot resume task because it is {}", statusText)))
		return
	}

	// 如果任务未暂停，直接返回成功
	if taskInfo.Pause == 0 {
		res.SetSuccess(public.LangCtx(ctx, "Task is already running"))
		return
	}

	// 获取任务执行器
	executor := batch_mail.GetTaskExecutor(req.TaskId)
	if executor == nil {
		// 没有活动执行器，直接更新数据库状态
		if err := batch_mail.UpdateTaskPauseStatus(ctx, req.TaskId, false); err != nil {
			res.Code = 500
			res.SetError(gerror.New(public.LangCtx(ctx, "Failed to resume task: {}", err.Error())))
			return nil, err
		}

		// 触发新一轮的任务处理
		go func() {
			// 创建新的执行器并处理任务
			newExecutor := batch_mail.GetOrCreateTaskExecutor(context.Background(), req.TaskId)
			batch_mail.RegisterTaskExecutor(req.TaskId, newExecutor)
			if err := newExecutor.ProcessTask(context.Background()); err != nil {
				g.Log().Error(context.Background(), "Failed to process resumed task: %v", err)
			}
		}()
	} else {
		// 通过执行器恢复任务
		if err := executor.ResumeTask(req.TaskId); err != nil {
			res.Code = 500
			res.SetError(gerror.New(public.LangCtx(ctx, "Failed to resume task: {}", err.Error())))
			return nil, err
		}
	}

	res.SetSuccess(public.LangCtx(ctx, "Task resumed successfully"))
	return
}
