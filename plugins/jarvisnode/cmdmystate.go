package pluginjarvisnode

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// cmdMyState - mystate
type cmdMyState struct {
}

// RunCommand - run command
func (cmd *cmdMyState) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	coredb := params.ChatBot.GetJarvisNodeCoreDB()

	str, _ := coredb.GetMyState()
	strret, err := chatbot.FormatJSON(str)
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), str, params.Msg)
		// params.ChatBot.SendMsg(params.Msg.GetFrom(), str)
	} else {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret, params.Msg)
		// params.ChatBot.SendMsg(params.Msg.GetFrom(), strret)
	}

	return true
}

// Parse - parse command line
func (cmd *cmdMyState) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) == 1 && params.LstStr[0] == "mystate" {
		return chatbot.NewEmptyCommandLine("mystate"), nil
	}

	return nil, chatbot.ErrMsgNotMine
}
