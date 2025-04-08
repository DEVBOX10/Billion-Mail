package entity

// Permission defines the permission entity
type Permission struct {
	PermissionId   int64  `json:"permission_id"   dc:"Permission ID"`
	PermissionName string `json:"permission_name" dc:"Permission name"`
	Description    string `json:"description"     dc:"Permission description"`
	Module         string `json:"module"          dc:"Module name"`
	Action         string `json:"action"          dc:"Action type: create/read/update/delete"`
	Resource       string `json:"resource"        dc:"Resource type"`
	Status         int    `json:"status"          dc:"Status: 1-active, 0-disabled"`
	CreateTime     int64  `json:"create_time"     dc:"Creation time"`
	UpdateTime     int64  `json:"update_time"     dc:"Update time"`
}
