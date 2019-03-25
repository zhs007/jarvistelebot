package plugincrawler

import (
	"testing"
)

func TestParseURL(t *testing.T) {

	urlParser := NewURLParser()

	type data struct {
		url string
		ret *URLResult
	}

	lst := []data{
		data{
			url: "https://post.smzdm.com/p/ag89qoem/",
			ret: &URLResult{
				URLType: "article",
				URL:     "https://post.smzdm.com/p/ag89qoem/",
				PDF:     "smzdm.ag89qoem.pdf",
			},
		},
		data{
			url: "https://36kr.com/p/5187453.html",
			ret: &URLResult{
				URLType: "article",
				URL:     "https://36kr.com/p/5187453.html",
				PDF:     "36kr.5187453.pdf",
			},
		},
		data{
			url: "https://www.huxiu.com/article/290658.html",
			ret: &URLResult{
				URLType: "article",
				URL:     "https://www.huxiu.com/article/290658.html",
				PDF:     "huxiu.290658.pdf",
			},
		},
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
		ret := urlParser.ParseURL(lst[i].url)
		if !isSame(ret, lst[i].ret) {
			t.Fatalf("TestParseURL %v %v", lst[i], ret)

			return
		}
	}

	t.Logf("TestParseURL OK")
}
