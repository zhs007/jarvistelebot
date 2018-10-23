package pluginnormal

import (
	"strings"

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
func (p *normalPlugin) OnMessage(bot chatbot.ChatBot, mgr chatbot.PluginsMgr, msg chatbot.Message) (bool, error) {
	from := msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if from.IsMaster() {
		arr := strings.Fields(msg.GetText())
		if arr[0] == "comein" {
			p := mgr.GetComeInPlugin(arr[1])
			if p != nil {
				mgr.ComeInPlugin(p)
			}
		} else if arr[0] == "exit" {
			mgr.ComeInPlugin(nil)
		} else if arr[0] == "getmystate" {
			p := mgr.GetCurPlugin()
			if p != nil {
				bot.SendMsg(from, "Your are in "+p.GetComeInCode())
			} else {
				bot.SendMsg(from, "nil.")
			}
		} else {
			bot.SendMsg(from, "Yes, master.")
		}
	} else {
		bot.SendMsg(from, "sorry, you are not my master.")
	}

	return true, nil
}

// GetComeInCode - if return is empty string, it means not comein
func (p *normalPlugin) GetComeInCode() string {
	return ""
}
