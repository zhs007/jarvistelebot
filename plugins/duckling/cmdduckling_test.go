package pluginduckling

import (
	"testing"

	"github.com/zhs007/jarvistelebot/chatbot"

	"github.com/zhs007/jarvistelebot/plugins/duckling/proto"
)

func TestParseDucklingCmd(t *testing.T) {

	cmdDuckling := cmdRequestDuckling{}

	type data struct {
		cmd string
		ret *pluginducklingpb.RequestDuckling
	}

	lst := []data{
		data{
			cmd: "duckling -l zh_CN -t \"这个星期四要去看复联\"",
			ret: &pluginducklingpb.RequestDuckling{
				Lang: "zh_CN",
				Text: "这个星期四要去看复联",
			},
		},
		data{
			cmd: "duckling -l en_GB -t \"I am going to see the Union this Thursday.\"",
			ret: &pluginducklingpb.RequestDuckling{
				Lang: "en_GB",
				Text: "I am going to see the Union this Thursday.",
			},
		},
	}

	for i := 0; i < len(lst); i++ {
		lstStr := chatbot.SplitString(lst[i].cmd)
		ret, err := cmdDuckling.parse(lstStr)
		if err != nil {
			t.Fatalf("TestParseDucklingCmd err %v %v", lst[i], err)
		} else {
			if ret.Lang != lst[i].ret.Lang || ret.Text != lst[i].ret.Text {
				t.Fatalf("TestParseDucklingCmd err %v %-v", lst[i], ret)
			}
		}
	}

	t.Logf("TestParseDucklingCmd OK")
}
