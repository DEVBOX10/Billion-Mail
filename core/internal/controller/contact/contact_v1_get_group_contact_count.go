package contact

import (
	"billionmail-core/api/contact/v1"
	"billionmail-core/internal/service/contact"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

// GetGroupContactCount handles the request to calculate the total number of active contacts in multiple groups
func (c *ControllerV1) GetGroupContactCount(ctx context.Context, req *v1.GetGroupContactCountReq) (res *v1.GetGroupContactCountRes, err error) {
	res = &v1.GetGroupContactCountRes{}

	// Validate input
	if len(req.GroupIds) == 0 {
		res.Code = 400
		res.SetError(gerror.New(public.LangCtx(ctx, "Group IDs cannot be empty")))
		return
	}

	// Call service function to get contact count
	total, err := contact.GetGroupContactCount(ctx, req.GroupIds)
	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get contact count: {}", err.Error())))
		return
	}

	// Set response data
	res.Data.Total = total
	res.SetSuccess(public.LangCtx(ctx, "Successfully retrieved contact count"))
	return
}
