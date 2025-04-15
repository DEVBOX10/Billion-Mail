package contact

import (
	"billionmail-core/api/contact/v1"
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/contact"
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"os"
)

func (c *ControllerV1) CreateGroup(ctx context.Context, req *v1.CreateGroupReq) (res *v1.CreateGroupRes, err error) {
	res = &v1.CreateGroupRes{}

	// Check if group name exists
	exists, err := contact.CheckGroupNameExists(ctx, req.Name)
	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to check group name existence {}", err.Error())))
		return
	}
	if exists {
		res.Code = 400
		res.SetError(gerror.New(public.LangCtx(ctx, "The group name already exists")))
		return
	}

	// Create contact group
	groupId, err := contact.CreateGroup(ctx, req.Name, req.Description)
	if err != nil {
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to create group {}", err.Error())))
		return
	}

	// Handle contacts based on creation type
	switch req.CreateType {
	case 1: // Create group only
		break
	case 2: // Create group and add contacts
		if len(req.Contacts) > 0 {
			var contacts []*entity.Contact
			for _, email := range req.Contacts {
				contacts = append(contacts, &entity.Contact{
					Email:   email,
					GroupId: groupId,
					Active:  1, // Default subscription status
				})
			}
			// Batch create contacts, ignore duplicate entries
			err = contact.BatchCreateContacts(ctx, contacts)
			if err != nil {
				res.Code = 500
				res.SetError(gerror.New(public.LangCtx(ctx, "Failed to create contacts {}", err.Error())))
				return
			}
		}
	case 3: // Create group and import contacts
		if len(req.ImportFile) > 0 {
			fmt.Println("Creating group and importing contacts")
			// Read file content
			var fileContent []byte
			fileContent, err = os.ReadFile(req.ImportFile)
			if err != nil {
				fmt.Printf("Failed to read file: %v, file path: %s\n", err, req.ImportFile)
				res.Code = 500
				res.SetError(gerror.New(public.LangCtx(ctx, "Failed to read file {}", err.Error())))
				return
			}

			// Parse file and import contacts
			var contacts []*entity.Contact
			contacts, err = parseContactFile(fileContent, req.FileType)
			if err != nil {
				fmt.Printf("Failed to parse file: %v, file type: %s, content length: %d\n", err, req.FileType, len(fileContent))
				res.Code = 500
				res.SetError(gerror.New(public.LangCtx(ctx, "Failed to parse contact file {}", err.Error())))
				return
			}

			// Set group ID
			for _, contactGroupId := range contacts {
				contactGroupId.GroupId = groupId
			}

			fmt.Printf("File parsed successfully, contact count: %d\n", len(contacts))
			err = contact.BatchCreateContacts(ctx, contacts)
			if err != nil {
				res.Code = 500
				res.SetError(gerror.New(public.LangCtx(ctx, "Failed to import contacts {}", err.Error())))
				return
			}
		}
	}
	res.SetSuccess(public.LangCtx(ctx, "Group created successfully"))
	return
}
