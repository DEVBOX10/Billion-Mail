package mail_boxes

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/mail_boxes/v1"
	"billionmail-core/internal/service/mail_boxes"
)

func (c *ControllerV1) DeleteMailbox(ctx context.Context, req *v1.DeleteMailboxReq) (res *v1.DeleteMailboxRes, err error) {
	res = &v1.DeleteMailboxRes{}

	if req.Email == "" {
		return nil, gerror.New("Email address cannot be empty")
	}

	if err = mail_boxes.Delete(ctx, req.Email); err != nil {
		return nil, err
	}

	res.SetSuccess("Mailbox deleted successfully")
	return
}
