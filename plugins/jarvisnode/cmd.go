package pluginjarvisnode

import (
	"context"
	"fmt"
	"path"
	"path/filepath"

	"github.com/zhs007/jarvistelebot/chatbot"
)

// cmdHelp - help
func cmdHelp(ctx context.Context, params *chatbot.MessageParams) bool {
	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "This is jarvisnode plugin help.")
	// params.ChatBot.SendMsg(params.Msg.GetFrom(), "This is jarvisnode plugin help.")

	return true
}

// cmdMyState - mystate
func cmdMyState(ctx context.Context, params *chatbot.MessageParams) bool {
	coredb := params.ChatBot.GetJarvisNodeCoreDB()

	str, _ := coredb.GetMyState()
	strret, err := chatbot.FormatJSON(str)
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), str)
		// params.ChatBot.SendMsg(params.Msg.GetFrom(), str)
	} else {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret)
		// params.ChatBot.SendMsg(params.Msg.GetFrom(), strret)
	}

	return true
}

// cmdRun - run
func cmdRun(ctx context.Context, params *chatbot.MessageParams) bool {
	// node := params.ChatBot.GetJarvisNode()

	// err := node.SendCtrl(ctx, "1JJaKpZGhYPuVHc1EKiiHZEswPAB5SybW5", "shell", "haha")
	// if err != nil {
	// 	params.ChatBot.SendMsg(params.Msg.GetFrom(), err.Error())
	// } else {
	// 	params.ChatBot.SendMsg(params.Msg.GetFrom(), "OK!")
	// }

	return true
}

// cmdNodes - nodes
func cmdNodes(ctx context.Context, params *chatbot.MessageParams) bool {
	coredb := params.ChatBot.GetJarvisNodeCoreDB()

	str, _ := coredb.GetNodes(100)
	strret, err := chatbot.FormatJSON(str)
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), str)
		// params.ChatBot.SendMsg(params.Msg.GetFrom(), str)
	} else {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret)
		// params.ChatBot.SendMsg(params.Msg.GetFrom(), strret)
	}

	return true
}

// cmdScripts - scripts
func cmdScripts(ctx context.Context, params *chatbot.MessageParams) bool {
	files, _ := filepath.Glob(path.Join(params.ChatBot.GetConfig().DownloadPath, "scripts", "*.sh"))

	strret, err := chatbot.FormatJSONObj(files)
	if err != nil {
		str := fmt.Sprintf("%+v", files)
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), str)
		// params.ChatBot.SendMsg(params.Msg.GetFrom(), str)
	} else {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret)
		// params.ChatBot.SendMsg(params.Msg.GetFrom(), strret)
	}

	return true
}

// cmdVersion - version
func cmdVersion(ctx context.Context, params *chatbot.MessageParams) bool {
	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), params.ChatBot.GetVersion())

	return true
}
