package entity

// Letsencrypt certificate info
type Letsencrypt struct {
	CertId      int64    `json:"cert_id"`     // cert id
	AccountId   int64    `json:"account_id"`  // account id
	Certificate string   `json:"certificate"` // cert pem
	PrivateKey  string   `json:"private_key"` // key pem
	ErrorInfo   string   `json:"error_info"`  // error info
	Progress    string   `json:"progress"`    // progress logs
	Status      int      `json:"status"`      // status
	NotAfter    string   `json:"not_after"`   // end time of cert
	NotBefore   string   `json:"not_before"`  // start time of cert
	DnsNames    []string `json:"dns"`         // domain names
	EndTime     int64    `json:"endtime"`     // end time of cert (Unix timestamp)
	Subject     string   `json:"subject"`     // subject
}
