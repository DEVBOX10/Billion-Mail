package database_initialization

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// Initialize RBAC related table structures
func init() {
	registerHandler(func() {
		rbacSQLList := []string{
			// Account table
			`CREATE TABLE IF NOT EXISTS account (
				account_id SERIAL PRIMARY KEY,
				username VARCHAR(64) NOT NULL UNIQUE,
				password VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				status INT NOT NULL DEFAULT 1, -- 1: enabled, 0: disabled
				language VARCHAR(50) NOT NULL DEFAULT 'en',
				last_login_time INT NOT NULL DEFAULT 0,
				create_time INT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
				update_time INT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
			)`,

			// Role table
			`CREATE TABLE IF NOT EXISTS role (
				role_id SERIAL PRIMARY KEY,
				role_name VARCHAR(64) NOT NULL UNIQUE,
				description TEXT,
				status INT NOT NULL DEFAULT 1, -- 1: enabled, 0: disabled
				create_time INT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
				update_time INT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
			)`,

			// Permission table
			`CREATE TABLE IF NOT EXISTS permission (
				permission_id SERIAL PRIMARY KEY,
				permission_name VARCHAR(64) NOT NULL UNIQUE,
				description TEXT,
				module VARCHAR(64) NOT NULL, -- Module name
				action VARCHAR(64) NOT NULL, -- Action type: create, read, update, delete
				resource VARCHAR(64) NOT NULL, -- Resource type
				status INT NOT NULL DEFAULT 1, -- 1: enabled, 0: disabled
				create_time INT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
				update_time INT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
				UNIQUE(module, action, resource)
			)`,

			// Account-Role mapping table
			`CREATE TABLE IF NOT EXISTS account_role (
				id SERIAL PRIMARY KEY,
				account_id INT NOT NULL,
				role_id INT NOT NULL,
				create_time INT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
				UNIQUE(account_id, role_id),
				FOREIGN KEY (account_id) REFERENCES account(account_id) ON DELETE CASCADE,
				FOREIGN KEY (role_id) REFERENCES role(role_id) ON DELETE CASCADE
			)`,

			// Role-Permission mapping table
			`CREATE TABLE IF NOT EXISTS role_permission (
				id SERIAL PRIMARY KEY,
				role_id INT NOT NULL,
				permission_id INT NOT NULL,
				create_time INT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
				UNIQUE(role_id, permission_id),
				FOREIGN KEY (role_id) REFERENCES role(role_id) ON DELETE CASCADE,
				FOREIGN KEY (permission_id) REFERENCES permission(permission_id) ON DELETE CASCADE
			)`,

			// Initialize basic roles
			`INSERT INTO role (role_name, description)
			VALUES 
				('admin', 'System administrator with full access'),
				('manager', 'Manager with limited administrative privileges'),
				('user', 'Regular user with basic access')
			ON CONFLICT (role_name) DO NOTHING`,

			// Initialize system permissions
			`INSERT INTO permission (permission_name, description, module, action, resource)
			VALUES 
				('account.create', 'Create user accounts', 'account', 'create', 'account'),
				('account.read', 'View user accounts', 'account', 'read', 'account'),
				('account.update', 'Update user accounts', 'account', 'update', 'account'),
				('account.delete', 'Delete user accounts', 'account', 'delete', 'account'),
				
				('role.create', 'Create roles', 'role', 'create', 'role'),
				('role.read', 'View roles', 'role', 'read', 'role'),
				('role.update', 'Update roles', 'role', 'update', 'role'),
				('role.delete', 'Delete roles', 'role', 'delete', 'role'),
				
				('permission.create', 'Create permissions', 'permission', 'create', 'permission'),
				('permission.read', 'View permissions', 'permission', 'read', 'permission'),
				('permission.update', 'Update permissions', 'permission', 'update', 'permission'),
				('permission.delete', 'Delete permissions', 'permission', 'delete', 'permission'),
				
				('domain.create', 'Create domains', 'domain', 'create', 'domain'),
				('domain.read', 'View domains', 'domain', 'read', 'domain'),
				('domain.update', 'Update domains', 'domain', 'update', 'domain'),
				('domain.delete', 'Delete domains', 'domain', 'delete', 'domain'),
				
				('mailbox.create', 'Create mailboxes', 'mailbox', 'create', 'mailbox'),
				('mailbox.read', 'View mailboxes', 'mailbox', 'read', 'mailbox'),
				('mailbox.update', 'Update mailboxes', 'mailbox', 'update', 'mailbox'),
				('mailbox.delete', 'Delete mailboxes', 'mailbox', 'delete', 'mailbox')
			ON CONFLICT (permission_name) DO NOTHING`,

			// Assign all permissions to admin role
			`INSERT INTO role_permission (role_id, permission_id)
			SELECT r.role_id, p.permission_id
			FROM role r, permission p
			WHERE r.role_name = 'admin'
			ON CONFLICT DO NOTHING`,

			// Assign partial permissions to manager role
			`INSERT INTO role_permission (role_id, permission_id)
			SELECT r.role_id, p.permission_id
			FROM role r, permission p
			WHERE r.role_name = 'manager' AND p.permission_name NOT LIKE '%delete'
			AND p.permission_name NOT IN ('permission.create', 'permission.update', 'role.create', 'role.update')
			ON CONFLICT DO NOTHING`,

			// Assign basic permissions to user role
			`INSERT INTO role_permission (role_id, permission_id)
			SELECT r.role_id, p.permission_id
			FROM role r, permission p
			WHERE r.role_name = 'user' AND p.permission_name IN (
				'account.read', 'domain.read', 'mailbox.read'
			)
			ON CONFLICT DO NOTHING`,

			// Create initial admin account (if not exists)
			`INSERT INTO account (username, password, email, status)
			VALUES ('admin', '$2a$10$RMENVjdmVbjauXOgbBQtlu8uyjV/cjTNIX58ORkuYTIOlKUMtUE5m', 'admin@example.com', 1)
			ON CONFLICT (username) DO NOTHING`,

			// Assign admin role to admin account
			`INSERT INTO account_role (account_id, role_id)
			SELECT a.account_id, r.role_id
			FROM account a, role r
			WHERE a.username = 'admin' AND r.role_name = 'admin'
			ON CONFLICT DO NOTHING`,
		}

		// Execute SQL statements
		for _, sql := range rbacSQLList {
			_, err := g.DB().Exec(context.Background(), sql)
			if err != nil {
				g.Log().Error(context.Background(), "Failed to execute RBAC SQL:", err, sql)
				return
			}
		}

		g.Log().Info(context.Background(), "RBAC tables initialized successfully")
	})
}
