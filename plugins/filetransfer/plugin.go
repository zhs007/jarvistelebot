package pluginfiletransfer

import (
	"context"
	"strings"

	"github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// filetransferPlugin - timestamp plugin
type filetransferPlugin struct {
}

// RegPlugin - reg timestamp plugin
func RegPlugin(cfgPath string, mgr chatbot.PluginsMgr) error {
	chatbot.Info("RegPlugin - filetransferPlugin")

	mgr.RegPlugin(&filetransferPlugin{})

	return nil
}

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

			params.ChatBot.GetJarvisNode().SendFile(ctx, curnode.Addr, fd)

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

// GetComeInCode - if return is empty string, it means not comein
func (p *filetransferPlugin) GetComeInCode() string {
	return "filetransfer"
}

// IsMyMessage
func (p *filetransferPlugin) IsMyMessage(params *chatbot.MessageParams) bool {
	return false
}

// OnStart - on start
func (p *filetransferPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *filetransferPlugin) GetPluginType() int {
	return chatbot.PluginTypeCommand
}
