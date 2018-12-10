package plugincore

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/core/proto"
)

// cmdUsers - users
type cmdUsers struct {
}

// RunCommand - run command
func (cmd *cmdUsers) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	if params.CommandLine != nil {
		userscmd, ok := params.CommandLine.(*plugincorepb.UsersCommand)
		if !ok {
			return false
		}

		lst, err := params.ChatBot.GetChatBotDB().GetUsers(int(userscmd.Nums))
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return true
		}

		strret, err := chatbot.FormatJSONObj(lst)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
		} else {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret, params.Msg)
		}

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdUsers) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 2 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("users", pflag.ContinueOnError)

	var nums = flagset.Int32P("nums", "n", 128, "you need see numbers")

	err := flagset.Parse(params.LstStr[2:])
	if err != nil {
		return nil, err
	}

	return &plugincorepb.UsersCommand{
		Nums: *nums,
	}, nil
}
