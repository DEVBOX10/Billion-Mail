package batch_mail

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sync"
	"time"
)

// SimpleRateController 简化的速率控制器
type SimpleRateController struct {
	maxPerMinute     int           // 每分钟最大发送数
	sentInLastMinute int           // 最近一分钟已发送数量
	lastResetTime    time.Time     // 上次重置计数时间
	waitTime         time.Duration // 每次发送等待时间
	mu               sync.Mutex    // 互斥锁
}

// NewSimpleRateController 创建简单速率控制器
func NewSimpleRateController(maxPerMinute int) *SimpleRateController {
	// 初始等待时间计算 (默认按照均匀分布)
	waitTime := time.Duration(60*1000/maxPerMinute) * time.Millisecond

	return &SimpleRateController{
		maxPerMinute:  maxPerMinute,
		lastResetTime: time.Now(),
		waitTime:      waitTime,
	}
}

// Wait 等待一个发送间隔
func (r *SimpleRateController) Wait(ctx context.Context) error {
	r.mu.Lock()
	waitTime := r.waitTime
	r.mu.Unlock()

	// 使用定时器进行等待
	timer := time.NewTimer(waitTime)
	defer timer.Stop()

	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RecordSend 记录一次发送
func (r *SimpleRateController) RecordSend() {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 检查是否需要重置计数
	now := time.Now()
	if now.Sub(r.lastResetTime) >= time.Minute {
		r.sentInLastMinute = 0
		r.lastResetTime = now
	}

	// 增加发送计数
	r.sentInLastMinute++
}

// AdjustRate 调整发送速率
func (r *SimpleRateController) AdjustRate() {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	elapsedSeconds := now.Sub(r.lastResetTime).Seconds()
	if elapsedSeconds < 10 {
		// 至少收集10秒数据再调整
		return
	}

	// 预计整分钟发送量
	projectedRate := float64(r.sentInLastMinute) / elapsedSeconds * 60
	targetRate := float64(r.maxPerMinute)

	// 根据实际发送率调整等待时间
	if projectedRate > targetRate*1.1 {
		// 发送过快，增加等待时间 (降低20%)
		r.waitTime = time.Duration(float64(r.waitTime) * 1.2)
		g.Log().Debug(context.Background(), "发送速率过高 (%.2f/分钟)，增加等待时间至 %v", projectedRate, r.waitTime)
	} else if projectedRate < targetRate*0.7 {
		// 发送过慢，减少等待时间 (提高20%)
		r.waitTime = time.Duration(float64(r.waitTime) * 0.8)
		g.Log().Debug(context.Background(), "发送速率过低 (%.2f/分钟)，减少等待时间至 %v", projectedRate, r.waitTime)
	}

	// 设置最小等待时间，防止过快发送
	minWait := time.Duration(5) * time.Millisecond
	if r.waitTime < minWait {
		r.waitTime = minWait
	}
}

// GetCurrentRate 获取当前每分钟发送速率
func (r *SimpleRateController) GetCurrentRate() float64 {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	elapsedSeconds := now.Sub(r.lastResetTime).Seconds()
	if elapsedSeconds <= 0 {
		return 0
	}

	return float64(r.sentInLastMinute) / elapsedSeconds * 60
}

// GetMaxRate 获取设置的最大速率
func (r *SimpleRateController) GetMaxRate() int {
	return r.maxPerMinute
}
