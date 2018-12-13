package pluginusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarviscore/base"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
	"go.uber.org/zap"
)

type fileTemplateInfo struct {
	filetemplatename string
	jarvisnodename   string
	fullpath         string
}

// cmdExpFileTemplates - expfiletemplates
type cmdExpFileTemplates struct {
}

// RunCommand - run command
func (cmd *cmdExpFileTemplates) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		eftcmd, ok := params.CommandLine.(*pluginusermgrpb.ExpFileTemplatesCommand)
		if !ok {
			return false
		}

		var user *chatbotdbpb.User
		if eftcmd.UserID != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUser(eftcmd.UserID)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

				return false
			}

			user = userbyuid
		} else if eftcmd.UserName != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(eftcmd.UserName)
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

		lst, err := params.ChatBot.GetChatBotDB().GetFileTemplates(user.UserID, "")
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		var lstobj []interface{}

		for _, v := range lst.Templates {
			o := fileTemplateInfo{
				filetemplatename: v.FileTemplateName,
			}

			ft, err := params.ChatBot.GetChatBotDB().GetFileTemplate(user.UserID, v.FileTemplateName)
			if err != nil {
				jarvisbase.Warn("cmdExpFileTemplates.RunCommand:GetUserScript", zap.Error(err))

				continue
			}

			o.jarvisnodename = ft.JarvisNodeName
			o.fullpath = ft.FullPath

			lstobj = append(lstobj, o)
		}

		buf, err := chatbot.Array2xlsx(lstobj)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
		} else {
			chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
				Filename: "filetemplates.xlsx",
				Data:     buf,
			})
		}

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdExpFileTemplates) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("updscripts", pflag.ContinueOnError)

	var uid = flagset.StringP("userid", "i", "", "you can use userid")
	var uname = flagset.StringP("username", "u", "", "you can use username")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *uid != "" || *uname != "" {

		return &pluginusermgrpb.ExpFileTemplatesCommand{
			UserID:   *uid,
			UserName: *uname,
		}, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
