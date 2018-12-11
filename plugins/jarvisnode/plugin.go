package pluginjarvisnode

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// PluginName - plugin name
const PluginName = "jarvisnode"

// jarvisnodePlugin - jarvisnode plugin
type jarvisnodePlugin struct {
	cmd *chatbot.CommandMap
}

// NewPlugin - new jarvisnode plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - jarvisnodePlugin")

	cmd := chatbot.NewCommandMap()

	// cmd.RegFunc("help", cmdHelp)
	cmd.AddCommand("mystate", &cmdMyState{})
	// cmd.RegFunc("run", cmdRun)
	cmd.AddCommand("nodes", &cmdNodes{})
	// cmd.RegFunc("scripts", cmdScripts)
	// cmd.RegFunc("version", cmdVersion)
	cmd.AddCommand("requestfile", &cmdRequestFile{})

	p := &jarvisnodePlugin{
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
func (p *jarvisnodePlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	// jarvisbase.Debug("jarvisnodePlugin.OnMessage", zap.String("params", fmt.Sprintf("%+v", params)))

	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if !params.ChatBot.IsMaster(from) {
		return false, nil
	}

	// file := params.Msg.GetFile()
	// if file != nil {
	// 	if file.FileType == chatbot.FileTypeShellScript {
	// 		ci, err := jarviscore.BuildCtrlInfoForScriptFile(1, file.Filename, file.Data, "")
	// 		if err != nil {
	// 			jarvisbase.Warn("jarvisnodePlugin.OnMessage", zap.Error(err))

	// 			return false, err
	// 		}

	// 		curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(params.Msg.GetText())
	// 		if curnode == nil {
	// 			return false, nil
	// 		}

	// 		params.ChatBot.GetJarvisNode().RequestCtrl(ctx, curnode.Addr, ci)

	// 		params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
	// 			cr := msg.GetCtrlResult()

	// 			chatbot.SendTextMsg(params.ChatBot, from, cr.CtrlResult)

	// 			return nil
	// 		})

	// 		return true, nil
	// 	}

	// 	return false, nil
	// }

	if len(params.LstStr) >= 1 {
		p.cmd.Run(ctx, params.LstStr[0], params)

		return true, nil
	}

	return false, nil
}

// GetPluginName - get plugin name
func (p *jarvisnodePlugin) GetPluginName() string {
	return PluginName
}

// OnStart - on start
func (p *jarvisnodePlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *jarvisnodePlugin) GetPluginType() int {
	return chatbot.PluginTypeCommand
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *jarvisnodePlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	file := params.Msg.GetFile()
	if file != nil {
		if file.FileType == chatbot.FileTypeShellScript {
			if len(params.LstStr) == 1 {
				return nil, nil
			}
		}
	}

	if len(params.LstStr) >= 1 {
		if p.cmd.HasCommand(params.LstStr[0]) {
			return p.cmd.ParseCommandLine(params.LstStr[0], params)
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
