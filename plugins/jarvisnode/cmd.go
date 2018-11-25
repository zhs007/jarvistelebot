package pluginjarvisnode

import (
	"context"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
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

// cmdRequestFile - request file
func cmdRequestFile(ctx context.Context, params *chatbot.MessageParams) bool {
	fn := strings.Join(params.LstStr[2:], " ")

	arr := strings.Split(fn, ":")
	if len(arr) < 2 {
		return false
	}

	curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(arr[0])
	if curnode == nil {
		return false
	}

	rf := &jarviscorepb.RequestFile{
		Filename: strings.Join(arr[1:], ":"),
	}

	params.ChatBot.GetJarvisNode().RequestFile(ctx, curnode.Addr, rf)

	params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
		if msg.MsgType == jarviscorepb.MSGTYPE_REPLY_REQUEST_FILE {
			fd := msg.GetFile()

			chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
				Filename: fd.Filename,
				Data:     fd.File,
			})
		}

		return nil
	})

	// chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), params.ChatBot.GetVersion())

	return true
}
