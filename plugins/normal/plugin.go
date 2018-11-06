package pluginnormal

import (
	"github.com/zhs007/jarvistelebot/chatbot"
)

// normalPlugin - normal plugin
type normalPlugin struct {
}

// RegPlugin - reg normal plugin
func RegPlugin(mgr chatbot.PluginsMgr) {
	chatbot.Info("RegPlugin - normalPlugin")

	mgr.RegPlugin(&normalPlugin{})
}

// OnMessage - get message
func (p *normalPlugin) OnMessage(params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if from.IsMaster() {
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
				params.ChatBot.SendMsg(from, "Your are in "+p.GetComeInCode())
			} else {
				params.ChatBot.SendMsg(from, "nil.")
			}
		} else {
			params.ChatBot.SendMsg(from, "Yes, master.")
		}
	} else {
		params.ChatBot.SendMsg(from, "sorry, you are not my master.")
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
