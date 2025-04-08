package entity

// RolePermission defines the role-permission relationship entity
type RolePermission struct {
	Id           int64 `json:"id"            dc:"Relation ID"`
	RoleId       int64 `json:"role_id"       dc:"Role ID"`
	PermissionId int64 `json:"permission_id" dc:"Permission ID"`
	CreateTime   int64 `json:"create_time"   dc:"Creation time"`
}
