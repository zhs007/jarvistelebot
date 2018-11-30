package plugincore

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// cmdUsers - users
type cmdUsers struct {
}

// RunCommand - run command
func (cmd *cmdUsers) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	lst, err := params.ChatBot.GetChatBotDB().GetUsers(100)
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())

		return true
	}

	strret, err := chatbot.FormatJSONObj(lst)
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
	} else {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret)
	}

	return true
}

// Parse - parse command line
func (cmd *cmdUsers) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	return nil, nil
}
