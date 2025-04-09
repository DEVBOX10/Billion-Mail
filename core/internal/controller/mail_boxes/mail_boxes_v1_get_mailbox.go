package mail_boxes

import (
	"billionmail-core/internal/service/mail_boxes"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/mail_boxes/v1"
)

func (c *ControllerV1) GetMailbox(ctx context.Context, req *v1.GetMailboxReq) (res *v1.GetMailboxRes, err error) {
	res = &v1.GetMailboxRes{}

	if req.Domain == "" {
		return nil, gerror.New("Domain cannot be empty")
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	mailboxList, total, err := mail_boxes.Get(ctx, req.Domain, req.Keyword, page, pageSize)
	if err != nil {
		return nil, err
	}

	// Handle sensitive information
	for i := range mailboxList {
		mailboxList[i].Password, _ = mail_boxes.PasswdDecode(ctx, mailboxList[i].PasswordEncode)
	}

	res.Data.Total = total
	res.Data.List = mailboxList

	res.SetSuccess("Success")

	return
}
