package plugincrawler

import (
	"testing"
)

func TestParseArticletmtpost(t *testing.T) {

	type data struct {
		url string
		ret *URLResult
	}

	lst := []data{
		data{
			url: "https://www.tmtpost.com/3839950.html",
			ret: &URLResult{
				URLType: "article",
				URL:     "https://www.tmtpost.com/3839950.html",
				PDF:     "tmtpost.3839950.pdf",
			},
		},
	}

	for i := 0; i < len(lst); i++ {
		ret := parseArticletmtpost(lst[i].url)
		if !isSame(ret, lst[i].ret) {
			t.Fatalf("TestParseArticletmtpost %v %v", lst[i], ret)

			return
		}
	}

	t.Logf("TestParseArticletmtpost OK")
}
