package timers

import (
	"billionmail-core/internal/service/domains"
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

	g.Log().Debug(ctx, "Start timers complete")

	return
}
