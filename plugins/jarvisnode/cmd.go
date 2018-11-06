package pluginjarvisnode

import (
	"github.com/zhs007/jarvistelebot/chatbot"
)

func cmdHelp(params *chatbot.MessageParams) bool {
	params.ChatBot.SendMsg(params.Msg.GetFrom(), "This is jarvisnode plugin help.")

	return true
}
