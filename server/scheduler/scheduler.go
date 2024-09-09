package scheduler

import (
	"context"
	"time"
)

func RunPeriodically(ctx context.Context, interval time.Duration, task func()) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			task()
		case <-ctx.Done():
			return
		}
	}
}