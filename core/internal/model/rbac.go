package model

import "billionmail-core/internal/model/entity"

// Account defines the account model
type Account entity.Account

// Role defines the role model
type Role entity.Role

// Permission defines the permission model
type Permission entity.Permission

// AccountRole defines the account-role relationship model
type AccountRole entity.AccountRole

// RolePermission defines the role-permission relationship model
type RolePermission entity.RolePermission
