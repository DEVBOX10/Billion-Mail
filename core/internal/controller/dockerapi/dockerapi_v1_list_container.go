package dockerapi

import (
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/dockerapi/v1"
)

func (c *ControllerV1) ListContainer(ctx context.Context, req *v1.ListContainerReq) (res *v1.ListContainerRes, err error) {
	res = &v1.ListContainerRes{}

	res.Data, err = public.DockerApiFromCtx(ctx).ListContainers(ctx)

	if err != nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "获取容器列表失败")
	}

	res.SetSuccess("Success")

	return
}
