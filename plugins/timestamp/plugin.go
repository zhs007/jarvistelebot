package plugintimestamp

import (
	"context"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// PluginName - plugin name
const PluginName = "timestamp"

// timestampPlugin - timestamp plugin
type timestampPlugin struct {
}

// NewPlugin - new timestamp plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - timestampPlugin")

	return &timestampPlugin{}, nil
}

// // RegPlugin - reg timestamp plugin
// func RegPlugin(cfgPath string, mgr chatbot.PluginsMgr) error {
// 	chatbot.Info("RegPlugin - timestampPlugin")

// 	mgr.RegPlugin(&timestampPlugin{})

// 	return nil
// }

// OnMessage - get message
func (p *timestampPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if params.Msg.GetText() == "" {
		return false, nil
	}

	if params.ChatBot.IsMaster(from) {
		arr := chatbot.SplitString(params.Msg.GetText())

		ts, err := strconv.ParseInt(arr[0], 10, 64)
		if err == nil {
			tm := time.Unix(ts, 0)
			// params.ChatBot.SendMsg(from, tm.Format("2006-01-02 15:04:05"))
			chatbot.SendTextMsg(params.ChatBot, from, tm.Format("2006-01-02 15:04:05"), params.Msg)

			return true, nil
		}

		tm2, err := time.Parse("2006-01-02 15:04:05", params.Msg.GetText())
		if err == nil {
			chatbot.SendTextMsg(params.ChatBot, from, strconv.FormatInt(tm2.Unix(), 10), params.Msg)
			// params.ChatBot.SendMsg(from, strconv.FormatInt(tm2.Unix(), 10))

			return true, nil
		}

	} else {
		chatbot.SendTextMsg(params.ChatBot, from, "sorry, you are not my master.", params.Msg)
		// params.ChatBot.SendMsg(from, "sorry, you are not my master.")
	}

	return false, nil
}

// GetPluginName - get plugin name
func (p *timestampPlugin) GetPluginName() string {
	return PluginName
}

// // IsMyMessage
// func (p *timestampPlugin) IsMyMessage(params *chatbot.MessageParams) bool {
// 	return false
// }

// OnStart - on start
func (p *timestampPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *timestampPlugin) GetPluginType() int {
	return chatbot.PluginTypeCommand
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *timestampPlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	return nil, chatbot.ErrMsgNotMine
}
