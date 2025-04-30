package abnormal_recipient

import (
	"billionmail-core/api/abnormal_recipient/v1"
	"billionmail-core/internal/service/abnormal_recipient"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) ListAbnormalRecipient(ctx context.Context, req *v1.ListAbnormalRecipientReq) (res *v1.ListAbnormalRecipientRes, err error) {
	res = &v1.ListAbnormalRecipientRes{}

	total, list, err := abnormal_recipient.GetListWithPage(ctx, req.Page, req.PageSize, req.Keyword, req.AddType)
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to get the exception recipient list: {}", err.Error())))
		return
	}

	apiList := make([]*v1.AbnormalRecipient, 0, len(list))
	for _, item := range list {
		apiList = append(apiList, &v1.AbnormalRecipient{
			Id:          item.Id,
			Recipient:   item.Recipient,
			CreateTime:  item.CreateTime,
			Description: item.Description,
			Count:       item.Count,
			AddType:     item.AddType,
		})
	}

	res.Data.Total = total
	res.Data.List = apiList

	res.SetSuccess(public.LangCtx(ctx, "The list of exception recipients was successfully obtained"))
	return
}
