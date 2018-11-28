package plugincore

import (
	"context"

	"github.com/zhs007/jarvistelebot/chatbot"
)

// cmdVersion - version
func cmdVersion(ctx context.Context, params *chatbot.MessageParams) bool {
	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), params.ChatBot.GetVersion())

	return true
}
