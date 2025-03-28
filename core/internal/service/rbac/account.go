package rbac

import (
	"billionmail-core/internal/model"
	"context"
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
func (s *accountService) Create(ctx context.Context, username, password, email string, status int) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := g.DB().Model("account").Data(g.Map{
		"username":    username,
		"password":    string(hashedPassword),
		"email":       email,
		"status":      status,
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
func (s *accountService) Update(ctx context.Context, accountId int64, username, email string, status int) error {
	_, err := g.DB().Model("account").Data(g.Map{
		"username":    username,
		"email":       email,
		"status":      status,
		"update_time": time.Now().Unix(),
	}).Where("account_id = ?", accountId).Update()
	return err
}

// Delete deletes an account
func (s *accountService) Delete(ctx context.Context, accountId int64) error {
	_, err := g.DB().Model("account").Where("account_id = ?", accountId).Delete()
	return err
}

// UpdatePassword updates account password
func (s *accountService) UpdatePassword(ctx context.Context, accountId int64, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = g.DB().Model("account").Data(g.Map{
		"password":    string(hashedPassword),
		"update_time": time.Now().Unix(),
	}).Where("account_id = ?", accountId).Update()
	return err
}

// VerifyPassword verifies account password
func (s *accountService) VerifyPassword(ctx context.Context, accountId int64, password string) (bool, error) {
	var account model.Account
	err := g.DB().Model("account").Where("account_id = ?", accountId).Scan(&account)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	return err == nil, nil
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
			"create_time": time.Now().Unix(),
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
		LeftJoin("account_role", "role.role_id=account_role.role_id").
		Where("account_role.account_id = ?", accountId).
		Scan(&roles)
	return roles, err
}

// GetPermissions gets permissions assigned to an account
func (s *accountService) GetPermissions(ctx context.Context, accountId int64) ([]model.Permission, error) {
	var permissions []model.Permission
	err := g.DB().Model("permission").
		LeftJoin("role_permission", "permission.permission_id=role_permission.permission_id").
		LeftJoin("account_role", "role_permission.role_id=account_role.role_id").
		Where("account_role.account_id = ?", accountId).
		Scan(&permissions)
	return permissions, err
}

// Login handles user login
func (s *accountService) Login(ctx context.Context, username, password string) (*model.Account, string, error) {
	var account model.Account
	err := g.DB().Model("account").Where("username = ?", username).Scan(&account)
	if err != nil {
		return nil, "", err
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		return nil, "", err
	}

	// Update last login time
	_, err = g.DB().Model("account").Data(g.Map{
		"last_login_time": time.Now().Unix(),
	}).Where("account_id = ?", account.Id).Update()
	if err != nil {
		return nil, "", err
	}

	return &account, account.Password, nil
}

// IsAdmin checks if account is admin
func (s *accountService) IsAdmin(ctx context.Context, accountId int64) (bool, error) {
	roles, err := s.GetRoles(ctx, accountId)
	if err != nil {
		return false, err
	}

	for _, role := range roles {
		if role.Name == "admin" {
			return true, nil
		}
	}

	return false, nil
}

// CountAdmins counts admin accounts
func (s *accountService) CountAdmins(ctx context.Context) (int, error) {
	count, err := g.DB().Model("account_role").
		LeftJoin("role", "account_role.role_id=role.role_id").
		Where("role.role_name = ?", "admin").
		Count()
	return count, err
}
