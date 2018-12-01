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

	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), fmt.Sprintf("%+v", lst))

	return true
}

// Parse - parse command line
func (cmd *cmdPlugins) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	return nil, nil
}
