package pluginusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
)

// cmdUpdFileTemplate - updfiletemplate
type cmdUpdFileTemplate struct {
}

// RunCommand - run command
func (cmd *cmdUpdFileTemplate) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		uftcmd, ok := params.CommandLine.(*pluginusermgrpb.UpdFileTemplateCommand)
		if !ok {

			return false
		}

		var user *chatbotdbpb.User
		if uftcmd.UserID != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUser(uftcmd.UserID)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

				return false
			}

			user = userbyuid
		} else if uftcmd.UserName != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(uftcmd.UserName)
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

		curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(uftcmd.JarvisNodeName)
		if curnode == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "Sorry, I can't find this node.", params.Msg)

			return false
		}

		if user != nil {
			uft := &chatbotdbpb.UserFileTemplate{
				FileTemplateName: uftcmd.FileTemplateName,
				JarvisNodeName:   uftcmd.JarvisNodeName,
				FullPath:         uftcmd.FullPath,
			}

			err := params.ChatBot.GetChatBotDB().SaveFileTemplate(user.UserID, uft)

			// strret, err := chatbot.FormatJSONObj(user)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
			} else {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "OK, It is done.", params.Msg)
			}

			params.ChatBot.OnUserEvent(ctx, params.ChatBot, chatbot.UserEventOnChgUserFileTemplate, user.UserID)

			return true
		}

		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "Sorry, I can't find this user.", params.Msg)

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdUpdFileTemplate) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("updfiletemplate", pflag.ContinueOnError)

	var uid = flagset.StringP("userid", "i", "", "you can use userid")
	var uname = flagset.StringP("username", "u", "", "you can use username")
	var filetemplatename = flagset.StringP("filetemplatename", "f", "", "you can use file template name")
	var nodename = flagset.StringP("nodename", "n", "", "you can use jarvis node name")
	var fullpath = flagset.StringP("path", "p", "", "you can use full path")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if (*uid != "" || *uname != "") && *filetemplatename != "" && *nodename != "" && *fullpath != "" {
		return &pluginusermgrpb.UpdFileTemplateCommand{
			UserID:           *uid,
			UserName:         *uname,
			FileTemplateName: *filetemplatename,
			JarvisNodeName:   *nodename,
			FullPath:         *fullpath,
		}, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
