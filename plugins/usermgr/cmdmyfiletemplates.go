package pluginusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
)

// cmdMyFileTemplates - myfiletemplates
type cmdMyFileTemplates struct {
}

// RunCommand - run command
func (cmd *cmdMyFileTemplates) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		myftcmd, ok := params.CommandLine.(*pluginusermgrpb.MyFileTemplatesCommand)
		if !ok {
			return false
		}

		from := params.Msg.GetFrom()
		if from != nil {
			lst, err := params.ChatBot.GetChatBotDB().GetFileTemplates(from.GetUserID(), myftcmd.JarvisNodeName, false)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

				return false
			}

			strret, err := chatbot.FormatJSONObj(lst)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
			} else {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret, params.Msg)
			}

			return true
		}
	}

	return false
}

// Parse - parse command line
func (cmd *cmdMyFileTemplates) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("myfiletemplates", pflag.ContinueOnError)

	var nodename = flagset.StringP("nodename", "n", "", "you can use jarvis node name")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *nodename != "" {
		return &pluginusermgrpb.MyFileTemplatesCommand{
			JarvisNodeName: *nodename,
		}, nil
	}

	return &pluginusermgrpb.MyFileTemplatesCommand{}, nil
}
