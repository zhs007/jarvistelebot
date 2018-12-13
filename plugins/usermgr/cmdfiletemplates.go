package pluginusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
)

// cmdFileTemplates - filetemplates
type cmdFileTemplates struct {
}

// RunCommand - run command
func (cmd *cmdFileTemplates) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		ftcmd, ok := params.CommandLine.(*pluginusermgrpb.FileTemplatesCommand)
		if !ok {
			return false
		}

		var user *chatbotdbpb.User
		if ftcmd.UserID != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUser(ftcmd.UserID)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

				return false
			}

			user = userbyuid
		} else if ftcmd.UserName != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(ftcmd.UserName)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

				return false
			}

			user = userbyuid
		}

		if user == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "Sorry, could not find this user.", params.Msg)

			return false
		}

		lst, err := params.ChatBot.GetChatBotDB().GetFileTemplates(user.UserID, ftcmd.JarvisNodeName, false)
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

	return false
}

// Parse - parse command line
func (cmd *cmdFileTemplates) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("updscripts", pflag.ContinueOnError)

	var uid = flagset.StringP("userid", "i", "", "you can use userid")
	var uname = flagset.StringP("username", "u", "", "you can use username")
	var nodename = flagset.StringP("nodename", "n", "", "you can use jarvis node name")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *uid != "" || *uname != "" {
		if *nodename != "" {
			return &pluginusermgrpb.FileTemplatesCommand{
				UserID:         *uid,
				UserName:       *uname,
				JarvisNodeName: *nodename,
			}, nil
		}

		return &pluginusermgrpb.FileTemplatesCommand{
			UserID:   *uid,
			UserName: *uname,
		}, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
