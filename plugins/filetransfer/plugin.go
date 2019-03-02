package pluginfiletransfer

import (
	"context"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// PluginName - plugin name
const PluginName = "filetransfer"

// filetransferPlugin - timestamp plugin
type filetransferPlugin struct {
}

// NewPlugin - new file transfer plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - filetransferPlugin")

	return &filetransferPlugin{}, nil
}

// // RegPlugin - reg timestamp plugin
// func RegPlugin(cfgPath string, mgr chatbot.PluginsMgr) error {
// 	chatbot.Info("RegPlugin - filetransferPlugin")

// 	mgr.RegPlugin(&filetransferPlugin{})

// 	return nil
// }

// OnMessage - get message
func (p *filetransferPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if !params.ChatBot.IsMaster(from) {
		return false, nil
	}

	file := params.Msg.GetFile()
	if file != nil {
		if file.FileType != chatbot.FileTypeShellScript {
			arr := strings.Split(params.Msg.GetText(), ":")
			if len(arr) < 2 {
				return false, nil
			}

			curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(arr[0])
			if curnode == nil {
				return false, nil
			}

			fd := &jarviscorepb.FileData{
				File:     file.Data,
				Filename: strings.Join(arr[1:], ":"),
			}

			params.ChatBot.GetJarvisNode().SendFile(ctx, curnode.Addr, fd,
				func(ctx context.Context, jarvisnode jarviscore.JarvisNode, request *jarviscorepb.JarvisMsg,
					reply *jarviscorepb.JarvisMsg) (bool, error) {

					return true, nil
				})

			// params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
			// 	cr := msg.GetCtrlResult()

			// 	chatbot.SendTextMsg(params.ChatBot, from, cr.CtrlResult)

			// 	return nil
			// })

			return true, nil
		}

		return false, nil
	}

	return false, nil
}

// GetPluginName - get plugin name
func (p *filetransferPlugin) GetPluginName() string {
	return PluginName
}

// // IsMyMessage
// func (p *filetransferPlugin) IsMyMessage(params *chatbot.MessageParams) bool {
// 	file := params.Msg.GetFile()
// 	if file != nil {
// 		if len(params.LstStr) == 1 {
// 			arr := strings.Split(params.Msg.GetText(), ":")
// 			if len(arr) == 2 {
// 				curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(arr[0])
// 				if curnode != nil {
// 					return true
// 				}
// 			}
// 		}
// 	}

// 	return false
// }

// OnStart - on start
func (p *filetransferPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *filetransferPlugin) GetPluginType() int {
	return chatbot.PluginTypeWritableCommand
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *filetransferPlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	file := params.Msg.GetFile()
	if file != nil {
		if len(params.LstStr) == 1 {
			arr := strings.Split(params.Msg.GetText(), ":")
			if len(arr) == 2 {
				curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(arr[0])
				if curnode != nil {
					return nil, nil
				}
			}
		}

		return nil, chatbot.ErrMsgNotMine
	}

	return nil, chatbot.ErrMsgNotMine
}
