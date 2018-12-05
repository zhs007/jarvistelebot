package pluginusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
)

// cmdUserScripts - userscripts
type cmdUserScripts struct {
}

// RunCommand - run command
func (cmd *cmdUserScripts) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		usscmd, ok := params.CommandLine.(*pluginusermgrpb.UserScriptsCommand)
		if !ok {
			return false
		}

		var user *chatbotdbpb.User
		if usscmd.UserID != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUser(usscmd.UserID)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())

				return false
			}

			user = userbyuid
		} else if usscmd.UserName != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(usscmd.UserName)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())

				return false
			}

			user = userbyuid
		}

		if user == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "Sorry, could not find this user.")

			return false
		}

		lst, err := params.ChatBot.GetChatBotDB().GetUserScripts(user.UserID, usscmd.JarvisNodeName)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())

			return false
		}

		strret, err := chatbot.FormatJSONObj(lst)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
		} else {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret)
		}

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdUserScripts) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 2 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("updscripts", pflag.ContinueOnError)

	var uid = flagset.StringP("userid", "i", "", "you can use userid")
	var uname = flagset.StringP("username", "u", "", "you can use username")
	var nodename = flagset.StringP("nodename", "n", "", "you can use jarvis node name")

	err := flagset.Parse(params.LstStr[2:])
	if err != nil {
		return nil, err
	}

	if *uid != "" || *uname != "" {
		if *nodename != "" {
			return &pluginusermgrpb.UserScriptsCommand{
				UserID:         *uid,
				UserName:       *uname,
				JarvisNodeName: *nodename,
			}, nil
		}

		return &pluginusermgrpb.UserScriptsCommand{
			UserID:   *uid,
			UserName: *uname,
		}, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
