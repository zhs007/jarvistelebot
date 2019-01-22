package pluginfiletemplate

import (
	"context"
	"fmt"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/filetemplate/proto"
)

// PluginName - plugin name
const PluginName = "filetemplate"

// filetemplatePlugin - filetemplate plugin
type filetemplatePlugin struct {
}

// NewPlugin - new normal plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - filetemplatePlugin")

	return &filetemplatePlugin{}, nil
}

// OnMessage - get message
func (p *filetemplatePlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if params.CommandLine != nil {
		ftcmd, ok := params.CommandLine.(*pluginfiletemplatepb.FileTemplateCommand)
		if !ok {
			return false, chatbot.ErrInvalidCommandLine
		}

		ft, err := params.ChatBot.GetChatBotDB().GetFileTemplate(from.GetUserID(), ftcmd.FileTemplateName)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false, err
		}

		file := params.Msg.GetFile()
		if file != nil {
			chatbot.SendTextMsg(params.ChatBot, from,
				fmt.Sprintf("I will send %v to %v.", ft.FullPath, ft.JarvisNodeName), params.Msg)

			curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(ft.JarvisNodeName)
			if curnode == nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrNoJarvisNode.Error(), params.Msg)

				return false, chatbot.ErrNoJarvisNode
			}

			fd := &jarviscorepb.FileData{
				File:     file.Data,
				Filename: ft.FullPath,
			}

			params.ChatBot.GetJarvisNode().SendFile(ctx, curnode.Addr, fd)

			if ft.SubfilesPath != "" {
				chatbot.SendTextMsg(params.ChatBot, from,
					fmt.Sprintf("I processing subfiles."), params.Msg)

				err = procSubfiles(params.ChatBot.GetChatBotDB(), from.GetUserID(), ft.JarvisNodeName, file.Data, ft.SubfilesPath)
				if err != nil {
					chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

					return false, err
				}
			}

			return true, nil
		}

		chatbot.SendTextMsg(params.ChatBot, from,
			fmt.Sprintf("I will get %v:%v.", ft.JarvisNodeName, ft.FullPath), params.Msg)

		curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(ft.JarvisNodeName)
		if curnode == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrNoJarvisNode.Error(), params.Msg)

			return false, chatbot.ErrNoJarvisNode
		}

		rf := &jarviscorepb.RequestFile{
			Filename: ft.FullPath,
		}

		params.ChatBot.GetJarvisNode().RequestFile(ctx, curnode.Addr, rf)

		params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
			if msg.MsgType == jarviscorepb.MSGTYPE_REPLY_REQUEST_FILE {
				fd := msg.GetFile()

				chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
					Filename: ft.FileTemplateName,
					Data:     fd.File,
				})
			}

			return nil
		})
	}

	return true, nil
}

// GetPluginName - get plugin name
func (p *filetemplatePlugin) GetPluginName() string {
	return PluginName
}

// OnStart - on start
func (p *filetemplatePlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *filetemplatePlugin) GetPluginType() int {
	return chatbot.PluginTypeNormal
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *filetemplatePlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	file := params.Msg.GetFile()
	if file != nil {
		if params.Msg.GetText() == "" {
			mgrFileTemplates := params.ChatBot.GetFileTemplatesMgr()
			ft, err := mgrFileTemplates.Get(params.ChatBot.GetChatBotDB(), params.Msg.GetFrom().GetUserID())
			if err != nil {
				jarvisbase.Warn("filetemplatePlugin.ParseMessage:mgrFileTemplates.Get", zap.Error(err))

				return nil, chatbot.ErrMsgNotMine
			}

			for _, v := range ft.Templates {
				if file.Filename == v {
					return &pluginfiletemplatepb.FileTemplateCommand{
						FileTemplateName: v,
					}, nil
				}
			}
		}

		return nil, chatbot.ErrMsgNotMine
	}

	if params.Msg.GetFrom() != nil && len(params.LstStr) == 1 {
		mgrFileTemplates := params.ChatBot.GetFileTemplatesMgr()
		ft, err := mgrFileTemplates.Get(params.ChatBot.GetChatBotDB(), params.Msg.GetFrom().GetUserID())
		if err != nil {
			jarvisbase.Warn("filetemplatePlugin.ParseMessage:mgrFileTemplates.Get", zap.Error(err))

			return nil, chatbot.ErrMsgNotMine
		}

		for _, v := range ft.Templates {
			if params.LstStr[0] == v {
				return &pluginfiletemplatepb.FileTemplateCommand{
					FileTemplateName: v,
				}, nil
			}
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
