package contact

import (
	"billionmail-core/internal/service/contact"
	"billionmail-core/internal/service/public"
	"context"

	"billionmail-core/api/contact/v1"
)

func (c *ControllerV1) ListContactsGroups(ctx context.Context, req *v1.ListContactsGroupsReq) (res *v1.ListContactsGroupsRes, err error) {
	res = &v1.ListContactsGroupsRes{}

	groups, err := contact.GetAllGroups(ctx, req.Keyword)
	if err != nil {
		return nil, err
	}

	res.Data.List = groups
	res.SetSuccess(public.LangCtx(ctx, "Get all contact groups successfully"))
	return
}
