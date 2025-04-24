package v1

import (
	"billionmail-core/utility/types/api_v1"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/panjf2000/ants/v2"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type TaskExecutor struct {
	ctx context.Context

	rateLimiter *rate.Limiter
	speedMeter  *SpeedMeter
	taskConfig  *TaskConfig
	pauseChan   chan struct{}
	isPaused    bool
	pauseMutex  sync.Mutex
	pool        *ants.Pool
}

// TaskConfig task config
type TaskConfig struct {
	Threads      int     // pool count
	MaxPerMinute int     // max emails per minute
	CurrentSpeed float64 // current speed
	MinSpeed     float64 // min speed
	MaxSpeed     float64 // max speed
	SpeedAdjust  float64 // speed adjust factor
}

// SpeedMeter speed meter
type SpeedMeter struct {
	sentCount    int64
	startTime    time.Time
	lastCheck    time.Time
	currentSpeed float64
	mu           sync.Mutex
}

// UpdateTaskSpeedReq update task speed request
type UpdateTaskSpeedReq struct {
	g.Meta        `path:"/batch_mail/task/speed" method:"post" tags:"BatchMail" summary:"update task speed"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	TaskId        int    `json:"task_id" v:"required" dc:"task id"`
	Threads       int    `json:"threads" v:"required|min:1|max:100" dc:"thread count (1-100), decide the number of emails sent concurrently"`
}

// UpdateTaskSpeedRes update task speed response
type UpdateTaskSpeedRes struct {
	api_v1.StandardRes
}

// PauseTaskReq pause task request
type PauseTaskReq struct {
	g.Meta        `path:"/batch_mail/task/pause" method:"post" tags:"BatchMail" summary:"pause task"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	TaskId        int    `json:"task_id" v:"required" dc:"task id"`
}

// PauseTaskRes pause task response
type PauseTaskRes struct {
	api_v1.StandardRes
}

// ResumeTaskReq resume task request
type ResumeTaskReq struct {
	g.Meta        `path:"/batch_mail/task/resume" method:"post" tags:"BatchMail" summary:"resume task"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	TaskId        int    `json:"task_id" v:"required" dc:"task id"`
}

// ResumeTaskRes resume task response
type ResumeTaskRes struct {
	api_v1.StandardRes
}
