package dockerapi

import (
	"billionmail-core/internal/service/public"
	"context"

	"billionmail-core/api/dockerapi/v1"
)

func (c *ControllerV1) RestartContainer(ctx context.Context, req *v1.RestartContainerReq) (res *v1.RestartContainerRes, err error) {
	res = &v1.RestartContainerRes{}

	err = public.DockerApiFromCtx(ctx).RestartContainer(ctx, req.ContainerID)

	if err != nil {
		res.SetError(err)
		return
	}

	res.SetSuccess("Success")

	return
}
