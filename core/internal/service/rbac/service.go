package rbac

import (
	"billionmail-core/internal/model"
	"context"
)

// Account service interface
type IAccount interface {
	// Get account list
	GetList(ctx context.Context, page, pageSize int, username, email string, status int) ([]model.Account, int, error)
	// Get account by ID
	GetById(ctx context.Context, accountId int64) (*model.Account, error)
	// Create account
	Create(ctx context.Context, accountData *model.Account) (int64, error)
	// Update account
	Update(ctx context.Context, accountData *model.Account) error
	// Delete account
	Delete(ctx context.Context, accountId int64) error
	// Update password
	UpdatePassword(ctx context.Context, accountId int64, newPassword string) error
	// Verify password
	VerifyPassword(hashedPassword, plainPassword string) bool
	// GeneratePasswordHash generates a hashed password
	GeneratePasswordHash(password string) (string, error)
	// Check if username exists
	UsernameExists(ctx context.Context, username string) (bool, error)
	// Check if email exists
	EmailExists(ctx context.Context, email string) (bool, error)
	// Get account roles
	GetAccountRoles(ctx context.Context, accountId int64) ([]model.Role, error)
	// Assign role to account
	AssignRole(ctx context.Context, accountId int64, roleId int64) error
	// Clear all roles from account
	ClearRoles(ctx context.Context, accountId int64) error
	// Get all roles
	GetAll(ctx context.Context) ([]model.Role, error)
	// Bind roles to account
	BindRoles(ctx context.Context, accountId int64, roleIds []int64) error
	// Get roles assigned to account
	GetRoles(ctx context.Context, accountId int64) ([]model.Role, error)
	// Get permissions assigned to account
	GetAccountPermissions(ctx context.Context, accountId int64) ([]model.Permission, error)
	// Handle user login
	Login(ctx context.Context, username, password string) (*model.Account, error)
	// Check if account is admin
	IsAdmin(ctx context.Context, accountId int64) (bool, error)
	// Count admin accounts
	CountAdmins(ctx context.Context) (int, error)
}

// Role service interface
type IRole interface {
	// Get role list with pagination
	GetList(ctx context.Context, page, pageSize int, name string, status int) ([]model.Role, int, error)
	// Get all roles
	GetAll(ctx context.Context) ([]model.Role, error)
	// Get role details by ID
	GetById(ctx context.Context, roleId int64) (*model.Role, error)
	// Create new role
	Create(ctx context.Context, name, description string, status int) (int64, error)
	// Update role information
	Update(ctx context.Context, roleId int64, name, description string, status int) error
	// Delete role
	Delete(ctx context.Context, roleId int64) error
	// Bind permissions to role
	BindPermissions(ctx context.Context, roleId int64, permissionIds []int64) error
	// Get permissions assigned to role
	GetPermissions(ctx context.Context, roleId int64) ([]model.Permission, error)
	// Check if role name exists
	NameExists(ctx context.Context, name string) (bool, error)
	// Assign a single permission to a role
	AssignPermission(ctx context.Context, roleId int64, permissionId int64) error
	// Clear all permissions assigned to a role
	ClearPermissions(ctx context.Context, roleId int64) error
	// Check if role has associated accounts
	HasAccounts(ctx context.Context, roleId int64) (bool, error)
	// Get role permissions (alias for GetPermissions)
	GetRolePermissions(ctx context.Context, roleId int64) ([]model.Permission, error)
}

// Permission service interface
type IPermission interface {
	// Get permission list with pagination
	GetList(ctx context.Context, page, pageSize int, module, action string, status int) ([]model.Permission, int, error)
	// Get all permissions
	GetAll(ctx context.Context) ([]model.Permission, error)
	// Get permission details by ID
	GetById(ctx context.Context, permissionId int64) (*model.Permission, error)
	// Create new permission
	Create(ctx context.Context, name, description, module, action, resource string, status int) (int64, error)
	// Update permission information
	Update(ctx context.Context, permissionId int64, name, description, module, action, resource string, status int) error
	// Delete permission
	Delete(ctx context.Context, permissionId int64) error
	// Check permission
	Check(ctx context.Context, accountId int64, module, action, resource string) (bool, error)
}

// Service instance getter functions
var (
	iAccountService    IAccount
	iRoleService       IRole
	iPermissionService IPermission
	iJWTService        *JWTService
)

// Get JWT service instance
func JWT() *JWTService {
	if iJWTService == nil {
		iJWTService = newJWTService()
	}
	return iJWTService
}

// Get account service instance
func Account() IAccount {
	if iAccountService == nil {
		iAccountService = newAccountService()
	}
	return iAccountService
}

// Get role service instance
func Role() IRole {
	if iRoleService == nil {
		iRoleService = newRoleService()
	}
	return iRoleService
}

// Get permission service instance
func Permission() IPermission {
	if iPermissionService == nil {
		iPermissionService = newPermissionService()
	}
	return iPermissionService
}
