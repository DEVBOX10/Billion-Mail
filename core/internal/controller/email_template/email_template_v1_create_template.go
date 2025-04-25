package email_template

import (
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/email_template/v1"
	"billionmail-core/internal/service/email_template"
)

func (c *ControllerV1) CreateTemplate(ctx context.Context, req *v1.CreateTemplateReq) (res *v1.CreateTemplateRes, err error) {
	res = &v1.CreateTemplateRes{}

	// Check if template name exists
	exists, err := email_template.CheckTemplateNameExists(ctx, req.TempName)
	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to check template name")))
		return
	}
	if exists {
		res.Code = 400
		res.SetError(gerror.New(public.LangCtx(ctx, "Template name already exists")))
		return
	}

	var content, render string
	if req.AddType == 0 { // upload
		if req.FileData == "" {
			res.Code = 400
			res.SetError(gerror.New(public.LangCtx(ctx, "File data is required for upload type")))
			return
		}
		// 直接使用传入的文件内容
		content = req.FileData

		// 验证内容不为空
		if content == "" {
			res.Code = 400
			res.SetError(gerror.New(public.LangCtx(ctx, "File content cannot be empty")))
			return
		}

		// Add unsubscribe button
		content = addUnsubscribeButton(content)
	} else { // Drag to generate
		if req.Content == "" || req.Render == "" {
			res.Code = 400
			res.SetError(gerror.New(public.LangCtx(ctx, "Content and render data are required for drag type")))
			return
		}
		content = req.Content
		render = req.Render
	}

	// Create template
	id, err := email_template.CreateTemplate(ctx, req.TempName, req.AddType, content, render)
	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to create template")))
		return
	}

	res.Data.Id = id
	res.SetSuccess(public.LangCtx(ctx, "Template created successfully"))
	return
}

// addUnsubscribeButton
func addUnsubscribeButton(content string) string {
	// Unsubscribe button HTML
	unsubscribeButton := `<div style="padding: 16px 0; text-align: center"><a href="__UNSUBSCRIBE_URL__" style="color: #ccc; font-size: 12px">Unsubscribe</a></div>`

	// If content already contains unsubscribe link, return directly
	if strings.Contains(content, "__UNSUBSCRIBE_URL__") {
		return content
	}

	// check if content has body tag
	hasBodyTag := strings.Contains(content, "</body>")
	hasHtmlTag := strings.Contains(content, "</html>")

	if hasBodyTag {
		return strings.Replace(content, "</body>", unsubscribeButton+"</body>", 1)
	}

	if hasHtmlTag {
		return strings.Replace(content, "</html>", unsubscribeButton+"</html>", 1)
	}

	// if content has no body tag or html tag, add a complete html framework
	return fmt.Sprintf(`<!DOCTYPE html><html><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"></head><body><div style="width: 100%%; max-width: 600px; margin: 0 auto;">%s%s</div></body></html>`, content, unsubscribeButton)
}
