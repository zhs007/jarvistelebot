package pluginjarvisnodeex

import (
	"context"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarviscore/base"
	"github.com/zhs007/jarviscore/proto"
	"go.uber.org/zap"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// PluginName - plugin name
const PluginName = "jarvisnodeex"

// jarvisnodeexPlugin - jarvisnodeex plugin
type jarvisnodeexPlugin struct {
	cmd *chatbot.CommandMap
}

// NewPlugin - new jarvisnode plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - jarvisnodeexPlugin")

	cmd := chatbot.NewCommandMap()

	// cmd.RegFunc("help", cmdHelp)
	// cmd.RegFunc("mystate", cmdMyState)
	// cmd.RegFunc("run", cmdRun)
	// cmd.RegFunc("nodes", cmdNodes)
	// cmd.RegFunc("scripts", cmdScripts)
	// // cmd.RegFunc("version", cmdVersion)
	// cmd.RegFunc("requestfile", cmdRequestFile)

	p := &jarvisnodeexPlugin{
		cmd: cmd,
	}

	return p, nil
}

// OnMessage - get message
func (p *jarvisnodeexPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	// jarvisbase.Debug("jarvisnodePlugin.OnMessage", zap.String("params", fmt.Sprintf("%+v", params)))

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
				jarvisbase.Warn("jarvisnodeexPlugin.OnMessage", zap.Error(err))

				return false, err
			}

			curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(params.Msg.GetText())
			if curnode == nil {
				return false, nil
			}

			params.ChatBot.GetJarvisNode().RequestCtrl(ctx, curnode.Addr, ci)

			params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
				cr := msg.GetCtrlResult()

				chatbot.SendTextMsg(params.ChatBot, from, cr.CtrlResult, params.Msg)

				return nil
			})

			return true, nil
		}

		return false, nil
	}

	if len(params.LstStr) > 1 && params.LstStr[0] == ">>" {
		p.cmd.Run(ctx, params.LstStr[1], params)

		return true, nil
	}

	return false, nil
}

// GetPluginName - get plugin name
func (p *jarvisnodeexPlugin) GetPluginName() string {
	return PluginName
}

// // IsMyMessage
// func (p *jarvisnodeexPlugin) IsMyMessage(params *chatbot.MessageParams) bool {
// 	file := params.Msg.GetFile()
// 	if file != nil {
// 		if file.FileType == chatbot.FileTypeShellScript {
// 			if len(params.LstStr) == 1 {
// 				arr := strings.Split(params.Msg.GetText(), ":")
// 				if len(arr) == 1 {
// 					return true
// 				}
// 			}
// 		}
// 	}

// 	if len(params.LstStr) >= 2 && params.LstStr[0] == ">>" {
// 		return p.cmd.HasCommand(params.LstStr[1])
// 	}

// 	return false
// }

// OnStart - on start
func (p *jarvisnodeexPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *jarvisnodeexPlugin) GetPluginType() int {
	return chatbot.PluginTypeWritableCommand
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *jarvisnodeexPlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	file := params.Msg.GetFile()
	if file != nil {
		if file.FileType == chatbot.FileTypeShellScript {
			if len(params.LstStr) == 1 {
				arr := strings.Split(params.Msg.GetText(), ":")
				if len(arr) == 1 {
					return nil, nil
				}
			}
		}

		return nil, chatbot.ErrMsgNotMine
	}

	if len(params.LstStr) >= 2 && params.LstStr[0] == ">" {
		if p.cmd.HasCommand(params.LstStr[1]) {
			return p.cmd.ParseCommandLine(params.LstStr[1], params)
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
