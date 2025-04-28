package rbac

import (
	"billionmail-core/internal/model"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/crypto/bcrypt"
)

type accountService struct{}

func newAccountService() *accountService {
	return &accountService{}
}

// GetList gets account list with pagination
func (s *accountService) GetList(ctx context.Context, page, pageSize int, username, email string, status int) ([]model.Account, int, error) {
	var accounts []model.Account
	var total int

	query := g.DB().Model("account")
	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}
	if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	// Get total count
	count, err := query.Count()
	if err != nil {
		return nil, 0, err
	}
	total = count

	// Get paginated data
	err = query.Page(page, pageSize).Scan(&accounts)
	if err != nil {
		return nil, 0, err
	}

	return accounts, total, nil
}

// GetById gets account details by ID
func (s *accountService) GetById(ctx context.Context, accountId int64) (*model.Account, error) {
	var account model.Account
	err := g.DB().Model("account").Where("account_id = ?", accountId).Scan(&account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// Create creates a new account
func (s *accountService) Create(ctx context.Context, accountData *model.Account) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(accountData.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := g.DB().Model("account").Data(g.Map{
		"username":    accountData.Username,
		"password":    string(hashedPassword),
		"email":       accountData.Email,
		"status":      accountData.Status,
		"language":    accountData.Language,
		"create_time": time.Now().Unix(),
		"update_time": time.Now().Unix(),
	}).Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update updates account information
func (s *accountService) Update(ctx context.Context, accountData *model.Account) error {
	_, err := g.DB().Model("account").Data(g.Map{
		"username":    accountData.Username,
		"email":       accountData.Email,
		"status":      accountData.Status,
		"language":    accountData.Language,
		"update_time": time.Now().Unix(),
	}).Where("account_id = ?", accountData.AccountId).Update()
	return err
}

// Delete deletes an account
func (s *accountService) Delete(ctx context.Context, accountId int64) error {
	_, err := g.DB().Model("account").Where("account_id = ?", accountId).Delete()
	return err
}

// UpdatePassword updates account password
func (s *accountService) UpdatePassword(ctx context.Context, accountId int64, newPassword string) error {
	hashedPassword, err := s.GeneratePasswordHash(newPassword)
	if err != nil {
		return err
	}

	_, err = g.DB().Model("account").Data(g.Map{
		"password":    hashedPassword,
		"update_time": time.Now().Unix(),
	}).Where("account_id = ?", accountId).Update()
	return err
}

// VerifyPassword verifies account password
func (s *accountService) VerifyPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

// GeneratePasswordHash generates a hashed password
func (s *accountService) GeneratePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// UsernameExists checks if a username already exists
func (s *accountService) UsernameExists(ctx context.Context, username string) (bool, error) {
	count, err := g.DB().Model("account").Where("username = ?", username).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// EmailExists checks if an email already exists
func (s *accountService) EmailExists(ctx context.Context, email string) (bool, error) {
	count, err := g.DB().Model("account").Where("email = ?", email).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetAccountRoles gets the role list for an account
func (s *accountService) GetAccountRoles(ctx context.Context, accountId int64) ([]model.Role, error) {
	var roles []model.Role
	err := g.DB().Model("role r").
		LeftJoin("account_role ar", "r.role_id = ar.role_id").
		Where("ar.account_id = ?", accountId).
		Fields("r.*").
		Scan(&roles)
	return roles, err
}

// AssignRole assigns a single role to an account
func (s *accountService) AssignRole(ctx context.Context, accountId int64, roleId int64) error {
	_, err := g.DB().Model("account_role").Data(g.Map{
		"account_id":  accountId,
		"role_id":     roleId,
		"create_time": time.Now(),
	}).Insert()
	return err
}

// ClearRoles clears all roles of an account
func (s *accountService) ClearRoles(ctx context.Context, accountId int64) error {
	_, err := g.DB().Model("account_role").Where("account_id = ?", accountId).Delete()
	return err
}

// GetAll gets all roles
func (s *accountService) GetAll(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	err := g.DB().Model("role").Scan(&roles)
	return roles, err
}

// BindRoles binds roles to an account
func (s *accountService) BindRoles(ctx context.Context, accountId int64, roleIds []int64) error {
	// Delete existing roles
	_, err := g.DB().Model("account_role").Where("account_id = ?", accountId).Delete()
	if err != nil {
		return err
	}

	// Add new roles
	for _, roleId := range roleIds {
		_, err = g.DB().Model("account_role").Data(g.Map{
			"account_id":  accountId,
			"role_id":     roleId,
			"create_time": time.Now(),
		}).Insert()
		if err != nil {
			return err
		}
	}

	return nil
}

// GetRoles gets roles assigned to an account
func (s *accountService) GetRoles(ctx context.Context, accountId int64) ([]model.Role, error) {
	var roles []model.Role
	err := g.DB().Model("role").
		LeftJoin("account_role", "role.id=account_role.role_id").
		Where("account_role.account_id = ?", accountId).
		Scan(&roles)
	return roles, err
}

// GetPermissions gets permissions assigned to an account
func (s *accountService) GetPermissions(ctx context.Context, accountId int64) ([]model.Permission, error) {
	var permissions []model.Permission
	err := g.DB().Model("permission").
		LeftJoin("role_permission", "permission.id=role_permission.permission_id").
		LeftJoin("account_role", "role_permission.role_id=account_role.role_id").
		Where("account_role.account_id = ?", accountId).
		Scan(&permissions)
	return permissions, err
}

// GetAccountPermissions gets the permission list for an account (alias for GetPermissions)
func (s *accountService) GetAccountPermissions(ctx context.Context, accountId int64) ([]model.Permission, error) {
	return s.GetPermissions(ctx, accountId)
}

// Login handles user login
func (s *accountService) Login(ctx context.Context, username, password string) (*model.Account, error) {
	var account model.Account
	err := g.DB().Model("account").Where("username", username).Scan(&account)
	if err != nil {
		return nil, err
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	// Update last login time
	_, err = g.DB().Model("account").Data(g.Map{
		"last_login_time": time.Now().Unix(),
	}).Where("account_id", account.AccountId).Update()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// IsAdmin checks if account is admin
func (s *accountService) IsAdmin(ctx context.Context, accountId int64) (bool, error) {
	roles, err := s.GetRoles(ctx, accountId)
	if err != nil {
		return false, err
	}

	for _, role := range roles {
		if role.RoleName == "admin" {
			return true, nil
		}
	}

	return false, nil
}

// CountAdmins counts admin accounts
func (s *accountService) CountAdmins(ctx context.Context) (int, error) {
	count, err := g.DB().Model("account_role").
		LeftJoin("role", "account_role.role_id=role.id").
		Where("role.name = ?", "admin").
		Count()
	return count, err
}

// GetCurrentAccountId gets the current user ID from context
func GetCurrentAccountId(ctx context.Context) int64 {
	value := ctx.Value("accountId")
	if value == nil {
		return 0
	}

	accountId, ok := value.(int64)
	if !ok {
		return 0
	}

	return accountId
}

// GetCurrentRoles gets the current user roles from context
func GetCurrentRoles(ctx context.Context) []string {
	value := ctx.Value("roles")
	if value == nil {
		return []string{}
	}

	roles, ok := value.([]string)
	if !ok {
		return []string{}
	}

	return roles
}

// GetCurrentAccount gets the current user account from context
func GetCurrentAccount(ctx context.Context) (acc *model.Account, err error) {
	accountId := GetCurrentAccountId(ctx)

	if accountId == 0 {
		return nil, fmt.Errorf("account ID not found in context")
	}

	if err = g.DB().Model("account").Where("account_id = ?", accountId).Scan(&acc); err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	return
}
