package rbac

import (
	service "billionmail-core/internal/service/rbac"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"billionmail-core/api/rbac/v1"
)

// Login handles user login
func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = &v1.LoginRes{}

	// Verify username and password
	account, err := c.AccountService.Verify(ctx, req.Username, req.Password)
	if err != nil {
		res.SetError(gerror.New("Invalid username or password"))
		return res, nil
	}

	// Get account roles
	roles, err := c.AccountService.GetAccountRoles(ctx, account.Id)
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
	token, expiresAt, err := c.JWTService.GenerateToken(account.Id, account.Username, roleNames)
	if err != nil {
		res.SetError(gerror.New("Failed to generate token"))
		return res, nil
	}

	// Generate refresh token
	refreshToken, err := c.JWTService.GenerateRefreshToken(account.Id, account.Username)
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
	username, err := c.JWTService.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		res.SetError(gerror.New("无效或过期的刷新令牌"))
		return res, nil
	}

	// 通过用户名获取账户
	account, err := c.AccountService.GetByUsername(ctx, username)
	if err != nil {
		res.SetError(gerror.New("账户未找到"))
		return res, nil
	}

	// 获取账户角色
	roles, err := c.AccountService.GetAccountRoles(ctx, account.Id)
	if err != nil {
		res.SetError(gerror.New("获取账户角色失败"))
		return res, nil
	}

	// 转换角色对象为角色名称
	roleNames := make([]string, 0, len(roles))
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}

	// 生成新的JWT令牌
	token, expiresAt, err := c.JWTService.GenerateToken(account.Id, account.Username, roleNames)
	if err != nil {
		res.SetError(gerror.New("生成令牌失败"))
		return res, nil
	}

	// 生成新的刷新令牌
	refreshToken, err := c.JWTService.GenerateRefreshToken(account.Id, account.Username)
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

// Register 用户注册
func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	res = &v1.RegisterRes{}

	// 检查用户名是否已存在
	exists, err := c.AccountService.UsernameExists(ctx, req.Username)
	if err != nil {
		res.SetError(gerror.New("检查用户名失败"))
		return res, nil
	}
	if exists {
		res.SetError(gerror.New("用户名已存在"))
		return res, nil
	}

	// 检查邮箱是否已存在
	exists, err = c.AccountService.EmailExists(ctx, req.Email)
	if err != nil {
		res.SetError(gerror.New("检查邮箱失败"))
		return res, nil
	}
	if exists {
		res.SetError(gerror.New("邮箱已存在"))
		return res, nil
	}

	// 创建账户
	accountData := &service.Account{
		Username: req.Username,
		Password: req.Password, // 密码将在服务层中哈希
		Email:    req.Email,
		Status:   1, // 默认激活
		Lang:     req.Lang,
	}

	accountId, err := c.AccountService.Create(ctx, accountData)
	if err != nil {
		res.SetError(gerror.New("创建账户失败"))
		return res, nil
	}

	// 分配默认用户角色
	userRole, err := c.RoleService.GetByName(ctx, "user")
	if err == nil && userRole != nil {
		err = c.AccountService.AssignRole(ctx, accountId, userRole.Id)
		if err != nil {
			g.Log().Warning(ctx, "分配默认用户角色失败:", err)
		}
	}

	res.Success = true
	res.Code = 0
	res.Msg = "注册成功"
	res.Data.AccountId = accountId

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
	account, err := c.AccountService.GetById(ctx, accountId)
	if err != nil {
		res.SetError(gerror.New("获取账户详情失败"))
		return res, nil
	}

	// 获取账户角色
	roles, err := c.AccountService.GetAccountRoles(ctx, accountId)
	if err != nil {
		res.SetError(gerror.New("获取账户角色失败"))
		return res, nil
	}

	// 获取账户权限
	permissions, err := c.AccountService.GetAccountPermissions(ctx, accountId)
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
