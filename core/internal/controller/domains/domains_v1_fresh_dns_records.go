package domains

import (
	"billionmail-core/internal/service/domains"
	"context"

	"billionmail-core/api/domains/v1"
)

func (c *ControllerV1) FreshDNSRecords(ctx context.Context, req *v1.FreshDNSRecordsReq) (res *v1.FreshDNSRecordsRes, err error) {
	res = &v1.FreshDNSRecordsRes{}

	domains.FreshRecords(ctx, req.Domain)

	res.Data = domains.GetRecordsInCache(req.Domain)

	res.SetSuccess("Success")

	return
}
