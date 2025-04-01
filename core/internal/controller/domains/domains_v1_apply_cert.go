package domains

import (
	"billionmail-core/internal/service/domains"
	"billionmail-core/internal/service/rbac"
	"context"

	"billionmail-core/api/domains/v1"
)

func (c *ControllerV1) ApplyCert(ctx context.Context, req *v1.ApplyCertReq) (res *v1.ApplyCertRes, err error) {
	res = &v1.ApplyCertRes{}

	acc, err := rbac.GetCurrentAccount(ctx)

	if err != nil {
		res.SetError(err)
		return
	}

	// Apply for the certificate
	err = domains.ApplyLetsEncryptCertWithHttp(ctx, req.Domain, acc.Email)

	if err != nil {
		res.SetError(err)
		return
	}

	res.SetSuccess("Certificate applied successfully")
	return
}
