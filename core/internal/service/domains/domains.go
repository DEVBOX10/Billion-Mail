package domains

import (
	v1 "billionmail-core/api/domains/v1"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func Add(ctx context.Context, domain *v1.Domain) error {
	domain.CreateTime = time.Now().Unix()
	domain.Active = 1

	_, err := g.DB().Model("domain").Ctx(ctx).Insert(domain)
	return err
}

func Update(ctx context.Context, domain *v1.Domain) error {
	_, err := g.DB().Model("domain").
		Ctx(ctx).
		Where("domain", domain.Domain).
		Update(domain)
	return err
}

func Delete(ctx context.Context, domainName string) error {
	_, err := g.DB().Model("domain").
		Ctx(ctx).
		Where("domain", domainName).
		Delete()
	return err
}

func Get(ctx context.Context, keyword string, page, pageSize int) ([]v1.Domain, int, error) {
	m := g.DB().Model("domain")

	if keyword != "" {
		m = m.WhereLike("domain", fmt.Sprintf("%%%s%%", keyword))
	}

	count, err := m.Count()
	if err != nil {
		return nil, 0, err
	}

	var domains []v1.Domain
	err = m.Page(page, pageSize).Scan(&domains)

	return domains, count, err
}
