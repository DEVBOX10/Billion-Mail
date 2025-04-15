package contact

import (
	"billionmail-core/internal/service/contact"
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/contact/v1"
)

func (c *ControllerV1) ListContacts(ctx context.Context, req *v1.ListContactsReq) (res *v1.ListContactsRes, err error) {
	res = &v1.ListContactsRes{}

	if req.Page < 1 {
		req.Page = DefaultPage
	}
	if req.PageSize < 1 || req.PageSize > MaxPageSize {
		req.PageSize = DefaultPageSize
	}
	if req.Status < -1 || req.Status > 1 {
		req.Status = 1 // Default get all status
	}

	// Get contacts list
	total, list, err := contact.GetContactsWithPage(ctx, req.Page, req.PageSize, req.GroupId, req.Keyword, req.Status)
	if err != nil {
		fmt.Printf("Failed to get contacts list: %v, group ID: %d, keyword: %s, status: %d\n", err, req.GroupId, req.Keyword, req.Status)
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get contacts list: {}", err.Error())))
		return
	}

	// Convert to API response format
	apiList := make([]*v1.Contact, len(list))
	for i, contactOne := range list {
		apiList[i] = &v1.Contact{
			Id:         contactOne.Id,
			Email:      contactOne.Email,
			GroupId:    contactOne.GroupId,
			Active:     contactOne.Active,
			TaskId:     contactOne.TaskId,
			CreateTime: contactOne.CreateTime,
		}
	}

	res.Data.Total = total
	res.Data.List = apiList
	res.SetSuccess(public.LangCtx(ctx, "Get contacts list successfully"))
	return
}
