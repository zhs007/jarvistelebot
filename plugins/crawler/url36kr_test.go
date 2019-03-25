package plugincrawler

import (
	"testing"
)

func TestParseArticle36kr(t *testing.T) {

	type data struct {
		url string
		ret *URLResult
	}

	lst := []data{
		data{
			url: "https://36kr.com/p/5187453.html",
			ret: &URLResult{
				URLType: "article",
				URL:     "https://36kr.com/p/5187453.html",
				PDF:     "36kr.5187453.pdf",
			},
		},
	}

	for i := 0; i < len(lst); i++ {
		ret := parseArticle36kr(lst[i].url)
		if !isSame(ret, lst[i].ret) {
			t.Fatalf("TestParseArticle36kr %v %v", lst[i], ret)

			return
		}
	}

	t.Logf("TestParseArticle36kr OK")
}
