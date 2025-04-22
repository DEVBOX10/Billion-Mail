package email_template

import (
	"billionmail-core/internal/service/email_template"
	"billionmail-core/internal/service/public"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/email_template/v1"
)

func (c *ControllerV1) GetTemplate(ctx context.Context, req *v1.GetTemplateReq) (res *v1.GetTemplateRes, err error) {
	res = &v1.GetTemplateRes{}

	TempInfo, err := email_template.GetTemplatesByID(ctx, req.Id)
	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get template")))
		return
	}

	if TempInfo == nil || TempInfo.Id == 0 {
		res.Code = 404
		res.SetError(gerror.New(public.LangCtx(ctx, "Template not found")))
		return
	}

	res.Data = TempInfo
	res.SetSuccess(public.LangCtx(ctx, "Get template successfully"))
	return
}
