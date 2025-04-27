package files

import (
	"billionmail-core/api/files/v1"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"path/filepath"
)

func (c *ControllerV1) DownloadFile(ctx context.Context, req *v1.DownloadFileReq) (res *v1.DownloadFileRes, err error) {
	if err = c.validateFilePath(req.FilePath); err != nil {
		return nil, err
	}

	if !gfile.Exists(req.FilePath) {
		return nil, gerror.New(public.LangCtx(ctx, "File not found"))
	}

	fileName := filepath.Base(req.FilePath)

	r := g.RequestFromCtx(ctx)

	// set response headers
	r.Response.Header().Set("Content-Type", "application/octet-stream")
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+fileName)

	r.Response.ServeFileDownload(req.FilePath, fileName)

	return nil, nil
}
