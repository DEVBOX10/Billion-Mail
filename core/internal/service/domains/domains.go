package domains

import (
	v2 "billionmail-core/api/dockerapi/v1"
	v1 "billionmail-core/api/domains/v1"
	"billionmail-core/internal/consts"
	docker "billionmail-core/internal/service/dockerapi"
	"billionmail-core/internal/service/mail_service"
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	mutex sync.Mutex
)

func Add(ctx context.Context, domain *v1.Domain) error {
	domain.CreateTime = time.Now().Unix()
	domain.Active = 1

	_, err := g.DB().Model("domain").Ctx(ctx).Insert(domain)
	return err
}

func Update(ctx context.Context, domain *v1.Domain) error {
	_, err := g.DB().Model("domain").
		Ctx(ctx).
		Where("domain", domain.Domain).
		Update(domain)
	return err
}

func Delete(ctx context.Context, domainName string) error {
	_, err := g.DB().Model("domain").
		Ctx(ctx).
		Where("domain", domainName).
		Delete()
	return err
}

func Get(ctx context.Context, keyword string, page, pageSize int) ([]v1.Domain, int, error) {
	m := g.DB().Model("domain").Order("create_time", "desc")

	if keyword != "" {
		m = m.WhereLike("domain", fmt.Sprintf("%%%s%%", keyword))
	}

	count, err := m.Count()
	if err != nil {
		return nil, 0, err
	}

	var domains []v1.Domain
	err = m.Page(page, pageSize).Scan(&domains)

	if err != nil {
		return nil, 0, err
	}

	crt := mail_service.NewCertificate()

	defer crt.Close()

	wg := sync.WaitGroup{}

	for i, domain := range domains {
		wg.Add(1)
		go func(i int, domain v1.Domain) {
			defer wg.Done()
			domains[i].DNSRecords = GetRecordsInCache(domain.Domain)
		}(i, domain)

		// Retrieve Domains SSL certificate information
		wg.Add(1)
		go func(i int, domain v1.Domain) {
			defer wg.Done()
			domains[i].CertInfo, _ = crt.GetSSLInfo(domain.Domain)
		}(i, domain)
	}

	wg.Wait()

	return domains, count, err
}

func All(ctx context.Context) ([]v1.Domain, error) {
	m := g.DB().Model("domain")

	var domains []v1.Domain
	err := m.Scan(&domains)

	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}

	for i, domain := range domains {
		wg.Add(1)
		go func(i int, domain v1.Domain) {
			defer wg.Done()
			domains[i].DNSRecords.A, _ = GetARecord(domain.Domain)
		}(i, domain)

		wg.Add(1)
		go func(i int, domain v1.Domain) {
			defer wg.Done()
			domains[i].DNSRecords.MX, _ = GetMXRecord(domain.Domain)
		}(i, domain)

		wg.Add(1)
		go func(i int, domain v1.Domain) {
			defer wg.Done()
			domains[i].DNSRecords.SPF, _ = GetSPFRecord(domain.Domain)
		}(i, domain)

		wg.Add(1)
		go func(i int, domain v1.Domain) {
			defer wg.Done()
			domains[i].DNSRecords.DKIM, _ = GetDKIMRecord(domain.Domain)
		}(i, domain)

		wg.Add(1)
		go func(i int, domain v1.Domain) {
			defer wg.Done()
			domains[i].DNSRecords.DMARC, _ = GetDMARCRecord(domain.Domain)
		}(i, domain)
	}

	wg.Wait()

	return domains, err
}

func Exists(ctx context.Context, domainName string) (bool, error) {
	count, err := g.DB().Model("domain").
		Ctx(ctx).
		Where("domain", domainName).
		Count()

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func buildCacheKey(domain, recordType string) string {
	return fmt.Sprintf("DOMAIN_DNS_RECORDS_:%s:_%s", domain, recordType)
}

func FreshRecords(ctx context.Context) {
	g.Log().Debug(ctx, "Fresh DNS records...")

	domains, err := All(ctx)

	if err != nil {
		return
	}

	cacheSeconds := 600

	wg := sync.WaitGroup{}

	for _, domain := range domains {
		// goroutine for each record type
		wg.Add(6) // total of 6 record types

		// retrieve A record
		go func(domain v1.Domain) {
			defer wg.Done()
			dr, err := GetARecord(domain.Domain)
			if err != nil {
				g.Log().Error(ctx, "Failed to get A record for domain %s: %v", domain.Domain, err)
			} else {
				public.SetCache(buildCacheKey(domain.Domain, "A"), dr, cacheSeconds)
			}
		}(domain)

		// retrieve MX record
		go func(domain v1.Domain) {
			defer wg.Done()
			dr, err := GetMXRecord(domain.Domain)
			if err != nil {
				g.Log().Error(ctx, "Failed to get MX record for domain %s: %v", domain.Domain, err)
			} else {
				public.SetCache(buildCacheKey(domain.Domain, "MX"), dr, cacheSeconds)
			}
		}(domain)

		// retrieve SPF record
		go func(domain v1.Domain) {
			defer wg.Done()
			dr, err := GetSPFRecord(domain.Domain)
			if err != nil {
				g.Log().Error(ctx, "Failed to get SPF record for domain %s: %v", domain.Domain, err)
			} else {
				public.SetCache(buildCacheKey(domain.Domain, "SPF"), dr, cacheSeconds)
			}
		}(domain)

		// retrieve DKIM record
		go func(domain v1.Domain) {
			defer wg.Done()
			dr, err := GetDKIMRecord(domain.Domain)
			if err != nil {
				g.Log().Error(ctx, "Failed to get DKIM record for domain %s: %v", domain.Domain, err)
			} else {
				public.SetCache(buildCacheKey(domain.Domain, "DKIM"), dr, cacheSeconds)
			}
		}(domain)

		// retrieve DMARC record
		go func(domain v1.Domain) {
			defer wg.Done()
			dr, err := GetDMARCRecord(domain.Domain)
			if err != nil {
				g.Log().Error(ctx, "Failed to get DMARC record for domain %s: %v", domain.Domain, err)
			} else {
				public.SetCache(buildCacheKey(domain.Domain, "DMARC"), dr, cacheSeconds)
			}
		}(domain)

		//retrieve PTR record
		go func(domain v1.Domain) {
			defer wg.Done()
			dr, err := GetPTRRecord(domain.Domain)
			if err != nil {
				g.Log().Error(ctx, "Failed to get PTR record for domain %s: %v", domain.Domain, err)
			} else {
				public.SetCache(buildCacheKey(domain.Domain, "PTR"), dr, cacheSeconds)
			}
		}(domain)
	}

	wg.Wait()

	g.Log().Debug(ctx, "Fresh DNS records completed.")
}

// GetDKIMRecord retrieves the DKIM record for a given domain.
func GetDKIMRecord(domain string) (record v1.DNSRecord, err error) {
	// Create DKIM directory
	dkimPath := public.AbsPath(filepath.Join(consts.RSPAMD_LIB_PATH, "dkim", domain))

	// Check if directory exists
	if !public.IsDir(dkimPath) {
		_ = os.MkdirAll(dkimPath, 0644)
	}

	// Check if DKIM private and public key files exist
	dkimPriPath := filepath.Join(dkimPath, "default.private")
	dkimPubPath := filepath.Join(dkimPath, "default.pub")

	var dk *docker.DockerAPI

	dk, err = docker.NewDockerAPI()

	if err != nil {
		err = fmt.Errorf("Failed to connect to Docker API: %v", err)
		return
	}

	defer dk.Close()

	// Generate new keys if they don't exist
	if !public.FileExists(dkimPriPath) || !public.FileExists(dkimPubPath) {
		mutex.Lock()
		defer mutex.Unlock()

		var res *v2.ExecResult
		res, err = dk.ExecCommandByName(context.Background(), "billionmail-rspamd-billionmail-1", []string{"rspamadm", "dkim_keygen", "-s", "'default'", "-b", "1024", "-d", domain, "-k", fmt.Sprintf("/var/lib/rspamd/dkim/%s/default.private", domain)}, "root")

		if err != nil {
			err = fmt.Errorf("Failed to generate DKIM key pair: %v", err)
			return
		}

		if res != nil {
			_, err = public.WriteFile(dkimPubPath, res.Output)

			if err != nil {
				err = fmt.Errorf("Failed to write DKIM public key: %v", err)
				return
			}
		}

		// build DKIM Sign config
		signConf := fmt.Sprintf(`#%s_DKIM_BEGIN
%s {
   selectors [
    {
      path: "%s";
      selector: "default"
    }
  ]
}
#%s_DKIM_END
`, domain, domain, dkimPriPath, domain)

		// Write DKIM sign config to file
		signConfPath := public.AbsPath(filepath.Join(consts.RSPAMD_LOCAL_D_PATH, "dkim_signing.conf"))
		signContent := `domain {
#BT_DOMAIN_DKIM_BEGIN
#BT_DOMAIN_DKIM_END
}`

		if public.FileExists(signConfPath) {
			signContent, err = public.ReadFile(signConfPath)

			if err != nil {
				err = fmt.Errorf("Failed to read DKIM sign config: %v", err)
				return
			}
		}

		signContent = strings.Replace(signContent, "#BT_DOMAIN_DKIM_END", signConf+"\n#BT_DOMAIN_DKIM_END", 1)

		_, err = public.WriteFile(signConfPath, signContent)

		if err != nil {
			err = fmt.Errorf("Failed to write DKIM sign config: %v", err)
			return
		}

		// Restart rspamd service
		err = dk.RestartContainerByName(context.Background(), "billionmail-rspamd-billionmail-1")

		if err != nil {
			err = fmt.Errorf("Failed to restart rspamd container: %v", err)
			return
		}
	}

	// DKIM public key is typically stored in a specific location in the container or host
	// Assuming we use docker exec to read the DKIM public key from the rspamd container
	dkimPub, err := public.ReadFile(dkimPubPath)

	if err != nil {
		err = fmt.Errorf("Cannot read DKIM public key: %v", err)
		return
	}

	// Format DKIM record
	// Expected format is a pre-formatted TXT record value like "v=DKIM1; k=rsa; p=MIIBIjANBg..."
	dkimRecord := strings.TrimSpace(dkimPub)

	// If the raw public key is read, format it into DNS TXT record format
	if !strings.Contains(dkimRecord, "v=DKIM1") && !strings.Contains(dkimRecord, "k=rsa") && !strings.Contains(dkimRecord, "p=") {
		// Remove possible header/footer markers and newlines
		dkimRecord = strings.ReplaceAll(dkimRecord, "-----BEGIN PUBLIC KEY-----", "")
		dkimRecord = strings.ReplaceAll(dkimRecord, "-----END PUBLIC KEY-----", "")
		dkimRecord = strings.ReplaceAll(dkimRecord, "\n", "")
		dkimRecord = strings.TrimSpace(dkimRecord)
		dkimRecord = fmt.Sprintf("v=DKIM1; k=rsa; p=%s", dkimRecord)
	} else {
		seps := strings.Split(strings.ReplaceAll(strings.ReplaceAll(dkimRecord, " ", ""), "\n", ""), "\"")
		if len(seps) < 4 {
			err = fmt.Errorf("Invalid DKIM record format")
			return
		}

		dkimRecord = seps[1] + seps[3]
	}

	record = v1.DNSRecord{
		Type:  "TXT",
		Host:  "default._domainkey",
		Value: dkimRecord,
	}

	// Validate the DKIM record
	record.Valid = ValidateTXTRecord(record)

	return
}

// GetDMARCRecord retrieves the DMARC record for a given domain.
func GetDMARCRecord(domain string) (record v1.DNSRecord, err error) {
	record = v1.DNSRecord{
		Type:  "TXT",
		Host:  "_dmarc",
		Value: fmt.Sprintf("v=DMARC1;p=quarantine;rua=mailto:admin@%s", domain),
	}

	// Validate the DMARC record
	record.Valid = ValidateTXTRecord(record)

	return
}

// GetSPFRecord retrieves the SPF record for a given domain.
func GetSPFRecord(domain string) (record v1.DNSRecord, err error) {
	serverIP, err := public.GetServerIP()

	if err != nil {
		err = fmt.Errorf("Failed to get server IP: %v", err)
		return
	}

	ipType := "ip4"

	if strings.Contains(serverIP, ":") {
		ipType = "ip6"
	}

	// Format SPF record
	record = v1.DNSRecord{
		Type:  "TXT",
		Host:  "@",
		Value: fmt.Sprintf("v=spf1 +a +mx +%s:%s -all", ipType, serverIP),
	}

	// Validate the SPF record
	record.Valid = ValidateTXTRecord(record)

	return
}

// GetMXRecord retrieves the MX record for a given domain.
func GetMXRecord(domain string) (record v1.DNSRecord, err error) {
	record = v1.DNSRecord{
		Type:  "MX",
		Host:  "@",
		Value: "mail." + domain,
	}

	// Validate the MX record
	record.Valid = ValidateMXRecord(record, domain)

	return
}

// GetARecord retrieves the A record for a given domain.
func GetARecord(domain string) (record v1.DNSRecord, err error) {
	serverIP, err := public.GetServerIP()

	if err != nil {
		err = fmt.Errorf("Failed to get server IP: %v", err)
		return
	}

	recordType := "A"

	if strings.Contains(serverIP, ":") {
		recordType = "AAAA"
	}

	record = v1.DNSRecord{
		Type:  recordType,
		Host:  "mail." + domain,
		Value: serverIP,
	}

	// Validate the A record
	record.Valid = ValidateARecord(record)

	return
}

// GetPTRRecord retrieves the PTR record for a given domain.
func GetPTRRecord(domain string) (record v1.DNSRecord, err error) {
	serverIP, err := public.GetServerIP()

	if err != nil {
		err = fmt.Errorf("Failed to get server IP: %v", err)
		return
	}

	record = v1.DNSRecord{
		Type:  "PTR",
		Host:  serverIP,
		Value: domain,
	}

	// Validate the PTR record
	// record.Valid = ValidatePTRRecord(record)

	return
}

func GetRecordsInCache(domain string) (records v1.DNSRecords) {
	// Get A record from cache
	aRecord := public.GetCache(buildCacheKey(domain, "A"))

	if aRecord != nil {
		if v, ok := aRecord.(v1.DNSRecord); ok {
			records.A = v
		}
	}

	// Get MX record from cache
	mxRecord := public.GetCache(buildCacheKey(domain, "MX"))

	if mxRecord != nil {
		if v, ok := mxRecord.(v1.DNSRecord); ok {
			records.MX = v
		}
	}

	// Get SPF record from cache
	spfRecord := public.GetCache(buildCacheKey(domain, "SPF"))

	if spfRecord != nil {
		if v, ok := spfRecord.(v1.DNSRecord); ok {
			records.SPF = v
		}
	}

	// Get DKIM record from cache
	dkimRecord := public.GetCache(buildCacheKey(domain, "DKIM"))

	if dkimRecord != nil {
		if v, ok := dkimRecord.(v1.DNSRecord); ok {
			records.DKIM = v
		}
	}

	// Get DMARC record from cache
	dmarcRecord := public.GetCache(buildCacheKey(domain, "DMARC"))

	if dmarcRecord != nil {
		if v, ok := dmarcRecord.(v1.DNSRecord); ok {
			records.DMARC = v
		}
	}

	// Get PTR record from cache
	ptrRecord := public.GetCache(buildCacheKey(domain, "PTR"))

	if ptrRecord != nil {
		if v, ok := ptrRecord.(v1.DNSRecord); ok {
			records.PTR = v
		}
	}

	return
}
