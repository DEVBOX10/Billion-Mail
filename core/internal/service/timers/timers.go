package timers

import (
	"billionmail-core/internal/service/domains"
	"billionmail-core/internal/service/maillog_stat"
	"context"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func Start(ctx context.Context) (err error) {
	g.Log().Debug(ctx, "Staring timers...")

	// Start DNS Records checking
	// Ensure the DNS records are fresh on startup
	gtimer.AddOnce(5*time.Second, func() {
		domains.FreshRecords(ctx)
	})
	// Check DNS Records every 5 minutes
	gtimer.Add(5*time.Minute, func() {
		domains.FreshRecords(ctx)
	})

	// Start maillog analysis
	gtimer.AddOnce(5*time.Second, func() {
		me := maillog_stat.NewMallogEventHandler("", 1*time.Second)
		me.Start()
	})

	// Start maillog aggregation
	gtimer.AddOnce(5*time.Second, func() {
		maillog_stat.AggregateMaillogsTask(1 * time.Minute)
	})

	// Start console panel baseurl update timer
	gtimer.AddOnce(5*time.Second, domains.UpdateBaseURL)
	gtimer.Add(5*time.Minute, domains.UpdateBaseURL)

	g.Log().Debug(ctx, "Start timers complete")

	return
}
