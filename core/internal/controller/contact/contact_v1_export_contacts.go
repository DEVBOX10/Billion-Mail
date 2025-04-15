// core/internal/controller/contact/contact_v1_export_contacts.go
package contact

import (
	"archive/zip"
	"billionmail-core/api/contact/v1"
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/contact"
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"io"
	"os"
	"path/filepath"
)

func (c *ControllerV1) ExportContacts(ctx context.Context, req *v1.ExportContactsReq) (res *v1.ExportContactsRes, err error) {
	res = &v1.ExportContactsRes{}

	// Set default value for including inactive contacts
	if req.IncludeInactive == false {
		// Default to true if not explicitly set in request
		req.IncludeInactive = true
	}

	// Create temporary directory
	tempDir := filepath.Join(gfile.MainPkgPath(), "temp", "exports")
	if err = gfile.Mkdir(tempDir); err != nil {
		fmt.Printf("Failed to create directory: %v, path: %s\n", err, tempDir)
		res.Code = 500
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to create temp directory")))
		return
	}

	if req.ExportType == 1 { // Merged export
		//fmt.Println("Exporting merged contacts")
		// Get contacts from all groups
		var allContacts []*entity.Contact
		for _, groupId := range req.GroupIds {
			contacts, err := contact.GetContactsByGroup(ctx, groupId)
			if err != nil {
				continue
			}
			// Filter contacts based on IncludeInactive
			for _, c := range contacts {
				if c.Active == 1 || req.IncludeInactive {
					allContacts = append(allContacts, c)
				}
			}
		}

		// Generate filename
		timestamp := gtime.Now().Format("YmdHis")
		fileName := fmt.Sprintf("contacts_export_%s.%s", timestamp, req.Format)
		filePath := filepath.Join(tempDir, fileName)
		fmt.Println("File path: ", filePath)

		// Export file based on format
		var fileContent []byte
		fileContent, err = exportContactFile(allContacts, req.Format)
		if err != nil {
			fmt.Println("Export file error: ", err)
			res.Code = 500
			res.SetError(gerror.New(public.LangCtx(ctx, "Failed to export contacts")))
			return
		}

		// Write file
		if err = gfile.PutBytes(filePath, fileContent); err != nil {
			fmt.Println("Write file error: ", err)
			res.Code = 500
			res.SetError(gerror.New(public.LangCtx(ctx, "Failed to save file")))
			return
		}

		res.Data.FileUrl = filePath

	} else { // Separate export
		var files []ExportFile
		totalContacts := 0

		// Create file for each group
		for _, groupId := range req.GroupIds {
			group, err := contact.GetGroup(ctx, groupId)
			if err != nil {
				continue
			}

			contacts, err := contact.GetContactsByGroup(ctx, groupId)
			if err != nil {
				continue
			}

			// Filter contacts based on IncludeInactive
			var filteredContacts []*entity.Contact
			for _, c := range contacts {
				if c.Active == 1 || req.IncludeInactive {
					filteredContacts = append(filteredContacts, c)
				}
			}

			if len(filteredContacts) == 0 {
				continue
			}

			// Generate filename
			fileName := fmt.Sprintf("%s.%s", group.Name, req.Format)
			filePath := filepath.Join(tempDir, fileName)

			// Export file
			fileContent, err := exportContactFile(filteredContacts, req.Format)
			if err != nil {
				continue
			}

			// Write file
			if err = gfile.PutBytes(filePath, fileContent); err != nil {
				continue
			}

			files = append(files, ExportFile{
				Path:     filePath,
				Name:     fileName,
				Contacts: len(filteredContacts),
			})
			totalContacts += len(filteredContacts)
		}

		if len(files) == 0 {
			// fmt.Println("No contacts found")
			res.Code = 400
			res.SetError(gerror.New(public.LangCtx(ctx, "No contacts to export")))
			return
		}

		// Create ZIP file
		timestamp := gtime.Now().Format("YmdHis")
		zipFileName := fmt.Sprintf("contacts_export_%s.zip", timestamp)
		zipPath := filepath.Join(tempDir, zipFileName)

		if err = createZipFile(files, zipPath); err != nil {
			fmt.Println("Create ZIP file error: ", err)
			res.Code = 500
			res.SetError(gerror.New(public.LangCtx(ctx, "Failed to create zip file")))
			return
		}

		// Clean up temporary files
		for _, file := range files {
			os.Remove(file.Path)
		}

		res.Data.FileUrl = zipPath
	}

	res.SetSuccess(public.LangCtx(ctx, "Contacts exported successfully"))
	return
}

// ExportFile Export file information
type ExportFile struct {
	Path     string // File path
	Name     string // File name
	Contacts int    // Contact count
}

// createZipFile Create ZIP file
func createZipFile(files []ExportFile, zipPath string) error {
	// Create zip file
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %v", err)
	}
	defer zipFile.Close()

	// Create zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Iterate through source files
	for _, file := range files {
		// Open source file
		srcFile, err := os.Open(file.Path)
		if err != nil {
			return fmt.Errorf("failed to open source file %s: %v", file.Path, err)
		}
		defer srcFile.Close()

		// Create zip entry
		zipEntry, err := zipWriter.Create(file.Name)
		if err != nil {
			return fmt.Errorf("failed to create zip entry for %s: %v", file.Name, err)
		}

		// Copy file content to zip
		if _, err := io.Copy(zipEntry, srcFile); err != nil {
			return fmt.Errorf("failed to copy file content to zip for %s: %v", file.Name, err)
		}
	}

	return nil
}
