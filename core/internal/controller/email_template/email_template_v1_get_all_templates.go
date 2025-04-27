package email_template

import (
	"billionmail-core/internal/service/email_template"
	"billionmail-core/internal/service/public"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/email_template/v1"
)

func (c *ControllerV1) GetAllTemplates(ctx context.Context, req *v1.GetAllTemplatesReq) (res *v1.GetAllTemplatesRes, err error) {
	res = &v1.GetAllTemplatesRes{}

	TempAll, err := email_template.GetTemplatesAll(ctx)
	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get template list")))
		return
	}

	res.Data = TempAll
	res.SetSuccess(public.LangCtx(ctx, "Get template list successfully"))
	return
}
