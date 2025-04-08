package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionInfoItem defines the permission information structure
type PermissionInfoItem struct {
	Id          int64  `json:"id" dc:"Permission ID"`
	Name        string `json:"name" dc:"Permission name"`
	Description string `json:"description" dc:"Permission description"`
	Module      string `json:"module" dc:"Module name"`
	Action      string `json:"action" dc:"Action name (create/read/update/delete)"`
	Resource    string `json:"resource" dc:"Resource name"`
	Status      int    `json:"status" dc:"Status (0:disabled, 1:enabled)"`
	CreateTime  int64  `json:"create_time" dc:"Creation time"`
	UpdateTime  int64  `json:"update_time" dc:"Update time"`
}

// PermissionListReq defines the request for getting permission list
type PermissionListReq struct {
	g.Meta        `path:"/permission/list" method:"get" tags:"RBAC" summary:"Get permission list" sm:"Get permission list" in:"query"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Page          int    `p:"page" d:"1" v:"min:1#Page number must be greater than 0" dc:"Page number"`
	PageSize      int    `p:"pageSize" d:"10" v:"min:1#Items per page must be greater than 0" dc:"Items per page"`
	Module        string `p:"module" dc:"Module name filter"`
	Action        string `p:"action" dc:"Action name filter"`
	Status        int    `p:"status" dc:"Status filter"`
}

// PermissionListRes defines the response for getting permission list
type PermissionListRes struct {
	api_v1.StandardRes
	Data struct {
		List  []PermissionInfoItem `json:"list" dc:"Permission list"`
		Total int                  `json:"total" dc:"Total count"`
		Page  int                  `json:"page" dc:"Current page number"`
	} `json:"data"`
}

// PermissionDetailReq defines the request for getting permission details
type PermissionDetailReq struct {
	g.Meta        `path:"/permission/detail" method:"get" tags:"RBAC" summary:"Get permission details" sm:"Get permission details" in:"query"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	PermissionId  int64  `p:"permissionId" v:"required#Please provide permission ID" dc:"Permission ID"`
}

// PermissionDetailRes defines the response for getting permission details
type PermissionDetailRes struct {
	api_v1.StandardRes
	Data struct {
		Permission PermissionInfoItem `json:"permission" dc:"Permission information"`
	} `json:"data"`
}

// PermissionCreateReq defines the request for creating permission
type PermissionCreateReq struct {
	g.Meta        `path:"/permission/create" method:"post" tags:"RBAC" summary:"Create permission" sm:"Create permission" in:"body"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Name          string `p:"name" v:"required#Permission name cannot be empty" dc:"Permission name"`
	Module        string `p:"module" v:"required#Module name cannot be empty" dc:"Module name"`
	Action        string `p:"action" v:"required#Action name cannot be empty" dc:"Action name"`
	Resource      string `p:"resource" v:"required#Resource name cannot be empty" dc:"Resource name"`
	Description   string `p:"description" dc:"Permission description"`
	Status        int    `p:"status" d:"1" dc:"Status (0:disabled, 1:enabled)"`
}

// PermissionCreateRes defines the response for creating permission
type PermissionCreateRes struct {
	api_v1.StandardRes
	Data struct {
		PermissionId int64 `json:"permissionId" dc:"New permission ID"`
	} `json:"data"`
}

// PermissionUpdateReq defines the request for updating permission
type PermissionUpdateReq struct {
	g.Meta        `path:"/permission/update" method:"post" tags:"RBAC" summary:"Update permission" sm:"Update permission" in:"body"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	PermissionId  int64  `p:"permissionId" v:"required#Permission ID cannot be empty" dc:"Permission ID"`
	Name          string `p:"name" dc:"Permission name"`
	Module        string `p:"module" dc:"Module name"`
	Action        string `p:"action" dc:"Action name"`
	Resource      string `p:"resource" dc:"Resource name"`
	Description   string `p:"description" dc:"Permission description"`
	Status        int    `p:"status" dc:"Status (0:disabled, 1:enabled)"`
}

// PermissionUpdateRes defines the response for updating permission
type PermissionUpdateRes struct {
	api_v1.StandardRes
}

// PermissionDeleteReq defines the request for deleting permission
type PermissionDeleteReq struct {
	g.Meta        `path:"/permission/delete" method:"post" tags:"RBAC" summary:"Delete permission" sm:"Delete permission" in:"body"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	PermissionId  int64  `p:"permissionId" v:"required#Permission ID cannot be empty" dc:"Permission ID"`
}

// PermissionDeleteRes defines the response for deleting permission
type PermissionDeleteRes struct {
	api_v1.StandardRes
}

// PermissionCheckReq defines the request for checking permission
type PermissionCheckReq struct {
	g.Meta        `path:"/permission/check" method:"post" tags:"RBAC" summary:"Check permission" sm:"Check permission" in:"body"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Module        string `p:"module" v:"required#Module name cannot be empty" dc:"Module name"`
	Action        string `p:"action" v:"required#Action name cannot be empty" dc:"Action name"`
	Resource      string `p:"resource" v:"required#Resource name cannot be empty" dc:"Resource name"`
}

// PermissionCheckRes defines the response for checking permission
type PermissionCheckRes struct {
	api_v1.StandardRes
	Data struct {
		HasPermission bool `json:"hasPermission" dc:"Whether has permission"`
	} `json:"data"`
}
