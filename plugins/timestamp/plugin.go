package plugintimestamp

import (
	"strconv"
	"strings"
	"time"

	"github.com/zhs007/jarvistelebot/chatbot"
)

// timestampPlugin - timestamp plugin
type timestampPlugin struct {
}

// RegPlugin - reg timestamp plugin
func RegPlugin(mgr chatbot.PluginsMgr) {
	chatbot.Info("RegPlugin - timestampPlugin")

	mgr.RegPlugin(&timestampPlugin{})
}

// OnMessage - get message
func (p *timestampPlugin) OnMessage(bot chatbot.ChatBot, mgr chatbot.PluginsMgr, msg chatbot.Message) (bool, error) {
	from := msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if from.IsMaster() {
		arr := strings.Fields(msg.GetText())

		ts, err := strconv.ParseInt(arr[0], 10, 64)
		if err == nil {
			tm := time.Unix(ts, 0)
			bot.SendMsg(from, tm.Format("2006-01-02 15:04:05"))

			return true, nil
		} else {
			tm2, err := time.Parse("2006-01-02 15:04:05", arr[0])
			if err == nil {
				bot.SendMsg(from, string(tm2.Unix()))

				return true, nil
			}
		}
	} else {
		bot.SendMsg(from, "sorry, you are not my master.")
	}

	return false, nil
}

// GetComeInCode - if return is empty string, it means not comein
func (p *timestampPlugin) GetComeInCode() string {
	return "timestamp"
}
