package mail_boxes

import (
	"context"

	"billionmail-core/api/mail_boxes/v1"
	"billionmail-core/internal/service/mail_boxes"
)

func (c *ControllerV1) AddMailbox(ctx context.Context, req *v1.AddMailboxReq) (res *v1.AddMailboxRes, err error) {
	res = &v1.AddMailboxRes{}

	mailbox := &v1.Mailbox{
		Username:  req.Username + "@" + req.Domain,
		Password:  req.Password,
		FullName:  req.Username,
		IsAdmin:   req.IsAdmin,
		Quota:     int64(req.Quota),
		LocalPart: req.Username,
		Domain:    req.Domain,
	}

	if err = mail_boxes.Add(ctx, mailbox); err != nil {
		return nil, err
	}

	res.SetSuccess("Mailbox added successfully")
	return
}
