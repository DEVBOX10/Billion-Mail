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

// TaskExecutor 任务执行器
type TaskExecutor struct {
	ctx context.Context

	// 添加速率控制相关字段
	rateLimiter *rate.Limiter // 速率限制器
	speedMeter  *SpeedMeter   // 速度计量器
	taskConfig  *TaskConfig   // 任务配置
	pauseChan   chan struct{}
	isPaused    bool
	pauseMutex  sync.Mutex
	pool        *ants.Pool // 工作池
}

// TaskConfig 任务配置
type TaskConfig struct {
	Threads      int     // 线程数
	MaxPerMinute int     // 每分钟最大发送数
	CurrentSpeed float64 // 当前发送速度
	MinSpeed     float64 // 最小发送速度
	MaxSpeed     float64 // 最大发送速度
	SpeedAdjust  float64 // 速度调整因子
}

// SpeedMeter 速度计量器
type SpeedMeter struct {
	sentCount    int64
	startTime    time.Time
	lastCheck    time.Time
	currentSpeed float64
	mu           sync.Mutex
}

// UpdateTaskSpeedReq 更新任务速度请求
type UpdateTaskSpeedReq struct {
	g.Meta        `path:"/batch_mail/task/speed" method:"post" tags:"BatchMail" summary:"更新任务发送速度"`
	Authorization string  `json:"authorization" dc:"Authorization" in:"header"`
	TaskId        int     `json:"task_id" v:"required" dc:"任务ID"`
	Speed         float64 `json:"speed" v:"required|min:0.1|max:10" dc:"发送速度因子(0.1-10)"`
}

// UpdateTaskSpeedRes 更新任务速度响应
type UpdateTaskSpeedRes struct {
	api_v1.StandardRes
}

// PauseTaskReq 暂停任务请求
type PauseTaskReq struct {
	g.Meta        `path:"/batch_mail/task/pause" method:"post" tags:"BatchMail" summary:"暂停任务"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	TaskId        int    `json:"task_id" v:"required" dc:"任务ID"`
}

// PauseTaskRes 暂停任务响应
type PauseTaskRes struct {
	api_v1.StandardRes
}

// ResumeTaskReq 恢复任务请求
type ResumeTaskReq struct {
	g.Meta        `path:"/batch_mail/task/resume" method:"post" tags:"BatchMail" summary:"恢复任务"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	TaskId        int    `json:"task_id" v:"required" dc:"任务ID"`
}

// ResumeTaskRes 恢复任务响应
type ResumeTaskRes struct {
	api_v1.StandardRes
}
