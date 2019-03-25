package plugincrawler

import (
	"testing"
)

func TestParseArticlegeekpark(t *testing.T) {

	type data struct {
		url string
		ret *URLResult
	}

	lst := []data{
		data{
			url: "https://www.geekpark.net/news/239623",
			ret: &URLResult{
				URLType: "article",
				URL:     "https://www.geekpark.net/news/239623",
				PDF:     "geekpark.239623.pdf",
			},
		},
	}

	for i := 0; i < len(lst); i++ {
		ret := parseArticlegeekpark(lst[i].url)
		if !isSame(ret, lst[i].ret) {
			t.Fatalf("TestParseArticlegeekpark %v %-v", lst[i], ret)

			return
		}
	}

	t.Logf("TestParseArticlegeekpark OK")
}
