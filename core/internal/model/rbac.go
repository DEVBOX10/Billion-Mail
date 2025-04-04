package model

import (
	"time"
)

// Account defines the account model
type Account struct {
	Id        int64     `json:"account_id"` // ID
	Username  string    `json:"username"`   // Username
	Email     string    `json:"email"`      // Email
	Password  string    `json:"-"`          // Password (hidden in JSON output)
	Status    int       `json:"status"`     // Status 0:disabled 1:enabled
	Lang      string    `json:"lang"`       // Language setting
	CreatedAt time.Time `json:"created_at"` // Creation time
	UpdatedAt time.Time `json:"updated_at"` // Update time
}

// Role defines the role model
type Role struct {
	Id          int64     `json:"role_id"`     // ID
	Name        string    `json:"name"`        // Role name
	Description string    `json:"description"` // Description
	Status      int       `json:"status"`      // Status 0:disabled 1:enabled
	CreatedAt   time.Time `json:"created_at"`  // Creation time
	UpdatedAt   time.Time `json:"updated_at"`  // Update time
}

// Permission defines the permission model
type Permission struct {
	Id          int64     `json:"id"`          // ID
	Name        string    `json:"name"`        // Permission name
	Description string    `json:"description"` // Description
	Module      string    `json:"module"`      // Module
	Action      string    `json:"action"`      // Action
	Resource    string    `json:"resource"`    // Resource
	Status      int       `json:"status"`      // Status 0:disabled 1:enabled
	CreatedAt   time.Time `json:"created_at"`  // Creation time
	UpdatedAt   time.Time `json:"updated_at"`  // Update time
}

// AccountRole defines the account-role relationship model
type AccountRole struct {
	Id        int64     `json:"id"`         // ID
	AccountId int64     `json:"account_id"` // Account ID
	RoleId    int64     `json:"role_id"`    // Role ID
	CreatedAt time.Time `json:"created_at"` // Creation time
	UpdatedAt time.Time `json:"updated_at"` // Update time
}

// RolePermission defines the role-permission relationship model
type RolePermission struct {
	Id           int64     `json:"id"`            // ID
	RoleId       int64     `json:"role_id"`       // Role ID
	PermissionId int64     `json:"permission_id"` // Permission ID
	CreatedAt    time.Time `json:"created_at"`    // Creation time
	UpdatedAt    time.Time `json:"updated_at"`    // Update time
}
