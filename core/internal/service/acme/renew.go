package acme

import (
	"billionmail-core/internal/service/public"
	"context"
	"encoding/json"
	"errors"
	"path/filepath"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

/**
 * @brief Update certificate files
 * @param certificate Certificate content
 * @param privateKey Private key content
 * @param certPath Path to save certificate files
 * @return bool Success status
 */
func UpdateCertFiles(certificate string, privateKey string, certPath string) bool {
	if certificate == "" || privateKey == "" || certPath == "" {
		return false
	}

	// Create directory if it doesn't exist
	if !public.FileExists(certPath) {
		err := public.MkDirAll(certPath, "www")
		if err != nil {
			return false
		}
	}

	certificateFile := filepath.Join(certPath, "certificate.pem")
	privateKeyFile := filepath.Join(certPath, "private_key.pem")

	// Write certificate files
	_, err1 := public.WriteFile(certificateFile, certificate)
	_, err2 := public.WriteFile(privateKeyFile, privateKey)

	return err1 == nil && err2 == nil
}

/**
 * @brief Start certificate renewal
 * @param certId Certificate ID
 * @param accountId Account ID
 * @param certPath Path to save certificate files
 * @param domains Domain list
 * @param email Email address
 * @param vtype Verification type
 * @param dnsProvider DNS provider
 * @param dnsProviderToken DNS provider credentials
 * @return error Error message
 */
func StartRenew(certId int, accountId int, certPath string, domains []string, email string, vtype string, dnsProvider string, dnsProviderToken map[string]string) error {
	// Create renewal log
	g.Log().Info(context.Background(), "Start renew certificate for domains: ", domains)
	renewId, err := public.MR("ssl", "renew_logs").InsertAndGetId(g.Map{
		"cert_id":    certId,
		"account_id": accountId,
		"status":     0,
		"progress":   "start renew",
	})

	// Try to apply for certificate
	ctx := gctx.New()
	certificate, privateKey, err := ApplySSLWithExistingServer(ctx, domains, email, vtype, dnsProvider, dnsProviderToken, certPath)
	progress := GetAcmeLogBody(ctx)
	if err != nil {
		// Update renewal log
		public.M("renew_logs").Where("renew_id = ?", renewId).Update(g.Map{"status": 2, "progress": progress, "error_info": err.Error()})
		g.Log().Error(context.Background(), "Renew certificate failed: ", err.Error())
		return err
	}

	// Certificate is empty
	if certificate == "" || privateKey == "" {
		// Update renewal log
		public.M("renew_logs").Where("renew_id = ?", renewId).Update(g.Map{"status": 2, "progress": progress, "error_info": "let's encrypt certificate is empty"})
		g.Log().Error(context.Background(), "Renew certificate failed: let's encrypt certificate is empty")
		return errors.New("let's encrypt certificate is empty")
	}

	// Update renewal log
	public.M("renew_logs").Where("renew_id = ?", renewId).Update(g.Map{"status": 1, "progress": progress, "error_info": ""})
	g.Log().Info(context.Background(), "Renew certificate success for domains: ", domains)

	dnsNames := "[]"
	status := 1
	certInfo := GetCertInfo(certificate)
	notAfter := certInfo.NotAfter
	notBefore := certInfo.NotBefore
	subject := certInfo.Subject
	endTime := certInfo.Endtime
	dnsNamesBytes, err := json.Marshal(domains)
	if err == nil {
		dnsNames = string(dnsNamesBytes)
	}

	// Get application progress log
	pdata := g.Map{
		"certificate": certificate,
		"private_key": privateKey,
		"progress":    progress,
		"status":      status,
		"not_after":   notAfter,
		"not_before":  notBefore,
		"dns":         dnsNames,
		"endtime":     endTime,
		"subject":     subject,
	}

	// Save certificate to database
	public.M("letsencrypts").Where("cert_id = ?", certId).Update(pdata)
	g.Log().Info(context.Background(), "Update certificate to database for domains: ", domains)

	// Update certificate files
	UpdateCertFiles(certificate, privateKey, certPath)

	return nil
}

/**
 * @brief Renew certificates
 * @return void
 */
func Renew() {
	// Auto-renew certificates that expire within 20 days
	renewTime := public.GetNowTime() + 86400*20
	certList, err := public.M("letsencrypts").Where("endtime < ?", renewTime).Where("status = 1").All()
	if err != nil {
		g.Log().Warning(context.Background(), "Renew certificate failed: ", err.Error())
		return
	}

	renewExpireTime := public.GetNowTime() - 86400

	for _, v := range certList {
		// Skip non-Let's Encrypt certificates
		issuer := v["issuer"].String()
		if issuer != "Let's Encrypt" {
			g.Log().Warning(context.Background(), "Renew certificate skipped: issuer is not Let's Encrypt")
			continue
		}

		certId := v["cert_id"].Int()
		accountId := v["account_id"].Int()
		if certId == 0 || accountId == 0 {
			g.Log().Warning(context.Background(), "Renew certificate failed: cert_id or account_id is empty", certId, accountId)
			continue
		}

		// Get certificate path
		certPath := v["cert_path"].String()
		if certPath == "" {
			// Use default path if not specified
			// Get the subject from certificate info
			subject := v["subject"].String()
			certPath = filepath.Join(public.ROOT_PATH, "ssl", "certs", subject)
		}

		domains := []string{}
		err = json.Unmarshal([]byte(v["dns"].String()), &domains)
		if err != nil {
			g.Log().Warning(context.Background(), "Renew certificate failed: parse domains failed", v["dns"].String())
			continue
		}

		// Get account email
		accountVal, err := public.M("account").Where("account_id = ?", accountId).Fields("email").Value()
		if err != nil {
			g.Log().Warning(context.Background(), "Renew certificate failed: get email failed", accountId)
			continue
		}

		email := accountVal.String()

		// Only renew once within 24 hours
		renewCount, _ := public.MR("ssl", "renew_logs").Where("cert_id = ? AND renew_time > ?", certId, renewExpireTime).Count()
		if renewCount > 0 {
			g.Log().Warning(context.Background(), "Renew certificate skipped: 24 hours only renew once", domains)
			continue
		}

		// Failed 3 times within 7 days
		renewFailCount, _ := public.MR("ssl", "renew_logs").Where("cert_id = ? AND renew_time > ?", certId, renewExpireTime-86400*6).Where("status = 2").Count()
		if renewFailCount > 2 {
			g.Log().Warning(context.Background(), "Renew certificate skipped: 7 days only renew 3 times", domains)
			continue
		}

		vtype := v["auth_type"].String()
		dnsProviderToken := map[string]string{}
		dnsProvider := v["dns_provider"].String()
		if vtype == "dns" {
			// Skip manual DNS verification
			if dnsProvider == "" {
				continue
			}
			// Parse DNS provider token
			err = json.Unmarshal([]byte(v["dns_provider_token"].String()), &dnsProviderToken)
			if err != nil {
				continue
			}
		}

		// Start renewal
		err = StartRenew(certId, accountId, certPath, domains, email, vtype, dnsProvider, dnsProviderToken)
		if err != nil {
			continue
		}
	}
}

/**
 * @brief Create a timer
 * @param CallBack Callback function
 * @param Sec Execution period (seconds)
 * @return *time.Ticker
 */
func NewTimerEvery(CallBack func(), Sec int) *time.Ticker {
	t := time.NewTicker(time.Duration(Sec) * time.Second)
	go func() {
		for {
			<-t.C
			CallBack()
		}
	}()
	return t
}

/**
 * @brief Start worker process
 * @return void
 */
func WorkerStart() {
	NewTimerEvery(Renew, 7200) // Renew every 2 hours
}
