package plugincrawler

import (
	"regexp"

	"github.com/zhs007/jarviscore"
)

const article36kr = "article.36kr"

var article36krRegexp *regexp.Regexp

func init() {
	article36krRegexp = regexp.MustCompile(`^https://36kr.com/p/([\w]+).html$`)
}

func parseArticle36kr(url string) *URLResult {
	params := article36krRegexp.FindStringSubmatch(url)

	// fmt.Printf("%v", params)

	if len(params) == 2 {
		return &URLResult{
			URLType: "article",
			URL:     url,
			PDF:     jarviscore.AppendString("36kr.", params[1], ".pdf"),
		}
	}

	return nil
}
