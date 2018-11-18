package pluginnormal

import (
	"context"

	"github.com/zhs007/jarvistelebot/chatbot"
)

// normalPlugin - normal plugin
type normalPlugin struct {
}

// RegPlugin - reg normal plugin
func RegPlugin(cfgPath string, mgr chatbot.PluginsMgr) error {
	chatbot.Info("RegPlugin - normalPlugin")

	mgr.RegPlugin(&normalPlugin{})

	return nil
}

// OnMessage - get message
func (p *normalPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if params.ChatBot.IsMaster(from) {
		// arr := strings.Fields(params.Msg.GetText())
		if params.LstStr[0] == "comein" {
			p := params.MgrPlugins.GetComeInPlugin(params.LstStr[1])
			if p != nil {
				params.MgrPlugins.ComeInPlugin(p)
			}
		} else if params.LstStr[0] == "exit" {
			params.MgrPlugins.ComeInPlugin(nil)
		} else if params.LstStr[0] == "getmystate" {
			p := params.MgrPlugins.GetCurPlugin()
			if p != nil {
				chatbot.SendTextMsg(params.ChatBot, from, "Your are in "+p.GetComeInCode())
				// params.ChatBot.SendMsg(from, "Your are in "+p.GetComeInCode())
			} else {
				chatbot.SendTextMsg(params.ChatBot, from, "nil.")
				// params.ChatBot.SendMsg(from, "nil.")
			}
		} else {
			lstOp := []string{"yes", "no"}
			chatbot.SendMsgWithOptions(params.ChatBot, from, "Yes, master.", lstOp,
				func(ctx context.Context, msg chatbot.Message, id int) error {
					chatbot.SendTextMsg(params.ChatBot, from, "you choice.")
					return nil
				})
			// chatbot.SendTextMsg(params.ChatBot, from, "Yes, master.")
			// params.ChatBot.SendMsg(from, "Yes, master.")
		}
	} else {
		chatbot.SendTextMsg(params.ChatBot, from, "sorry, you are not my master.")
		// params.ChatBot.SendMsg(from, "sorry, you are not my master.")
	}

	return true, nil
}

// GetComeInCode - if return is empty string, it means not comein
func (p *normalPlugin) GetComeInCode() string {
	return ""
}

// IsMyMessage
func (p *normalPlugin) IsMyMessage(params *chatbot.MessageParams) bool {
	return false
}

// OnStart - on start
func (p *normalPlugin) OnStart(ctx context.Context) error {
	return nil
}
