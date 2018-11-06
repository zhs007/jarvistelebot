package pluginjarvisnode

import (
	"github.com/zhs007/jarvistelebot/chatbot"
)

// jarvisnodePlugin - jarvisnode plugin
type jarvisnodePlugin struct {
	cmd *chatbot.CommandMap
}

func newPlugin() *jarvisnodePlugin {
	cmd := chatbot.NewCommandMap()

	cmd.RegFunc("help", cmdHelp)
	cmd.RegFunc("mystate", cmdMyState)

	p := &jarvisnodePlugin{
		cmd: cmd,
	}

	return p
}

// RegPlugin - reg timestamp plugin
func RegPlugin(mgr chatbot.PluginsMgr) {
	chatbot.Info("RegPlugin - jarvisnodePlugin")

	mgr.RegPlugin(newPlugin())
}

// OnMessage - get message
func (p *jarvisnodePlugin) OnMessage(params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if len(params.LstStr) > 1 && params.LstStr[0] == ">" {
		p.cmd.Run(params.LstStr[1], params)

		return true, nil
	}

	// arr := strings.Fields(params.Msg.GetText())
	// if len(arr) > 1 && arr[0] == ">" {
	// 	return true, nil
	// }

	// if from.IsMaster() {
	// 	arr := strings.Fields(msg.GetText())

	// 	ts, err := strconv.ParseInt(arr[0], 10, 64)
	// 	if err == nil {
	// 		tm := time.Unix(ts, 0)
	// 		bot.SendMsg(from, tm.Format("2006-01-02 15:04:05"))

	// 		return true, nil
	// 	} else {
	// 		tm2, err := time.Parse("2006-01-02 15:04:05", msg.GetText())
	// 		if err == nil {
	// 			bot.SendMsg(from, strconv.FormatInt(tm2.Unix(), 10))

	// 			return true, nil
	// 		}
	// 	}
	// } else {
	// 	bot.SendMsg(from, "sorry, you are not my master.")
	// }

	return false, nil
}

// GetComeInCode - if return is empty string, it means not comein
func (p *jarvisnodePlugin) GetComeInCode() string {
	return "jarvisnode"
}

// IsMyMessage
func (p *jarvisnodePlugin) IsMyMessage(params *chatbot.MessageParams) bool {
	// arr := strings.Fields(msg.GetText())
	if len(params.LstStr) > 1 && params.LstStr[0] == ">" {
		return true
	}

	return false
}
