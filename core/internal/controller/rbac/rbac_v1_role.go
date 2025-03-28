package rbac

import (
	"billionmail-core/internal/service/middlewares"
	service "billionmail-core/internal/service/rbac"
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"billionmail-core/api/rbac/v1"
)

// ListRole gets role list
func (c *ControllerV1) ListRole(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	res = &v1.RoleListRes{}

	// Check permission
	if !middlewares.HasPermission(ctx, "role", "read", "role") {
		res.SetError(gerror.New("Insufficient permissions"))
		return res, nil
	}

	// Get role list
	roles, total, err := c.RoleService.GetList(ctx, req.Page, req.PageSize, req.Name, req.Status)
	if err != nil {
		res.SetError(gerror.New("Failed to get role list: " + err.Error()))
		return res, nil
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
			Id:          role.Id,
			Name:        role.Name,
			Description: role.Description,
			Status:      role.Status,
			CreatedAt:   role.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   role.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

// GetRole gets role details
func (c *ControllerV1) GetRole(ctx context.Context, req *v1.RoleDetailReq) (res *v1.RoleDetailRes, err error) {
	res = &v1.RoleDetailRes{}

	// Check permission
	if !middlewares.HasPermission(ctx, "role", "read", "role") {
		res.SetError(gerror.New("Insufficient permissions"))
		return res, nil
	}

	// 获取角色详情
	role, err := c.RoleService.GetById(ctx, req.RoleId)
	if err != nil {
		res.SetError(gerror.New("获取角色详情失败: " + err.Error()))
		return res, nil
	}

	// 获取角色权限
	permissions, err := c.RoleService.GetRolePermissions(ctx, req.RoleId)
	if err != nil {
		res.SetError(gerror.New("获取角色权限失败: " + err.Error()))
		return res, nil
	}

	// 获取所有权限
	allPermissions, err := c.PermissionService.GetAll(ctx)
	if err != nil {
		res.SetError(gerror.New("获取所有权限失败: " + err.Error()))
		return res, nil
	}

	// 准备响应数据
	res.Success = true
	res.Code = 0
	res.Msg = "获取成功"

	// 设置角色信息
	res.Data.Role = v1.RoleInfoItem{
		Id:          role.Id,
		Name:        role.Name,
		Description: role.Description,
		Status:      role.Status,
		CreatedAt:   role.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   role.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	// 设置角色权限
	res.Data.Permissions = make([]v1.PermissionInfoItem, 0, len(permissions))
	for _, perm := range permissions {
		res.Data.Permissions = append(res.Data.Permissions, v1.PermissionInfoItem{
			Id:          perm.Id,
			Name:        perm.Name,
			Description: perm.Description,
			Module:      perm.Module,
			Action:      perm.Action,
			Resource:    perm.Resource,
			Status:      perm.Status,
			CreatedAt:   perm.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   perm.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// 设置所有权限
	res.Data.AllPermissions = make([]v1.PermissionInfoItem, 0, len(allPermissions))
	for _, perm := range allPermissions {
		res.Data.AllPermissions = append(res.Data.AllPermissions, v1.PermissionInfoItem{
			Id:          perm.Id,
			Name:        perm.Name,
			Description: perm.Description,
			Module:      perm.Module,
			Action:      perm.Action,
			Resource:    perm.Resource,
			Status:      perm.Status,
			CreatedAt:   perm.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   perm.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

// CreateRole 创建角色
func (c *Controller) CreateRole(ctx context.Context, req *v1.RoleCreateReq) (res *v1.RoleCreateRes, err error) {
	res = &v1.RoleCreateRes{}

	// 检查权限
	if !middlewares.HasPermission(ctx, "role", "create", "role") {
		res.SetError(gerror.New("权限不足"))
		return res, nil
	}

	// 检查角色名称是否已存在
	exists, err := c.RoleService.NameExists(ctx, req.Name)
	if err != nil {
		res.SetError(gerror.New("检查角色名称失败: " + err.Error()))
		return res, nil
	}
	if exists {
		res.SetError(gerror.New("角色名称已存在"))
		return res, nil
	}

	// 创建角色
	roleData := &service.Role{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}

	roleId, err := c.RoleService.Create(ctx, roleData)
	if err != nil {
		res.SetError(gerror.New("创建角色失败: " + err.Error()))
		return res, nil
	}

	// 分配权限
	if len(req.PermissionIds) > 0 {
		for _, permId := range req.PermissionIds {
			err = c.RoleService.AssignPermission(ctx, roleId, permId)
			if err != nil {
				g.Log().Warning(ctx, "分配权限失败: ", err)
			}
		}
	}

	// 准备响应数据
	res.Success = true
	res.Code = 0
	res.Msg = "创建成功"
	res.Data.RoleId = roleId

	return res, nil
}

// UpdateRole 更新角色
func (c *Controller) UpdateRole(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error) {
	res = &v1.RoleUpdateRes{}

	// 检查权限
	if !middlewares.HasPermission(ctx, "role", "update", "role") {
		res.SetError(gerror.New("权限不足"))
		return res, nil
	}

	// 获取角色
	role, err := c.RoleService.GetById(ctx, req.RoleId)
	if err != nil {
		res.SetError(gerror.New("获取角色失败: " + err.Error()))
		return res, nil
	}

	// 检查是否是admin角色
	if role.Name == "admin" && req.Name != "admin" {
		res.SetError(gerror.New("不能修改admin角色的名称"))
		return res, nil
	}

	// 检查角色名称是否已存在
	if req.Name != "" && req.Name != role.Name {
		exists, err := c.RoleService.NameExists(ctx, req.Name)
		if err != nil {
			res.SetError(gerror.New("检查角色名称失败: " + err.Error()))
			return res, nil
		}
		if exists {
			res.SetError(gerror.New("角色名称已存在"))
			return res, nil
		}
	}

	// 更新角色信息
	updateData := &service.Role{
		Id:          req.RoleId,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		UpdatedAt:   time.Now(),
	}

	err = c.RoleService.Update(ctx, updateData)
	if err != nil {
		res.SetError(gerror.New("更新角色失败: " + err.Error()))
		return res, nil
	}

	// 更新权限
	if len(req.PermissionIds) > 0 {
		// 先移除所有权限
		err = c.RoleService.ClearPermissions(ctx, req.RoleId)
		if err != nil {
			res.SetError(gerror.New("清除权限失败: " + err.Error()))
			return res, nil
		}

		// 分配新权限
		for _, permId := range req.PermissionIds {
			err = c.RoleService.AssignPermission(ctx, req.RoleId, permId)
			if err != nil {
				g.Log().Warning(ctx, "分配权限失败: ", err)
			}
		}
	}

	// 准备响应数据
	res.Success = true
	res.Code = 0
	res.Msg = "更新成功"

	return res, nil
}

// DeleteRole 删除角色
func (c *Controller) DeleteRole(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	res = &v1.RoleDeleteRes{}

	// 检查权限
	if !middlewares.HasPermission(ctx, "role", "delete", "role") {
		res.SetError(gerror.New("权限不足"))
		return res, nil
	}

	// 获取角色
	role, err := c.RoleService.GetById(ctx, req.RoleId)
	if err != nil {
		res.SetError(gerror.New("获取角色失败: " + err.Error()))
		return res, nil
	}

	// 检查是否是保留角色
	if role.Name == "admin" || role.Name == "user" {
		res.SetError(gerror.New("不能删除系统保留角色"))
		return res, nil
	}

	// 检查角色是否有关联账户
	hasAccounts, err := c.RoleService.HasAccounts(ctx, req.RoleId)
	if err != nil {
		res.SetError(gerror.New("检查角色关联账户失败: " + err.Error()))
		return res, nil
	}
	if hasAccounts {
		res.SetError(gerror.New("角色已关联账户，请先解除关联"))
		return res, nil
	}

	// 删除角色
	err = c.RoleService.Delete(ctx, req.RoleId)
	if err != nil {
		res.SetError(gerror.New("删除角色失败: " + err.Error()))
		return res, nil
	}

	// 准备响应数据
	res.Success = true
	res.Code = 0
	res.Msg = "删除成功"

	return res, nil
}
