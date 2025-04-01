package mail_boxes

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/mail_boxes/v1"
	"billionmail-core/internal/service/mail_boxes"
)

func (c *ControllerV1) UpdateMailbox(ctx context.Context, req *v1.UpdateMailboxReq) (res *v1.UpdateMailboxRes, err error) {
	res = &v1.UpdateMailboxRes{}

	if req.Email == "" {
		return nil, gerror.New("Email address cannot be empty")
	}

	parts := strings.Split(req.Email, "@")
	if len(parts) != 2 {
		return nil, gerror.New("Invalid email format")
	}

	mailbox := &v1.Mailbox{
		Username:  req.Email,
		Password:  req.Password, // If empty, password won't be updated
		FullName:  req.Username,
		IsAdmin:   req.IsAdmin,
		Quota:     int64(req.Quota),
		LocalPart: parts[0],
		Domain:    parts[1],
		Active:    req.Active,
	}

	if err = mail_boxes.Update(ctx, mailbox); err != nil {
		return nil, err
	}

	res.SetSuccess("Mailbox updated successfully")
	return
}
