package plugincrawler

import (
	"regexp"

	"github.com/zhs007/jarviscore"
)

const articlegeekpark = "article.geekpark"

var articlegeekparkRegexp *regexp.Regexp

func init() {
	articlegeekparkRegexp = regexp.MustCompile(`^https://www.geekpark.net/news/([\w]+)$`)
}

func parseArticlegeekpark(url string) *URLResult {
	params := articlegeekparkRegexp.FindStringSubmatch(url)

	// fmt.Printf("%v", params)

	if len(params) == 2 {
		return &URLResult{
			URLType: "article",
			URL:     url,
			PDF:     jarviscore.AppendString("geekpark.", params[1], ".pdf"),
		}
	}

	return nil
}
