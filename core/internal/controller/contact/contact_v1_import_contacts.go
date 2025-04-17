// core/internal/controller/contact/contact_v1_import_contacts.go
package contact

import (
	"billionmail-core/api/contact/v1"
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/contact"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"strings"
)

// parseEmailContent parse email content
func parseEmailContent(content string) []*entity.Contact {
	var contacts []*entity.Contact
	lines := strings.Split(strings.TrimSpace(content), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// parse email and status
		parts := strings.Split(line, ",")
		email := strings.TrimSpace(parts[0])
		active := 1 // default subscription status

		// if there is status information
		if len(parts) > 1 {
			status := strings.TrimSpace(parts[1])
			if status == "0" {
				active = 0
			}
		}

		if email != "" {
			contacts = append(contacts, &entity.Contact{
				Email:  email,
				Active: active,
			})
		}
	}
	return contacts
}

func (c *ControllerV1) ImportContacts(ctx context.Context, req *v1.ImportContactsReq) (res *v1.ImportContactsRes, err error) {
	res = &v1.ImportContactsRes{}

	if len(req.GroupIds) == 0 {
		res.Code = 400
		res.SetError(gerror.New(public.LangCtx(ctx, "Group IDs cannot be empty")))
		return
	}

	for _, groupId := range req.GroupIds {
		_, err = contact.GetGroup(ctx, groupId)
		if err != nil {
			res.Code = 400
			res.SetError(gerror.New(public.LangCtx(ctx, "Group {} does not exist", groupId)))
			return
		}
	}

	var contactList []*entity.Contact

	// handle content according to import type
	switch req.ImportType {
	case 1: // file upload
		if req.FileData == "" {
			res.Code = 400
			res.SetError(gerror.New(public.LangCtx(ctx, "File content cannot be empty")))
			return
		}
		contactList = parseEmailContent(req.FileData)

	case 2: // paste content
		if req.Contacts == "" {
			res.Code = 400
			res.SetError(gerror.New(public.LangCtx(ctx, "Contacts content cannot be empty")))
			return
		}
		contactList = parseEmailContent(req.Contacts)

	default:
		res.Code = 400
		res.SetError(gerror.New(public.LangCtx(ctx, "Invalid import type")))
		return
	}

	if len(contactList) == 0 {
		res.Code = 400
		res.SetError(gerror.New(public.LangCtx(ctx, "No valid contacts found")))
		return
	}

	// import contacts to each group
	successCount := 0
	for _, groupId := range req.GroupIds {
		// prepare contacts data for current group
		var groupContacts []*entity.Contact
		for _, c := range contactList {
			contactInfo := *c // create a copy
			contactInfo.GroupId = groupId
			groupContacts = append(groupContacts, &contactInfo)
		}

		// batch create new contacts, ignore duplicate data
		if len(groupContacts) > 0 {
			count, err := contact.BatchCreateContactsIgnoreDuplicate(ctx, groupContacts)
			if err != nil {
				res.Code = 500
				res.SetError(gerror.New(public.LangCtx(ctx, "Failed to import contacts: {}", err.Error())))
				return nil, err
			}
			successCount += count
		}
	}

	res.Data.ImportedCount = successCount
	res.SetSuccess(public.LangCtx(ctx, "Successfully imported {} contacts", successCount))
	return
}
