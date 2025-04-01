package entity

// Mailbox defines the mailbox entity
type Mailbox struct {
	Username       string `json:"username"        dc:"Email address"`
	Password       string `json:"password"        dc:"Password"`
	PasswordEncode string `json:"password_encode" dc:"Encoded password"`
	FullName       string `json:"full_name"       dc:"Full name"`
	IsAdmin        int    `json:"is_admin"        dc:"Is administrator: 1-yes, 0-no"`
	Maildir        string `json:"maildir"         dc:"Mailbox directory"`
	Quota          int64  `json:"quota"           dc:"Mailbox quota"`
	LocalPart      string `json:"local_part"      dc:"Local part (username)"`
	Domain         string `json:"domain"          dc:"Domain name"`
	CreateTime     int64  `json:"create_time"     dc:"Creation time"`
	UpdateTime     int64  `json:"update_time"     dc:"Update time"`
	Active         int    `json:"active"          dc:"Status: 1-enabled, 0-disabled"`
}
