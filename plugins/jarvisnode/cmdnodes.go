package pluginjarvisnode

import (
	"context"
	"encoding/json"

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
	if params.CommandLine != nil {
		nodescmd, ok := params.CommandLine.(*pluginjarvisnodepb.NodesCommand)
		if !ok {
			return false
		}

		coredb := params.ChatBot.GetJarvisNodeCoreDB()

		lst, err := coredb.GetNodes(int(nodescmd.Nums))
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		jret, err := json.Marshal(lst)
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
	}

	return true
}

// Parse - parse command line
func (cmd *cmdNodes) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("nodes", pflag.ContinueOnError)

	var nums = flagset.Int32P("nums", "n", 128, "you need see numbers")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	return &pluginjarvisnodepb.NodesCommand{
		Nums: *nums,
	}, nil
}
