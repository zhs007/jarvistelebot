package plugincrawler

import (
	"testing"
)

func TestParseArticlelieyunwang(t *testing.T) {

	type data struct {
		url string
		ret *URLResult
	}

	lst := []data{
		data{
			url: "https://www.lieyunwang.com/archives/452856",
			ret: &URLResult{
				URLType: "article",
				URL:     "https://www.lieyunwang.com/archives/452856",
				PDF:     "lieyunwang.452856.pdf",
			},
		},
	}

	for i := 0; i < len(lst); i++ {
		ret := parseArticlelieyunwang(lst[i].url)
		if !isSame(ret, lst[i].ret) {
			t.Fatalf("TestParseArticlelieyunwang %v %v", lst[i], ret)

			return
		}
	}

	t.Logf("TestParseArticlelieyunwang OK")
}
