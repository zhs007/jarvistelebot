package pluginxlsx2json

import (
	"context"

	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

// PluginName - plugin name
const PluginName = "xlsx2json"

// xlsx2jsonPlugin - xlsx2json plugin
type xlsx2jsonPlugin struct {
}

// NewPlugin - new xlsx2json plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - xlsx2jsonPlugin")

	return &xlsx2jsonPlugin{}, nil
}

// // RegPlugin - reg xlsx2json plugin
// func RegPlugin(cfgPath string, mgr chatbot.PluginsMgr) error {
// 	chatbot.Info("RegPlugin - xlsx2jsonPlugin")

// 	mgr.RegPlugin(&xlsx2jsonPlugin{})

// 	return nil
// }

// OnMessage - get message
func (p *xlsx2jsonPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if !params.ChatBot.IsMaster(from) {
		return false, nil
	}

	file := params.Msg.GetFile()
	if file != nil {
		if file.FileType == chatbot.FileExcel {
			str, err := toJSON(file.Data)
			if err == nil {
				fd := &chatbotdbpb.File{
					Filename: chatbot.GetFileNameFromFullPathNoExt(file.Filename) + ".json",
					Data:     []byte(str),
				}

				chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), fd)
			}
			// arr := strings.Split(params.Msg.GetText(), ":")
			// if len(arr) < 2 {
			// 	return false, nil
			// }

			// curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(arr[0])
			// if curnode == nil {
			// 	return false, nil
			// }

			// fd := &jarviscorepb.FileData{
			// 	File:     file.Data,
			// 	Filename: strings.Join(arr[1:], ":"),
			// }

			// params.ChatBot.GetJarvisNode().SendFile(ctx, curnode.Addr, fd)

			return true, nil
		}

		return false, nil
	}

	return false, nil
}

// GetComeInCode - if return is empty string, it means not comein
func (p *xlsx2jsonPlugin) GetComeInCode() string {
	return PluginName
}

// IsMyMessage
func (p *xlsx2jsonPlugin) IsMyMessage(params *chatbot.MessageParams) bool {
	return false
}

// OnStart - on start
func (p *xlsx2jsonPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *xlsx2jsonPlugin) GetPluginType() int {
	return chatbot.PluginTypeCommand
}
