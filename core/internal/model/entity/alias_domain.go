package entity

// AliasDomain defines the domain alias entity
type AliasDomain struct {
	AliasDomain  string `json:"alias_domain"  dc:"Alias domain name"`
	TargetDomain string `json:"target_domain" dc:"Target domain name"`
	CreateTime   int64  `json:"create_time"   dc:"Creation time"`
	UpdateTime   int64  `json:"update_time"   dc:"Update time"`
	Active       int    `json:"active"        dc:"Status: 1-enabled, 0-disabled"`
}
