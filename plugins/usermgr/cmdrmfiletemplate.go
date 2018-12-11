package pluginusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
)

// cmdRmFileTemplate - rmfiletemplate
type cmdRmFileTemplate struct {
}

// RunCommand - run command
func (cmd *cmdRmFileTemplate) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		rmftcmd, ok := params.CommandLine.(*pluginusermgrpb.RemoveFileTemplateCommand)
		if !ok {
			return false
		}

		var user *chatbotdbpb.User
		if rmftcmd.UserID != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUser(rmftcmd.UserID)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

				return false
			}

			user = userbyuid
		} else if rmftcmd.UserName != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(rmftcmd.UserName)
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

		err := params.ChatBot.GetChatBotDB().RemoveFileTemplate(user.UserID, rmftcmd.FileTemplateName)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "OK, It is done.", params.Msg)

		params.ChatBot.OnUserEvent(ctx, params.ChatBot, chatbot.UserEventOnChgUserFileTemplate, user.UserID)

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdRmFileTemplate) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("rmscript", pflag.ContinueOnError)

	var uid = flagset.StringP("userid", "i", "", "you can use userid")
	var uname = flagset.StringP("username", "u", "", "you can use username")
	var filetemplatename = flagset.StringP("filetemplatename", "f", "", "you can remove file template name")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if (*uid != "" || *uname != "") && *filetemplatename != "" {
		return &pluginusermgrpb.RemoveFileTemplateCommand{
			UserID:           *uid,
			UserName:         *uname,
			FileTemplateName: *filetemplatename,
		}, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
