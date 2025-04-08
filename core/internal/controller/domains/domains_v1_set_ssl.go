package domains

import (
	"billionmail-core/internal/service/domains"
	"billionmail-core/internal/service/mail_service"
	"context"
	"fmt"

	"billionmail-core/api/domains/v1"
)

func (c *ControllerV1) SetSSL(ctx context.Context, req *v1.SetSSLReq) (res *v1.SetSSLRes, err error) {
	res = &v1.SetSSLRes{}

	if ex, err1 := domains.Exists(ctx, req.Domain); !ex {
		if err1 != nil {
			res.SetError(fmt.Errorf("fail to check domain: %w", err1))
			return
		}

		res.SetError(fmt.Errorf("domain %s does not exist", req.Domain))
		return
	}

	crt := mail_service.NewCertificate()

	defer crt.Close()

	err = crt.SetSNI(req.Domain, req.Certificate, req.Key)

	if err != nil {
		res.SetError(fmt.Errorf("fail to set ssl: %w", err))
		return
	}

	res.SetSuccess("Success")
	return
}
