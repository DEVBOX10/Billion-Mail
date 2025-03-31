package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

type OverviewReq struct {
	g.Meta        `path:"/overview" tags:"Overview" method:"get" summary:"Get overview information" in:"query"`
	Authorization string   `json:"authorization" dc:"Authorization" in:"header"`
	Domains       []string `v:"required" dc:"Domains" dc:"Domains"`
	StartTime     int      `v:"required" dc:"StartTime" dc:"StartTime"`
	EndTime       int      `v:"required" dc:"EndTime" dc:"EndTime"`
}

type OverviewRes struct {
	api_v1.StandardRes
}
