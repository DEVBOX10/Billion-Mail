// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package overview

import (
	"context"

	"billionmail-core/api/overview/v1"
)

type IOverviewV1 interface {
	Overview(ctx context.Context, req *v1.OverviewReq) (res *v1.OverviewRes, err error)
	FailedList(ctx context.Context, req *v1.FailedListReq) (res *v1.FailedListRes, err error)
}
