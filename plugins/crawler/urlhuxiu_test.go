package plugincrawler

import (
	"testing"
)

func TestParseArticlehuxiu(t *testing.T) {

	type data struct {
		url string
		ret *URLResult
	}

	lst := []data{
		data{
			url: "https://www.huxiu.com/article/290658.html",
			ret: &URLResult{
				URLType: "article",
				URL:     "https://www.huxiu.com/article/290658.html",
				PDF:     "huxiu.290658.pdf",
			},
		},
	}

	for i := 0; i < len(lst); i++ {
		ret := parseArticlehuxiu(lst[i].url)
		if !isSame(ret, lst[i].ret) {
			t.Fatalf("TestParseArticlehuxiu %v %v", lst[i], ret)

			return
		}
	}

	t.Logf("TestParseArticlehuxiu OK")
}
