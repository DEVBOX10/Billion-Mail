package rbac

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/rbac/v1"
)

func (c *ControllerV1) PermissionUpdate(ctx context.Context, req *v1.PermissionUpdateReq) (res *v1.PermissionUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
