package pluginassistant

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// cmdEndKey - endkey
type cmdEndKey struct {
}

// RunCommand - run command
func (cmd *cmdEndKey) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CurPlugin == nil {
		chatbot.SendTextMsg(params.ChatBot, from, chatbot.ErrInvalidParamsNoCurPlugin.Error(), params.Msg)

		return false
	}

	pluginAssistant, ok := params.CurPlugin.(*AssistantPlugin)
	if !ok {
		chatbot.SendTextMsg(params.ChatBot, from, chatbot.ErrInvalidParamsInvalidCurPlugin.Error(), params.Msg)

		return false
	}

	_, err := pluginAssistant.Mgr.SaveCurNote(from.GetUserID())
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, from, err.Error(), params.Msg)

		return false
	}

	params.MgrPlugins.SetCurPlugin(nil)

	chatbot.SendTextMsg(params.ChatBot, from, "I get it.", params.Msg)

	// if params.CommandLine != nil {
	// 	notecmd, ok := params.CommandLine.(*pluginassistanepb.NoteCommand)
	// 	if !ok {
	// 		chatbot.SendTextMsg(params.ChatBot, from, chatbot.ErrInvalidCommandLine.Error())

	// 		return false
	// 	}

	// 	cn, err := pluginAssistant.mgr.NewNote(from.GetUserID())
	// 	if err != nil {
	// 		chatbot.SendTextMsg(params.ChatBot, from, err.Error())

	// 		return false
	// 	}

	// 	if len(notecmd.Keys) > 0 {
	// 		for _, v := range notecmd.Keys {
	// 			cn.Keys = append(notecmd.Keys, v)
	// 		}
	// 	}

	// 	chatbot.SendTextMsg(params.ChatBot, from, "I get it, please tell me what to record.")
	// }

	return true
}

// Parse - parse command line
func (cmd *cmdEndKey) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) >= 1 {
		if params.LstStr[0] == "endkey" {
			return chatbot.NewEmptyCommandLine("endkey"), nil
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
