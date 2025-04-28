// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package batch_mail

import (
	"context"

	"billionmail-core/api/batch_mail/v1"
)

type IBatchMailV1 interface {
	ListTasks(ctx context.Context, req *v1.ListTasksReq) (res *v1.ListTasksRes, err error)
	DeleteTask(ctx context.Context, req *v1.DeleteTaskReq) (res *v1.DeleteTaskRes, err error)
	CreateTask(ctx context.Context, req *v1.CreateTaskReq) (res *v1.CreateTaskRes, err error)
	GetTaskSendCount(ctx context.Context, req *v1.GetTaskSendCountReq) (res *v1.GetTaskSendCountRes, err error)
	TaskMailProviderStat(ctx context.Context, req *v1.TaskMailProviderStatReq) (res *v1.TaskMailProviderStatRes, err error)
	GetTaskMailLogs(ctx context.Context, req *v1.GetTaskMailLogsReq) (res *v1.GetTaskMailLogsRes, err error)
	SendTestEmail(ctx context.Context, req *v1.SendTestEmailReq) (res *v1.SendTestEmailRes, err error)
	UpdateTaskSpeed(ctx context.Context, req *v1.UpdateTaskSpeedReq) (res *v1.UpdateTaskSpeedRes, err error)
	PauseTask(ctx context.Context, req *v1.PauseTaskReq) (res *v1.PauseTaskRes, err error)
	ResumeTask(ctx context.Context, req *v1.ResumeTaskReq) (res *v1.ResumeTaskRes, err error)
	Unsubscribe(ctx context.Context, req *v1.UnsubscribeReq) (res *v1.UnsubscribeRes, err error)
	GetUserGroups(ctx context.Context, req *v1.GetUserGroupsReq) (res *v1.GetUserGroupsRes, err error)
}
