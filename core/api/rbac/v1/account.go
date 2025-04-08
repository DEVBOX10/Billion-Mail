package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

// AccountInfoItem defines the account information structure
type AccountInfoItem struct {
	Id         int64  `json:"id" dc:"Account ID"`
	Username   string `json:"username" dc:"Username"`
	Email      string `json:"email" dc:"Email"`
	Status     int    `json:"status" dc:"Status(0:disabled, 1:enabled)"`
	Language   string `json:"language" dc:"Language setting"`
	CreateTime int64  `json:"create_time" dc:"Creation time"`
}

// AccountListReq defines the request for getting account list
type AccountListReq struct {
	g.Meta        `path:"/account/list" method:"get" tags:"RBAC" summary:"Get account list" sm:"Get account list" in:"query"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Page          int    `p:"page" d:"1" v:"min:1#Page number must be greater than 0" dc:"Page number"`
	PageSize      int    `p:"pageSize" d:"10" v:"min:1#Items per page must be greater than 0" dc:"Items per page"`
	Username      string `p:"username" dc:"Username filter"`
	Email         string `p:"email" dc:"Email filter"`
	Status        int    `p:"status" dc:"Status filter"`
}

// AccountListRes is the response for getting account list
type AccountListRes struct {
	api_v1.StandardRes
	Data struct {
		List  []AccountInfoItem `json:"list" dc:"Account list"`
		Total int               `json:"total" dc:"Total"`
		Page  int               `json:"page" dc:"Current page"`
	} `json:"data"`
}

// AccountDetailReq is the request for getting account details
type AccountDetailReq struct {
	g.Meta        `path:"/account/detail" method:"get" tags:"RBAC" summary:"Get account details" sm:"Get account details" in:"query"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	AccountId     int64  `p:"accountId" v:"required#Please provide account ID" dc:"Account ID"`
}

// AccountDetailRes is the response for getting account details
type AccountDetailRes struct {
	api_v1.StandardRes
	Data struct {
		Account  AccountInfoItem `json:"account" dc:"Account information"`
		Roles    []RoleInfoItem  `json:"roles" dc:"Account roles list"`
		AllRoles []RoleInfoItem  `json:"allRoles" dc:"All available roles"`
	} `json:"data"`
}

// AccountCreateReq is the request for creating account
type AccountCreateReq struct {
	g.Meta        `path:"/account/create" method:"post" tags:"RBAC" summary:"Create account" sm:"Create account" in:"body"`
	Authorization string  `json:"authorization" dc:"Authorization" in:"header"`
	Username      string  `p:"username" v:"required#User name can not be empty" dc:"User name"`
	Password      string  `p:"password" v:"required#Password can not be empty" dc:"Password"`
	Email         string  `p:"email" v:"required|email#Email can not be empty|Email format is not correct" dc:"Email"`
	RoleIds       []int64 `p:"roleIds" dc:"Role ID list"`
	Status        int     `p:"status" d:"1" dc:"Status(0:disabled, 1:enabled)"`
	Lang          string  `p:"lang" d:"en" dc:"Language setting"`
}

// AccountCreateRes is the response for creating account
type AccountCreateRes struct {
	api_v1.StandardRes
	Data struct {
		AccountId int64 `json:"accountId" dc:"New account ID"`
	} `json:"data"`
}

// AccountUpdateReq is the request for updating account
type AccountUpdateReq struct {
	g.Meta        `path:"/account/update" method:"post" tags:"RBAC" summary:"Update account" sm:"Update account" in:"body"`
	Authorization string  `json:"authorization" dc:"Authorization" in:"header"`
	AccountId     int64   `p:"accountId" v:"required#Account ID can not be empty" dc:"Account ID"`
	Username      string  `p:"username" dc:"User name"`
	Email         string  `p:"email" v:"email#Email format is not correct" dc:"Email"`
	RoleIds       []int64 `p:"roleIds" dc:"Role ID list"`
	Status        int     `p:"status" dc:"Status(0:disabled, 1:enabled)"`
	Lang          string  `p:"lang" dc:"Language setting"`
}

// AccountUpdateRes is the response for updating account
type AccountUpdateRes struct {
	api_v1.StandardRes
}

// AccountPasswordReq is the request for updating account password
type AccountPasswordReq struct {
	g.Meta        `path:"/account/password" method:"post" tags:"RBAC" summary:"Update account password" sm:"Update account password" in:"body"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	AccountId     int64  `p:"accountId" v:"required#Account ID can not be empty" dc:"Account ID"`
	OldPassword   string `p:"oldPassword" dc:"Old password(administrator optional)"`
	NewPassword   string `p:"newPassword" v:"required#New password can not be empty" dc:"New password"`
}

// AccountPasswordRes is the response for updating account password
type AccountPasswordRes struct {
	api_v1.StandardRes
}

// AccountDeleteReq is the request for deleting account
type AccountDeleteReq struct {
	g.Meta        `path:"/account/delete" method:"post" tags:"RBAC" summary:"Delete account" sm:"Delete account" in:"body"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	AccountId     int64  `p:"accountId" v:"required#Account ID can not be empty" dc:"Account ID"`
}

// AccountDeleteRes is the response for deleting account
type AccountDeleteRes struct {
	api_v1.StandardRes
}
