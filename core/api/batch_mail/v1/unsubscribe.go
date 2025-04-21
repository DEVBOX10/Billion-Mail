package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UnsubscribeReq struct {
	g.Meta  `path:"/unsubscribe" method:"post" tags:"Unsubscribe" summary:" unsubscribe request"`
	Jwt     string `json:"jwt" dc:" unsubscribe jwt"`
	GroupId []int  `json:"group_id" dc:"unsubscribe groupID"`
}

type UnsubscribeRes struct {
	Success bool   `json:"success" dc:"is success"`
	Message string `json:"message" dc:"response messageonse message"`
}

type GetUserGroupsReq struct {
	g.Meta `path:"/unsubscribe/user_group" method:"get" tags:"Unsubscribe" summary:"get user groups"`
	Email  string `json:"email" dc:"user email"`
}

type GetUserGroupsRes struct {
	Groups []*GroupInfo `json:"groups" dc:"user groups"`
}
