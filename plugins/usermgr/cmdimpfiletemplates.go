package pluginusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
)

// cmdImpFileTemplates - impfiletemplates
type cmdImpFileTemplates struct {
}

// RunCommand - run command
func (cmd *cmdImpFileTemplates) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		impscmd, ok := params.CommandLine.(*pluginusermgrpb.ImpFileTemplatesCommand)
		if !ok {
			return false
		}

		file := params.Msg.GetFile()
		if file == nil {
			return false
		}

		if file.FileType != chatbot.FileExcel {
			return false
		}

		arr, err := chatbot.Xlsx2ArrayMap(file.Data)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		var user *chatbotdbpb.User
		if impscmd.UserID != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUser(impscmd.UserID)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

				return false
			}

			user = userbyuid
		} else if impscmd.UserName != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(impscmd.UserName)
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

		for _, v := range lst.Templates {
			err := params.ChatBot.GetChatBotDB().RemoveFileTemplate(user.UserID, v.FileTemplateName)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

				return false
			}
		}

		for _, v := range arr {
			_, ok1 := v["filetemplatename"]
			_, ok2 := v["jarvisnodename"]
			_, ok3 := v["fullpath"]
			if ok1 && ok2 && ok3 {
				filetemplatename := v["filetemplatename"].(string)
				jarvisnodename := v["jarvisnodename"].(string)
				fullpath := v["fullpath"].(string)

				us := &chatbotdbpb.UserFileTemplate{
					FileTemplateName: filetemplatename,
					JarvisNodeName:   jarvisnodename,
					FullPath:         fullpath,
				}

				err := params.ChatBot.GetChatBotDB().SaveFileTemplate(user.UserID, us)
				if err != nil {
					chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

					return false
				}
			}
		}

		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "OK, I get it.", params.Msg)

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdImpFileTemplates) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	file := params.Msg.GetFile()
	if file == nil {
		return nil, chatbot.ErrInvalidCommandLine
	}

	if file.FileType != chatbot.FileExcel {
		return nil, chatbot.ErrInvalidCommandLine
	}

	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("impfiletemplates", pflag.ContinueOnError)

	var uid = flagset.StringP("userid", "i", "", "you can use userid")
	var uname = flagset.StringP("username", "u", "", "you can use username")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *uid != "" || *uname != "" {
		return &pluginusermgrpb.ImpFileTemplatesCommand{
			UserID:   *uid,
			UserName: *uname,
		}, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
