package contact

import (
	v1 "billionmail-core/api/contact/v1"
	"billionmail-core/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
	"time"
)

func CreateGroup(ctx context.Context, name, description string) (int, error) {
	now := time.Now().Unix()
	data := g.Map{
		"name":        name,
		"description": description,
		"create_time": int(now),
		"update_time": int(now),
	}
	lastInsertId, err := g.DB().Model("contact_groups").Ctx(ctx).Data(data).InsertAndGetId()
	return int(lastInsertId), err
}

func GetGroup(ctx context.Context, id int) (*entity.ContactGroup, error) {
	var group entity.ContactGroup
	err := g.DB().Model("contact_groups").Ctx(ctx).Where("id", id).Scan(&group)
	return &group, err
}

func GetAllGroups(ctx context.Context, keyword string) ([]*v1.ContactGroup, error) {
	model := g.DB().Model("contact_groups").
		Fields("id, name, description, create_time, update_time").
		Order("create_time desc")

	// Add keyword search (group name or description)
	if keyword != "" {
		model = model.WhereLike("name", "%"+keyword+"%").
			WhereOrLike("description", "%"+keyword+"%")
	}

	var groups []*v1.ContactGroup
	err := model.Scan(&groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func UpdateGroup(ctx context.Context, id int, data g.Map) error {
	data["update_time"] = time.Now().Unix()
	_, err := g.DB().Model("contact_groups").Ctx(ctx).Data(data).Where("id", id).Update()
	return err
}

// DeleteContactsByGroupId
func DeleteContactsByGroupId(ctx context.Context, groupId int) error {
	_, err := g.DB().Model("contacts").Ctx(ctx).Where("group_id", groupId).Delete()
	return err
}

func DeleteGroup(ctx context.Context, id int) error {
	_, err := g.DB().Model("contact_groups").Ctx(ctx).Where("id", id).Delete()
	return err
}

func CreateContacts(ctx context.Context, contact *entity.Contact) error {
	now := time.Now().Unix()
	data := g.Map{
		"email":       contact.Email,
		"group_id":    contact.GroupId,
		"active":      contact.Active,
		"task_id":     contact.TaskId,
		"create_time": int(now),
	}
	_, err := g.DB().Model("contacts").Ctx(ctx).Data(data).InsertAndGetId()
	return err
}

func BatchCreateContacts(ctx context.Context, contacts []*entity.Contact) error {
	if len(contacts) == 0 {
		return nil
	}
	now := time.Now().Unix()
	var data []g.Map
	for _, contact := range contacts {
		data = append(data, g.Map{
			"email":       contact.Email,
			"group_id":    contact.GroupId,
			"active":      contact.Active,
			"task_id":     contact.TaskId,
			"create_time": int(now),
		})
	}

	_, err := g.DB().Model("contacts").Ctx(ctx).Data(data).InsertIgnore()
	return err
}

func GetContact(ctx context.Context, id int) (*entity.Contact, error) {
	var contact entity.Contact
	err := g.DB().Model("contacts").Ctx(ctx).Where("id", id).Scan(&contact)
	return &contact, err
}

func GetContactsByGroup(ctx context.Context, groupId int) ([]*entity.Contact, error) {
	var contacts []*entity.Contact
	err := g.DB().Model("contacts").Ctx(ctx).Where("group_id", groupId).Scan(&contacts)
	return contacts, err
}

func UpdateContact(ctx context.Context, id int, data g.Map) error {
	_, err := g.DB().Model("contacts").Ctx(ctx).Data(data).Where("id", id).Update()
	return err
}

func DeleteContact(ctx context.Context, id int) error {
	_, err := g.DB().Model("contacts").Ctx(ctx).Where("id", id).Delete()
	return err
}

func BatchDeleteContacts(ctx context.Context, ids []int) error {
	_, err := g.DB().Model("contacts").Ctx(ctx).Where("id", ids).Delete()
	return err
}

func CountContactsByGroup(ctx context.Context, groupId int, active int) (int, error) {
	model := g.DB().Model("contacts").Ctx(ctx).Where("group_id", groupId)
	if active >= 0 {
		model = model.Where("active", active)
	}
	return model.Count()
}

func CheckGroupNameExists(ctx context.Context, name string) (bool, error) {
	count, err := g.DB().Model("contact_groups").Ctx(ctx).Where("name", name).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Get contact list (with pagination and search)
func ContactsGroupWithPage(ctx context.Context, page, pageSize int, keyword string) ([]*entity.ContactGroup, int, error) {

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	keyword = strings.TrimSpace(keyword)

	model := g.DB().Model("contact_groups").Ctx(ctx)

	if keyword != "" {
		model = model.WhereLike("name", "%"+keyword+"%").
			WhereOrLike("description", "%"+keyword+"%")
	}

	total, err := model.Count()
	if err != nil {
		return nil, 0, err
	}

	// limit
	var groups []*entity.ContactGroup
	err = model.Page(page, pageSize).
		OrderDesc("create_time").
		Scan(&groups)

	return groups, total, err
}

// Get contacts in group (with pagination)
func GetContactsByGroupWithPage(ctx context.Context, groupId int, page, pageSize int) ([]*entity.Contact, int, error) {
	model := g.DB().Model("contacts").Ctx(ctx).Where("group_id", groupId)

	// Get total
	total, err := model.Count()
	if err != nil {
		return nil, 0, err
	}

	// Pagination query
	var contacts []*entity.Contact
	err = model.Page(page, pageSize).Scan(&contacts)
	return contacts, total, err
}

// Batch update contact status
func BatchUpdateContactStatus(ctx context.Context, groupId int, active int) error {
	_, err := g.DB().Model("contacts").Ctx(ctx).
		Where("group_id", groupId).
		Data(g.Map{"active": active}).
		Update()
	return err
}

// GetContactsWithPage
func GetContactsWithPage(ctx context.Context, page, pageSize int, groupId int, keyword string, status int) (total int, list []*entity.Contact, err error) {
	model := g.DB().Model("contacts").Safe()

	// Add query conditions
	if groupId > 0 {
		model = model.Where("group_id", groupId)
	}

	if len(keyword) > 0 {
		keywordPattern := "%" + keyword + "%"
		model = model.WhereLike("email", keywordPattern).WhereOrLike("full_name", keywordPattern)
	}

	if status != -1 {
		model = model.Where("active", status)
	}

	// Get total
	total, err = model.Count()
	if err != nil {
		return 0, nil, err
	}

	// If no data, return directly
	if total == 0 {
		return 0, make([]*entity.Contact, 0), nil
	}

	// Pagination query, create time descending
	list = make([]*entity.Contact, 0)
	err = model.Page(page, pageSize).
		Order("create_time desc").
		Scan(&list)
	if err != nil {
		return 0, nil, err
	}

	return total, list, nil
}
