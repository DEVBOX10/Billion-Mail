package timers

import (
	"billionmail-core/internal/service/domains"
	"context"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/v2/os/glog"
	"time"
)

func Start(ctx context.Context) (err error) {
	glog.Debug(ctx, "Staring timers...")

	// Start DNS Records checking
	gtimer.Add(5*time.Minute, func() {
		domains.FreshRecords(ctx)
	})

	glog.Debug(ctx, "Start timers complete")

	return
}
