package plugincore

import (
	"strings"
	"testing"

	"github.com/zhs007/jarvistelebot/plugins/core/proto"

	"go.uber.org/zap/zapcore"

	"github.com/zhs007/jarvistelebot/chatbot"
)

func Test_cmdUser_Parse(t *testing.T) {
	chatbot.InitLogger(zapcore.InfoLevel, true, "./")

	cmd := &cmdUser{}

	type parseResult struct {
		cmd      string
		userid   string
		username string
	}

	arrOK := []parseResult{
		parseResult{
			cmd:      "123 --username haha",
			userid:   "",
			username: "haha",
		},
		parseResult{
			cmd:      "-n haha 123",
			userid:   "",
			username: "haha",
		},
		parseResult{
			cmd:      "--username haha 123",
			userid:   "",
			username: "haha",
		},
		parseResult{
			cmd:      "123 -n haha",
			userid:   "",
			username: "haha",
		},
		parseResult{
			cmd:      "123 --username=haha",
			userid:   "",
			username: "haha",
		},
		parseResult{
			cmd:      "   123   ",
			userid:   "123",
			username: "",
		},
	}

	for _, v := range arrOK {
		lstfullcmd := []string{">", "user"}
		lstcmd := strings.Fields(v.cmd)

		// t.Logf("%v", len(lstcmd))

		msg, err := cmd.ParseCommandLine(&chatbot.MessageParams{
			LstStr: append(lstfullcmd, lstcmd...),
		})
		if err != nil {
			t.Fatalf("Test_cmdUser_Parse arrOK %v %v", v, err)
		}

		uc, ok := msg.(*plugincorepb.UserCommand)
		if !ok {
			t.Fatalf("Test_cmdUser_Parse arrOK msg err %v", uc)
		}

		if !(uc.UserID == v.userid && uc.UserName == v.username) {
			t.Fatalf("Test_cmdUser_Parse arrOK result %v %v", v, uc)
		}
	}

	t.Log("Test_cmdUser_Parse OK")
}
