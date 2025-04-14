package redis_initialization

import (
	"billionmail-core/internal/service/public"
	"fmt"
	"github.com/gogf/gf/v2/database/gredis"
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

	return nil
}
