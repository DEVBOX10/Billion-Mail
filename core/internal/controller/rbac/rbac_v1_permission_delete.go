package rbac

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/rbac/v1"
)

func (c *ControllerV1) PermissionDelete(ctx context.Context, req *v1.PermissionDeleteReq) (res *v1.PermissionDeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
