package pluginusermgr

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
)

// cmdShowFileTemplate - showfiletemplate
type cmdShowFileTemplate struct {
}

// RunCommand - run command
func (cmd *cmdShowFileTemplate) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		sftcmd, ok := params.CommandLine.(*pluginusermgrpb.ShowFileTemplateCommand)
		if !ok {
			return false
		}

		var user *chatbotdbpb.User
		if sftcmd.UserID != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUser(sftcmd.UserID)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())

				return false
			}

			user = userbyuid
		} else if sftcmd.UserName != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(sftcmd.UserName)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())

				return false
			}

			user = userbyuid
		}

		ft, err := params.ChatBot.GetChatBotDB().GetFileTemplate(user.UserID, sftcmd.FileTemplateName)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())

			return false
		}

		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
			fmt.Sprintf("If you send %v to me, I will send the file to %v:%v",
				ft.FileTemplateName, ft.JarvisNodeName, ft.FullPath))

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdShowFileTemplate) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 2 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("showfiletemplate", pflag.ContinueOnError)

	var uid = flagset.StringP("userid", "i", "", "you can use userid")
	var uname = flagset.StringP("username", "u", "", "you can use username")
	var filetemplatename = flagset.StringP("filetemplatename", "f", "", "you can show file template name")

	err := flagset.Parse(params.LstStr[2:])
	if err != nil {
		return nil, err
	}

	if (*uid != "" || *uname != "") && *filetemplatename != "" {
		return &pluginusermgrpb.ShowFileTemplateCommand{
			UserID:           *uid,
			UserName:         *uname,
			FileTemplateName: *filetemplatename,
		}, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
