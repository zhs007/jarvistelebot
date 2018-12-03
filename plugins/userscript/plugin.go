package pluginuserscript

import (
	"context"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/userscript/proto"
)

// PluginName - plugin name
const PluginName = "userscript"

// userscriptPlugin - userscript plugin
type userscriptPlugin struct {
}

// NewPlugin - new normal plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - userscriptPlugin")

	return &userscriptPlugin{}, nil
}

// OnMessage - get message
func (p *userscriptPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if params.CommandLine != nil {
		rscmd, ok := params.CommandLine.(*pluginuserscriptpb.RunScriptCommand)
		if !ok {
			return false, chatbot.ErrInvalidCommandLine
		}

		us, err := params.ChatBot.GetChatBotDB().GetUserScript(from.GetUserID(), rscmd.ScriptName)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())

			return false, err
		}

		chatbot.SendTextMsg(params.ChatBot, from,
			"I will execute the script "+rscmd.ScriptName+" for "+us.JarvisNodeName)

		ci, err := jarviscore.BuildCtrlInfoForScriptFile(1, us.File.Filename, us.File.Data, "")
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
			// jarvisbase.Warn("userscriptPlugin.OnMessage", zap.Error(err))

			return false, err
		}

		curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(us.JarvisNodeName)
		if curnode == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrNoJarvisNode.Error())

			return false, chatbot.ErrNoJarvisNode
		}

		params.ChatBot.GetJarvisNode().RequestCtrl(ctx, curnode.Addr, ci)

		params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
			cr := msg.GetCtrlResult()

			chatbot.SendTextMsg(params.ChatBot, from, cr.CtrlResult)

			return nil
		})

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
func (p *userscriptPlugin) GetPluginName() string {
	return PluginName
}

// OnStart - on start
func (p *userscriptPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *userscriptPlugin) GetPluginType() int {
	return chatbot.PluginTypeNormal
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *userscriptPlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	if params.Msg.GetFrom() != nil && len(params.LstStr) == 1 {
		mgrUserScripts := params.ChatBot.GetUserScriptsMgr()
		us, err := mgrUserScripts.Get(params.ChatBot.GetChatBotDB(), params.Msg.GetFrom().GetUserID())
		if err != nil {
			jarvisbase.Warn("userscriptPlugin.ParseMessage:mgrUserScripts.Get", zap.Error(err))

			return nil, chatbot.ErrMsgNotMine
		}

		for _, v := range us.Scripts {
			if params.LstStr[0] == v {
				return &pluginuserscriptpb.RunScriptCommand{
					ScriptName: v,
				}, nil
			}
		}
	}

	return nil, chatbot.ErrMsgNotMine
}