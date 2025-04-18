package batch_mail

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"billionmail-core/api/batch_mail/v1"
)

func (c *ControllerV1) GetUserGroups(ctx context.Context, req *v1.GetUserGroupsReq) (res *v1.GetUserGroupsRes, err error) {
	res = &v1.GetUserGroupsRes{}

	var groups []*v1.GroupInfo
	err = g.DB().Model("contacts c").
		LeftJoin("contact_groups cg", "c.group_id = cg.id").
		Fields("cg.id, cg.name").
		Where("c.email", req.Email).
		Where("c.active", 1).
		Scan(&groups)
	if err != nil {
		return nil, err
	}

	res.Groups = groups
	return res, nil
}
