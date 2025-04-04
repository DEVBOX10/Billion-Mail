package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

// Mailbox defines the mailbox entity
type Mailbox struct {
	Username       string `json:"username"        dc:"Email address"`
	Password       string `json:"password"        dc:"Password"`
	PasswordEncode string `json:"password_encode" dc:"Encoded password"`
	FullName       string `json:"full_name"       dc:"Full name"`
	IsAdmin        int    `json:"is_admin"        dc:"Is administrator: 1-yes, 0-no"`
	Maildir        string `json:"maildir"         dc:"Mailbox directory"`
	Quota          int64  `json:"quota"           dc:"Mailbox quota"`
	LocalPart      string `json:"local_part"      dc:"Local part (username)"`
	Domain         string `json:"domain"          dc:"Domain name"`
	CreateTime     int64  `json:"create_time"     dc:"Creation time"`
	UpdateTime     int64  `json:"update_time"     dc:"Update time"`
	Active         int    `json:"active"          dc:"Status: 1-enabled, 0-disabled"`
}

type AddMailboxReq struct {
	g.Meta        `path:"/mailbox/create" tags:"MailBox" method:"post" summary:"Create mailbox" in:"body"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Domain        string `json:"domain" v:"required|domain" dc:"Domain"`
	FullName      string `json:"full_name" v:"required|passport" dc:"username"`
	Password      string `json:"password" v:"required|password" dc:"Password"`
	Active        int    `json:"active" v:"required" dc:"Status" d:"1"`
	IsAdmin       int    `json:"isAdmin" v:"required" dc:"IsAdmin" d:"0"`
	Quota         int    `json:"quota" v:"required|min:1" dc:"Quota" d:"5242880"`
}

type AddMailboxRes struct {
	api_v1.StandardRes
}

type UpdateMailboxReq struct {
	g.Meta        `path:"/mailbox/update" tags:"MailBox" method:"post" summary:"Update mailbox" in:"body"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Domain        string `json:"domain" v:"required|domain" dc:"Domain"`
	FullName      string `json:"full_name" v:"required|passport" dc:"username"`
	Password      string `json:"password" v:"password" dc:"Password"`
	Active        int    `json:"active" v:"required" dc:"Status" d:"1"`
	IsAdmin       int    `json:"isAdmin" v:"required" dc:"IsAdmin" d:"0"`
	Quota         int    `json:"quota" v:"required|min:1" dc:"Quota" d:"5242880"`
}

type UpdateMailboxRes struct {
	api_v1.StandardRes
}

type DeleteMailboxReq struct {
	g.Meta        `path:"/mailbox/delete" tags:"MailBox" method:"post" summary:"Delete mailbox" in:"body"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Email         string `json:"email" v:"required|email" dc:"Email"`
}

type DeleteMailboxRes struct {
	api_v1.StandardRes
}

type GetMailboxReq struct {
	g.Meta        `path:"/mailbox/list" tags:"MailBox" method:"get" summary:"Get mailbox" in:"query"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Keyword       string `json:"keyword" dc:"Keyword"`
	Domain        string `json:"domain" v:"required|domain" dc:"Domain"`
	Page          int    `json:"page" v:"min:1" dc:"Page" d:"1"`
	PageSize      int    `json:"page_size" v:"min:1" dc:"Page size" d:"20"`
}

type GetMailboxRes struct {
	api_v1.StandardRes
	Data struct {
		Total int       `json:"total"`
		List  []Mailbox `json:"list"`
	} `json:"data"`
}
