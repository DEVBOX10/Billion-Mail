package domains

import (
	"billionmail-core/api/domains/v1"
	"billionmail-core/internal/service/mail_service"
	"context"
)

func (c *ControllerV1) GetSSL(ctx context.Context, req *v1.GetSSLReq) (res *v1.GetSSLRes, err error) {
	res = &v1.GetSSLRes{}

	crt := mail_service.NewCertificate()
	defer crt.Close()

	// Get SSL certificate information
	res.Data, _ = crt.GetSSLInfo(req.Domain)

	res.SetSuccess("Success")
	return
}
