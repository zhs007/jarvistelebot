package pluginusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/usermgr/proto"
)

// cmdMyScripts - myscripts
type cmdMyScripts struct {
}

// RunCommand - run command
func (cmd *cmdMyScripts) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		msscmd, ok := params.CommandLine.(*pluginusermgrpb.MyScriptsCommand)
		if !ok {
			return false
		}

		// var user *chatbotdbpb.User
		// if usscmd.UserID != "" {
		// 	userbyuid, err := params.ChatBot.GetChatBotDB().GetUser(usscmd.UserID)
		// 	if err != nil {
		// 		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

		// 		return false
		// 	}

		// 	user = userbyuid
		// } else if usscmd.UserName != "" {
		// 	userbyuid, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(usscmd.UserName)
		// 	if err != nil {
		// 		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

		// 		return false
		// 	}

		// 	user = userbyuid
		// }

		// if user == nil {
		// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "Sorry, could not find this user.", params.Msg)

		// 	return false
		// }

		from := params.Msg.GetFrom()
		if from != nil {
			lst, err := params.ChatBot.GetChatBotDB().GetUserScripts(from.GetUserID(), msscmd.JarvisNodeName)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

				return false
			}

			strret, err := chatbot.FormatJSONObj(lst)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
			} else {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret, params.Msg)
			}

			return true
		}
	}

	return false
}

// Parse - parse command line
func (cmd *cmdMyScripts) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("myscripts", pflag.ContinueOnError)

	var nodename = flagset.StringP("nodename", "n", "", "you can use jarvis node name")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if nodename != nil {
		return &pluginusermgrpb.MyScriptsCommand{
			JarvisNodeName: *nodename,
		}, nil
	}

	return &pluginusermgrpb.MyScriptsCommand{}, nil

	// if *uid != "" || *uname != "" {
	// 	if *nodename != "" {
	// 		return &pluginusermgrpb.UserScriptsCommand{
	// 			UserID:         *uid,
	// 			UserName:       *uname,
	// 			JarvisNodeName: *nodename,
	// 		}, nil
	// 	}

	// 	return &pluginusermgrpb.UserScriptsCommand{
	// 		UserID:   *uid,
	// 		UserName: *uname,
	// 	}, nil
	// }

	// return chatbot.NewEmptyCommandLine("myscripts"), nil
}
