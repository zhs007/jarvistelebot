package plugincrawler

import (
	"regexp"

	"github.com/zhs007/jarviscore"
)

const articlebaijingapp = "article.baijingapp"

var articlebaijingappRegexp *regexp.Regexp

func init() {
	articlebaijingappRegexp = regexp.MustCompile(`^http://www.baijingapp.com/article/([\w]+)$`)
}

func parseArticlebaijingapp(url string) *URLResult {
	params := articlebaijingappRegexp.FindStringSubmatch(url)

	// fmt.Printf("%v", params)

	if len(params) == 2 {
		return &URLResult{
			URLType: "article",
			URL:     url,
			PDF:     jarviscore.AppendString("baijingapp.", params[1], ".pdf"),
		}
	}

	return nil
}
