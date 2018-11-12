package pluginjarvisnode

import (
	"context"

	"github.com/zhs007/jarvistelebot/chatbot"
)

// jarvisnodePlugin - jarvisnode plugin
type jarvisnodePlugin struct {
	cmd *chatbot.CommandMap
}

func newPlugin() *jarvisnodePlugin {
	cmd := chatbot.NewCommandMap()

	cmd.RegFunc("help", cmdHelp)
	cmd.RegFunc("mystate", cmdMyState)
	cmd.RegFunc("run", cmdRun)
	cmd.RegFunc("nodes", cmdNodes)

	p := &jarvisnodePlugin{
		cmd: cmd,
	}

	return p
}

// RegPlugin - reg timestamp plugin
func RegPlugin(cfgPath string, mgr chatbot.PluginsMgr) error {
	chatbot.Info("RegPlugin - jarvisnodePlugin")

	mgr.RegPlugin(newPlugin())

	return nil
}

// OnMessage - get message
func (p *jarvisnodePlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if len(params.LstStr) > 1 && params.LstStr[0] == ">" {
		p.cmd.Run(ctx, params.LstStr[1], params)

		return true, nil
	}

	return false, nil
}

// GetComeInCode - if return is empty string, it means not comein
func (p *jarvisnodePlugin) GetComeInCode() string {
	return "jarvisnode"
}

// IsMyMessage
func (p *jarvisnodePlugin) IsMyMessage(params *chatbot.MessageParams) bool {
	if len(params.LstStr) > 1 && params.LstStr[0] == ">" {
		return true
	}

	return false
}
