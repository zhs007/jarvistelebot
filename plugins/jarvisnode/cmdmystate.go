package pluginjarvisnode

import (
	"context"
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// cmdMyState - mystate
type cmdMyState struct {
}

// RunCommand - run command
func (cmd *cmdMyState) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	coredb := params.ChatBot.GetJarvisNodeCoreDB()

	pd, err := coredb.GetMyData()
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

		return false
	}

	jret, err := json.Marshal(pd)
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

		return false
	}

	strret, err := chatbot.FormatJSON(string(jret))
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), string(jret), params.Msg)
	} else {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret, params.Msg)
	}

	nodestatus := params.ChatBot.GetJarvisNode().BuildStatus()

	jret, err = json.MarshalIndent(nodestatus, "", "  ")
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
	} else {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), string(jret), params.Msg)
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
