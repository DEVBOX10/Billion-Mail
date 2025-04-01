package mail_boxes

import (
	v1 "billionmail-core/api/mail_boxes/v1"
	"context"
	"crypto/md5"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func Add(ctx context.Context, mailbox *v1.Mailbox) error {
	// 加密密码
	mailbox.PasswordEncode = fmt.Sprintf("%x", md5.Sum([]byte(mailbox.Password)))

	now := time.Now().Unix()
	mailbox.CreateTime = now
	mailbox.UpdateTime = now
	mailbox.Active = 1
	mailbox.Maildir = fmt.Sprintf("%s@%s/", mailbox.LocalPart, mailbox.Domain)

	_, err := g.DB().Model("mailbox").Ctx(ctx).Insert(mailbox)
	return err
}

func Update(ctx context.Context, mailbox *v1.Mailbox) error {
	mailbox.UpdateTime = time.Now().Unix()
	if mailbox.Password != "" {
		mailbox.PasswordEncode = fmt.Sprintf("%x", md5.Sum([]byte(mailbox.Password)))
	}

	_, err := g.DB().Model("mailbox").
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
	m := g.DB().Model("mailbox").Where("domain", domain)

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
