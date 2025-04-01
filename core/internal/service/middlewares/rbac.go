package middlewares

import (
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"

	"billionmail-core/internal/service/rbac"
)

// PathToRouteInfo converts path to module, action, and resource
func PathToRouteInfo(path string) (module, action, resource string) {
	// Extract module
	modules := []string{"account", "role", "permission"}
	for _, m := range modules {
		if strings.Contains(path, "/"+m+"/") || strings.HasSuffix(path, "/"+m) {
			module = m
			break
		}
	}

	// Extract action and resource
	pattern := `/api/(\w+)/(\w+)(?:/.*)?`
	match, err := gregex.MatchString(pattern, path)
	if err == nil && len(match) >= 3 {
		resource = match[1]
		actionName := match[2]

		// Map HTTP method to CRUD action if action is a standard CRUD
		switch actionName {
		case "list", "detail":
			action = "read"
		case "create":
			action = "create"
		case "update":
			action = "update"
		case "delete":
			action = "delete"
		default:
			action = actionName
		}
	}

	return
}

// RBACMiddleware handles permission verification for HTTP requests
type RBACMiddleware struct {
	PermissionService rbac.IPermission
}

// NewRBACMiddleware creates a new RBACMiddleware
func NewRBACMiddleware() *RBACMiddleware {
	return &RBACMiddleware{
		PermissionService: rbac.Permission(),
	}
}

// PermissionCheck checks if the current user has the required permission
func (m *RBACMiddleware) PermissionCheck(r *ghttp.Request) {
	// Skip permission check for authentication-related routes
	if r.URL.Path == "/api/v1/login" ||
		r.URL.Path == "/api/v1/refresh-token" {
		r.Middleware.Next()
		return
	}

	// Extract account ID from context
	accountIdVar := r.GetCtxVar("accountId")
	if accountIdVar == nil {
		r.Response.WriteJson(public.CodeMap[401])
		r.Exit()
		return
	}
	accountId := gconv.Int64(accountIdVar)

	// Get roles from context
	roles := r.GetCtxVar("roles", []string{}).Strings()

	// Check for admin role (has all permissions)
	for _, role := range roles {
		if role == "admin" {
			r.Middleware.Next()
			return
		}
	}

	// Extract module, action, and resource from request path
	module, action, resource := PathToRouteInfo(r.URL.Path)

	// If we couldn't determine the module, action, or resource, log it and allow the request
	if module == "" || action == "" || resource == "" {
		g.Log().Warning(context.Background(),
			fmt.Sprintf("Could not determine permission components for path: %s, allowing access", r.URL.Path))
		r.Middleware.Next()
		return
	}

	// Check if user has the required permission
	hasPermission, err := m.PermissionService.Check(r.GetCtx(), accountId, module, action, resource)
	if err != nil {
		g.Log().Error(r.GetCtx(), "Permission check error:", err)
		r.Response.WriteJson(g.Map{
			"code": 500,
			"msg":  "Error checking permissions",
		})
		r.Exit()
		return
	}

	if !hasPermission {
		r.Response.WriteJson(g.Map{
			"code": 403,
			"msg":  "Insufficient permissions",
		})
		r.Exit()
		return
	}

	r.Middleware.Next()
}

// HasPermission checks if the current user has a specific permission
func HasPermission(ctx context.Context, module, action, resource string) bool {
	accountId := rbac.GetCurrentAccountId(ctx)
	if accountId == 0 {
		return false
	}

	// Get roles from context
	rolesVar := ctx.Value("roles")
	roles := []string{}
	if rolesVar != nil {
		roles = rolesVar.([]string)
	}

	// Check for admin role (has all permissions)
	for _, role := range roles {
		if role == "admin" {
			return true
		}
	}

	// Check specific permission
	permissionService := rbac.Permission()
	hasPermission, err := permissionService.Check(ctx, accountId, module, action, resource)
	if err != nil {
		g.Log().Error(ctx, "Permission check error:", err)
		return false
	}

	return hasPermission
}

// RequirePermission returns a middleware handler that checks for a specific permission
// RequirePermission middleware checks if user has required permission
func RequirePermission(module, action, resource string) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		if !HasPermission(r.GetCtx(), module, action, resource) {
			r.Response.WriteJson(g.Map{
				"code": 403,
				"msg":  "Insufficient permissions",
			})
			r.Exit()
			return
		}
		r.Middleware.Next()
	}
}
