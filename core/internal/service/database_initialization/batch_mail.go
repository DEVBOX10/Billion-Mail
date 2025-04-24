package database_initialization

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	registerHandler(func() {
		batchMailSQLList := []string{

			`CREATE TABLE IF NOT EXISTS contact_groups (
                id SERIAL PRIMARY KEY,
                name VARCHAR(255) NOT NULL,
                description TEXT,
                create_time INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
                update_time INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
                UNIQUE(name)
            )`,

			`CREATE TABLE IF NOT EXISTS contacts (
                id SERIAL PRIMARY KEY,
                email VARCHAR(320) NOT NULL,
                group_id INTEGER,
                active INTEGER DEFAULT 1,
                task_id INTEGER DEFAULT 0,
                create_time INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
                FOREIGN KEY (group_id) REFERENCES contact_groups(id) ON DELETE SET NULL,
                UNIQUE(group_id, email)
            )`,

			`CREATE TABLE IF NOT EXISTS email_templates (
                id SERIAL PRIMARY KEY,
                temp_name VARCHAR(255),
                add_type SMALLINT NOT NULL DEFAULT 0,
                content TEXT NOT NULL DEFAULT '',
                render TEXT NOT NULL DEFAULT '',
                create_time INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
                update_time INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    			UNIQUE(temp_name)
            )`,

			`CREATE TABLE IF NOT EXISTS email_tasks (
                id SERIAL PRIMARY KEY,
                task_name VARCHAR(255) NOT NULL,
                addresser VARCHAR(320) NOT NULL,
                subject TEXT,
                full_name VARCHAR(255),
                remark TEXT,
                recipient_count INTEGER NOT NULL,
                task_process SMALLINT NOT NULL DEFAULT 0,
                pause SMALLINT NOT NULL DEFAULT 0,
                template_id INTEGER NOT NULL,
                is_record SMALLINT NOT NULL DEFAULT 0,
                unsubscribe SMALLINT NOT NULL DEFAULT 0,
                threads INTEGER NOT NULL DEFAULT 0,
                etypes VARCHAR(320) NOT NULL DEFAULT '1',
                track_open SMALLINT NOT NULL DEFAULT 1,
                track_click SMALLINT NOT NULL DEFAULT 1,
                start_time INTEGER NOT NULL DEFAULT 0,
                create_time INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
                update_time INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
                active SMALLINT NOT NULL DEFAULT 0,
    			add_type SMALLINT NOT NULL DEFAULT 0,	
                FOREIGN KEY (template_id) REFERENCES email_templates(id)
            )`,

			`CREATE TABLE IF NOT EXISTS recipient_info (
                id SERIAL PRIMARY KEY,
                task_id INTEGER NOT NULL,
                recipient VARCHAR(320) NOT NULL,
                is_sent SMALLINT DEFAULT 0,
                sent_time INTEGER NOT NULL DEFAULT 0,
                message_id TEXT NOT NULL,
                create_time INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
                FOREIGN KEY (task_id) REFERENCES email_tasks(id) ON DELETE CASCADE,
                UNIQUE(task_id, recipient)
            )`,

			`CREATE TABLE IF NOT EXISTS unsubscribe_records (
                id SERIAL PRIMARY KEY,
                email VARCHAR(320) NOT NULL,
                group_id INTEGER,
                template_id INTEGER,
                task_id INTEGER,
                unsubscribe_time INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
                FOREIGN KEY (group_id) REFERENCES contact_groups(id) ON DELETE SET NULL,
                FOREIGN KEY (template_id) REFERENCES email_templates(id) ON DELETE SET NULL,
                FOREIGN KEY (task_id) REFERENCES email_tasks(id) ON DELETE SET NULL
            )`,

			`CREATE INDEX IF NOT EXISTS idx_unsubscribe_email ON unsubscribe_records (email)`,
			`CREATE INDEX IF NOT EXISTS idx_unsubscribe_time ON unsubscribe_records (unsubscribe_time)`,
			`CREATE INDEX IF NOT EXISTS idx_contacts_email ON contacts (email)`,
			`CREATE INDEX IF NOT EXISTS idx_contacts_active ON contacts (active)`,
		}

		for _, sql := range batchMailSQLList {
			_, err := g.DB().Exec(context.Background(), sql)
			if err != nil {
				g.Log().Error(context.Background(), "Failed to execute batch mail SQL:", err, sql)
				return
			}
		}

		g.Log().Info(context.Background(), "Batch mail tables initialized successfully")
	})
}
