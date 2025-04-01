package domains

import (
	"billionmail-core/internal/service/domains"
	"context"

	"billionmail-core/api/domains/v1"
)

func (c *ControllerV1) AddDomain(ctx context.Context, req *v1.AddDomainReq) (res *v1.AddDomainRes, err error) {
	res = &v1.AddDomainRes{}

	domain := &v1.Domain{
		Domain:       req.Domain,
		ARecord:      "mail." + req.Domain,
		Mailboxes:    req.Mailboxes,
		MailboxQuota: int64(req.MailboxQuota),
		Quota:        int64(req.Quota),
		RateLimit:    req.RateLimit,
	}

	if err = domains.Add(ctx, domain); err != nil {
		return nil, err
	}

	res.SetSuccess("Domain added successfully")
	return
}
