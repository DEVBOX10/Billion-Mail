package email_template

import (
	"billionmail-core/internal/service/email_template"
	"billionmail-core/internal/service/public"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/email_template/v1"
)

func (c *ControllerV1) ListTemplates(ctx context.Context, req *v1.ListTemplatesReq) (res *v1.ListTemplatesRes, err error) {

	res = &v1.ListTemplatesRes{}

	total, list, err := email_template.GetTemplatesWithPage(ctx, req.Page, req.PageSize, req.Keyword, req.AddType)
	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get template list")))
		return
	}

	res.Data.Total = total
	res.Data.List = list
	res.SetSuccess(public.LangCtx(ctx, "Get template list successfully"))
	return
}
