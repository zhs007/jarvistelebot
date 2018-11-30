package plugincore

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// PluginName - plugin name
const PluginName = "core"

// jarvisnodePlugin - jarvisnode plugin
type corePlugin struct {
	cmd *chatbot.CommandMap
}

// NewPlugin - new jarvisnode plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - jarvisnodePlugin")

	cmd := chatbot.NewCommandMap()

	cmd.AddCommand("version", &cmdVersion{})
	cmd.AddCommand("users", &cmdUsers{})
	cmd.AddCommand("user", &cmdUser{})

	p := &corePlugin{
		cmd: cmd,
	}

	return p, nil
}

// // RegPlugin - reg timestamp plugin
// func RegPlugin(cfgPath string, mgr chatbot.PluginsMgr) error {
// 	chatbot.Info("RegPlugin - jarvisnodePlugin")

// 	mgr.RegPlugin(newPlugin())

// 	return nil
// }

// OnMessage - get message
func (p *corePlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	// jarvisbase.Debug("jarvisnodePlugin.OnMessage", zap.String("params", fmt.Sprintf("%+v", params)))

	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if !params.ChatBot.IsMaster(from) {
		return false, nil
	}

	if len(params.LstStr) > 1 && params.LstStr[0] == ">" {
		p.cmd.Run(ctx, params.LstStr[1], params)

		return true, nil
	}

	return false, nil
}

// GetPluginName - get plugin name
func (p *corePlugin) GetPluginName() string {
	return PluginName
}

// // IsMyMessage
// func (p *corePlugin) IsMyMessage(params *chatbot.MessageParams) bool {
// 	if len(params.LstStr) >= 2 && params.LstStr[0] == ">" {
// 		return p.cmd.HasCommand(params.LstStr[1])
// 	}

// 	return false
// }

// OnStart - on start
func (p *corePlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *corePlugin) GetPluginType() int {
	return chatbot.PluginTypeCommand
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *corePlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) >= 2 && params.LstStr[0] == ">" {
		if p.cmd.HasCommand(params.LstStr[1]) {
			return p.cmd.ParseCommandLine(params.LstStr[1], params)
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
