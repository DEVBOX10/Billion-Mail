package entity

// Domain defines the domain entity
type Domain struct {
	Domain       string `json:"domain"        dc:"Domain name"`
	ARecord      string `json:"a_record"      dc:"A record"`
	Mailboxes    int    `json:"mailboxes"     dc:"Number of mailboxes created"`
	MailboxQuota int64  `json:"mailbox_quota" dc:"Default mailbox space size"`
	Quota        int64  `json:"quota"         dc:"Domain quota"`
	RateLimit    int    `json:"rate_limit"    dc:"Rate limit for sending emails per second"`
	CreateTime   int64  `json:"create_time"   dc:"Creation time"`
	Active       int    `json:"active"        dc:"Status: 1-enabled, 0-disabled"`
}
