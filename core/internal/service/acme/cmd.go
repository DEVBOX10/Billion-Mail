package acme

import (
	"billionmail-core/internal/service/public"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

/**
 * @brief Run the ACME CLI
 * @param args Command line arguments
 * @return int Exit code (0 for success, non-zero for error)
 */
func RunCLI(args []string) int {
	if len(args) == 0 {
		PrintHelp()
		return 0
	}

	// Common flags
	emailFlag := flag.String("email", "", "Email address for registration")
	domainsFlag := flag.String("domains", "", "Comma-separated list of domains")
	outputPathFlag := flag.String("output-path", "", "Path to save certificate files")
	verifyTypeFlag := flag.String("verify-type", "http", "Verification type: http or dns")
	dnsProviderFlag := flag.String("dns-provider", "", "DNS provider (required for dns verification)")
	dnsConfigFlag := flag.String("dns-config", "", "DNS provider credentials in JSON format")
	configFileFlag := flag.String("config", "", "Load configuration from JSON file")
	saveConfigFlag := flag.String("save-config", "", "Save configuration to JSON file")
	certIdFlag := flag.Int("cert-id", 0, "Certificate ID for renewal")
	accountIdFlag := flag.Int("account-id", 0, "Account ID for saving to database")
	useOwnServerFlag := flag.Bool("use-own-server", true, "Use built-in HTTP server for verification")

	// Parse flags
	flag.CommandLine.Parse(args[1:])

	// Get command
	command := args[0]

	switch command {
	case "apply":
		return applyCommand(*emailFlag, *domainsFlag, *outputPathFlag, *verifyTypeFlag, *dnsProviderFlag, *dnsConfigFlag, *configFileFlag, *saveConfigFlag, *accountIdFlag, *useOwnServerFlag)
	case "renew":
		return renewCommand(*certIdFlag)
	case "list":
		return listCommand()
	case "help":
		PrintHelp()
		return 0
	default:
		fmt.Printf("Unknown command: %s\n", command)
		PrintHelp()
		return 1
	}
}

/**
 * @brief Apply for a new certificate
 * @return int Exit code
 */
func applyCommand(email, domains, outputPath, verifyType, dnsProvider, dnsConfigStr, configFile, saveConfigFile string, accountId int, useOwnServer bool) int {
	var cli *AcmeCLI
	var err error

	// Load from config file if specified
	if configFile != "" {
		cli, err = LoadConfigFromFile(configFile)
		if err != nil {
			fmt.Printf("Error loading config: %s\n", err)
			return 1
		}
	} else {
		// Parse domains
		domainList := []string{}
		if domains != "" {
			domainList = strings.Split(domains, ",")
		}

		// Create CLI instance
		if len(domainList) == 0 {
			fmt.Println("Error: at least one domain is required")
			return 1
		}

		cli = NewAcmeCLI(email, domainList, outputPath)

		// Set use own server option
		cli.UseOwnServer = useOwnServer

		// Set verification type
		if verifyType == "dns" {
			// Parse DNS config
			dnsConfig := map[string]string{}
			if dnsConfigStr != "" {
				err = json.Unmarshal([]byte(dnsConfigStr), &dnsConfig)
				if err != nil {
					fmt.Printf("Error parsing DNS config: %s\n", err)
					return 1
				}
			}

			cli.SetDNSVerification(dnsProvider, dnsConfig)
		} else {
			cli.SetHTTPVerification()
		}
	}

	// Save config to file if specified
	if saveConfigFile != "" {
		err = cli.SaveConfigToFile(saveConfigFile)
		if err != nil {
			fmt.Printf("Error saving config: %s\n", err)
			return 1
		}
		fmt.Printf("Configuration saved to: %s\n", saveConfigFile)
	}

	// Validate parameters
	if err = cli.Validate(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return 1
	}

	// Apply for certificate
	certificatePath, privateKeyPath, err := cli.Apply()
	if err != nil {
		fmt.Printf("Error applying for certificate: %s\n", err)
		return 1
	}

	// Save to database if account ID is provided
	if accountId > 0 {
		// Read certificate and private key from files
		certificate, err := public.ReadFile(certificatePath)
		if err != nil {
			fmt.Printf("Error reading certificate file: %s\n", err)
			return 1
		}

		privateKey, err := public.ReadFile(privateKeyPath)
		if err != nil {
			fmt.Printf("Error reading private key file: %s\n", err)
			return 1
		}

		// Save to database
		certId, err := cli.SaveToDatabase(accountId, certificate, privateKey)
		if err != nil {
			fmt.Printf("Error saving to database: %s\n", err)
			return 1
		}

		fmt.Printf("Certificate saved to database with ID: %d\n", certId)
	}

	return 0
}

/**
 * @brief Renew an existing certificate
 * @return int Exit code
 */
func renewCommand(certId int) int {
	if certId <= 0 {
		fmt.Println("Error: certificate ID is required")
		return 1
	}

	// Get certificate info from database
	certInfo, err := public.MR("ssl", "letsencrypts").Where("cert_id = ?", certId).One()
	if err != nil {
		fmt.Printf("Error retrieving certificate: %s\n", err)
		return 1
	}

	if certInfo == nil {
		fmt.Printf("Certificate not found with ID: %d\n", certId)
		return 1
	}

	// Parse domain list
	var domains []string
	err = json.Unmarshal([]byte(certInfo["dns"].String()), &domains)
	if err != nil {
		fmt.Printf("Error parsing domains: %s\n", err)
		return 1
	}

	// Parse DNS provider token
	dnsProviderToken := map[string]string{}
	if certInfo["dns_provider_token"].String() != "" {
		err = json.Unmarshal([]byte(certInfo["dns_provider_token"].String()), &dnsProviderToken)
		if err != nil {
			fmt.Printf("Error parsing DNS provider token: %s\n", err)
			return 1
		}
	}

	// Get account info
	accountId := certInfo["account_id"].Int()
	accountInfo, err := public.M("account").Where("account_id = ?", accountId).One()
	if err != nil {
		fmt.Printf("Error retrieving account: %s\n", err)
		return 1
	}

	if accountInfo == nil {
		fmt.Printf("Account not found with ID: %d\n", accountId)
		return 1
	}

	email := accountInfo["email"].String()

	// Get certificate path
	certPath := certInfo["cert_path"].String()
	if certPath == "" {
		// Use default path if not specified
		certPath = public.AbsPath(fmt.Sprintf("ssl/certs/%s", certInfo["subject"].String()))
	}

	// Start renewal process
	fmt.Printf("Renewing certificate for domains: %s\n", strings.Join(domains, ", "))
	err = StartRenew(
		certId,
		accountId,
		certPath,
		domains,
		email,
		certInfo["auth_type"].String(),
		certInfo["dns_provider"].String(),
		dnsProviderToken,
	)

	if err != nil {
		fmt.Printf("Error renewing certificate: %s\n", err)
		return 1
	}

	fmt.Println("Certificate renewed successfully")
	return 0
}

/**
 * @brief List all certificates
 * @return int Exit code
 */
func listCommand() int {
	// Get all certificates from database
	certList, err := public.MR("ssl", "letsencrypts").All()
	if err != nil {
		fmt.Printf("Error retrieving certificates: %s\n", err)
		return 1
	}

	if len(certList) == 0 {
		fmt.Println("No certificates found")
		return 0
	}

	// Print certificate list
	fmt.Println("ID\tSubject\tIssuer\tValid Until\tDomains")
	fmt.Println("--\t-------\t------\t-----------\t-------")

	for _, cert := range certList {
		// Parse domain list
		var domains []string
		err = json.Unmarshal([]byte(cert["dns"].String()), &domains)
		if err != nil {
			domains = []string{"<error>"}
		}

		// Format expiration date
		expireTime := cert["endtime"].Int64()
		expireDate := public.TimeToDate(expireTime, "2006-01-02")

		// Print certificate info
		fmt.Printf("%d\t%s\t%s\t%s\t%s\n",
			cert["cert_id"].Int(),
			cert["subject"].String(),
			cert["issuer"].String(),
			expireDate,
			strings.Join(domains, ", "),
		)
	}

	return 0
}

/**
 * @brief Main entry point for command line usage
 */
func main() {
	// Skip program name
	args := os.Args[1:]

	// Run CLI
	exitCode := RunCLI(args)

	// Exit with code
	os.Exit(exitCode)
}
