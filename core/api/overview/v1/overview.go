package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

type OverviewReq struct {
	g.Meta        `path:"/overview" tags:"Overview" method:"get" summary:"Get overview information" in:"query"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Domain        string `v:"required|domain" dc:"Domains" dc:"Domain"`
	StartTime     int64  `v:"required" dc:"StartTime" dc:"StartTime"`
	EndTime       int64  `v:"required" dc:"EndTime" dc:"EndTime"`
}

type OverviewRes struct {
	api_v1.StandardRes
	Data struct {
		Dashboard       interface{} `json:"dashboard" dc:"Dashboard"`
		MailProviders   interface{} `json:"mail_providers" dc:"Mail Providers"`
		SendMailChart   interface{} `json:"send_mail_chart" dc:"Send Mail Chart"`
		BounceRateChart interface{} `json:"bounce_rate_chart" dc:"Bounce Rate Chart"`
		OpenRateChart   interface{} `json:"open_rate_chart" dc:"Open Rate Chart"`
		ClickRateChart  interface{} `json:"click_rate_chart" dc:"Click Rate Chart"`
	} `json:"data" dc:"Data" dc:"Data"`
}
