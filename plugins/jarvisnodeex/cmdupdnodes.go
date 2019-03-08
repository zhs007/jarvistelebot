package pluginjarvisnodeex

import (
	"context"
	"encoding/json"
	"fmt"

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

		lastend := 0
		firstlog := false
		params.ChatBot.GetJarvisNode().UpdateAllNodes(ctx, updnodes.NodeType, updnodes.NodeTypeVer,
			func(ctx context.Context, jarvisnode jarviscore.JarvisNode, request *jarviscorepb.JarvisMsg,
				reply *jarviscorepb.JarvisMsg) (bool, error) {

				return true, nil
			}, func(ctx context.Context, jarvisnode jarviscore.JarvisNode, numsNode int, lstResult []*jarviscore.ClientGroupProcMsgResults) error {
				if !firstlog {
					chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
						fmt.Sprintf("The total number of nodes is %v.", numsNode), params.Msg)

					firstlog = true
				}

				curend := jarviscore.CountClientGroupProcMsgResultsEnd(lstResult)
				if curend == numsNode {
					str, err := json.MarshalIndent(lstResult, "", "\t")
					if err != nil {
						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
					} else {
						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), string(str), params.Msg)
					}

					chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "It's done.", params.Msg)
				} else {
					if curend > lastend {
						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							fmt.Sprintf("The %vth node has been completed.", curend), params.Msg)

						lastend = curend
					}
				}

				return nil
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
