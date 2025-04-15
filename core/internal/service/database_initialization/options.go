package database_initialization

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	registerHandler(func() {
		sqlList := []string{
			`-- BillionMail Options
			CREATE TABLE IF NOT EXISTS bm_options (
				name VARCHAR(255) NOT NULL PRIMARY KEY,
				value TEXT NOT NULL
			)`,
		}

		for _, sql := range sqlList {
			_, err := g.DB().Exec(context.Background(), sql)
			if err != nil {
				g.Log().Error(context.Background(), "Failed to create billion-mail options table:", err)
				return
			}
		}
	})
}
