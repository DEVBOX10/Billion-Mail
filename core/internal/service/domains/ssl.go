package domains

import (
	"billionmail-core/internal/consts"
	"billionmail-core/internal/model"
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/acme"
	"billionmail-core/internal/service/mail_service"
	"billionmail-core/internal/service/public"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

// ApplyLetsEncryptCertWithHttp applies for a Let's Encrypt certificate for the given domain.
func ApplyLetsEncryptCertWithHttp(ctx context.Context, domain string, accountInfo *model.Account) error {
	// Find the existing certificate for the domain
	if crt, err := FindSSLByDomain(domain); err == nil && crt != nil {
		if crt.Status == 1 && crt.EndTime > time.Now().Unix() {
			return ApplyCertToService(domain, crt.Certificate, crt.PrivateKey)
		}
	}

	certificate, privateKey, err := acme.ApplySSLWithExistingServer(ctx, []string{domain}, accountInfo.Email, "http", "", nil, public.AbsPath(consts.SSL_PATH))

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

	// get the progress of the certificate application
	progress := acme.GetAcmeLogBody(ctx)
	pdata := g.Map{
		"account_id":  accountInfo.AccountId,
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

	// Apply the certificate to the service, postfix, dovecot, etc.
	return ApplyCertToService(domain, certificate, privateKey)
}

// FindSSLByDomain retrieves the SSL certificate information for a given domain.
func FindSSLByDomain(domain string) (crt *entity.Letsencrypt, err error) {
	err = g.DB().Model("letsencrypts").Where("dns::jsonb ? $1", domain).Order("endtime desc").Limit(1).Scan(&crt)
	return
}

// ApplyCertToService applies the SSL certificate to the service.
func ApplyCertToService(domain, crtPem, keyPem string) (err error) {
	crt := mail_service.NewCertificate()

	defer crt.Close()

	err = crt.SetSNI(domain, crtPem, keyPem)

	return
}
