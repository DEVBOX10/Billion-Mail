package entity

// Alias defines the email alias entity
type Alias struct {
	Address    string `json:"address"     dc:"Alias address"`
	Goto       string `json:"goto"        dc:"Forwarding target"`
	Domain     string `json:"domain"      dc:"Domain name"`
	CreateTime int64  `json:"create_time" dc:"Creation time"`
	UpdateTime int64  `json:"update_time" dc:"Update time"`
	Active     int    `json:"active"      dc:"Status: 1-enabled, 0-disabled"`
}
