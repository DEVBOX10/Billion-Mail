package batch_mail

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"time"

	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/service/batch_mail"
)

func (c *ControllerV1) Unsubscribe(ctx context.Context, req *v1.UnsubscribeReq) (res *v1.UnsubscribeRes, err error) {
	res = &v1.UnsubscribeRes{}

	claims, err := batch_mail.ParseUnsubscribeJWT(req.Jwt)
	if err != nil {
		res.Success = false
		res.Message = "Invalid token"
		return res, nil
	}

	// 更新联系人状态
	_, err = g.DB().Model("contacts").
		Where("email", claims.Email).
		Where("group_id", claims.GroupId).
		Data(g.Map{"active": 0}).
		Update()
	if err != nil {
		res.Success = false
		res.Message = "Failed to update contact status"
		return res, nil
	}

	// 记录退订详情
	_, err = g.DB().Model("unsubscribe_records").Insert(g.Map{
		"email":            claims.Email,
		"group_id":         claims.GroupId,
		"template_id":      claims.TemplateId,
		"task_id":          claims.TaskId,
		"unsubscribe_time": time.Now().Unix(),
	})
	if err != nil {
		res.Success = false
		res.Message = "Failed to record unsubscribe"
		return res, nil
	}

	res.Success = true
	res.Message = "Successfully unsubscribed"
	return res, nil
}
