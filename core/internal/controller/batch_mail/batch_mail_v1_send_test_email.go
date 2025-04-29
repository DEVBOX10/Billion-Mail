package batch_mail

import (
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/batch_mail"
	"billionmail-core/internal/service/domains"
	"billionmail-core/internal/service/mail_service"
	"strings"

	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"billionmail-core/api/batch_mail/v1"
)

func (c *ControllerV1) SendTestEmail(ctx context.Context, req *v1.SendTestEmailReq) (res *v1.SendTestEmailRes, err error) {

	res = &v1.SendTestEmailRes{}
	sender, err := mail_service.NewEmailSenderWithLocal(req.Addresser)
	if err != nil {
		res.Code = 400
		res.SetError(gerror.New(public.LangCtx(ctx, "create email sender failed: : {}", err.Error())))
		return
	}
	defer sender.Close()

	var template entity.EmailTemplate

	err = g.DB().Model("email_templates").
		Where("id", req.TemplateId).
		Scan(&template)

	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "failed to get template {}", err)))
		return
	}

	// --------------------------------------
	if !strings.Contains(template.Content, "__UNSUBSCRIBE_URL__") {
		template.Content = public.AddUnsubscribeButton(template.Content)
	}

	domain := domains.GetBaseURL()

	// generate unsubscribe URL placeholder
	unsubscribeURL := fmt.Sprintf("%s/api/unsubscribe", domain)
	groupURL := fmt.Sprintf("%s/api/unsubscribe/user_group", domain)

	unsubscribeJumpURL := fmt.Sprintf("%s/unsubscribe.html?jwt=__JWT__&email=__EMAIL__&url_type=%s&url_unsubscribe=%s",
		domain, groupURL, unsubscribeURL)

	// replace unsubscribe link placeholder
	content := strings.ReplaceAll(template.Content, "__UNSUBSCRIBE_URL__", unsubscribeJumpURL)

	jwtToken, err := batch_mail.GenerateUnsubscribeJWT(
		req.Recipient,
		req.TemplateId,
		0,
	)

	if err != nil {
		g.Log().Error(ctx, "Failed to generate unsubscribe JWT: %v", err)
	} else {
		// replace JWT and email
		content = strings.ReplaceAll(content, "__JWT__", jwtToken)
		content = strings.ReplaceAll(content, "__EMAIL__", req.Recipient)
	}
	// --------------------------------------

	personalizedContent := content

	message := mail_service.NewMessage(req.Subject, personalizedContent)

	message.SetMessageID(sender.GenerateMessageID())

	// send email
	err = sender.Send(message, []string{req.Recipient})
	if err != nil {
		g.Log().Error(ctx, "send email to %s failed: %v", req.Recipient, err)
		res.SetError(gerror.New(public.LangCtx(ctx, "send email to %s failed: %v", req.Recipient, err)))
		return
	}
	res.SetSuccess(public.LangCtx(ctx, "send email successfully"))

	return
}
