package pluginjarvisnode

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/jarvisnode/proto"
)

// cmdNodes - nodes
type cmdNodes struct {
}

// RunCommand - run command
func (cmd *cmdNodes) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	coredb := params.ChatBot.GetJarvisNodeCoreDB()

	str, _ := coredb.GetNodes(100)
	strret, err := chatbot.FormatJSON(str)
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), str)
		// params.ChatBot.SendMsg(params.Msg.GetFrom(), str)
	} else {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret)
		// params.ChatBot.SendMsg(params.Msg.GetFrom(), strret)
	}

	return true
}

// Parse - parse command line
func (cmd *cmdNodes) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 2 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("nodes", pflag.ContinueOnError)

	var nums = flagset.Int32P("nums", "n", 128, "you need see numbers")

	err := flagset.Parse(params.LstStr[2:])
	if err != nil {
		return nil, err
	}

	return &pluginjarvisnodepb.NodesCommand{
		Nums: *nums,
	}, nil
}