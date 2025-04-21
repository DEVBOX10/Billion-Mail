package batch_mail

import (
	"billionmail-core/api/batch_mail/v1"
	"billionmail-core/internal/service/batch_mail"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func (c *ControllerV1) Unsubscribe(ctx context.Context, req *v1.UnsubscribeReq) (res *v1.UnsubscribeRes, err error) {
	res = &v1.UnsubscribeRes{}

	claims, err := batch_mail.ParseUnsubscribeJWT(req.Jwt)
	if err != nil {
		res.Success = false
		res.Message = "Invalid token"
		return res, nil
	}
	// 获取当前用户的组ID
	GroupId := req.GroupId

	// 更新联系人状态
	_, err = g.DB().Model("contacts").
		Where("email", claims.Email).
		WhereIn("group_id", GroupId).
		Data(g.Map{"active": 0}).
		Update()
	if err != nil {
		res.Success = false
		res.Message = "Failed to update contact status"
		return res, nil
	}

	// 记录退订详情 每个组ID都要记录
	for _, groupId := range GroupId {
		_, err = g.DB().Model("unsubscribe_records").Insert(g.Map{
			"email":            claims.Email,
			"group_id":         groupId,
			"template_id":      claims.TemplateId,
			"task_id":          claims.TaskId,
			"unsubscribe_time": time.Now().Unix(),
		})
		if err != nil {
			res.Success = false
			res.Message = "Failed to record unsubscribe"
			return res, nil
		}
	}

	//_, err = g.DB().Model("unsubscribe_records").Insert(g.Map{
	//	"email":            claims.Email,
	//	"group_id":         req.GroupId,
	//	"template_id":      claims.TemplateId,
	//	"task_id":          claims.TaskId,
	//	"unsubscribe_time": time.Now().Unix(),
	//})

	res.Success = true
	res.Message = "Successfully unsubscribed"
	return res, nil
}
