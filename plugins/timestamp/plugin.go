package plugintimestamp

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/zhs007/jarvistelebot/chatbot"
)

// timestampPlugin - timestamp plugin
type timestampPlugin struct {
}

// RegPlugin - reg timestamp plugin
func RegPlugin(cfgPath string, mgr chatbot.PluginsMgr) error {
	chatbot.Info("RegPlugin - timestampPlugin")

	mgr.RegPlugin(&timestampPlugin{})

	return nil
}

// OnMessage - get message
func (p *timestampPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if from.IsMaster() {
		arr := strings.Fields(params.Msg.GetText())

		ts, err := strconv.ParseInt(arr[0], 10, 64)
		if err == nil {
			tm := time.Unix(ts, 0)
			params.ChatBot.SendMsg(from, tm.Format("2006-01-02 15:04:05"))

			return true, nil
		} else {
			tm2, err := time.Parse("2006-01-02 15:04:05", params.Msg.GetText())
			if err == nil {
				params.ChatBot.SendMsg(from, strconv.FormatInt(tm2.Unix(), 10))

				return true, nil
			}
		}
	} else {
		params.ChatBot.SendMsg(from, "sorry, you are not my master.")
	}

	return false, nil
}

// GetComeInCode - if return is empty string, it means not comein
func (p *timestampPlugin) GetComeInCode() string {
	return "timestamp"
}

// IsMyMessage
func (p *timestampPlugin) IsMyMessage(params *chatbot.MessageParams) bool {
	return false
}
