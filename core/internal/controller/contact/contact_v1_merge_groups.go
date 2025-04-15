// core/internal/controller/contact/contact_v1_merge_groups.go
package contact

import (
	"billionmail-core/api/contact/v1"
	"billionmail-core/internal/service/contact"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) MergeGroups(ctx context.Context, req *v1.MergeGroupsReq) (res *v1.MergeGroupsRes, err error) {
	res = &v1.MergeGroupsRes{}

	// Get target group
	targetGroup, err := contact.GetGroup(ctx, req.TargetGroupId)
	if err != nil {
		res.Code = 400
		res.SetError(gerror.New(public.LangCtx(ctx, "Target group not found")))
		return
	}

	// Merge contacts from each source group
	for _, sourceGroupId := range req.SourceGroupIds {
		// Get contacts from source group
		contacts, err := contact.GetContactsByGroup(ctx, sourceGroupId)
		if err != nil {
			continue
		}

		// Update contact group ID
		for _, email := range contacts {
			email.GroupId = targetGroup.Id
		}

		// Batch update contacts
		err = contact.BatchCreateContacts(ctx, contacts)
		if err != nil {
			res.Data.Duplicate += len(contacts)
			continue
		}

		res.Data.TotalMerged += len(contacts)
	}
	return
}
