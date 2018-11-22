package pluginjarvisnode

import (
	"context"
	"fmt"

	"github.com/zhs007/jarviscore/base"
	"github.com/zhs007/jarviscore/proto"
	"go.uber.org/zap"

	"github.com/zhs007/jarviscore"
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
	cmd.RegFunc("scripts", cmdScripts)

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
	jarvisbase.Debug("jarvisnodePlugin.OnMessage", zap.String("params", fmt.Sprintf("%+v", params)))

	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if !params.ChatBot.IsMaster(from) {
		return false, nil
	}

	file := params.Msg.GetFile()
	if file != nil {
		if file.FileType == chatbot.FileTypeShellScript {
			ci, err := jarviscore.BuildCtrlInfoForScriptFile(1, file.Filename, file.Data, "")
			if err != nil {
				jarvisbase.Warn("jarvisnodePlugin.OnMessage", zap.Error(err))

				return false, err
			}

			curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(params.Msg.GetText())
			if curnode == nil {
				return false, nil
			}

			params.ChatBot.GetJarvisNode().RequestCtrl(ctx, curnode.Addr, ci)

			params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
				cr := msg.GetCtrlResult()

				chatbot.SendTextMsg(params.ChatBot, from, cr.CtrlResult)

				return nil
			})

			return true, nil
		}

		return false, nil
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
	file := params.Msg.GetFile()
	if file != nil {
		if file.FileType == chatbot.FileTypeShellScript {
			return true
		}
	}

	if len(params.LstStr) > 1 && params.LstStr[0] == ">" {
		return true
	}

	return false
}

// OnStart - on start
func (p *jarvisnodePlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *jarvisnodePlugin) GetPluginType() int {
	return chatbot.PluginTypeCommand
}
