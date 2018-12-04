package pluginnotekeyword

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarviscore/base"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/assistant"
	"github.com/zhs007/jarvistelebot/plugins/notekeyword/proto"
)

// PluginName - plugin name
const PluginName = "notekeyword"

// userscriptPlugin - notekeyword plugin
type notekeywordPlugin struct {
}

// NewPlugin - new normal plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - notekeywordPlugin")

	return &notekeywordPlugin{}, nil
}

// OnMessage - get message
func (p *notekeywordPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if params.CommandLine != nil {
		nkcmd, ok := params.CommandLine.(*pluginnotekeywordpb.NoteKeywordCommand)
		if !ok {
			return false, chatbot.ErrInvalidCommandLine
		}

		plugin := params.MgrPlugins.FindPlugin(pluginassistant.PluginName)
		if plugin != nil {
			pluginAssistant, ok := plugin.(*pluginassistant.AssistantPlugin)
			if ok {
				lst, err := pluginAssistant.Mgr.FindNoteWithKeyword(from.GetUserID(), nkcmd.Keyword)
				if err != nil {
					chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
				}

				if lst != nil {
					if len(lst) == 1 {
						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "I get a note.")

						for _, v := range lst[0].Data {
							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), v)
						}
					} else {
						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							fmt.Sprintf("I got %v notes.", len(lst)))

						for i, v := range lst {
							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
								fmt.Sprintf("note %v is ", i+1))

							for _, vd := range v.Data {
								chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), vd)
							}
						}
					}

					return true, nil
				}
			}
		}

		return false, nil

		// us, err := params.ChatBot.GetChatBotDB().GetUserScript(from.GetUserID(), rscmd.ScriptName)
		// if err != nil {
		// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())

		// 	return false, err
		// }

		// chatbot.SendTextMsg(params.ChatBot, from,
		// 	"I will execute the script "+rscmd.ScriptName+" for "+us.JarvisNodeName)

		// ci, err := jarviscore.BuildCtrlInfoForScriptFile(1, us.File.Filename, us.File.Data, "")
		// if err != nil {
		// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
		// 	// jarvisbase.Warn("userscriptPlugin.OnMessage", zap.Error(err))

		// 	return false, err
		// }

		// curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(us.JarvisNodeName)
		// if curnode == nil {
		// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrNoJarvisNode.Error())

		// 	return false, chatbot.ErrNoJarvisNode
		// }

		// params.ChatBot.GetJarvisNode().RequestCtrl(ctx, curnode.Addr, ci)

		// params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
		// 	cr := msg.GetCtrlResult()

		// 	chatbot.SendTextMsg(params.ChatBot, from, cr.CtrlResult)

		// 	return nil
		// })

		// chatbot.SendTextMsg(params.ChatBot, from, rscmd.ScriptName)
	}

	// if params.Msg.GetText() == "" {
	// 	return false, chatbot.ErrEmptyMsgText
	// }

	// if params.ChatBot.IsMaster(from) {
	// 	chatbot.SendTextMsg(params.ChatBot, from, "Sorry, I can't understand.")
	// } else {
	// 	chatbot.SendTextMsg(params.ChatBot, from, "Sorry, you are not my master.")
	// }

	return true, nil
}

// GetPluginName - get plugin name
func (p *notekeywordPlugin) GetPluginName() string {
	return PluginName
}

// OnStart - on start
func (p *notekeywordPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *notekeywordPlugin) GetPluginType() int {
	return chatbot.PluginTypeNormal
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *notekeywordPlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	jarvisbase.Debug("notekeywordPlugin.ParseMessage")
	from := params.Msg.GetFrom()
	if from != nil && len(params.LstStr) == 1 {
		jarvisbase.Debug("notekeywordPlugin.ParseMessage:1")
		plugin := params.MgrPlugins.FindPlugin(pluginassistant.PluginName)
		if plugin != nil {
			jarvisbase.Debug("notekeywordPlugin.ParseMessage:2")
			pluginAssistant, ok := plugin.(*pluginassistant.AssistantPlugin)
			if ok {
				jarvisbase.Debug("notekeywordPlugin.ParseMessage:3")
				if pluginAssistant.Mgr.HasKeyword(from.GetUserID(), params.LstStr[0]) {
					jarvisbase.Debug("notekeywordPlugin.ParseMessage:4")
					return &pluginnotekeywordpb.NoteKeywordCommand{
						Keyword: params.LstStr[0],
					}, nil
				}
			}
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
