package entity

// AccountRole defines the account-role relationship entity
type AccountRole struct {
	Id         int `json:"id"          dc:"Relation ID"`
	AccountId  int `json:"account_id"  dc:"Account ID"`
	RoleId     int `json:"role_id"     dc:"Role ID"`
	CreateTime int `json:"create_time" dc:"Creation time"`
}
