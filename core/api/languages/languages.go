// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package languages

import (
	"context"

	"billionmail-core/api/languages/v1"
)

type ILanguagesV1 interface {
	SetLanguage(ctx context.Context, req *v1.SetLanguageReq) (res *v1.SetLanguageRes, err error)
	GetLanguage(ctx context.Context, req *v1.GetLanguageReq) (res *v1.GetLanguageRes, err error)
}
