package plugincrawler

import (
	"regexp"

	"github.com/zhs007/jarviscore"
)

const articleSMZDM = "article.smzdm"

var articleSMZDMRegexp *regexp.Regexp

func init() {
	articleSMZDMRegexp = regexp.MustCompile(`^https://post.smzdm.com/p/([\w]+)/$`)
}

func parseArticleSMZDM(url string) *URLResult {
	params := articleSMZDMRegexp.FindStringSubmatch(url)

	// fmt.Printf("%v", params)

	if len(params) == 2 {
		return &URLResult{
			URLType: "article",
			URL:     url,
			PDF:     jarviscore.AppendString("smzdm.", params[1], ".pdf"),
		}
	}

	return nil
}
