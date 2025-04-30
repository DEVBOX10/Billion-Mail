package redis_initialization

import (
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

// InitRedis initialize redis configuration
func InitRedis() (err error) {
	// get redis password from environment variable
	passwd, err := public.DockerEnv("REDISPASS")

	if err != nil {
		return fmt.Errorf("redis env init error: %v", err)
	}

	address := "127.0.0.1:26379"

	if public.IsRunningInContainer() {
		address = "redis:6379"
	}

	// Initialize Redis configuration
	gredis.SetConfig(&gredis.Config{
		Address: address,
		Db:      1,
		Pass:    passwd,
	})

	// Testing redis connection until it successful
	connectionOK := false
	for i := 0; i < 10; i++ {
		k := "bm_test_connection"
		if err := g.Redis().SetEX(context.Background(), k, k, 1); err != nil {
			g.Log().Error(context.Background(), "Redis connection test failed: ", err, " Waiting for 5 seconds before retrying...")
			time.Sleep(5 * time.Second)
			continue
		}

		if _, err := g.Redis().Del(context.Background(), k); err != nil {
			g.Log().Error(context.Background(), "Redis connection test failed: ", err, " Waiting for 5 seconds before retrying...")
			time.Sleep(5 * time.Second)
			continue
		}

		connectionOK = true
		g.Log().Debug(context.Background(), "Redis connection test successful")
		break
	}

	if !connectionOK {
		return fmt.Errorf("redis connection test failed after 10 attempts")
	}

	return nil
}
