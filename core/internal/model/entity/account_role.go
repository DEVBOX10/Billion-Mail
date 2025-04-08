package entity

// AccountRole defines the account-role relationship entity
type AccountRole struct {
	Id         int64 `json:"id"          dc:"Relation ID"`
	AccountId  int64 `json:"account_id"  dc:"Account ID"`
	RoleId     int64 `json:"role_id"     dc:"Role ID"`
	CreateTime int64 `json:"create_time" dc:"Creation time"`
}
