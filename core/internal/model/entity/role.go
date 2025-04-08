package entity

// Role defines the role entity
type Role struct {
	RoleId      int64  `json:"role_id"       dc:"Role ID"`
	RoleName    string `json:"role_name"     dc:"Role name"`
	Description string `json:"description"   dc:"Role description"`
	Status      int    `json:"status"        dc:"Status: 1-active, 0-disabled"`
	CreateTime  int64  `json:"create_time"   dc:"Creation time"`
	UpdateTime  int64  `json:"update_time"   dc:"Update time"`
}
