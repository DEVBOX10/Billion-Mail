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
	lastInsertId, err := g.DB().Model("bm_contact_groups").Ctx(ctx).Data(data).InsertAndGetId()
	return int(lastInsertId), err
}

func GetGroup(ctx context.Context, id int) (*entity.ContactGroup, error) {
	var group entity.ContactGroup
	err := g.DB().Model("bm_contact_groups").Ctx(ctx).Where("id", id).Scan(&group)
	return &group, err
}

func GetAllGroups(ctx context.Context, keyword string) ([]*v1.ContactGroup, error) {
	model := g.DB().Model("bm_contact_groups").
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
	_, err := g.DB().Model("bm_contact_groups").Ctx(ctx).Data(data).Where("id", id).Update()
	return err
}

func DeleteContactsByGroupId(ctx context.Context, groupId int) error {
	_, err := g.DB().Model("bm_contacts").Ctx(ctx).Where("group_id", groupId).Delete()
	return err
}

func DeleteGroup(ctx context.Context, id int) error {
	_, err := g.DB().Model("bm_contact_groups").Ctx(ctx).Where("id", id).Delete()
	return err
}

func BatchCreateContactsIgnoreDuplicate(ctx context.Context, contacts []*entity.Contact) (int, error) {
	if len(contacts) == 0 {
		return 0, nil
	}

	// Set batch size to avoid PostgreSQL limitations
	const batchSize = 1000

	// Calculate total batches
	totalBatches := (len(contacts) + batchSize - 1) / batchSize
	g.Log().Debug(ctx, "Batch creating %d contacts in %d batches (batch size: %d)",
		len(contacts), totalBatches, batchSize)

	// Process in batches
	totalAffected := 0
	now := time.Now().Unix()

	for i := 0; i < totalBatches; i++ {
		// Calculate start and end indices for current batch
		startIdx := i * batchSize
		endIdx := (i + 1) * batchSize
		if endIdx > len(contacts) {
			endIdx = len(contacts)
		}

		// Get current batch
		currentBatch := contacts[startIdx:endIdx]

		// Prepare data for current batch
		var data []g.Map
		for _, contact := range currentBatch {
			data = append(data, g.Map{
				"email":       contact.Email,
				"group_id":    contact.GroupId,
				"active":      contact.Active,
				"task_id":     contact.TaskId,
				"create_time": int(now),
			})
		}

		// Insert current batch
		result, err := g.DB().Model("bm_contacts").Ctx(ctx).Data(data).InsertIgnore()
		if err != nil {
			g.Log().Error(ctx, "Failed to insert batch %d/%d: %v", i+1, totalBatches, err)
			return totalAffected, err
		}

		// Count affected rows
		affected, err := result.RowsAffected()
		if err != nil {
			g.Log().Warning(ctx, "Could not get affected rows for batch %d/%d: %v", i+1, totalBatches, err)
		} else {
			totalAffected += int(affected)
			g.Log().Debug(ctx, "Batch %d/%d inserted %d records successfully", i+1, totalBatches, affected)
		}
	}

	g.Log().Info(ctx, "Total %d contacts inserted successfully", totalAffected)
	return totalAffected, nil
}

// BatchCreateContacts creates contacts in batch
func BatchCreateContacts(ctx context.Context, contacts []*entity.Contact) error {
	_, err := BatchCreateContactsIgnoreDuplicate(ctx, contacts)
	return err
}

func GetContactsByGroup(ctx context.Context, groupId int) ([]*entity.Contact, error) {
	var contacts []*entity.Contact
	err := g.DB().Model("bm_contacts").Ctx(ctx).Where("group_id", groupId).Scan(&contacts)
	return contacts, err
}

func CountContactsByGroup(ctx context.Context, groupId int, active int) (int, error) {
	model := g.DB().Model("bm_contacts").Ctx(ctx).Where("group_id", groupId)
	if active >= 0 {
		model = model.Where("active", active)
	}
	return model.Count()
}

func CheckGroupNameExists(ctx context.Context, name string) (bool, error) {
	count, err := g.DB().Model("bm_contact_groups").Ctx(ctx).Where("name", name).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func ContactsGroupWithPage(ctx context.Context, page, pageSize int, keyword string) ([]*entity.ContactGroup, int, error) {

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	keyword = strings.TrimSpace(keyword)

	model := g.DB().Model("bm_contact_groups").Ctx(ctx)

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

func GetContactsWithPage(ctx context.Context, page, pageSize int, groupId int, keyword string, status int) (total int, list []*entity.Contact, err error) {
	model := g.DB().Model("bm_contacts").Safe()

	// Add query conditions
	if groupId > 0 {
		model = model.Where("group_id", groupId)
	}

	if len(keyword) > 0 {
		keywordPattern := "%" + keyword + "%"
		model = model.WhereLike("email", keywordPattern)
	}

	if status != -1 {
		model = model.Where("active", status)
	}

	// Get total with distinct email
	value, err := model.Fields("COUNT(DISTINCT email) as count").Value()
	if err != nil {
		return 0, nil, err
	}
	total = value.Int()

	// If no data, return directly
	if total == 0 {
		return 0, make([]*entity.Contact, 0), nil
	}

	// Construct the query conditions
	conditions := g.Map{}
	if groupId > 0 {
		conditions["group_id"] = groupId
	}
	if status != -1 {
		conditions["active"] = status
	}

	subQuery := g.DB().Model("bm_contacts").
		Fields("DISTINCT ON (email) id")

	if len(conditions) > 0 {
		subQuery = subQuery.Where(conditions)
	}
	if len(keyword) > 0 {
		keywordPattern := "%" + keyword + "%"
		subQuery = subQuery.WhereLike("email", keywordPattern)
	}

	subQuery = subQuery.Order("email, create_time DESC")

	list = make([]*entity.Contact, 0)
	err = g.DB().Model("contacts c").
		Where("c.id IN (?)", subQuery).
		Order("c.create_time DESC").
		Page(page, pageSize).
		Scan(&list)
	if err != nil {
		return 0, nil, err
	}

	return total, list, nil
}

// Gets the group information to which the contact belongs
func GetContactGroupsInfo(ctx context.Context, email string, status int) ([]*entity.ContactGroup, error) {
	var groups []*entity.ContactGroup

	err := g.DB().Model("contacts c").
		LeftJoin("bm_contact_groups g", "c.group_id = g.id").
		Fields("DISTINCT g.id, g.name, g.description, g.create_time, g.update_time").
		Where("c.email", email).
		Where("c.active", status).
		Scan(&groups)

	if err != nil {
		return nil, err
	}

	return groups, nil
}

func GetContactsTrend(ctx context.Context, startTime, endTime time.Time) ([]*struct {
	Month            string `json:"month"`
	SubscribeCount   int    `json:"subscribe_count"`
	UnsubscribeCount int    `json:"unsubscribe_count"`
}, error) {
	var trends []*struct {
		Month            string `json:"month"`
		SubscribeCount   int    `json:"subscribe_count"`
		UnsubscribeCount int    `json:"unsubscribe_count"`
	}

	err := g.DB().Model("bm_contacts").
		Fields(
			"to_char(to_timestamp(create_time), 'YYYY-MM') as month",
			"SUM(CASE WHEN active = 1 THEN 1 ELSE 0 END) as subscribe_count",
			"SUM(CASE WHEN active = 0 THEN 1 ELSE 0 END) as unsubscribe_count",
		).
		Where("create_time BETWEEN ? AND ?", startTime.Unix(), endTime.Unix()).
		Group("month").
		Order("month ASC").
		Scan(&trends)

	return trends, err
}

func UpdateContactsGroups(ctx context.Context, emails []string, status int, newGroupIds []int) (int, error) {
	// 1. delete old group relations
	_, err := g.DB().Model("bm_contacts").
		Where("email IN(?)", emails).
		Where("active", status).
		Delete()
	if err != nil {
		return 0, err
	}

	// 2. create new group relations
	now := time.Now().Unix()
	var data []g.Map
	for _, email := range emails {
		for _, groupId := range newGroupIds {
			data = append(data, g.Map{
				"email":       email,
				"group_id":    groupId,
				"active":      status,
				"create_time": now,
			})
		}
	}

	if len(data) == 0 {
		return 0, nil
	}

	result, err := g.DB().Model("bm_contacts").
		Data(data).
		Insert()
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	return int(affected), err
}

// GetGroupContactCount calculates the total number of active contacts in multiple groups
func GetGroupContactCount(ctx context.Context, groupIds []int) (int, error) {
	if len(groupIds) == 0 {
		return 0, nil
	}

	// Query the total number of unique contacts across the specified groups
	var result struct {
		Count int `json:"count"`
	}

	err := g.DB().Model("bm_contacts").
		Fields("COUNT(DISTINCT email) as count").
		WhereIn("group_id", groupIds).
		Where("active", 1). // Only count active contacts
		Scan(&result)

	if err != nil {
		return 0, err
	}

	return result.Count, nil
}
