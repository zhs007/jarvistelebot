package plugincore

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/core/proto"
)

// cmdUser - user
type cmdUser struct {
}

// RunCommand - run command
func (cmd *cmdUser) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		usercmd, ok := params.CommandLine.(*plugincorepb.UserCommand)
		if !ok {
			return false
		}

		var user *chatbotdbpb.User
		if usercmd.UserID != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUser(usercmd.UserID)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())

				return false
			}

			user = userbyuid
		} else if usercmd.UserName != "" {
			userbyuid, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(usercmd.UserName)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())

				return false
			}

			user = userbyuid
		}

		if user != nil {
			strret, err := chatbot.FormatJSONObj(user)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
			} else {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret)
			}

			return true
		}

		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "Sorry, I can't find this user.")

		return true
	}

	// if len(params.LstStr) == 3 {
	// 	lst, err := params.ChatBot.GetChatBotDB().GetUser(params.LstStr[2])
	// 	if err != nil {
	// 		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
	// 	}

	// 	strret, err := chatbot.FormatJSONObj(lst)
	// 	if err != nil {
	// 		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
	// 	} else {
	// 		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret)
	// 	}

	// 	return true
	// }

	return false
}

// Parse - parse command line
func (cmd *cmdUser) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 2 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("user", pflag.ContinueOnError)

	var uname = flagset.StringP("username", "n", "", "you can use username")

	err := flagset.Parse(params.LstStr[2:])
	if err != nil {
		return nil, err
	}

	if *uname == "" {
		args := flagset.Args()
		if len(args) == 1 {
			return &plugincorepb.UserCommand{
				UserID: args[0],
			}, nil
		}

		return nil, chatbot.ErrInvalidCommandLine
	}

	return &plugincorepb.UserCommand{
		UserName: *uname,
	}, nil
}
