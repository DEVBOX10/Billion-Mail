package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

type AbnormalRecipient struct {
	Id          int    `json:"id"          dc:"Group ID"`
	Recipient   string `json:"recipient"   dc:"Recipient Email"`
	CreateTime  int    `json:"create_time" dc:"Create Time"`
	Description string `json:"description" dc:"Description"`
	Count       int    `json:"count"       dc:"Count"`
	AddType     int    `json:"add_type"    dc:"Add Type"`
	// AddType:Description     1: Manually added, 2: Automatically scanned, 3: Manually scanned
}

type ListAbnormalRecipientReq struct {
	g.Meta        `path:"/abnormal_recipient/list" method:"get" tags:"Abnormal Recipient" summary:"List abnormal recipients"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Page          int    `json:"page" v:"required|min:1" dc:"Page Number"`
	PageSize      int    `json:"page_size" v:"required|min:1" dc:"Page Size"`
	Keyword       string `json:"keyword" dc:"Search Keyword"`
	AddType       int    `json:"add_type" dc:"Add Type"`
}

type ListAbnormalRecipientRes struct {
	api_v1.StandardRes
	Data struct {
		Total int                  `json:"total" dc:"Total Count"`
		List  []*AbnormalRecipient `json:"list" dc:"Abnormal Recipient List"`
	} `json:"data"`
}

type AddAbnormalRecipientReq struct {
	g.Meta        `path:"/abnormal_recipient/add" method:"post" tags:"Abnormal Recipient" summary:"Add abnormal recipient"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Recipient     string `json:"recipient" v:"required|email" dc:"Recipient Email"`
}

type AddAbnormalRecipientRes struct {
	api_v1.StandardRes
}

type DeleteAbnormalRecipientReq struct {
	g.Meta        `path:"/abnormal_recipient/delete" method:"post" tags:"Abnormal Recipient" summary:"Delete abnormal recipient"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Id            int    `json:"id" v:"required|min:1" dc:"ID"`
}

type DeleteAbnormalRecipientRes struct {
	api_v1.StandardRes
}
