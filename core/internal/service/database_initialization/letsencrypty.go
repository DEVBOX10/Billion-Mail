package database_initialization

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	registerHandler(func() {
		row_sql_list := []string{
			`-- Let's Encrypt Certificate Table
			CREATE TABLE IF NOT EXISTS letsencrypts (
				cert_id SERIAL PRIMARY KEY,
				account_id INTEGER NOT NULL DEFAULT 0,
				issuer TEXT NOT NULL DEFAULT '',
				not_after TEXT NOT NULL DEFAULT '',
				not_before TEXT NOT NULL DEFAULT '',
				dns TEXT NOT NULL DEFAULT '',
				auth_type TEXT NOT NULL DEFAULT '',
				dns_provider TEXT NOT NULL DEFAULT '',
				dns_provider_token TEXT NOT NULL DEFAULT '{}',
				subject TEXT NOT NULL DEFAULT '',
				status INTEGER NOT NULL DEFAULT 0, -- 0=applying 1=success 2=fail
				certificate TEXT NOT NULL DEFAULT '', -- PEM
				private_key TEXT NOT NULL DEFAULT '',
				endtime INTEGER NOT NULL DEFAULT 0,
				create_time INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
				progress TEXT NOT NULL DEFAULT '',
				error_info TEXT NOT NULL DEFAULT '' -- Error message (when status=2)
			)`,

			`-- Let's Encrypt Certificate Renewal Log Table
			CREATE TABLE IF NOT EXISTS renew_logs (
				renew_id SERIAL PRIMARY KEY,
				cert_id INTEGER NOT NULL DEFAULT 0,
				account_id INTEGER NOT NULL DEFAULT 0,
				status INTEGER NOT NULL DEFAULT 0, -- 0=renewing 1=success 2=fail
				progress TEXT NOT NULL DEFAULT '',
				error_info TEXT NOT NULL DEFAULT '',
				renew_time INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
			)`,
		}

		for _, sql := range row_sql_list {
			_, err := g.DB().Exec(context.Background(), sql)
			if err != nil {
				g.Log().Error(context.Background(), "Failed to create table:", err)
				return
			}
		}
	})
}
