package rbac

import (
	service "billionmail-core/internal/service/rbac"
	"context"

	"billionmail-core/api/rbac/v1"
	"github.com/gogf/gf/v2/errors/gerror"
)

// Login handles user login
func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = &v1.LoginRes{}

	// Verify username and password
	account, err := service.Account().Login(ctx, req.Username, req.Password)
	if err != nil {
		res.SetError(gerror.New("Invalid username or password"))
		return res, nil
	}

	// Get account roles
	roles, err := service.Account().GetAccountRoles(ctx, account.Id)
	if err != nil {
		res.SetError(gerror.New("Failed to get account roles"))
		return res, nil
	}

	// Convert roles to role names
	roleNames := make([]string, 0, len(roles))
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}

	// Generate JWT token
	token, expiresAt, err := service.JWT().GenerateToken(account.Id, account.Username, roleNames)
	if err != nil {
		res.SetError(gerror.New("Failed to generate token"))
		return res, nil
	}

	// Generate refresh token
	refreshToken, err := service.JWT().GenerateRefreshToken(account.Id, account.Username)
	if err != nil {
		res.SetError(gerror.New("Failed to generate refresh token"))
		return res, nil
	}

	// Prepare response
	res.Success = true
	res.Code = 0
	res.Msg = "Login successful"
	res.Data.Token = token
	res.Data.RefreshToken = refreshToken
	res.Data.ExpiresAt = expiresAt

	// 设置账户基本信息
	res.Data.AccountInfo.Id = account.Id
	res.Data.AccountInfo.Username = account.Username
	res.Data.AccountInfo.Email = account.Email
	res.Data.AccountInfo.Status = account.Status
	res.Data.AccountInfo.Lang = account.Lang

	return res, nil
}

// Logout 用户登出
func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	res = &v1.LogoutRes{}

	// 在无状态JWT认证系统中，登出通常在客户端处理
	// 通过删除存储的令牌。服务器端可以实现令牌黑名单
	// 以增强安全性，但这超出了本示例的范围。

	res.Success = true
	res.Code = 0
	res.Msg = "登出成功"

	return res, nil
}

// RefreshToken 刷新令牌
func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	res = &v1.RefreshTokenRes{}

	// 验证刷新令牌
	claims, err := service.JWT().ParseToken(req.RefreshToken)
	if err != nil {
		res.SetError(gerror.New("无效或过期的刷新令牌"))
		return res, nil
	}

	// 生成新的JWT令牌
	token, expiresAt, err := service.JWT().GenerateToken(claims.AccountId, claims.Username, []string{})
	if err != nil {
		res.SetError(gerror.New("生成令牌失败"))
		return res, nil
	}

	// 生成新的刷新令牌
	refreshToken, err := service.JWT().GenerateRefreshToken(claims.AccountId, claims.Username)
	if err != nil {
		res.SetError(gerror.New("生成刷新令牌失败"))
		return res, nil
	}

	// 准备响应
	res.Success = true
	res.Code = 0
	res.Msg = "令牌刷新成功"
	res.Data.Token = token
	res.Data.RefreshToken = refreshToken
	res.Data.ExpiresAt = expiresAt

	return res, nil
}

// CurrentUser 获取当前登录用户信息
func (c *ControllerV1) CurrentUser(ctx context.Context, req *v1.CurrentUserReq) (res *v1.CurrentUserRes, err error) {
	res = &v1.CurrentUserRes{}

	// 从上下文获取账户ID
	accountId := service.GetCurrentAccountId(ctx)
	if accountId == 0 {
		res.SetError(gerror.New("未认证"))
		return res, nil
	}

	// 获取账户详情
	account, err := service.Account().GetById(ctx, accountId)
	if err != nil {
		res.SetError(gerror.New("获取账户详情失败"))
		return res, nil
	}

	// 获取账户角色
	roles, err := service.Account().GetAccountRoles(ctx, accountId)
	if err != nil {
		res.SetError(gerror.New("获取账户角色失败"))
		return res, nil
	}

	// 获取账户权限
	permissions, err := service.Account().GetAccountPermissions(ctx, accountId)
	if err != nil {
		res.SetError(gerror.New("获取账户权限失败"))
		return res, nil
	}

	// 准备响应
	res.Success = true
	res.Code = 0
	res.Msg = "获取成功"

	// 设置账户信息
	res.Data.Account.Id = account.Id
	res.Data.Account.Username = account.Username
	res.Data.Account.Email = account.Email
	res.Data.Account.Status = account.Status
	res.Data.Account.Lang = account.Lang
	res.Data.Account.CreatedAt = account.CreatedAt.Format("2006-01-02 15:04:05")

	// 设置角色
	res.Data.Roles = make([]string, 0, len(roles))
	for _, role := range roles {
		res.Data.Roles = append(res.Data.Roles, role.Name)
	}

	// 设置权限
	res.Data.Permissions = make([]string, 0, len(permissions))
	for _, perm := range permissions {
		permStr := perm.Module + ":" + perm.Action + ":" + perm.Resource
		res.Data.Permissions = append(res.Data.Permissions, permStr)
	}

	return res, nil
}
