package plugincore

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// cmdPlugins - plugins
type cmdPlugins struct {
}

// RunCommand - run command
func (cmd *cmdPlugins) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	lst := params.ChatBot.GetPluginsMgr().GetPlugins()

	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), fmt.Sprintf("%+v", lst), params.Msg)

	return true
}

// Parse - parse command line
func (cmd *cmdPlugins) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) == 1 && params.LstStr[0] == "plugins" {
		return chatbot.NewEmptyCommandLine("plugins"), nil
	}

	return nil, chatbot.ErrMsgNotMine
}
