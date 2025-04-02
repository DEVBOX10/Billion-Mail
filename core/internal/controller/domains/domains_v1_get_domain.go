package domains

import (
	"billionmail-core/internal/service/domains"
	"context"

	"billionmail-core/api/domains/v1"
)

func (c *ControllerV1) GetDomain(ctx context.Context, req *v1.GetDomainReq) (res *v1.GetDomainRes, err error) {
	res = &v1.GetDomainRes{}

	page := req.Page
	if page <= 0 {
		page = 1
	}

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	domainList, total, err := domains.Get(ctx, req.Keyword, page, pageSize)
	if err != nil {
		return nil, err
	}

	res.Data.Total = total
	res.Data.List = domainList

	res.SetSuccess("Success")

	return
}
