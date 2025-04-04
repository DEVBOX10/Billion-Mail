package mail_boxes

import (
	"billionmail-core/api/mail_boxes/v1"
	"billionmail-core/internal/service/mail_boxes"
	"context"
)

func (c *ControllerV1) UpdateMailbox(ctx context.Context, req *v1.UpdateMailboxReq) (res *v1.UpdateMailboxRes, err error) {
	res = &v1.UpdateMailboxRes{}

	mailbox := &v1.Mailbox{
		Username:  req.FullName + "@" + req.Domain,
		Password:  req.Password, // If empty, password won't be updated
		FullName:  req.FullName,
		IsAdmin:   req.IsAdmin,
		Quota:     int64(req.Quota),
		LocalPart: req.FullName,
		Domain:    req.Domain,
		Active:    req.Active,
	}

	if err = mail_boxes.Update(ctx, mailbox); err != nil {
		return nil, err
	}

	res.SetSuccess("Mailbox updated successfully")
	return
}
