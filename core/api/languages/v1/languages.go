package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

type SetLanguageReq struct {
	g.Meta   `path:"/languages/set" method:"post" tags:"Languages" summary:"Set Language"`
	Language string `json:"language"  v:"required" dc:"Language"`
}

type SetLanguageRes struct {
	api_v1.StandardRes
}

type LanguageInfo struct {
	Name      string `json:"name"`
	Google    string `json:"google"`
	Title     string `json:"title"`
	LocalName string `json:"cn"`
}

type GetLanguageReq struct {
	g.Meta `path:"/languages/get" method:"get" tags:"Languages" summary:"Get Language"`
}

type GetLanguageRes struct {
	api_v1.StandardRes
	Data struct {
		CurrentLanguage    string         `json:"current_language"`
		AvailableLanguages []LanguageInfo `json:"available_languages"`
	} `json:"data"`
}
