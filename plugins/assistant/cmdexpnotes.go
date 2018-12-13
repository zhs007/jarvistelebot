package pluginassistant

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

// cmdExpNotes - expnotes
type cmdExpNotes struct {
}

// RunCommand - run command
func (cmd *cmdExpNotes) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
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

	arr, err := pluginAssistant.Mgr.Export(from.GetUserID())
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

		return false
	}

	buf, err := chatbot.ArrayMap2xlsx(arr)
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
	} else {
		chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
			Filename: "notes.xlsx",
			Data:     buf,
		})
	}

	// uai, err := pluginAssistant.Mgr.GetUserAssistantInfo(from.GetUserID())
	// if err != nil {
	// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

	// 	return false
	// }

	// strret, err := chatbot.FormatJSONObj(uai)
	// if err != nil {
	// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
	// } else {
	// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret, params.Msg)
	// }

	// chatbot.SendTextMsg(params.ChatBot, from, "I get it, please tell me the keywords of this note, one at a time.")
	// chatbot.SendTextMsg(params.ChatBot, from, "If you want to stop inputing keywords, you can send ``endkey``.")

	return true
}

// Parse - parse command line
func (cmd *cmdExpNotes) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) >= 1 {
		if params.LstStr[0] == "expnotes" {
			return chatbot.NewEmptyCommandLine("expnotes"), nil
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
