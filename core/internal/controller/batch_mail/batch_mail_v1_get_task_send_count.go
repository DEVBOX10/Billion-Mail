package batch_mail

import (
	"billionmail-core/internal/service/batch_mail"
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/service/public"
)

// GetTaskSendCount
func (c *ControllerV1) GetTaskSendCount(ctx context.Context, req *v1.GetTaskSendCountReq) (res *v1.GetTaskSendCountRes, err error) {
	res = &v1.GetTaskSendCountRes{}

	// validate task id
	if req.TaskId <= 0 {
		res.Code = 400
		res.SetError(gerror.New(public.LangCtx(ctx, "Invalid task ID")))
		return
	}

	// query task info
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

	// set time range
	startTime := time.Now().Add(-24 * time.Hour).Unix() // default query last 24 hours
	if req.StartTime > 0 {
		startTime = req.StartTime
	}

	endTime := time.Now().Unix()
	if req.EndTime > 0 {
		endTime = req.EndTime
	}

	// query send count per minute
	type StatItem struct {
		Minute    int64 `json:"minute"`     // minute timestamp (rounded to the nearest minute)
		SendCount int   `json:"send_count"` // send count
		StartTime int64 `json:"start_time"` // start time of the minute
		EndTime   int64 `json:"end_time"`   // end time of the minute
	}

	var statList []StatItem

	// SQL query, group by minute

	err = g.DB().Model("recipient_info").
		Where("task_id", req.TaskId).
		Where("is_sent", 1).
		Where("sent_time BETWEEN ? AND ?", startTime, endTime).
		Fields("(sent_time - (sent_time % 60)) AS minute, COUNT(1) AS send_count"). // group by minute
		Group("minute").
		Order("minute ASC").
		Scan(&statList)

	if err != nil {
		g.Log().Error(ctx, "Error querying statistics: %v", err)
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get statistics data")))
		return
	}

	for i := range statList {
		statList[i].StartTime = statList[i].Minute
		statList[i].EndTime = statList[i].Minute + 60
	}

	// calculate send rate metrics
	var totalSent int
	var maxPerMinute int
	var avgPerMinute float64

	if len(statList) > 0 {
		for _, item := range statList {
			totalSent += item.SendCount
			if item.SendCount > maxPerMinute {
				maxPerMinute = item.SendCount
			}
		}

		avgPerMinute = float64(totalSent) / float64(len(statList))
	}

	res.Data = &v1.TaskSendCountData{
		TaskId:        req.TaskId,
		StartTime:     startTime,
		EndTime:       endTime,
		TotalSent:     totalSent,
		MaxPerMinute:  maxPerMinute,
		AvgPerMinute:  avgPerMinute,
		MinuteDetails: statList,
	}

	res.SetSuccess(public.LangCtx(ctx, "Get task send count successfully"))
	return
}
