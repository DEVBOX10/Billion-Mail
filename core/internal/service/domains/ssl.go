package domains

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

// ApplyLetsEncryptCertWithHttp applies for a Let's Encrypt certificate for the given domain.
func ApplyLetsEncryptCertWithHttp(ctx context.Context, domain string, email string) error {
	return gerror.New("Not implemented")

	//crt, key, err := acme.ApplySSLWithExistingServer(ctx, []string{domain}, email, "http", "", nil, public.AbsPath("../ssl"))
	//
	//if err != nil {
	//	return err
	//}
	//
	//// Save the certificate and key to the database
	//g.DB().Model("letsencrypts").Insert()
}
