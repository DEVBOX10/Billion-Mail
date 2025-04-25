package contact

import (
	"billionmail-core/api/contact/v1"
	"billionmail-core/internal/service/contact"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) ListGroups(ctx context.Context, req *v1.ListGroupsReq) (res *v1.ListGroupsRes, err error) {

	res = &v1.ListGroupsRes{}
	// Handle request parameters
	page := req.Page
	if page <= 0 {
		page = DefaultPage
	}

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}

	// Get group list
	groups, total, err := contact.ContactsGroupWithPage(ctx, page, pageSize, req.Keyword)
	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get group list {}", err.Error())))
		return
	}

	// Convert group information
	groupInfos := make([]*v1.ContactGroupInfo, 0, len(groups))
	for _, group := range groups {
		// Get statistics information
		activeCount, err := contact.CountContactsByGroup(ctx, group.Id, 1)
		if err != nil {
			activeCount = 0
		}

		unsubscribeCount, err := contact.CountContactsByGroup(ctx, group.Id, 0)
		if err != nil {
			unsubscribeCount = 0
		}

		// Convert to group information
		groupInfo := &v1.ContactGroupInfo{
			Id:               group.Id,
			Name:             group.Name,
			Description:      group.Description,
			CreateTime:       group.CreateTime,
			UpdateTime:       group.UpdateTime,
			TotalCount:       activeCount + unsubscribeCount,
			ActiveCount:      activeCount,
			UnsubscribeCount: unsubscribeCount,
		}
		groupInfos = append(groupInfos, groupInfo)
	}

	res.Data.Total = total
	res.Data.List = groupInfos
	res.SetSuccess(public.LangCtx(ctx, "Group list retrieved successfully"))
	return

}
