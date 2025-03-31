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

	// Initialize Redis configuration
	gredis.SetConfig(&gredis.Config{
		Address: "127.0.0.1:26379",
		Db:      1,
		Pass:    passwd,
	})

	return nil
}
