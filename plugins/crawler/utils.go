package plugincrawler

import (
	"bytes"
	"html/template"

	"go.uber.org/zap"

	"github.com/zhs007/jarviscore/base"
)

// UpdateCrawlerParam - the parameter for update crawler
type UpdateCrawlerParam struct {
	CrawlerPath string
}

// updateCrawler - update crawler
func updateCrawler(params *UpdateCrawlerParam, scriptUpd string) ([]byte, error) {
	tpl, err := template.New("updcrawler").Parse(scriptUpd)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	tpl.Execute(&b, params)

	jarvisbase.Info("updateCrawler script", zap.String("script", b.String()))

	return b.Bytes(), nil
}

// ExpArticleParam - the parameter for export article
type ExpArticleParam struct {
	CrawlerPath string
	URL         string
	PDF         string
}

// expArticle - export article
func expArticle(params *ExpArticleParam, scriptExp string) ([]byte, error) {
	tpl, err := template.New("exparticle").Parse(scriptExp)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	tpl.Execute(&b, params)

	jarvisbase.Info("exparticle script", zap.String("script", b.String()))

	return b.Bytes(), nil
}
