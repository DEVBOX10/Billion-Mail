package rbac

import (
	"billionmail-core/internal/model"
	"billionmail-core/internal/service/middlewares"
	service "billionmail-core/internal/service/rbac"
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"billionmail-core/api/rbac/v1"
)

// AccountList gets account list
func (c *ControllerV1) AccountList(ctx context.Context, req *v1.AccountListReq) (res *v1.AccountListRes, err error) {
	res = &v1.AccountListRes{}

	// Check permission
	if !middlewares.HasPermission(ctx, "account", "read", "account") {
		res.SetError(gerror.New("Insufficient permissions"))
		return
	}

	// Get account list
	accounts, total, err := service.Account().GetList(ctx, req.Page, req.PageSize, req.Username, req.Email, req.Status)
	if err != nil {
		res.SetError(gerror.New("Failed to get account list: " + err.Error()))
		return
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
			Id:         account.AccountId,
			Username:   account.Username,
			Email:      account.Email,
			Status:     account.Status,
			Language:   account.Language,
			CreateTime: account.CreateTime,
		})
	}

	return
}

// AccountDetail gets account details
func (c *ControllerV1) AccountDetail(ctx context.Context, req *v1.AccountDetailReq) (res *v1.AccountDetailRes, err error) {
	res = &v1.AccountDetailRes{}

	// Check permission
	if !middlewares.HasPermission(ctx, "account", "read", "account") {
		res.SetError(gerror.New("Insufficient permissions"))
		return
	}

	// Get account details
	account, err := service.Account().GetById(ctx, req.AccountId)
	if err != nil {
		res.SetError(gerror.New("Failed to get account details: " + err.Error()))
		return
	}

	// Get account roles
	roles, err := service.Account().GetAccountRoles(ctx, req.AccountId)
	if err != nil {
		res.SetError(gerror.New("Failed to get account roles: " + err.Error()))
		return
	}

	// Get all roles
	allRoles, err := service.Account().GetAll(ctx)
	if err != nil {
		res.SetError(gerror.New("Failed to get all roles: " + err.Error()))
		return
	}

	// Prepare response data
	res.Success = true
	res.Code = 0
	res.Msg = "Success"

	// Set account info
	res.Data.Account = v1.AccountInfoItem{
		Id:         account.AccountId,
		Username:   account.Username,
		Email:      account.Email,
		Status:     account.Status,
		Language:   account.Language,
		CreateTime: account.CreateTime,
	}

	// Set account roles
	res.Data.Roles = make([]v1.RoleInfoItem, 0, len(roles))
	for _, role := range roles {
		res.Data.Roles = append(res.Data.Roles, v1.RoleInfoItem{
			Id:          role.RoleId,
			Name:        role.RoleName,
			Description: role.Description,
			Status:      role.Status,
			CreateTime:  role.CreateTime,
		})
	}

	// Set all roles
	res.Data.AllRoles = make([]v1.RoleInfoItem, 0, len(allRoles))
	for _, role := range allRoles {
		res.Data.AllRoles = append(res.Data.AllRoles, v1.RoleInfoItem{
			Id:          role.RoleId,
			Name:        role.RoleName,
			Description: role.Description,
			Status:      role.Status,
			CreateTime:  role.CreateTime,
		})
	}

	return
}

// AccountCreate create account
func (c *ControllerV1) AccountCreate(ctx context.Context, req *v1.AccountCreateReq) (res *v1.AccountCreateRes, err error) {
	res = &v1.AccountCreateRes{}

	// Check permissions
	if !middlewares.HasPermission(ctx, "account", "create", "account") {
		res.SetError(gerror.New("Insufficient permissions"))
		return
	}

	// Check if username already exists
	exists, err := service.Account().UsernameExists(ctx, req.Username)
	if err != nil {
		res.SetError(gerror.New("Failed to check username: " + err.Error()))
		return
	}
	if exists {
		res.SetError(gerror.New("Username already exists"))
		return
	}

	// Check if email already exists
	exists, err = service.Account().EmailExists(ctx, req.Email)
	if err != nil {
		res.SetError(gerror.New("Failed to check email: " + err.Error()))
		return
	}
	if exists {
		res.SetError(gerror.New("Email already exists"))
		return
	}

	// Create account
	accountData := &model.Account{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Status:   req.Status,
		Language: req.Lang,
	}

	accountId, err := service.Account().Create(ctx, accountData)
	if err != nil {
		res.SetError(gerror.New("Failed to create account: " + err.Error()))
		return
	}

	// Assign roles
	if len(req.RoleIds) > 0 {
		for _, roleId := range req.RoleIds {
			err = service.Account().AssignRole(ctx, accountId, roleId)
			if err != nil {
				g.Log().Warning(ctx, "Failed to assign role: ", err)
			}
		}
	}

	// Prepare response data
	res.Success = true
	res.Code = 0
	res.Msg = "Created successfully"
	res.Data.AccountId = accountId

	return
}

// AccountUpdate update account
func (c *ControllerV1) AccountUpdate(ctx context.Context, req *v1.AccountUpdateReq) (res *v1.AccountUpdateRes, err error) {
	res = &v1.AccountUpdateRes{}

	// Check permissions
	if !middlewares.HasPermission(ctx, "account", "update", "account") {
		res.SetError(gerror.New("Insufficient permissions"))
		return
	}

	// Get account
	account, err := service.Account().GetById(ctx, req.AccountId)
	if err != nil {
		res.SetError(gerror.New("Failed to get account: " + err.Error()))
		return
	}

	var exists bool

	// Check if username already exists
	if req.Username != "" && req.Username != account.Username {
		exists, err = service.Account().UsernameExists(ctx, req.Username)
		if err != nil {
			res.SetError(gerror.New("Failed to check username: " + err.Error()))
			return
		}
		if exists {
			res.SetError(gerror.New("Username already exists"))
			return
		}
	}

	// Check if email already exists
	if req.Email != "" && req.Email != account.Email {
		exists, err = service.Account().EmailExists(ctx, req.Email)
		if err != nil {
			res.SetError(gerror.New("Failed to check email: " + err.Error()))
			return
		}
		if exists {
			res.SetError(gerror.New("Email already exists"))
			return
		}
	}

	// Update account info
	updateData := &model.Account{
		AccountId:  req.AccountId,
		Username:   req.Username,
		Email:      req.Email,
		Status:     req.Status,
		Language:   req.Lang,
		UpdateTime: time.Now().Unix(),
	}

	err = service.Account().Update(ctx, updateData)
	if err != nil {
		res.SetError(gerror.New("Failed to update account: " + err.Error()))
		return
	}

	// Update roles
	if len(req.RoleIds) > 0 {
		// Remove all roles first
		err = service.Account().ClearRoles(ctx, req.AccountId)
		if err != nil {
			res.SetError(gerror.New("Failed to clear roles: " + err.Error()))
			return
		}

		// Assign new roles
		for _, roleId := range req.RoleIds {
			err = service.Account().AssignRole(ctx, req.AccountId, roleId)
			if err != nil {
				g.Log().Warning(ctx, "Failed to assign role: ", err)
			}
		}
	}

	// Prepare response data
	res.Success = true
	res.Code = 0
	res.Msg = "Updated successfully"

	return
}

// AccountPassword update account password
func (c *ControllerV1) AccountPassword(ctx context.Context, req *v1.AccountPasswordReq) (res *v1.AccountPasswordRes, err error) {
	res = &v1.AccountPasswordRes{}

	// Get current user
	currentAccountId := service.GetCurrentAccountId(ctx)
	currentRoles := service.GetCurrentRoles(ctx)

	// If not admin, can only modify own password
	isAdmin := false
	for _, role := range currentRoles {
		if role == "admin" {
			isAdmin = true
			break
		}
	}

	if !isAdmin && currentAccountId != req.AccountId {
		res.SetError(gerror.New("Insufficient permissions"))
		return
	}

	// If not admin, need to verify old password
	if !isAdmin && req.OldPassword == "" {
		res.SetError(gerror.New("Please provide the old password"))
		return
	}

	// If old password provided, verify it
	if req.OldPassword != "" {
		var account *model.Account
		account, err = service.Account().GetById(ctx, req.AccountId)
		if err != nil {
			res.SetError(gerror.New("Failed to get account: " + err.Error()))
			return
		}

		// Verify old password
		if !service.Account().VerifyPassword(account.Password, req.OldPassword) {
			res.SetError(gerror.New("Incorrect old password"))
			return
		}
	}

	// Update password
	err = service.Account().UpdatePassword(ctx, req.AccountId, req.NewPassword)
	if err != nil {
		res.SetError(gerror.New("Failed to update password: " + err.Error()))
		return
	}

	// Prepare response data
	res.Success = true
	res.Code = 0
	res.Msg = "Password updated successfully"

	return
}

// AccountDelete delete account
func (c *ControllerV1) AccountDelete(ctx context.Context, req *v1.AccountDeleteReq) (res *v1.AccountDeleteRes, err error) {
	res = &v1.AccountDeleteRes{}

	// Check permissions
	if !middlewares.HasPermission(ctx, "account", "delete", "account") {
		res.SetError(gerror.New("Insufficient permissions"))
		return
	}

	// Get account
	account, err := service.Account().GetById(ctx, req.AccountId)
	if err != nil {
		res.SetError(gerror.New("Failed to get account: " + err.Error()))
		return
	}

	// Cannot delete yourself
	currentAccountId := service.GetCurrentAccountId(ctx)
	if currentAccountId == req.AccountId {
		res.SetError(gerror.New("Cannot delete your own account"))
		return
	}

	// Check if admin account
	if account.Username == "admin" {
		res.SetError(gerror.New("Cannot delete admin account"))
		return
	}

	// Delete account
	err = service.Account().Delete(ctx, req.AccountId)
	if err != nil {
		res.SetError(gerror.New("Failed to delete account: " + err.Error()))
		return
	}

	// Prepare response data
	res.Success = true
	res.Code = 0
	res.Msg = "Deleted successfully"

	return
}
