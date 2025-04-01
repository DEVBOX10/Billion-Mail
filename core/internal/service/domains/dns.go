package domains

import (
	v1 "billionmail-core/api/domains/v1"
	"net"
	"strings"
)

// ValidateARecord checks if the given A record is valid
func ValidateARecord(record v1.DNSRecord) bool {
	if strings.ToUpper(record.Type) != "A" && strings.ToUpper(record.Type) != "AAAA" {
		return false
	}

	// Query A records
	ips, err := net.LookupIP(record.Host)
	if err != nil {
		return false
	}

	// Check if any IP matches the expected server IP
	for _, ip := range ips {
		if ip.String() == record.Value {
			return true
		}
	}

	return false
}

// ValidateTXTRecord checks if the given TXT record is valid
func ValidateTXTRecord(record v1.DNSRecord) bool {
	if strings.ToUpper(record.Type) != "TXT" {
		return false
	}

	// Query TXT records
	txtRecords, err := net.LookupTXT(record.Host)
	if err != nil {
		return false
	}

	// Check if any TXT record matches the expected value
	for _, txt := range txtRecords {
		if txt == record.Value {
			return true
		}
	}

	return false
}

// ValidateMXRecord checks if the given MX record is valid
func ValidateMXRecord(record v1.DNSRecord, domain string) bool {
	if strings.ToUpper(record.Type) != "MX" {
		return false
	}

	// Query MX records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return false
	}

	// Check if any MX record matches the expected value
	for _, mx := range mxRecords {
		if mx.Host == record.Value {
			return true
		}
	}

	return false
}
