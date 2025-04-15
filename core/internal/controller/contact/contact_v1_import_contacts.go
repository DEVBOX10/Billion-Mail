// core/internal/controller/contact/contact_v1_import_contacts.go
package contact

import (
	"billionmail-core/api/contact/v1"
	"billionmail-core/internal/service/contact"
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"os"
)

func (c *ControllerV1) ImportContacts(ctx context.Context, req *v1.ImportContactsReq) (res *v1.ImportContactsRes, err error) {
	res = &v1.ImportContactsRes{}

	// Handle group-related logic
	groupId, err := c.handleGroup(ctx, req)
	if err != nil {
		return res, err
	}

	// Return directly if only creating new group
	if req.ImportType == 3 {
		res.SetSuccess(public.LangCtx(ctx, "Group created successfully"))
		return res, nil
	}

	// Read file content
	fileContent, err := os.ReadFile(req.File)
	if err != nil {
		fmt.Println("Failed to read file: ", err.Error(), "Path: ", req.File, "File type: ", req.FileType, "Group ID: ", groupId)
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to read file: {}", err.Error())))
		return
	}

	// Parse file content
	contacts, err := parseContactFile(fileContent, req.FileType)
	if err != nil {
		fmt.Println("Failed to parse file: ", err.Error(), "Path: ", req.File, "File type: ", req.FileType, "Group ID: ", groupId)
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to parse contact file: {}", err.Error())))
		return res, nil
	}

	// Set default status
	for _, email := range contacts {
		email.GroupId = groupId
		if email.Active < 0 {
			email.Active = req.DefaultActive
		}
	}

	// Import contacts
	err = contact.BatchCreateContacts(ctx, contacts)
	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to import contacts: {}", err.Error())))
		return
	}

	res.SetSuccess(public.LangCtx(ctx, "Contacts imported successfully"))
	return res, nil
}

// handleGroup Process group-related logic
func (c *ControllerV1) handleGroup(ctx context.Context, req *v1.ImportContactsReq) (groupId int, err error) {
	switch req.ImportType {
	case 1, 3: // Create new group and import OR Create group only
		// Check if group name exists
		exists, err := contact.CheckGroupNameExists(ctx, req.Name)
		if err != nil {
			return 0, gerror.New(public.LangCtx(ctx, "Failed to check group name: {}", err.Error()))
		}
		if exists {
			return 0, gerror.New(public.LangCtx(ctx, "The group name already exists"))
		}

		groupId, err = contact.CreateGroup(ctx, req.Name, req.Description)
		if err != nil {
			return 0, gerror.New(public.LangCtx(ctx, "Failed to create group: {}", err.Error()))
		}
		return groupId, nil

	case 2: // Import to existing group
		groupId = req.GroupId
		// Check if group exists
		_, err = contact.GetGroup(ctx, groupId)
		if err != nil {
			return 0, gerror.New(public.LangCtx(ctx, "Group does not exist"))
		}
		return groupId, nil

	default:
		return 0, gerror.New(public.LangCtx(ctx, "Invalid import type"))
	}
}
