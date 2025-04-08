package rbac

import (
	"billionmail-core/api/rbac/v1"
	service "billionmail-core/internal/service/rbac"
	"context"
	"fmt"
	"github.com/gogf/gf/util/gconv"
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
	roles, err := service.Account().GetAccountRoles(ctx, account.AccountId)
	if err != nil {
		err = fmt.Errorf("Failed to get account roles: %w", err)
		return
	}

	// Convert roles to role names
	roleNames := make([]string, 0, len(roles))
	for _, role := range roles {
		roleNames = append(roleNames, role.RoleName)
	}

	// Generate JWT token
	token, _, err := service.JWT().GenerateToken(account.AccountId, account.Username, roleNames)
	if err != nil {
		res.SetError(gerror.New("Failed to generate token"))
		return res, nil
	}

	// Generate refresh token
	refreshToken, err := service.JWT().GenerateRefreshToken(account.AccountId, account.Username)
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
	res.Data.TTL = gconv.Int64(service.JWT().AccessExpiry.Seconds())

	// Set account basic information
	res.Data.AccountInfo.Id = account.AccountId
	res.Data.AccountInfo.Username = account.Username
	res.Data.AccountInfo.Email = account.Email
	res.Data.AccountInfo.Status = account.Status
	res.Data.AccountInfo.Lang = account.Language

	return res, nil
}

// Logout handles user logout
func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	res = &v1.LogoutRes{}

	// Parse the token from the request
	claims, err := service.JWT().ParseToken(req.Authorization)
	if err != nil {
		res.SetError(gerror.New("Invalid or expired refresh token"))
		return res, nil
	}

	// Add token to blacklist
	err = service.JWT().InvalidateToken(claims)
	if err != nil {
		res.SetError(gerror.New("Logout failed"))
		return res, nil
	}

	res.Success = true
	res.Code = 0
	res.Msg = "Logout successful"

	return res, nil
}

// RefreshToken handles token refresh
func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	res = &v1.RefreshTokenRes{}

	// Verify refresh token
	claims, err := service.JWT().ParseToken(req.RefreshToken)
	if err != nil {
		res.SetError(gerror.New("Invalid or expired refresh token"))
		return res, nil
	}

	// Generate new JWT token
	token, _, err := service.JWT().GenerateToken(claims.AccountId, claims.Username, []string{})
	if err != nil {
		res.SetError(gerror.New("Failed to generate token"))
		return res, nil
	}

	// Generate new refresh token
	refreshToken, err := service.JWT().GenerateRefreshToken(claims.AccountId, claims.Username)
	if err != nil {
		res.SetError(gerror.New("Failed to generate refresh token"))
		return res, nil
	}

	// Prepare response
	res.Success = true
	res.Code = 0
	res.Msg = "Token refreshed successfully"
	res.Data.Token = token
	res.Data.RefreshToken = refreshToken
	res.Data.TTL = gconv.Int64(service.JWT().RefreshExpiry.Seconds())

	return res, nil
}

// CurrentUser retrieves the current logged-in user information
func (c *ControllerV1) CurrentUser(ctx context.Context, req *v1.CurrentUserReq) (res *v1.CurrentUserRes, err error) {
	res = &v1.CurrentUserRes{}

	// Get account ID from context
	accountId := service.GetCurrentAccountId(ctx)
	if accountId == 0 {
		err = gerror.New("Unauthorized")
		return
	}

	// Get account details
	account, err := service.Account().GetById(ctx, accountId)
	if err != nil {
		err = gerror.New("Failed to get account details")
		return
	}

	// Get account roles
	roles, err := service.Account().GetAccountRoles(ctx, accountId)
	if err != nil {
		err = gerror.New("Failed to get account roles")
		return
	}

	// Get account permissions
	permissions, err := service.Account().GetAccountPermissions(ctx, accountId)
	if err != nil {
		err = gerror.New("Failed to get account permissions")
		return
	}

	// Prepare response
	res.Success = true
	res.Code = 0
	res.Msg = "Retrieved successfully"

	// Set account information
	res.Data.Account.Id = account.AccountId
	res.Data.Account.Username = account.Username
	res.Data.Account.Email = account.Email
	res.Data.Account.Status = account.Status
	res.Data.Account.Lang = account.Language

	// Set roles
	res.Data.Roles = make([]string, 0, len(roles))
	for _, role := range roles {
		res.Data.Roles = append(res.Data.Roles, role.RoleName)
	}

	// Set permissions
	res.Data.Permissions = make([]string, 0, len(permissions))
	for _, perm := range permissions {
		permStr := perm.Module + ":" + perm.Action + ":" + perm.Resource
		res.Data.Permissions = append(res.Data.Permissions, permStr)
	}

	return res, nil
}
