package domains

import (
	"billionmail-core/internal/service/mail_service"
	"context"
	"fmt"

	"billionmail-core/api/domains/v1"
)

func (c *ControllerV1) GetSSL(ctx context.Context, req *v1.GetSSLReq) (res *v1.GetSSLRes, err error) {
	res = &v1.GetSSLRes{}

	crt := mail_service.NewCertificate()
	defer crt.Close()

	// Get SSL certificate information
	res.Data, err = crt.GetSSLInfo(req.Domain)

	if err != nil {
		res.SetError(fmt.Errorf("error getting certificate: %v", err))
		return
	}

	res.SetSuccess("Success")
	return
}
