package domains

import (
	"billionmail-core/internal/model"
	"billionmail-core/internal/service/acme"
	"billionmail-core/internal/service/public"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
)

// ApplyLetsEncryptCertWithHttp applies for a Let's Encrypt certificate for the given domain.
func ApplyLetsEncryptCertWithHttp(ctx context.Context, domain string, accountInfo *model.Account) error {
	certificate, privateKey, err := acme.ApplySSLWithExistingServer(ctx, []string{domain}, accountInfo.Email, "http", "", nil, public.AbsPath("../ssl"))

	if err != nil {
		return err
	}

	dnsNames := "[]"
	notAfter := "0000-00-00"
	notBefore := "0000-00-00"
	endTime := 0
	status := -1
	subject := ""

	if certificate != "" {
		certInfo := acme.GetCertInfo(certificate)
		notAfter = certInfo.NotAfter
		notBefore = certInfo.NotBefore
		subject = certInfo.Subject
		endTime = certInfo.Endtime
		dnsNamesBytes, err := json.Marshal([]string{domain})
		if err == nil {
			dnsNames = string(dnsNamesBytes)
		}
		status = 1
	}

	// 获取申请进度日志
	progress := acme.GetAcmeLogBody(ctx)
	pdata := g.Map{
		"certificate": certificate,
		"private_key": privateKey,
		"error_info":  "",
		"progress":    progress,
		"status":      status,
		"not_after":   notAfter,
		"not_before":  notBefore,
		"dns":         dnsNames,
		"endtime":     endTime,
		"subject":     subject,
	}

	// Save the certificate and key to the database
	_, err = g.DB().Model("letsencrypts").Insert(pdata)

	if err != nil {
		return err
	}

	// TODO: Apply the certificate to the service, postfix, dovecot, etc.

	return nil
}
