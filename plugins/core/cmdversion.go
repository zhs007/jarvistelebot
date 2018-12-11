package plugincore

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// cmdVersion - user
type cmdVersion struct {
}

// RunCommand - run command
func (cmd *cmdVersion) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), params.ChatBot.GetVersion(), params.Msg)

	return true
}

// Parse - parse command line
func (cmd *cmdVersion) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) == 1 && params.LstStr[0] == "version" {
		return chatbot.NewEmptyCommandLine("version"), nil
	}

	return nil, chatbot.ErrMsgNotMine
}
