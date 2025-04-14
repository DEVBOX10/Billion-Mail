package rbac

import (
	"billionmail-core/internal/service/middlewares"
	service "billionmail-core/internal/service/rbac"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"billionmail-core/api/rbac/v1"
)

// RoleList gets role list
func (c *ControllerV1) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	res = &v1.RoleListRes{}

	// Check permission
	if !middlewares.HasPermission(ctx, "role", "read", "role") {
		res.SetError(gerror.New("Insufficient permissions"))
		return
	}

	// Get role list
	roles, total, err := service.Role().GetList(ctx, req.Page, req.PageSize, req.Name, req.Status)
	if err != nil {
		res.SetError(gerror.New("Failed to get role list: " + err.Error()))
		return
	}

	// Prepare response data
	res.Success = true
	res.Code = 0
	res.Msg = "Success"
	res.Data.Total = total
	res.Data.Page = req.Page

	// Convert to API response format
	res.Data.List = make([]v1.RoleInfoItem, 0, len(roles))
	for _, role := range roles {
		res.Data.List = append(res.Data.List, v1.RoleInfoItem{
			Id:          role.RoleId,
			Name:        role.RoleName,
			Description: role.Description,
			Status:      role.Status,
			CreateTime:  role.CreateTime,
		})
	}

	return
}

// RoleDetail gets role details
func (c *ControllerV1) RoleDetail(ctx context.Context, req *v1.RoleDetailReq) (res *v1.RoleDetailRes, err error) {
	res = &v1.RoleDetailRes{}

	// Check permission
	if !middlewares.HasPermission(ctx, "role", "read", "role") {
		res.SetError(gerror.New("Insufficient permissions"))
		return
	}

	// Get role details
	role, err := service.Role().GetById(ctx, req.RoleId)
	if err != nil {
		res.SetError(gerror.New("Failed to get role details: " + err.Error()))
		return
	}

	// Get role permissions
	permissions, err := service.Role().GetRolePermissions(ctx, req.RoleId)
	if err != nil {
		res.SetError(gerror.New("Failed to get role permissions: " + err.Error()))
		return
	}

	// Get all permissions
	allPermissions, err := service.Permission().GetAll(ctx)
	if err != nil {
		res.SetError(gerror.New("Failed to get all permissions: " + err.Error()))
		return
	}

	// Prepare response data
	res.Success = true
	res.Code = 0
	res.Msg = "Success"

	// Set role information
	res.Data.Role = v1.RoleInfoItem{
		Id:          role.RoleId,
		Name:        role.RoleName,
		Description: role.Description,
		Status:      role.Status,
		CreateTime:  role.CreateTime,
	}

	// Set role permissions
	res.Data.Permissions = make([]v1.PermissionInfoItem, 0, len(permissions))
	for _, perm := range permissions {
		res.Data.Permissions = append(res.Data.Permissions, v1.PermissionInfoItem{
			Id:          perm.PermissionId,
			Name:        perm.PermissionName,
			Description: perm.Description,
			Module:      perm.Module,
			Action:      perm.Action,
			Resource:    perm.Resource,
			Status:      perm.Status,
			CreateTime:  perm.CreateTime,
			UpdateTime:  perm.UpdateTime,
		})
	}

	// Set all permissions
	res.Data.AllPermissions = make([]v1.PermissionInfoItem, 0, len(allPermissions))
	for _, perm := range allPermissions {
		res.Data.AllPermissions = append(res.Data.AllPermissions, v1.PermissionInfoItem{
			Id:          perm.PermissionId,
			Name:        perm.PermissionName,
			Description: perm.Description,
			Module:      perm.Module,
			Action:      perm.Action,
			Resource:    perm.Resource,
			Status:      perm.Status,
			CreateTime:  perm.CreateTime,
			UpdateTime:  perm.UpdateTime,
		})
	}

	return
}

// RoleCreate creates a role
func (c *ControllerV1) RoleCreate(ctx context.Context, req *v1.RoleCreateReq) (res *v1.RoleCreateRes, err error) {
	res = &v1.RoleCreateRes{}

	// Check permission
	if !middlewares.HasPermission(ctx, "role", "create", "role") {
		res.SetError(gerror.New("Insufficient permissions"))
		return
	}

	// Check if role name already exists
	exists, err := service.Role().NameExists(ctx, req.Name)
	if err != nil {
		res.SetError(gerror.New("Failed to check role name: " + err.Error()))
		return
	}
	if exists {
		res.SetError(gerror.New("Role name already exists"))
		return
	}

	roleId, err := service.Role().Create(ctx, req.Name, req.Description, req.Status)
	if err != nil {
		res.SetError(gerror.New("Failed to create role: " + err.Error()))
		return
	}

	// Assign permissions
	if len(req.PermissionIds) > 0 {
		for _, permId := range req.PermissionIds {
			err = service.Role().AssignPermission(ctx, roleId, permId)
			if err != nil {
				g.Log().Warning(ctx, "Failed to assign permission: ", err)
			}
		}
	}

	// Prepare response data
	res.Success = true
	res.Code = 0
	res.Msg = "Created successfully"
	res.Data.RoleId = roleId

	return
}

// RoleUpdate updates a role
func (c *ControllerV1) RoleUpdate(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error) {
	res = &v1.RoleUpdateRes{}

	// Check permission
	if !middlewares.HasPermission(ctx, "role", "update", "role") {
		res.SetError(gerror.New("Insufficient permissions"))
		return
	}

	// Get role
	role, err := service.Role().GetById(ctx, req.RoleId)
	if err != nil {
		res.SetError(gerror.New("Failed to get role: " + err.Error()))
		return
	}

	// Check if it's admin role
	if role.RoleName == "admin" && req.Name != "admin" {
		res.SetError(gerror.New("Cannot modify the name of admin role"))
		return
	}

	// Check if role name already exists
	if req.Name != "" && req.Name != role.RoleName {
		var exists bool
		exists, err = service.Role().NameExists(ctx, req.Name)
		if err != nil {
			res.SetError(gerror.New("Failed to check role name: " + err.Error()))
			return
		}
		if exists {
			res.SetError(gerror.New("Role name already exists"))
			return
		}
	}

	// Update role information
	err = service.Role().Update(ctx, req.RoleId, req.Name, req.Description, req.Status)
	if err != nil {
		res.SetError(gerror.New("Failed to update role: " + err.Error()))
		return
	}

	// Update permissions
	if len(req.PermissionIds) > 0 {
		// Remove all permissions first
		err = service.Role().ClearPermissions(ctx, req.RoleId)
		if err != nil {
			res.SetError(gerror.New("Failed to clear permissions: " + err.Error()))
			return
		}

		// Assign new permissions
		for _, permId := range req.PermissionIds {
			err = service.Role().AssignPermission(ctx, req.RoleId, permId)
			if err != nil {
				g.Log().Warning(ctx, "Failed to assign permission: ", err)
			}
		}
	}

	// Prepare response data
	res.Success = true
	res.Code = 0
	res.Msg = "Updated successfully"

	return
}

// RoleDelete deletes a role
func (c *ControllerV1) RoleDelete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	res = &v1.RoleDeleteRes{}

	// Check permission
	if !middlewares.HasPermission(ctx, "role", "delete", "role") {
		res.SetError(gerror.New("Insufficient permissions"))
		return
	}

	// Get role
	role, err := service.Role().GetById(ctx, req.RoleId)
	if err != nil {
		res.SetError(gerror.New("Failed to get role: " + err.Error()))
		return
	}

	// Check if it's a reserved role
	if role.RoleName == "admin" || role.RoleName == "user" {
		res.SetError(gerror.New("Cannot delete system reserved roles"))
		return
	}

	// Check if the role has associated accounts
	hasAccounts, err := service.Role().HasAccounts(ctx, req.RoleId)
	if err != nil {
		res.SetError(gerror.New("Failed to check role associated accounts: " + err.Error()))
		return
	}
	if hasAccounts {
		res.SetError(gerror.New("Role has associated accounts, please remove associations first"))
		return
	}

	// Delete role
	err = service.Role().Delete(ctx, req.RoleId)
	if err != nil {
		res.SetError(gerror.New("Failed to delete role: " + err.Error()))
		return
	}

	// Prepare response data
	res.Success = true
	res.Code = 0
	res.Msg = "Deleted successfully"

	return
}
