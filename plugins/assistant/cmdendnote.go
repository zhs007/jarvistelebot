package pluginassistant

import (
	"context"

	"github.com/zhs007/jarvistelebot/assistant"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// cmdEndNote - endnote
type cmdEndNote struct {
}

// RunCommand - run command
func (cmd *cmdEndNote) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
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

	pluginAssistant.Mgr.ChgCurNoteMode(from.GetUserID(), assistant.ModeInputKey)

	chatbot.SendTextMsg(params.ChatBot, from, "I get it, please tell me the keywords of this note, one at a time.", params.Msg)
	chatbot.SendTextMsg(params.ChatBot, from, "If you want to stop inputing keywords, you can send ``endkey``.", params.Msg)

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
func (cmd *cmdEndNote) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) >= 1 {
		if params.LstStr[0] == "endnote" {
			return chatbot.NewEmptyCommandLine("endnote"), nil
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
