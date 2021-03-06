package pluginusermgr

import (
	"context"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
)

type scriptInfo struct {
	scriptname     string
	jarvisnodename string
	scriptinfo     string
}

// cmdExpScripts - expscripts
type cmdExpScripts struct {
}

// RunCommand - run command
func (cmd *cmdExpScripts) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		usscmd, ok := params.CommandLine.(*pluginusermgrpb.ExpScriptsCommand)
		if !ok {
			return false
		}

		var user *chatbotdbpb.User
		if usscmd.UserID != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUser(usscmd.UserID)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

				return false
			}

			user = userbyuid
		} else if usscmd.UserName != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(usscmd.UserName)
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

		var lstobj []interface{}

		for _, v := range lst.Scripts {
			o := scriptInfo{
				scriptname: v.ScriptName,
			}

			us, err := params.ChatBot.GetChatBotDB().GetUserScript(user.UserID, v.ScriptName)
			if err != nil {
				jarvisbase.Warn("cmdExpScripts.RunCommand:GetUserScript", zap.Error(err))

				continue
			}

			o.jarvisnodename = us.JarvisNodeName

			if us.File != nil {
				o.scriptinfo = string(us.File.Data)
			}

			lstobj = append(lstobj, o)
		}

		buf, err := chatbot.Array2xlsx(lstobj)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
		} else {
			chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
				Filename: "scripts.xlsx",
				Data:     buf,
			}, params.Msg)
		}

		// strret, err := chatbot.FormatJSONObj(lst)
		// if err != nil {
		// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
		// } else {
		// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret, params.Msg)
		// }

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdExpScripts) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("expscripts", pflag.ContinueOnError)

	var uid = flagset.StringP("userid", "i", "", "you can use userid")
	var uname = flagset.StringP("username", "u", "", "you can use username")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *uid != "" || *uname != "" {
		return &pluginusermgrpb.ExpScriptsCommand{
			UserID:   *uid,
			UserName: *uname,
		}, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
