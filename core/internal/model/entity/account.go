package entity

// Account defines the account entity
type Account struct {
	AccountId     int64  `json:"account_id"     dc:"Account ID"`
	Username      string `json:"username"       dc:"Username"`
	Password      string `json:"-"       dc:"Password"`
	Email         string `json:"email"          dc:"Email address"`
	Status        int    `json:"status"         dc:"Status: 1-active, 0-disabled"`
	Language      string `json:"language"       dc:"Language setting"`
	LastLoginTime int64  `json:"last_login_time" dc:"Last login time"`
	CreateTime    int64  `json:"create_time"    dc:"Creation time"`
	UpdateTime    int64  `json:"update_time"    dc:"Update time"`
}
