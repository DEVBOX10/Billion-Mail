// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dockerapi

import (
	"context"

	"billionmail-core/api/dockerapi/v1"
)

type IDockerapiV1 interface {
	ListContainer(ctx context.Context, req *v1.ListContainerReq) (res *v1.ListContainerRes, err error)
	ContainerState(ctx context.Context, req *v1.ContainerStateReq) (res *v1.ContainerStateRes, err error)
	RestartContainer(ctx context.Context, req *v1.RestartContainerReq) (res *v1.RestartContainerRes, err error)
}
