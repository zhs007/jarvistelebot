package pluginnormal

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// PluginName - plugin name
const PluginName = "normal"

// normalPlugin - normal plugin
type normalPlugin struct {
}

// NewPlugin - new normal plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - normalPlugin")

	return &normalPlugin{}, nil
}

// // RegPlugin - reg normal plugin
// func RegPlugin(cfgPath string, mgr chatbot.PluginsMgr) error {
// 	chatbot.Info("RegPlugin - normalPlugin")

// 	mgr.RegPlugin(&normalPlugin{})

// 	return nil
// }

// OnMessage - get message
func (p *normalPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if params.Msg.GetText() == "" {
		return false, chatbot.ErrEmptyMsgText
	}

	if params.ChatBot.IsMaster(from) {
		// arr := strings.Fields(params.Msg.GetText())
		// if params.LstStr[0] == "comein" {
		// 	p := params.MgrPlugins.GetComeInPlugin(params.LstStr[1])
		// 	if p != nil {
		// 		params.MgrPlugins.ComeInPlugin(p)
		// 	}
		// } else if params.LstStr[0] == "exit" {
		// 	params.MgrPlugins.ComeInPlugin(nil)
		// } else if params.LstStr[0] == "getmystate" {
		// 	p := params.MgrPlugins.GetCurPlugin()
		// 	if p != nil {
		// 		chatbot.SendTextMsg(params.ChatBot, from, "Your are in "+p.GetComeInCode())
		// 		// params.ChatBot.SendMsg(from, "Your are in "+p.GetComeInCode())
		// 	} else {
		// 		chatbot.SendTextMsg(params.ChatBot, from, "nil.")
		// 		// params.ChatBot.SendMsg(from, "nil.")
		// 	}
		// } else {
		if !params.Msg.InGroup() {
			chatbot.SendTextMsg(params.ChatBot, from, "Sorry, I can't understand.", params.Msg)
		}

		// lstOp := []string{"yes", "no"}
		// chatbot.SendMsgWithOptions(params.ChatBot, from, "Yes, master.", lstOp,
		// 	func(ctx context.Context, msg chatbot.Message, id int) error {
		// 		chatbot.SendTextMsg(params.ChatBot, from, "you choice."+msg.GetOption(id))
		// 		return nil
		// 	})

		// chatbot.SendTextMsg(params.ChatBot, from, "Yes, master.")
		// params.ChatBot.SendMsg(from, "Yes, master.")
		// }
	} else {
		if !params.Msg.InGroup() {
			chatbot.SendTextMsg(params.ChatBot, from, "Sorry, you are not my master.", params.Msg)
		}

		// params.ChatBot.SendMsg(from, "sorry, you are not my master.")
	}

	return true, nil
}

// GetPluginName - get plugin name
func (p *normalPlugin) GetPluginName() string {
	return PluginName
}

// // IsMyMessage
// func (p *normalPlugin) IsMyMessage(params *chatbot.MessageParams) bool {
// 	return false
// }

// OnStart - on start
func (p *normalPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *normalPlugin) GetPluginType() int {
	return chatbot.PluginTypeNormal
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *normalPlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	return nil, chatbot.ErrMsgNotMine
}
