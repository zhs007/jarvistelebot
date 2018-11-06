package pluginjarvisnode

import (
	"github.com/zhs007/jarvistelebot/chatbot"
)

func cmdHelp(params *chatbot.MessageParams) bool {
	params.ChatBot.SendMsg(params.Msg.GetFrom(), "This is jarvisnode plugin help.")

	return true
}

func cmdMyState(params *chatbot.MessageParams) bool {
	coredb := params.ChatBot.GetJarvisNodeCoreDB()

	str, _ := coredb.GetMyState()
	params.ChatBot.SendMsg(params.Msg.GetFrom(), str)

	return true
}
