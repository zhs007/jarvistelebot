package plugincrawler

import (
	"regexp"

	"github.com/zhs007/jarviscore"
)

const articlehuxiu = "article.huxiu"

var articlehuxiuRegexp *regexp.Regexp

func init() {
	articlehuxiuRegexp = regexp.MustCompile(`^https://www.huxiu.com/article/([\w]+).html$`)
}

func parseArticlehuxiu(url string) *URLResult {
	params := articlehuxiuRegexp.FindStringSubmatch(url)

	// fmt.Printf("%v", params)

	if len(params) == 2 {
		return &URLResult{
			URLType: "article",
			URL:     url,
			PDF:     jarviscore.AppendString("huxiu.", params[1], ".pdf"),
		}
	}

	return nil
}
