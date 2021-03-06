package plugintranslate

import (
	"context"
	"path"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarviscore/base"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/translate/proto"
)

// PluginName - plugin name
const PluginName = "translate"

// translatePlugin - translate plugin
type translatePlugin struct {
	cmd             *chatbot.CommandMap
	cfg             *config
	client          *translateClient
	translateParams *plugintranslatepb.StartTranslateCommand
	mapGroupInfo    *groupInfo
}

// NewPlugin - new xlsx2json plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - translatePlugin")

	cfg := loadConfig(path.Join(cfgPath, "translate.yaml"))
	err := checkConfig(cfg)
	if err != nil {
		jarvisbase.Warn("plugintranslate.NewPlugin:checkConfig")

		return nil, err
	}

	cmd := chatbot.NewCommandMap()

	cmd.AddCommand("starttranslate", &cmdStartTranslate{})
	cmd.AddCommand("stoptranslate", &cmdStopTranslate{})

	p := &translatePlugin{
		cmd:          cmd,
		cfg:          cfg,
		client:       newTranslateClient(cfg),
		mapGroupInfo: &groupInfo{},
	}

	return p, nil
}

// // RegPlugin - reg xlsx2json plugin
// func RegPlugin(cfgPath string, mgr chatbot.PluginsMgr) error {
// 	chatbot.Info("RegPlugin - xlsx2jsonPlugin")

// 	mgr.RegPlugin(&xlsx2jsonPlugin{})

// 	return nil
// }

// OnMessage - get message
func (p *translatePlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if len(params.LstStr) > 0 {
		if p.cmd.Run(ctx, params.LstStr[0], params) {
			return true, nil
		}
	}

	if params.CommandLine != nil {
		cmd, ok := params.CommandLine.(*plugintranslatepb.TextCommand)
		if ok {
			if p.translateParams != nil {
				str, err := p.client.translate(ctx, cmd.Text,
					p.translateParams.SrcLang, p.translateParams.DestLang)

				if err != nil {
					chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
						err.Error(), params.Msg)
				} else {
					chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
						str, params.Msg)

					if p.translateParams.Retranslate {
						restr, err := p.client.translate(ctx, str,
							p.translateParams.DestLang, p.translateParams.SrcLang)

						if err != nil {
							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
								err.Error(), params.Msg)
						} else {
							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
								restr, params.Msg)
						}
					}
				}

				return true, nil
			}

			if params.Msg.InGroup() {
				groupid := params.Msg.GetGroupID()
				uid := params.Msg.GetFrom().GetUserID()
				nickname := params.Msg.GetFrom().GetNickName()
				gui := p.mapGroupInfo.getGroupUserInfo(groupid, uid)
				if gui != nil {
					str, err := p.client.translate(ctx, cmd.Text,
						gui.srcLang, gui.destLang)

					if err != nil {
						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							err.Error(), params.Msg)
					} else {
						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							nickname+": "+str, params.Msg)

						if gui.retranslate {
							restr, err := p.client.translate(ctx, str,
								gui.destLang, gui.srcLang)

							if err != nil {
								chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
									err.Error(), params.Msg)
							} else {
								chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
									nickname+": "+restr, params.Msg)
							}
						}
					}

					return true, nil
				}
			}
		}
	}

	// if p.translateParams != nil && params.CommandLine != nil {

	// 	eacmd, ok := params.CommandLine.(*plugintranslatepb.TextCommand)
	// 	if !ok {
	// 		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
	// 			chatbot.ErrInvalidCommandLine.Error(), params.Msg)

	// 		return false, chatbot.ErrInvalidCommandLine
	// 	}

	// 	str, err := p.client.translate(ctx, eacmd.Text,
	// 		p.translateParams.SrcLang, p.translateParams.DestLang)

	// 	if err != nil {
	// 		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
	// 			err.Error(), params.Msg)
	// 	} else {
	// 		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
	// 			str, params.Msg)
	// 	}

	// 	return true, nil
	// }

	// if !params.ChatBot.IsMaster(from) {
	// 	return false, nil
	// }

	// file := params.Msg.GetFile()
	// if file != nil {
	// 	if file.FileType == chatbot.FileExcel {
	// 		str, err := toJSON(file.Data)
	// 		if err == nil {
	// 			fd := &chatbotdbpb.File{
	// 				Filename: chatbot.GetFileNameFromFullPathNoExt(file.Filename) + ".json",
	// 				Data:     []byte(str),
	// 			}

	// 			chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), fd, params.Msg)
	// 		}

	// 		return true, nil
	// 	}

	// 	return false, nil
	// }

	return false, nil
}

// GetPluginName - get plugin name
func (p *translatePlugin) GetPluginName() string {
	return PluginName
}

// // IsMyMessage
// func (p *xlsx2jsonPlugin) IsMyMessage(params *chatbot.MessageParams) bool {
// 	file := params.Msg.GetFile()
// 	if file != nil {
// 		if file.FileType == chatbot.FileExcel {
// 			if params.Msg.GetText() == "" {
// 				return true
// 			}
// 			// if len(params.LstStr) == 1 {
// 			// 	arr := strings.Split(params.Msg.GetText(), ":")
// 			// 	if len(arr) == 2 {
// 			// 		curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(arr[0])
// 			// 		if curnode != nil {
// 			// 			return true
// 			// 		}
// 			// 	}
// 			// }
// 		}
// 	}

// 	return false
// }

// OnStart - on start
func (p *translatePlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *translatePlugin) GetPluginType() int {
	return chatbot.PluginTypeWritableCommand
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *translatePlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) >= 1 && p.cmd.HasCommand(params.LstStr[0]) {
		msg, _ := p.cmd.ParseCommandLine(params.LstStr[0], params)
		if msg != nil {
			return msg, nil
		}
	}

	if p.translateParams != nil && len(params.LstStr) >= 1 {
		uac := &plugintranslatepb.TextCommand{
			Text: params.Msg.GetText(),
		}

		return uac, nil
	}

	if params.Msg.InGroup() {
		groupid := params.Msg.GetGroupID()
		uid := params.Msg.GetFrom().GetUserID()
		gui := p.mapGroupInfo.getGroupUserInfo(groupid, uid)
		if gui != nil {
			uac := &plugintranslatepb.TextCommand{
				Text: params.Msg.GetText(),
			}

			return uac, nil
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
