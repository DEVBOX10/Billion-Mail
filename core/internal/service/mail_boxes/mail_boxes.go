package mail_boxes

import (
	v1 "billionmail-core/api/mail_boxes/v1"
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/GehirnInc/crypt/md5_crypt"
	"github.com/gogf/gf/v2/frame/g"
)

func Add(ctx context.Context, mailbox *v1.Mailbox) (err error) {
	// Encode password
	mailbox.PasswordEncode = PasswdEncode(ctx, mailbox.Password)

	// Crypt password
	mailbox.Password, err = PasswdMD5Crypt(ctx, mailbox.Password)

	if err != nil {
		err = fmt.Errorf("Generate password md5-crypt failed: %w", err)
		return
	}

	now := time.Now().Unix()
	mailbox.CreateTime = now
	mailbox.UpdateTime = now
	mailbox.Active = 1
	mailbox.Maildir = fmt.Sprintf("%s@%s/", mailbox.LocalPart, mailbox.Domain)

	_, err = g.DB().Model("mailbox").Ctx(ctx).Insert(mailbox)
	return err
}

func Update(ctx context.Context, mailbox *v1.Mailbox) (err error) {
	mailbox.UpdateTime = time.Now().Unix()
	if mailbox.Password != "" {
		mailbox.Password, err = PasswdMD5Crypt(ctx, mailbox.Password)

		if err != nil {
			err = fmt.Errorf("Generate password md5-crypt failed: %w", err)
			return
		}

		mailbox.PasswordEncode = PasswdEncode(ctx, mailbox.Password)
	}

	_, err = g.DB().Model("mailbox").
		Ctx(ctx).
		Where("username", mailbox.Username).
		Update(mailbox)
	return err
}

func Delete(ctx context.Context, email string) error {
	_, err := g.DB().Model("mailbox").
		Ctx(ctx).
		Where("username", email).
		Delete()
	return err
}

func Get(ctx context.Context, domain, keyword string, page, pageSize int) ([]v1.Mailbox, int, error) {
	m := g.DB().Model("mailbox").Order("create_time", "desc")

	if domain != "" {
		m.Where("domain", domain)
	}

	if keyword != "" {
		m = m.WhereLike("username", fmt.Sprintf("%%%s%%", keyword))
	}

	count, err := m.Count()
	if err != nil {
		return nil, 0, err
	}

	var mailboxes []v1.Mailbox
	err = m.Page(page, pageSize).Scan(&mailboxes)

	return mailboxes, count, err
}

func PasswdEncode(ctx context.Context, password string) string {
	return hex.EncodeToString([]byte(base64.StdEncoding.EncodeToString([]byte(password))))
}

func PasswdDecode(ctx context.Context, password string) (string, error) {
	decoded, err := hex.DecodeString(password)
	if err != nil {
		return "", err
	}
	decodedStr, err := base64.StdEncoding.DecodeString(string(decoded))
	if err != nil {
		return "", err
	}
	return string(decodedStr), nil
}

func PasswdMD5Crypt(ctx context.Context, password string) (string, error) {
	// use md5_crypt package to generate MD5-CRYPT hash
	crypter := md5_crypt.New()

	// Generate generates a hash using the given password and salt.
	// The salt is optional and can be empty.
	// If the salt is empty, a random salt will be generated.
	result, err := crypter.Generate([]byte(password), []byte(""))
	if err != nil {
		return "", err
	}

	return result, nil
}
