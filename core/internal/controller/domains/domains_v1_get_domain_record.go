package domains

import (
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/domains/v1"
	"billionmail-core/internal/service/domains"
)

func (c *ControllerV1) GetDomainRecord(ctx context.Context, req *v1.GetDomainRecordReq) (res *v1.GetDomainRecordRes, err error) {
	res = &v1.GetDomainRecordRes{}

	if req.Domain == "" {
		return nil, gerror.New("Domain cannot be empty")
	}

	// Get server IP address
	ipv4Address, err := getServerIPv4()
	if err != nil {
		ipv4Address = "YOUR_SERVER_IPV4_ADDRESS"
	}

	// Get DKIM record
	dkimRecord, err := getDKIMRecord(req.Domain)
	if err != nil {
		// Log error but continue execution
		fmt.Printf("Failed to get DKIM record: %v\n", err)
	}

	res.Records = []v1.DomainRecord{
		{Type: "A", Host: "mail." + req.Domain, Value: ipv4Address},
		{Type: "MX", Host: "@", Value: "mail." + req.Domain, Priority: 10},
		{Type: "TXT", Host: "@", Value: fmt.Sprintf("v=spf1 +a +mx +ip4:%s -all", ipv4Address)},
		{Type: "TXT", Host: "_dmarc", Value: fmt.Sprintf("v=DMARC1;p=quarantine;rua=mailto:admin@%s", req.Domain)},
	}

	if dkimRecord != "" {
		res.Records = append(res.Records, v1.DomainRecord{
			Type:  "TXT",
			Host:  "default._domainkey",
			Value: dkimRecord,
		})
	}

	return
}

func getServerIPv4() (string, error) {
	resp, err := http.Get("https://ifconfig.me")
	if err != nil {
		// Try alternative service
		resp, err = http.Get("https://www.aapanel.com/api/common/getClientIP")
		if err != nil {
			return "", err
		}
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(ip)), nil
}

func getDKIMRecord(domain string) (string, error) {
	// Create DKIM directory
	dkimPath := public.AbsPath(filepath.Join("../rspamd-data", "dkim", domain))

	// Check if directory exists
	if !public.IsDir(dkimPath) {
		_ = os.MkdirAll(dkimPath, 0644)
	}

	// Check if DKIM private and public key files exist
	dkimPriPath := filepath.Join(dkimPath, "default.private")
	dkimPubPath := filepath.Join(dkimPath, "default.pub")

	// Generate new keys if they don't exist
	if !public.FileExists(dkimPriPath) || !public.FileExists(dkimPubPath) {
		err := exec.Command("docker", "exec", "biilionmail-rspamd-billionmail-1", "rspamadm", fmt.Sprintf("dkim_keygen -s 'default' -b 1024 -d {domain} -k \"/var/lib/rspamd/dkim/%s/default.private\" > \"/var/lib/rspamd/dkim/%s/default.pub", domain, domain)).Run()

		if err != nil {
			return "", fmt.Errorf("Failed to generate DKIM key pair: %v", err)
		}
	}

	// DKIM public key is typically stored in a specific location in the container or host
	// Assuming we use docker exec to read the DKIM public key from the rspamd container
	dkimPub, err := public.ReadFile(dkimPubPath)

	if err != nil {
		return "", fmt.Errorf("Cannot read DKIM public key: %v", err)
	}

	// Format DKIM record
	// Expected format is a pre-formatted TXT record value like "v=DKIM1; k=rsa; p=MIIBIjANBg..."
	dkimRecord := strings.TrimSpace(dkimPub)

	// If the raw public key is read, format it into DNS TXT record format
	if !strings.HasPrefix(dkimRecord, "v=DKIM1") {
		// Remove possible header/footer markers and newlines
		dkimRecord = strings.ReplaceAll(dkimRecord, "-----BEGIN PUBLIC KEY-----", "")
		dkimRecord = strings.ReplaceAll(dkimRecord, "-----END PUBLIC KEY-----", "")
		dkimRecord = strings.ReplaceAll(dkimRecord, "\n", "")
		dkimRecord = strings.TrimSpace(dkimRecord)
		dkimRecord = fmt.Sprintf("v=DKIM1; k=rsa; p=%s", dkimRecord)
	}

	return dkimRecord, nil
}
