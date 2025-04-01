package domains

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/domains/v1"
	"billionmail-core/internal/service/domains"
)

func (c *ControllerV1) DeleteDomain(ctx context.Context, req *v1.DeleteDomainReq) (res *v1.DeleteDomainRes, err error) {
	res = &v1.DeleteDomainRes{}

	if req.Domain == "" {
		return nil, gerror.New("Domain cannot be empty")
	}

	if err = domains.Delete(ctx, req.Domain); err != nil {
		return nil, err
	}

	res.SetSuccess("Domain deleted successfully")
	return
}
