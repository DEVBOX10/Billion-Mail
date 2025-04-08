package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

// RoleInfoItem defines the role information structure
type RoleInfoItem struct {
	Id          int64  `json:"id" dc:"Role ID"`
	Name        string `json:"name" dc:"Role name"`
	Description string `json:"description" dc:"Role description"`
	Status      int    `json:"status" dc:"Status (0:disabled, 1:enabled)"`
	CreateTime  int64  `json:"create_time" dc:"Creation time"`
}

// RoleListReq defines the request for getting role list
type RoleListReq struct {
	g.Meta        `path:"/role/list" method:"get" tags:"RBAC" summary:"Get role list" sm:"Get role list" in:"query"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Page          int    `p:"page" d:"1" v:"min:1#Page number must be greater than 0" dc:"Page number"`
	PageSize      int    `p:"pageSize" d:"10" v:"min:1#Items per page must be greater than 0" dc:"Items per page"`
	Name          string `p:"name" dc:"Role name filter"`
	Status        int    `p:"status" dc:"Status filter"`
}

// RoleListRes defines the response for getting role list
type RoleListRes struct {
	api_v1.StandardRes
	Data struct {
		List  []RoleInfoItem `json:"list" dc:"Role list"`
		Total int            `json:"total" dc:"Total count"`
		Page  int            `json:"page" dc:"Current page number"`
	} `json:"data"`
}

// RoleDetailReq defines the request for getting role details
type RoleDetailReq struct {
	g.Meta        `path:"/role/detail" method:"get" tags:"RBAC" summary:"Get role details" sm:"Get role details" in:"query"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	RoleId        int64  `p:"roleId" v:"required#Please provide role ID" dc:"Role ID"`
}

// RoleDetailRes defines the response for getting role details
type RoleDetailRes struct {
	api_v1.StandardRes
	Data struct {
		Role           RoleInfoItem         `json:"role" dc:"Role information"`
		Permissions    []PermissionInfoItem `json:"permissions" dc:"Role permissions list"`
		AllPermissions []PermissionInfoItem `json:"allPermissions" dc:"All available permissions"`
	} `json:"data"`
}

// RoleCreateReq defines the request for creating role
type RoleCreateReq struct {
	g.Meta        `path:"/role/create" method:"post" tags:"RBAC" summary:"Create role" sm:"Create role" in:"body"`
	Authorization string  `json:"authorization" dc:"Authorization" in:"header"`
	Name          string  `p:"name" v:"required#Role name cannot be empty" dc:"Role name"`
	Description   string  `p:"description" dc:"Role description"`
	PermissionIds []int64 `p:"permissionIds" dc:"Permission IDs list"`
	Status        int     `p:"status" d:"1" dc:"Status (0:disabled, 1:enabled)"`
}

// RoleCreateRes defines the response for creating role
type RoleCreateRes struct {
	api_v1.StandardRes
	Data struct {
		RoleId int64 `json:"roleId" dc:"New role ID"`
	} `json:"data"`
}

// RoleUpdateReq defines the request for updating role
type RoleUpdateReq struct {
	g.Meta        `path:"/role/update" method:"post" tags:"RBAC" summary:"Update role" sm:"Update role" in:"body"`
	Authorization string  `json:"authorization" dc:"Authorization" in:"header"`
	RoleId        int64   `p:"roleId" v:"required#Role ID cannot be empty" dc:"Role ID"`
	Name          string  `p:"name" dc:"Role name"`
	Description   string  `p:"description" dc:"Role description"`
	PermissionIds []int64 `p:"permissionIds" dc:"Permission IDs list"`
	Status        int     `p:"status" dc:"Status (0:disabled, 1:enabled)"`
}

// RoleUpdateRes defines the response for updating role
type RoleUpdateRes struct {
	api_v1.StandardRes
}

// RoleDeleteReq defines the request for deleting role
type RoleDeleteReq struct {
	g.Meta        `path:"/role/delete" method:"post" tags:"RBAC" summary:"Delete role" sm:"Delete role" in:"body"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	RoleId        int64  `p:"roleId" v:"required#Role ID cannot be empty" dc:"Role ID"`
}

// RoleDeleteRes defines the response for deleting role
type RoleDeleteRes struct {
	api_v1.StandardRes
}
