package pluginusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
)

// cmdShowScript - showscript
type cmdShowScript struct {
}

// RunCommand - run command
func (cmd *cmdShowScript) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		sscmd, ok := params.CommandLine.(*pluginusermgrpb.ShowScriptCommand)
		if !ok {
			return false
		}

		var user *chatbotdbpb.User
		if sscmd.UserID != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUser(sscmd.UserID)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

				return false
			}

			user = userbyuid
		} else if sscmd.UserName != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(sscmd.UserName)
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

		us, err := params.ChatBot.GetChatBotDB().GetUserScript(user.UserID, sscmd.ScriptName)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
			"This script is prepared for "+us.JarvisNodeName, params.Msg)
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
			"The script data:", params.Msg)
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
			string(us.File.Data), params.Msg)

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdShowScript) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("showscript", pflag.ContinueOnError)

	var uid = flagset.StringP("userid", "i", "", "you can use userid")
	var uname = flagset.StringP("username", "u", "", "you can use username")
	var scriptname = flagset.StringP("scriptname", "s", "", "you can remove script name")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if (*uid != "" || *uname != "") && *scriptname != "" {
		return &pluginusermgrpb.ShowScriptCommand{
			UserID:     *uid,
			UserName:   *uname,
			ScriptName: *scriptname,
		}, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
