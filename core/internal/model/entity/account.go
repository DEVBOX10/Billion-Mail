package entity

// Account defines the account entity
type Account struct {
	AccountId     int    `json:"account_id"     dc:"Account ID"`
	Username      string `json:"username"       dc:"Username"`
	Password      string `json:"password"       dc:"Password"`
	Email         string `json:"email"          dc:"Email address"`
	Status        int    `json:"status"         dc:"Status: 1-active, 0-disabled"`
	Language      string `json:"language"       dc:"Language setting"`
	LastLoginTime int    `json:"last_login_time" dc:"Last login time"`
	CreateTime    int    `json:"create_time"    dc:"Creation time"`
	UpdateTime    int    `json:"update_time"    dc:"Update time"`
}
