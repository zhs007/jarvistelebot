package plugincrawler

import (
	"regexp"

	"github.com/zhs007/jarviscore"
)

const articlelieyunwang = "article.lieyunwang"

var articlelieyunwangRegexp *regexp.Regexp

func init() {
	articlelieyunwangRegexp = regexp.MustCompile(`^https://www.lieyunwang.com/archives/([\w]+)$`)
}

func parseArticlelieyunwang(url string) *URLResult {
	params := articlelieyunwangRegexp.FindStringSubmatch(url)

	// fmt.Printf("%v", params)

	if len(params) == 2 {
		return &URLResult{
			URLType: "article",
			URL:     url,
			PDF:     jarviscore.AppendString("lieyunwang.", params[1], ".pdf"),
		}
	}

	return nil
}
