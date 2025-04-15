package domains

import (
	v1 "billionmail-core/api/domains/v1"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/frame/g"
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
func ValidateTXTRecord(record v1.DNSRecord, domain string) bool {
	if strings.ToUpper(record.Type) != "TXT" {
		return false
	}

	if record.Host != "@" {
		domain = record.Host + "." + domain
	}

	cacheKey := "txt_records_" + domain
	txtRecords := make([]string, 0)
	cacheTxtRecords := public.GetCache(cacheKey)

	if cacheTxtRecords != nil {
		txtRecords = cacheTxtRecords.([]string)
		g.Log().Debug(context.Background(), "get txt records from cache", cacheTxtRecords)
	} else {
		// Query TXT records
		var err error
		txtRecords, err = net.LookupTXT(domain)
		if err != nil {
			// g.Log().Error(context.Background(), "query txt records failed", err)
			return false
		}
		g.Log().Debug(context.Background(), "query txt records success", txtRecords)
		// Store TXT records in cache
		public.SetCache(cacheKey, txtRecords, 60)
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
		// g.Log().Error(context.Background(), "query mx record failed", err)
		return false
	}

	g.Log().Debug(context.Background(), "query mx record success", mxRecords)

	// Check if any MX record matches the expected value
	for _, mx := range mxRecords {
		if strings.TrimSuffix(mx.Host, ".") == record.Value {
			return true
		}
	}

	return false
}

// ValidatePTRRecord checks if the given PTR record is valid
func ValidatePTRRecord(record v1.DNSRecord) bool {
	if strings.ToUpper(record.Type) != "PTR" {
		return false
	}

	// Query PTR records
	ptrRecords, err := net.LookupAddr(record.Host)
	if err != nil {
		return false
	}

	g.Log().Debug(context.Background(), "query ptr record success", ptrRecords)

	// Check if any PTR record matches the expected value
	for _, ptr := range ptrRecords {
		if strings.TrimSuffix(ptr, ".") == record.Value {
			return true
		}
	}

	return false
}
