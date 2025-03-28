package rbac

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/rbac/v1"
)

func (c *ControllerV1) PermissionDetail(ctx context.Context, req *v1.PermissionDetailReq) (res *v1.PermissionDetailRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
