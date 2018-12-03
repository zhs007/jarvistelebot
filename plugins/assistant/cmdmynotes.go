package pluginassistant

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// cmdMyNotes - mynotes
type cmdMyNotes struct {
}

// RunCommand - run command
func (cmd *cmdMyNotes) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CurPlugin == nil {
		chatbot.SendTextMsg(params.ChatBot, from, chatbot.ErrInvalidParamsNoCurPlugin.Error())

		return false
	}

	pluginAssistant, ok := params.CurPlugin.(*assistantPlugin)
	if !ok {
		chatbot.SendTextMsg(params.ChatBot, from, chatbot.ErrInvalidParamsInvalidCurPlugin.Error())

		return false
	}

	uai, err := pluginAssistant.mgr.GetUserAssistantInfo(from.GetUserID())
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
	}

	strret, err := chatbot.FormatJSONObj(uai)
	if err != nil {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
	} else {
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret)
	}

	// chatbot.SendTextMsg(params.ChatBot, from, "I get it, please tell me the keywords of this note, one at a time.")
	// chatbot.SendTextMsg(params.ChatBot, from, "If you want to stop inputing keywords, you can send ``>> endkey``.")

	return true
}

// Parse - parse command line
func (cmd *cmdMyNotes) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) >= 2 && params.LstStr[0] == ">>" {
		if params.LstStr[1] == "mynotes" {
			return nil, nil
		}
	}

	return nil, nil
}
