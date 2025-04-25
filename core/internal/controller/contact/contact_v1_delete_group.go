package contact

import (
	"billionmail-core/api/contact/v1"
	"billionmail-core/internal/service/contact"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) DeleteGroup(ctx context.Context, req *v1.DeleteGroupReq) (res *v1.DeleteGroupRes, err error) {
	res = &v1.DeleteGroupRes{}

	for _, groupId := range req.GroupIds {

		err = contact.DeleteContactsByGroupId(ctx, groupId)
		if err != nil {
			continue
		}

		err = contact.DeleteGroup(ctx, groupId)
		if err != nil {
			res.Data.FailedCount++
			continue
		}
		res.Data.SuccessCount++
	}
	if res.Data.SuccessCount == 0 && res.Data.FailedCount == 0 {
		res.Code = 400
		res.SetError(gerror.New(public.LangCtx(ctx, "No groups were deleted")))
		return
	}
	res.SetSuccess(public.LangCtx(ctx, "Groups deleted successfully"))

	return
}
