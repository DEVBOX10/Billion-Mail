package acme

import (
	"billionmail-core/internal/service/public"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/**
 * @brief Command line interface for ACME certificate management
 * @struct AcmeCLI
 */
type AcmeCLI struct {
	// General settings
	Email      string   `json:"email"`       // Email address for registration
	Domains    []string `json:"domains"`     // Domain list
	OutputPath string   `json:"output_path"` // Path to save certificate files

	// Verification settings
	VerifyType  string            `json:"verify_type"`  // Verification type: "http" or "dns"
	DnsProvider string            `json:"dns_provider"` // DNS provider: tencentcloud/aliyun/cloudxns/azuredns/cloudflare/godaddy
	DnsConfig   map[string]string `json:"dns_config"`   // DNS provider credentials

	// Advanced options
	UseOwnServer bool `json:"use_own_server"` // Whether to use own HTTP server for verification
}

/**
 * @brief Create a new AcmeCLI instance
 * @param email Email address
 * @param domains Domain list
 * @param outputPath Path to save certificate files
 * @return *AcmeCLI
 */
func NewAcmeCLI(email string, domains []string, outputPath string) *AcmeCLI {
	if outputPath == "" {
		// Default output path
		outputPath = filepath.Join(public.ROOT_PATH, "ssl", "certs", domains[0])
	}

	return &AcmeCLI{
		Email:        email,
		Domains:      domains,
		OutputPath:   outputPath,
		VerifyType:   "http", // Default verification type
		UseOwnServer: true,   // Default to using our own HTTP server
	}
}

/**
 * @brief Set DNS verification
 * @param provider DNS provider
 * @param config DNS provider credentials
 * @return *AcmeCLI
 */
func (cli *AcmeCLI) SetDNSVerification(provider string, config map[string]string) *AcmeCLI {
	cli.VerifyType = "dns"
	cli.DnsProvider = provider
	cli.DnsConfig = config
	return cli
}

/**
 * @brief Set HTTP verification
 * @return *AcmeCLI
 */
func (cli *AcmeCLI) SetHTTPVerification() *AcmeCLI {
	cli.VerifyType = "http"
	cli.DnsProvider = ""
	cli.DnsConfig = nil
	return cli
}

/**
 * @brief Validate parameters
 * @return error
 */
func (cli *AcmeCLI) Validate() error {
	// Check email
	if cli.Email == "" {
		return fmt.Errorf("email is required")
	}

	// Check domains
	if len(cli.Domains) == 0 {
		return fmt.Errorf("at least one domain is required")
	}

	// Check verification type
	if cli.VerifyType != "http" && cli.VerifyType != "dns" {
		return fmt.Errorf("verification type must be either 'http' or 'dns'")
	}

	// Check DNS provider if using DNS verification
	if cli.VerifyType == "dns" {
		if cli.DnsProvider == "" {
			return fmt.Errorf("DNS provider is required for DNS verification")
		}

		// Check if DNS provider is supported
		supportedProviders := []string{"tencentcloud", "alidns", "cloudxns", "azuredns", "cloudflare", "godaddy"}
		isSupported := false
		for _, provider := range supportedProviders {
			if cli.DnsProvider == provider {
				isSupported = true
				break
			}
		}

		if !isSupported {
			return fmt.Errorf("unsupported DNS provider: %s, supported providers: %s",
				cli.DnsProvider, strings.Join(supportedProviders, ", "))
		}

		// Check DNS config
		if cli.DnsConfig == nil || len(cli.DnsConfig) == 0 {
			return fmt.Errorf("DNS provider credentials are required for DNS verification")
		}
	}

	return nil
}

/**
 * @brief Apply for certificate
 * @return certificatePath, privateKeyPath, error
 */
func (cli *AcmeCLI) Apply() (string, string, error) {
	// Validate parameters
	if err := cli.Validate(); err != nil {
		return "", "", err
	}

	// Create context
	ctx := context.Background()

	// Apply for certificate
	fmt.Printf("Applying for certificate for domains: %s\n", strings.Join(cli.Domains, ", "))
	fmt.Printf("Using verification type: %s\n", cli.VerifyType)
	if cli.VerifyType == "dns" {
		fmt.Printf("Using DNS provider: %s\n", cli.DnsProvider)
	} else if cli.VerifyType == "http" {
		if cli.UseOwnServer {
			fmt.Println("Using built-in HTTP server for verification")
		} else {
			fmt.Println("Using existing HTTP server for verification")
		}
	}

	var certificate, privateKey string
	var err error

	// Choose the appropriate method based on the UseOwnServer flag
	if cli.UseOwnServer || cli.VerifyType == "dns" {
		// Use the standard method that starts its own HTTP server if needed
		certificate, privateKey, err = ApplySSLWithExistingServer(
			ctx,
			cli.Domains,
			cli.Email,
			cli.VerifyType,
			cli.DnsProvider,
			cli.DnsConfig,
			cli.OutputPath,
		)
	} else {
		// Use the method that assumes an existing HTTP server
		certificate, privateKey, err = ApplySSLWithExistingServer(
			ctx,
			cli.Domains,
			cli.Email,
			cli.VerifyType,
			cli.DnsProvider,
			cli.DnsConfig,
			cli.OutputPath,
		)
	}

	if err != nil {
		return "", "", err
	}

	// Get certificate info
	certInfo := GetCertInfo(certificate)
	fmt.Printf("Certificate issued successfully:\n")
	fmt.Printf("  Subject: %s\n", certInfo.Subject)
	fmt.Printf("  Issuer: %s\n", certInfo.Issuer)
	fmt.Printf("  Valid from: %s\n", certInfo.NotBefore)
	fmt.Printf("  Valid to: %s\n", certInfo.NotAfter)
	fmt.Printf("  Domains: %s\n", strings.Join(certInfo.DNSNames, ", "))

	// Certificate files
	certificatePath := filepath.Join(cli.OutputPath, "certificate.pem")
	privateKeyPath := filepath.Join(cli.OutputPath, "private_key.pem")

	// Save certificate and private key to files
	_, err = public.WriteFile(certificatePath, certificate)

	if err != nil {
		return "", "", fmt.Errorf("failed to save certificate: %v", err)
	}

	_, err = public.WriteFile(privateKeyPath, privateKey)

	if err != nil {
		return "", "", fmt.Errorf("failed to save private key: %v", err)
	}

	fmt.Printf("Certificate saved to: %s\n", certificatePath)
	fmt.Printf("Private key saved to: %s\n", privateKeyPath)

	return certificatePath, privateKeyPath, nil
}

/**
 * @brief Save certificate to database
 * @return error
 */
func (cli *AcmeCLI) SaveToDatabase(accountId int, certificate, privateKey string) (int, error) {
	// Get certificate info
	certInfo := GetCertInfo(certificate)
	if certInfo.Subject == "" {
		return 0, fmt.Errorf("invalid certificate")
	}

	// Prepare DNS names
	dnsNames, _ := json.Marshal(cli.Domains)

	// Prepare data
	data := map[string]interface{}{
		"account_id":   accountId,
		"certificate":  certificate,
		"private_key":  privateKey,
		"subject":      certInfo.Subject,
		"dns":          string(dnsNames),
		"not_before":   certInfo.NotBefore,
		"not_after":    certInfo.NotAfter,
		"endtime":      certInfo.Endtime,
		"issuer":       certInfo.Issuer,
		"auth_type":    cli.VerifyType,
		"dns_provider": cli.DnsProvider,
		"cert_path":    cli.OutputPath,
		"create_time":  time.Now().Unix(),
		"status":       1,
	}

	// Save DNS provider token if using DNS verification
	if cli.VerifyType == "dns" && cli.DnsConfig != nil {
		dnsProviderToken, _ := json.Marshal(cli.DnsConfig)
		data["dns_provider_token"] = string(dnsProviderToken)
	}

	// Insert into database
	certId, err := public.MR("ssl", "letsencrypts").InsertAndGetId(data)
	if err != nil {
		return 0, err
	}

	return int(certId), nil
}

/**
 * @brief Load configuration from JSON file
 * @param filepath Path to JSON configuration file
 * @return *AcmeCLI, error
 */
func LoadConfigFromFile(filepath string) (*AcmeCLI, error) {
	// Check if file exists
	if !public.FileExists(filepath) {
		return nil, fmt.Errorf("configuration file does not exist: %s", filepath)
	}

	// Read file
	content, err := public.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// Parse JSON
	var config AcmeCLI
	err = json.Unmarshal([]byte(content), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

/**
 * @brief Save configuration to JSON file
 * @param filepath Path to JSON configuration file
 * @return error
 */
func (cli *AcmeCLI) SaveConfigToFile(p string) error {
	// Convert to JSON
	content, err := json.MarshalIndent(cli, "", "  ")
	if err != nil {
		return err
	}

	// Create directory if it doesn't exist
	dir := public.AbsPath(p)
	dir = filepath.Dir(dir)
	if !public.FileExists(dir) {
		err = os.MkdirAll(dir, 0750)
		if err != nil {
			return err
		}
	}

	// Write file
	_, err = public.WriteFile(p, string(content))
	return err
}

/**
 * @brief Print help message
 */
func PrintHelp() {
	fmt.Println("ACME Certificate Management CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  acme apply --email <email> --domains <domains> [options]")
	fmt.Println("  acme renew --cert-id <cert_id>")
	fmt.Println("  acme list")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --email <email>               Email address for registration")
	fmt.Println("  --domains <domains>           Comma-separated list of domains")
	fmt.Println("  --output-path <path>          Path to save certificate files (default: ./ssl/certs/<domain>)")
	fmt.Println("  --verify-type <type>          Verification type: http or dns (default: http)")
	fmt.Println("  --dns-provider <provider>     DNS provider (required for dns verification)")
	fmt.Println("  --dns-config <json>           DNS provider credentials in JSON format")
	fmt.Println("  --config <file>               Load configuration from JSON file")
	fmt.Println("  --save-config <file>          Save configuration to JSON file")
	fmt.Println("  --account-id <id>             Account ID for saving to database")
	fmt.Println("  --use-own-server <bool>       Use built-in HTTP server for verification (default: true)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  acme apply --email admin@example.com --domains example.com,www.example.com")
	fmt.Println("  acme apply --email admin@example.com --domains example.com --use-own-server=false")
	fmt.Println("  acme apply --config /path/to/config.json")
	fmt.Println("  acme renew --cert-id 123")
	fmt.Println("  acme list")
}
