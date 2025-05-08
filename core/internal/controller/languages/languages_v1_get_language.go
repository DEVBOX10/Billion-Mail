package languages

import (
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"billionmail-core/api/languages/v1"
)

func (c *ControllerV1) GetLanguage(ctx context.Context, req *v1.GetLanguageReq) (res *v1.GetLanguageRes, err error) {
	res = &v1.GetLanguageRes{}

	val, err := g.DB().Model("bm_options").
		Where("name", "CurrentLanguage").
		Value("value")

	if err == nil && val != nil && val.String() != "" {
		res.Data.CurrentLanguage = val.String()
	} else {

		res.Data.CurrentLanguage = public.GetLanguageFromCtx(ctx)
	}

	languageList := public.GetLanguageList()
	res.Data.AvailableLanguages = make([]v1.LanguageInfo, 0, len(languageList))

	for _, langInfo := range languageList {
		langCode, hasName := langInfo["name"].(string)
		langTitle, hasTitle := langInfo["title"].(string)
		langCN, hasCN := langInfo["cn"].(string)
		langGoogle, hasGoogle := langInfo["google"].(string)

		if hasName {
			info := v1.LanguageInfo{
				Name: langCode,
			}

			if hasTitle {
				info.Title = langTitle
			}

			if hasCN {
				info.LocalName = langCN
			}

			if hasGoogle {
				info.Google = langGoogle
			}

			res.Data.AvailableLanguages = append(res.Data.AvailableLanguages, info)
		}
	}

	res.SetSuccess(public.LangCtx(ctx, "Language information retrieved successfully"))
	return
}
