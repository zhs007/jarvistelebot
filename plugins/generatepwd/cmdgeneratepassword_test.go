package plugingeneratepwd

import (
	"testing"

	"github.com/zhs007/jarvistelebot/chatbot"

	"github.com/zhs007/jarvistelebot/plugins/generatepwd/proto"
)

func TestParseGeneratePasswordCmd(t *testing.T) {

	cmdDuckling := cmdGeneratePassword{}

	type data struct {
		cmd string
		ret *plugingeneratepwdpb.GeneratePassword
	}

	lst := []data{
		data{
			cmd: "generatepassword -m normal -l 16",
			ret: &plugingeneratepwdpb.GeneratePassword{
				Length: 16,
				Mode:   "normal",
			},
		},
		data{
			cmd: "generatepassword -m normal -l 32",
			ret: &plugingeneratepwdpb.GeneratePassword{
				Length: 32,
				Mode:   "normal",
			},
		},
		data{
			cmd: "generatepassword",
			ret: &plugingeneratepwdpb.GeneratePassword{
				Length: 16,
				Mode:   "normal",
			},
		},
	}

	for i := 0; i < len(lst); i++ {
		lstStr := chatbot.SplitString(lst[i].cmd)
		ret, err := cmdDuckling.parse(lstStr)
		if err != nil {
			t.Fatalf("TestParseGeneratePasswordCmd err %v %v", lst[i], err)
		} else {
			if ret.Length != lst[i].ret.Length || ret.Mode != lst[i].ret.Mode {
				t.Fatalf("TestParseGeneratePasswordCmd err %v %-v", lst[i], ret)
			}
		}
	}

	t.Logf("TestParseGeneratePasswordCmd OK")
}
