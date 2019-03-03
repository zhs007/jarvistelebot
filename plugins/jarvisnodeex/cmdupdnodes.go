package pluginjarvisnodeex

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/jarvisnodeex/proto"
)

// cmdUpdNodes - updnodes
type cmdUpdNodes struct {
}

// RunCommand - run command
func (cmd *cmdUpdNodes) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		updnodes, ok := params.CommandLine.(*pluginjarvisnodeexpb.UpdNodesCommand)
		if !ok {
			return false
		}

		params.ChatBot.GetJarvisNode().UpdateAllNodes(ctx, updnodes.NodeType, updnodes.NodeTypeVer,
			func(ctx context.Context, jarvisnode jarviscore.JarvisNode, request *jarviscorepb.JarvisMsg,
				reply *jarviscorepb.JarvisMsg) (bool, error) {

				return true, nil
			})

		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "I get it.", params.Msg)

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdUpdNodes) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("updnodes", pflag.ContinueOnError)

	var nodetype = flagset.StringP("nodetype", "t", "", "you want update nodetype")
	var ver = flagset.StringP("version", "v", "", "you want update to the version")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *nodetype == "" || *ver == "" {
		return nil, chatbot.ErrInvalidCommandLine
	}

	return &pluginjarvisnodeexpb.UpdNodesCommand{
		NodeType:    *nodetype,
		NodeTypeVer: *ver,
	}, nil
}