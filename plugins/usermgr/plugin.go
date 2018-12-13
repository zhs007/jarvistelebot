package pluginusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// PluginName - plugin name
const PluginName = "usermgr"

// usermgrPlugin - usermgr plugin
type usermgrPlugin struct {
	cmd *chatbot.CommandMap
}

// NewPlugin - new jarvisnode plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - usermgrPlugin")

	cmd := chatbot.NewCommandMap()

	cmd.AddCommand("updscript", &cmdUpdScript{})
	cmd.AddCommand("userscripts", &cmdUserScripts{})
	cmd.AddCommand("rmscript", &cmdRmScript{})
	cmd.AddCommand("showscript", &cmdShowScript{})
	cmd.AddCommand("updfiletemplate", &cmdUpdFileTemplate{})
	cmd.AddCommand("filetemplates", &cmdFileTemplates{})
	cmd.AddCommand("rmfiletemplate", &cmdRmFileTemplate{})
	cmd.AddCommand("showfiletemplate", &cmdShowFileTemplate{})
	cmd.AddCommand("myscripts", &cmdMyScripts{})
	cmd.AddCommand("myfiletemplates", &cmdMyFileTemplates{})
	cmd.AddCommand("expscripts", &cmdExpScripts{})
	cmd.AddCommand("expfiletemplates", &cmdExpFileTemplates{})

	p := &usermgrPlugin{
		cmd: cmd,
	}

	return p, nil
}

// OnMessage - get message
func (p *usermgrPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {

	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if !params.ChatBot.IsMaster(from) {
		return false, nil
	}

	if len(params.LstStr) >= 1 {
		p.cmd.Run(ctx, params.LstStr[0], params)

		return true, nil
	}

	return false, nil
}

// GetPluginName - get plugin name
func (p *usermgrPlugin) GetPluginName() string {
	return PluginName
}

// OnStart - on start
func (p *usermgrPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *usermgrPlugin) GetPluginType() int {
	return chatbot.PluginTypeWritableCommand
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *usermgrPlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) >= 1 {
		if p.cmd.HasCommand(params.LstStr[0]) {
			return p.cmd.ParseCommandLine(params.LstStr[0], params)
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
