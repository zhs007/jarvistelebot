package pluginusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
)

// cmdImpScripts - impscripts
type cmdImpScripts struct {
}

// RunCommand - run command
func (cmd *cmdImpScripts) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		impscmd, ok := params.CommandLine.(*pluginusermgrpb.ImpScriptsCommand)
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

		lst, err := params.ChatBot.GetChatBotDB().GetUserScripts(user.UserID, "")
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		for _, v := range lst.Scripts {
			err := params.ChatBot.GetChatBotDB().RemoveUserScripts(user.UserID, v.ScriptName)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

				return false
			}
		}

		for _, v := range arr {
			_, ok1 := v["scriptname"]
			_, ok2 := v["jarvisnodename"]
			_, ok3 := v["scriptinfo"]
			if ok1 && ok2 && ok3 {
				scriptname := v["scriptname"].(string)
				jarvisnodename := v["jarvisnodename"].(string)
				scriptinfo := v["scriptinfo"].(string)

				us := &chatbotdbpb.UserScript{
					ScriptName:     scriptname,
					JarvisNodeName: jarvisnodename,
					File: &chatbotdbpb.File{
						Filename: scriptname,
						FileType: chatbot.FileTypeShellScript,
						Data:     []byte(scriptinfo),
					},
				}

				err := params.ChatBot.GetChatBotDB().SaveUserScript(user.UserID, us)
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
func (cmd *cmdImpScripts) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
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

	flagset := pflag.NewFlagSet("impscripts", pflag.ContinueOnError)

	var uid = flagset.StringP("userid", "i", "", "you can use userid")
	var uname = flagset.StringP("username", "u", "", "you can use username")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *uid != "" || *uname != "" {
		return &pluginusermgrpb.ImpScriptsCommand{
			UserID:   *uid,
			UserName: *uname,
		}, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
