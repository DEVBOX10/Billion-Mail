// =====================================================
// Define the standard response format for the API.
// =====================================================

package api_v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// StandardRes response format
type StandardRes struct {
	g.Meta  `mime:"application/json"`
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg" dc:"Success message"`
	Data    interface{} `json:"data"`
}

func (resp *StandardRes) SetSuccess(msg string) {
	resp.Success = true
	resp.Msg = msg
}

func (resp *StandardRes) SetError(err error) {
	resp.Success = false
	resp.Msg = err.Error()
}

// StandardDownloadRes download response format
type StandardDownloadRes struct {
	g.Meta   `mime:"application/octet-stream"`
	Filename string `json:"filename" dc:"Full path of the downloaded file"`
}

// StandardPagination pagination data structure
type StandardPagination struct {
	Total int         `json:"total" dc:"count all"`
	List  interface{} `json:"list" dc:"data list for pagination"`
}
