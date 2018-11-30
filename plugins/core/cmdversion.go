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
	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), params.ChatBot.GetVersion())

	return true
}

// Parse - parse command line
func (cmd *cmdVersion) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	return nil, nil
}
