package plugintranslate

import (
	"testing"

	"github.com/zhs007/jarvistelebot/plugins/translate/proto"
)

func TestParseTranslateCmd(t *testing.T) {
	type data struct {
		lststr []string
		cmd    *plugintranslatepb.TranslateCommand
		err    error
	}

	lst := []data{
		data{
			lststr: []string{"-s", "zh-CN", "-d", "en", "-p", "google", "-r", "true"},
			cmd: &plugintranslatepb.TranslateCommand{
				Platform: "google",
				SrcLang:  "zh-CN",
				DestLang: "en",
				Run:      true,
			},
			err: nil,
		},
	}

	for i := 0; i < len(lst); i++ {
		_, err := parseTranslateCmd(lst[i].lststr)
		if err != lst[i].err {
			t.Fatalf("TestParseTranslateCmd parseTranslateCmd %v - %v", lst[i], err)
		}

		// if cmd != lst[i].cmd {
		// 	t.Fatalf("TestParseTranslateCmd parseTranslateCmd %v - %v", lst[i], cmd)
		// }
	}

	t.Log("TestParseTranslateCmd OK")
}
