package database_initialization

import (
	"billionmail-core/internal/consts"
	"billionmail-core/internal/service/public"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
)

var registeredHandlers = make([]func(), 0, 256)

// InitDatabase initializes the database configuration
func InitDatabase() (err error) {
	dbPass, err := public.DockerEnv("DBPASS")

	if err != nil {
		return fmt.Errorf("Read database password failed: %v", err)
	}

	// Setting database configuration
	err = gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				// Debug: true,
				Host:             public.AbsPath(consts.POSTGRESQL_SOCK),
				User:             "billionmail",
				Pass:             dbPass,
				Name:             "billionmail",
				Type:             "pgsql",
				Role:             "master",
				MaxOpenConnCount: 100,
			},
		},
	})

	if err != nil {
		return fmt.Errorf("Set database configuration failed: %v", err)
	}

	// Execute registered handlers
	for _, handler := range registeredHandlers {
		if handler != nil {
			handler()
		}
	}

	// Empty the registered handlers
	registeredHandlers = registeredHandlers[:0]

	return nil
}

// registerHandler registers a handler for the database initialization
func registerHandler(handler func()) {
	if handler == nil {
		return
	}

	registeredHandlers = append(registeredHandlers, handler)
}
