package contact

import (
	"billionmail-core/internal/service/contact"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"billionmail-core/api/contact/v1"
)

func (c *ControllerV1) UpdateContactsGroup(ctx context.Context, req *v1.UpdateContactsGroupReq) (res *v1.UpdateContactsGroupRes, err error) {
	res = &v1.UpdateContactsGroupRes{}

	// Start transaction
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// Verify if all target groups exist
		for _, groupId := range req.GroupIds {
			if _, err := contact.GetGroup(ctx, groupId); err != nil {
				return gerror.New(public.LangCtx(ctx, "Group {} not found", groupId))
			}
		}

		// Update contact group
		updatedCount, err := contact.UpdateContactsGroups(ctx, req.Emails, req.Status, req.GroupIds)
		if err != nil {
			return err
		}

		res.Data.UpdatedCount = updatedCount
		return nil
	})

	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to update contacts group: {}", err.Error())))
		return
	}

	res.SetSuccess(public.LangCtx(ctx, "Contacts group updated successfully"))
	return
}
