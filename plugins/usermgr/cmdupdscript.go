package pluginusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
)

// cmdUpdScript - updscript
type cmdUpdScript struct {
}

// RunCommand - run command
func (cmd *cmdUpdScript) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	file := params.Msg.GetFile()
	if file != nil && params.CommandLine != nil {
		uscmd, ok := params.CommandLine.(*pluginusermgrpb.UpdScriptCommand)
		if !ok {
			return false
		}

		if file.FileType != chatbot.FileTypeShellScript {
			return false
		}

		var user *chatbotdbpb.User
		if uscmd.UserID != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUser(uscmd.UserID)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())

				return false
			}

			user = userbyuid
		} else if uscmd.UserName != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(uscmd.UserName)
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

		curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(uscmd.JarvisNodeName)
		if curnode == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "Sorry, I can't find this node.")

			return false
		}

		if user != nil {
			us := &chatbotdbpb.UserScript{
				ScriptName:     uscmd.ScriptName,
				JarvisNodeName: uscmd.JarvisNodeName,
				File:           file,
			}

			err := params.ChatBot.GetChatBotDB().SaveUserScript(user.UserID, us)

			// strret, err := chatbot.FormatJSONObj(user)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
			} else {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "OK, It is done.")
			}

			params.ChatBot.OnUserEvent(ctx, params.ChatBot, chatbot.UserEventOnChgUserScript, user.UserID)

			return true
		}

		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "Sorry, I can't find this user.")

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdUpdScript) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 2 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("updscript", pflag.ContinueOnError)

	var uid = flagset.StringP("userid", "i", "", "you can use userid")
	var uname = flagset.StringP("username", "u", "", "you can use username")
	var scriptname = flagset.StringP("scriptname", "s", "", "you can use scriptname")
	var nodename = flagset.StringP("nodename", "n", "", "you can use jarvis node name")

	err := flagset.Parse(params.LstStr[2:])
	if err != nil {
		return nil, err
	}

	if (*uid != "" || *uname != "") && *scriptname != "" && *nodename != "" {
		return &pluginusermgrpb.UpdScriptCommand{
			UserID:         *uid,
			UserName:       *uname,
			ScriptName:     *scriptname,
			JarvisNodeName: *nodename,
		}, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
