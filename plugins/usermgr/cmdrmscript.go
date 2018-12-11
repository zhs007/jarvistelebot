package pluginusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
)

// cmdRmScript - rmscript
type cmdRmScript struct {
}

// RunCommand - run command
func (cmd *cmdRmScript) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		rmscmd, ok := params.CommandLine.(*pluginusermgrpb.RemoveScriptCommand)
		if !ok {
			return false
		}

		var user *chatbotdbpb.User
		if rmscmd.UserID != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUser(rmscmd.UserID)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

				return false
			}

			user = userbyuid
		} else if rmscmd.UserName != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(rmscmd.UserName)
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

		err := params.ChatBot.GetChatBotDB().RemoveUserScripts(user.UserID, rmscmd.ScriptName)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "OK, It is done.", params.Msg)

		params.ChatBot.OnUserEvent(ctx, params.ChatBot, chatbot.UserEventOnChgUserScript, user.UserID)

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdRmScript) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("rmscript", pflag.ContinueOnError)

	var uid = flagset.StringP("userid", "i", "", "you can use userid")
	var uname = flagset.StringP("username", "u", "", "you can use username")
	var scriptname = flagset.StringP("scriptname", "s", "", "you can remove script name")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if (*uid != "" || *uname != "") && *scriptname != "" {
		return &pluginusermgrpb.RemoveScriptCommand{
			UserID:     *uid,
			UserName:   *uname,
			ScriptName: *scriptname,
		}, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
