package rbac

import (
	"billionmail-core/internal/service/middlewares"
	"billionmail-core/internal/service/rbac"
	"github.com/gogf/gf/v2/net/ghttp"
)

// RegisterRouter registers all RBAC related routes
func RegisterRouter(group *ghttp.RouterGroup) {
	// Create JWT middleware
	jwtService := rbac.NewJWTService()

	// Create RBAC middleware
	rbacMiddleware := middlewares.NewRBACMiddleware()

	// Apply JWT authentication middleware to all routes
	group.Middleware(jwtService.JWTAuthMiddleware)

	// Apply RBAC permission middleware after JWT
	group.Middleware(rbacMiddleware.PermissionCheck)

	// Register authentication controllers
	loginController := NewLoginController()
	group.Bind(
		loginController.Login,
		loginController.Logout,
		loginController.RefreshToken,
		loginController.Register,
		loginController.CurrentUser,
	)

	// Register account management controllers
	accountController := NewAccountController()
	group.Bind(
		accountController.List,
		accountController.Detail,
		accountController.Create,
		accountController.Update,
		accountController.Delete,
		accountController.UpdatePassword,
	)

	// Register role management controllers
	roleController := NewRoleController()
	group.Bind(
		roleController.List,
		roleController.Detail,
		roleController.Create,
		roleController.Update,
		roleController.Delete,
	)

	// Register permission management controllers
	permissionController := NewPermissionController()
	group.Bind(
		permissionController.List,
		permissionController.Detail,
		permissionController.Create,
		permissionController.Update,
		permissionController.Delete,
		permissionController.Check,
	)
}
