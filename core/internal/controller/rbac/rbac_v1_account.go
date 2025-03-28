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

// ListAccount gets account list
func (c *ControllerV1) ListAccount(ctx context.Context, req *v1.AccountListReq) (res *v1.AccountListRes, err error) {
	res = &v1.AccountListRes{}

	// Check permission
	if !middlewares.HasPermission(ctx, "account", "read", "account") {
		res.SetError(gerror.New("Insufficient permissions"))
		return res, nil
	}

	// Get account list
	accounts, total, err := c.AccountService.GetList(ctx, req.Page, req.PageSize, req.Username, req.Email, req.Status)
	if err != nil {
		res.SetError(gerror.New("Failed to get account list: " + err.Error()))
		return res, nil
	}

	// Prepare response data
	res.Success = true
	res.Code = 0
	res.Msg = "Success"
	res.Data.Total = total
	res.Data.Page = req.Page

	// Convert to API response format
	res.Data.List = make([]v1.AccountInfoItem, 0, len(accounts))
	for _, account := range accounts {
		res.Data.List = append(res.Data.List, v1.AccountInfoItem{
			Id:        account.Id,
			Username:  account.Username,
			Email:     account.Email,
			Status:    account.Status,
			Lang:      account.Lang,
			CreatedAt: account.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: account.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

// GetAccount gets account details
func (c *ControllerV1) GetAccount(ctx context.Context, req *v1.AccountDetailReq) (res *v1.AccountDetailRes, err error) {
	res = &v1.AccountDetailRes{}

	// Check permission
	if !middlewares.HasPermission(ctx, "account", "read", "account") {
		res.SetError(gerror.New("Insufficient permissions"))
		return res, nil
	}

	// Get account details
	account, err := c.AccountService.GetById(ctx, req.AccountId)
	if err != nil {
		res.SetError(gerror.New("获取账户详情失败: " + err.Error()))
		return res, nil
	}

	// 获取账户角色
	roles, err := c.AccountService.GetAccountRoles(ctx, req.AccountId)
	if err != nil {
		res.SetError(gerror.New("获取账户角色失败: " + err.Error()))
		return res, nil
	}

	// 获取所有角色
	allRoles, err := c.RoleService.GetAll(ctx)
	if err != nil {
		res.SetError(gerror.New("获取所有角色失败: " + err.Error()))
		return res, nil
	}

	// 准备响应数据
	res.Success = true
	res.Code = 0
	res.Msg = "获取成功"

	// 设置账户信息
	res.Data.Account = v1.AccountInfoItem{
		Id:        account.Id,
		Username:  account.Username,
		Email:     account.Email,
		Status:    account.Status,
		Lang:      account.Lang,
		CreatedAt: account.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: account.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	// 设置账户角色
	res.Data.Roles = make([]v1.RoleInfoItem, 0, len(roles))
	for _, role := range roles {
		res.Data.Roles = append(res.Data.Roles, v1.RoleInfoItem{
			Id:          role.Id,
			Name:        role.Name,
			Description: role.Description,
			Status:      role.Status,
			CreatedAt:   role.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   role.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// 设置所有角色
	res.Data.AllRoles = make([]v1.RoleInfoItem, 0, len(allRoles))
	for _, role := range allRoles {
		res.Data.AllRoles = append(res.Data.AllRoles, v1.RoleInfoItem{
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

// CreateAccount 创建账户
func (c *ControllerV1) CreateAccount(ctx context.Context, req *v1.AccountCreateReq) (res *v1.AccountCreateRes, err error) {
	res = &v1.AccountCreateRes{}

	// 检查权限
	if !middlewares.HasPermission(ctx, "account", "create", "account") {
		res.SetError(gerror.New("权限不足"))
		return res, nil
	}

	// 检查用户名是否已存在
	exists, err := c.AccountService.UsernameExists(ctx, req.Username)
	if err != nil {
		res.SetError(gerror.New("检查用户名失败: " + err.Error()))
		return res, nil
	}
	if exists {
		res.SetError(gerror.New("用户名已存在"))
		return res, nil
	}

	// 检查邮箱是否已存在
	exists, err = c.AccountService.EmailExists(ctx, req.Email)
	if err != nil {
		res.SetError(gerror.New("检查邮箱失败: " + err.Error()))
		return res, nil
	}
	if exists {
		res.SetError(gerror.New("邮箱已存在"))
		return res, nil
	}

	// 创建账户
	accountData := &service.Account{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Status:   req.Status,
		Lang:     req.Lang,
	}

	accountId, err := c.AccountService.Create(ctx, accountData)
	if err != nil {
		res.SetError(gerror.New("创建账户失败: " + err.Error()))
		return res, nil
	}

	// 分配角色
	if len(req.RoleIds) > 0 {
		for _, roleId := range req.RoleIds {
			err = c.AccountService.AssignRole(ctx, accountId, roleId)
			if err != nil {
				g.Log().Warning(ctx, "分配角色失败: ", err)
			}
		}
	}

	// 准备响应数据
	res.Success = true
	res.Code = 0
	res.Msg = "创建成功"
	res.Data.AccountId = accountId

	return res, nil
}

// UpdateAccount 更新账户
func (c *ControllerV1) UpdateAccount(ctx context.Context, req *v1.AccountUpdateReq) (res *v1.AccountUpdateRes, err error) {
	res = &v1.AccountUpdateRes{}

	// 检查权限
	if !middlewares.HasPermission(ctx, "account", "update", "account") {
		res.SetError(gerror.New("权限不足"))
		return res, nil
	}

	// 获取账户
	account, err := c.AccountService.GetById(ctx, req.AccountId)
	if err != nil {
		res.SetError(gerror.New("获取账户失败: " + err.Error()))
		return res, nil
	}

	// 检查用户名是否已存在
	if req.Username != "" && req.Username != account.Username {
		exists, err := c.AccountService.UsernameExists(ctx, req.Username)
		if err != nil {
			res.SetError(gerror.New("检查用户名失败: " + err.Error()))
			return res, nil
		}
		if exists {
			res.SetError(gerror.New("用户名已存在"))
			return res, nil
		}
	}

	// 检查邮箱是否已存在
	if req.Email != "" && req.Email != account.Email {
		exists, err := c.AccountService.EmailExists(ctx, req.Email)
		if err != nil {
			res.SetError(gerror.New("检查邮箱失败: " + err.Error()))
			return res, nil
		}
		if exists {
			res.SetError(gerror.New("邮箱已存在"))
			return res, nil
		}
	}

	// 更新账户信息
	updateData := &service.Account{
		Id:        req.AccountId,
		Username:  req.Username,
		Email:     req.Email,
		Status:    req.Status,
		Lang:      req.Lang,
		UpdatedAt: time.Now(),
	}

	err = c.AccountService.Update(ctx, updateData)
	if err != nil {
		res.SetError(gerror.New("更新账户失败: " + err.Error()))
		return res, nil
	}

	// 更新角色
	if len(req.RoleIds) > 0 {
		// 先移除所有角色
		err = c.AccountService.ClearRoles(ctx, req.AccountId)
		if err != nil {
			res.SetError(gerror.New("清除角色失败: " + err.Error()))
			return res, nil
		}

		// 分配新角色
		for _, roleId := range req.RoleIds {
			err = c.AccountService.AssignRole(ctx, req.AccountId, roleId)
			if err != nil {
				g.Log().Warning(ctx, "分配角色失败: ", err)
			}
		}
	}

	// 准备响应数据
	res.Success = true
	res.Code = 0
	res.Msg = "更新成功"

	return res, nil
}

// UpdateAccountPassword 更新账户密码
func (c *ControllerV1) UpdateAccountPassword(ctx context.Context, req *v1.AccountPasswordReq) (res *v1.AccountPasswordRes, err error) {
	res = &v1.AccountPasswordRes{}

	// 获取当前用户
	currentAccountId := service.GetCurrentAccountId(ctx)
	currentRoles := service.GetCurrentRoles(ctx)

	// 如果不是管理员，只能修改自己的密码
	isAdmin := false
	for _, role := range currentRoles {
		if role == "admin" {
			isAdmin = true
			break
		}
	}

	if !isAdmin && currentAccountId != req.AccountId {
		res.SetError(gerror.New("权限不足"))
		return res, nil
	}

	// 如果不是管理员，需要验证旧密码
	if !isAdmin && req.OldPassword == "" {
		res.SetError(gerror.New("请提供旧密码"))
		return res, nil
	}

	// 如果提供了旧密码，验证旧密码
	if req.OldPassword != "" {
		account, err := c.AccountService.GetById(ctx, req.AccountId)
		if err != nil {
			res.SetError(gerror.New("获取账户失败: " + err.Error()))
			return res, nil
		}

		// 验证旧密码
		if !c.AccountService.VerifyPassword(account.Password, req.OldPassword) {
			res.SetError(gerror.New("旧密码不正确"))
			return res, nil
		}
	}

	// 更新密码
	err = c.AccountService.UpdatePassword(ctx, req.AccountId, req.NewPassword)
	if err != nil {
		res.SetError(gerror.New("更新密码失败: " + err.Error()))
		return res, nil
	}

	// 准备响应数据
	res.Success = true
	res.Code = 0
	res.Msg = "密码更新成功"

	return res, nil
}

// DeleteAccount 删除账户
func (c *ControllerV1) DeleteAccount(ctx context.Context, req *v1.AccountDeleteReq) (res *v1.AccountDeleteRes, err error) {
	res = &v1.AccountDeleteRes{}

	// 检查权限
	if !middlewares.HasPermission(ctx, "account", "delete", "account") {
		res.SetError(gerror.New("权限不足"))
		return res, nil
	}

	// 获取账户
	account, err := c.AccountService.GetById(ctx, req.AccountId)
	if err != nil {
		res.SetError(gerror.New("获取账户失败: " + err.Error()))
		return res, nil
	}

	// 不能删除自己
	currentAccountId := service.GetCurrentAccountId(ctx)
	if currentAccountId == req.AccountId {
		res.SetError(gerror.New("不能删除自己的账户"))
		return res, nil
	}

	// 检查是否是admin账户
	if account.Username == "admin" {
		res.SetError(gerror.New("不能删除管理员账户"))
		return res, nil
	}

	// 删除账户
	err = c.AccountService.Delete(ctx, req.AccountId)
	if err != nil {
		res.SetError(gerror.New("删除账户失败: " + err.Error()))
		return res, nil
	}

	// 准备响应数据
	res.Success = true
	res.Code = 0
	res.Msg = "删除成功"

	return res, nil
}
