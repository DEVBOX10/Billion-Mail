package entity

// RolePermission defines the role-permission relationship entity
type RolePermission struct {
	Id           int `json:"id"            dc:"Relation ID"`
	RoleId       int `json:"role_id"       dc:"Role ID"`
	PermissionId int `json:"permission_id" dc:"Permission ID"`
	CreateTime   int `json:"create_time"   dc:"Creation time"`
}
