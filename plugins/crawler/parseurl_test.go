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
