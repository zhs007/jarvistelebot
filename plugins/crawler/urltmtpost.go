package plugincrawler

import (
	"regexp"

	"github.com/zhs007/jarviscore"
)

const articletmtpost = "article.tmtpost"

var articletmtpostRegexp *regexp.Regexp

func init() {
	articletmtpostRegexp = regexp.MustCompile(`^https://www.tmtpost.com/([\w]+).html$`)
}

func parseArticletmtpost(url string) *URLResult {
	params := articletmtpostRegexp.FindStringSubmatch(url)

	// fmt.Printf("%v", params)

	if len(params) == 2 {
		return &URLResult{
			URLType: "article",
			URL:     url,
			PDF:     jarviscore.AppendString("tmtpost.", params[1], ".pdf"),
		}
	}

	return nil
}
