package pluginjarvisnode

import (
	"context"

	"github.com/zhs007/jarvistelebot/chatbot"
)

// cmdHelp - help
func cmdHelp(ctx context.Context, params *chatbot.MessageParams) bool {
	params.ChatBot.SendMsg(params.Msg.GetFrom(), "This is jarvisnode plugin help.")

	return true
}

// cmdMyState - mystate
func cmdMyState(ctx context.Context, params *chatbot.MessageParams) bool {
	coredb := params.ChatBot.GetJarvisNodeCoreDB()

	str, _ := coredb.GetMyState()
	strret, err := chatbot.FormatJSON(str)
	if err != nil {
		params.ChatBot.SendMsg(params.Msg.GetFrom(), str)
	} else {
		params.ChatBot.SendMsg(params.Msg.GetFrom(), strret)
	}

	return true
}

// cmdRun - run
func cmdRun(ctx context.Context, params *chatbot.MessageParams) bool {
	node := params.ChatBot.GetJarvisNode()

	err := node.SendCtrl(ctx, "1JJaKpZGhYPuVHc1EKiiHZEswPAB5SybW5", "shell", "haha")
	if err != nil {
		params.ChatBot.SendMsg(params.Msg.GetFrom(), err.Error())
	} else {
		params.ChatBot.SendMsg(params.Msg.GetFrom(), "OK!")
	}

	return true
}

// cmdNodes - nodes
func cmdNodes(ctx context.Context, params *chatbot.MessageParams) bool {
	coredb := params.ChatBot.GetJarvisNodeCoreDB()

	str, _ := coredb.GetNodes(100)
	strret, err := chatbot.FormatJSON(str)
	if err != nil {
		params.ChatBot.SendMsg(params.Msg.GetFrom(), str)
	} else {
		params.ChatBot.SendMsg(params.Msg.GetFrom(), strret)
	}

	return true
}
