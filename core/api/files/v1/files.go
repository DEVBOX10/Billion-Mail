package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

//// UploadFileReq File upload request
//type UploadFileReq struct {
//	g.Meta        `path:"/file/upload" method:"post" tags:"Files" summary:"Upload file with support for chunked upload and base64 data"`
//	Authorization string            `json:"authorization" dc:"Authorization" in:"header" v:"required"`
//	FName         string            `json:"f_name"       v:"required" dc:"File name, no special characters allowed, max length 256 bytes"`
//	FPath         string            `json:"f_path"       v:"required" dc:"File save path, root path '/' and './' not allowed"`
//	FSize         int64             `json:"f_size"       v:"required" dc:"Total file size in bytes"`
//	FStart        int64             `json:"f_start"      dc:"Start position for resumable upload in bytes, used for validation"`
//	B64Data       string            `json:"b64_data"     dc:"Base64 encoded file content, mutually exclusive with blob parameter"`
//	Blob          *ghttp.UploadFile `json:"blob"         dc:"File content, mutually exclusive with b64_data parameter"`
//	DirMode       string            `json:"dir_mode"     dc:"Directory permission mode, e.g.: '0755'"`
//	FileMode      string            `json:"file_mode"    dc:"File permission mode, e.g.: '0644'"`
//}
//
//// UploadFileRes File upload response
//type UploadFileRes struct {
//	api_v1.StandardRes
//	Data struct {
//		FileId     string `json:"file_id"      dc:"Unique file identifier"`
//		UploadSize int64  `json:"upload_size"  dc:"Number of bytes uploaded"`
//		FileUrl    string `json:"file_url"     dc:"File access URL"`
//	} `json:"data"`
//}

// DownloadFileReq File download request
type DownloadFileReq struct {
	g.Meta        `path:"/file/download" method:"get" tags:"Files" summary:"Download file by path"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header" v:"required"`
	FilePath      string `json:"file_path"    v:"required" dc:"File path for download"`
}

// DownloadFileRes File download response
type DownloadFileRes struct {
	api_v1.StandardRes
}
