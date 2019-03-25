package plugincrawler

import (
	"testing"
)

func TestParseArticlebaijingapp(t *testing.T) {

	type data struct {
		url string
		ret *URLResult
	}

	lst := []data{
		data{
			url: "http://www.baijingapp.com/article/22008",
			ret: &URLResult{
				URLType: "article",
				URL:     "http://www.baijingapp.com/article/22008",
				PDF:     "baijingapp.22008.pdf",
			},
		},
	}

	for i := 0; i < len(lst); i++ {
		ret := parseArticlebaijingapp(lst[i].url)
		if !isSame(ret, lst[i].ret) {
			t.Fatalf("TestParseArticlebaijingapp %v %v", lst[i], ret)

			return
		}
	}

	t.Logf("TestParseArticlebaijingapp OK")
}
