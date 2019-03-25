package plugincrawler

import (
	"testing"
)

func TestParseArticleSMZDM(t *testing.T) {

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
		ret := parseArticleSMZDM(lst[i].url)
		if !isSame(ret, lst[i].ret) {
			t.Fatalf("TestParseArticleSMZDM %v %v", lst[i], ret)

			return
		}
	}

	t.Logf("TestParseArticleSMZDM OK")
}
