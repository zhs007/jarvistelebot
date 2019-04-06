package plugintranslate

import (
	"testing"

	"github.com/zhs007/jarvistelebot/plugins/translate/proto"
)

func TestParseStartTranslateCmd(t *testing.T) {
	type data struct {
		lststr []string
		cmd    *plugintranslatepb.StartTranslateCommand
		err    error
	}

	lst := []data{
		data{
			lststr: []string{"-s", "zh-CN", "-d", "en", "-p", "google"},
			cmd: &plugintranslatepb.StartTranslateCommand{
				Platform: "google",
				SrcLang:  "zh-CN",
				DestLang: "en",
			},
			err: nil,
		},
		data{
			lststr: []string{"-s", "zh-CN", "-d", "en"},
			cmd: &plugintranslatepb.StartTranslateCommand{
				Platform: "google",
				SrcLang:  "zh-CN",
				DestLang: "en",
			},
			err: nil,
		},
	}

	for i := 0; i < len(lst); i++ {
		cmd, err := parseStartTranslateCmd(lst[i].lststr)
		if err != lst[i].err {
			t.Fatalf("TestParseStartTranslateCmd parseStartTranslateCmd %v - %v", lst[i], err)
		}

		if lst[i].cmd != nil && cmd == nil || lst[i].cmd == nil && cmd != nil {
			t.Fatalf("TestParseStartTranslateCmd parseTranslateCmd %v - %v", lst[i], cmd)
		} else if cmd != nil && lst[i].cmd != nil {
			if cmd.Platform != lst[i].cmd.Platform {
				t.Fatalf("TestParseStartTranslateCmd parseTranslateCmd %v %v - %v",
					lst[i].lststr, lst[i].cmd.Platform, cmd.Platform)
			}

			if cmd.SrcLang != lst[i].cmd.SrcLang {
				t.Fatalf("TestParseStartTranslateCmd parseTranslateCmd %v %v - %v",
					lst[i].lststr, lst[i].cmd.SrcLang, cmd.SrcLang)
			}

			if cmd.DestLang != lst[i].cmd.DestLang {
				t.Fatalf("TestParseStartTranslateCmd parseTranslateCmd %v %v - %v",
					lst[i].lststr, lst[i].cmd.DestLang, cmd.DestLang)
			}
		}

		// if cmd != lst[i].cmd {
		// 	t.Fatalf("TestParseTranslateCmd parseTranslateCmd %v - %v", lst[i], cmd)
		// }
	}

	t.Log("TestParseStartTranslateCmd OK")
}
