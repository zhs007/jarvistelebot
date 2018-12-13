package pluginassistant

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// cmdImpNotes - impnotes
type cmdImpNotes struct {
}

// RunCommand - run command
func (cmd *cmdImpNotes) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		if params.CurPlugin == nil {
			chatbot.SendTextMsg(params.ChatBot, from, chatbot.ErrInvalidParamsNoCurPlugin.Error(), params.Msg)

			return false
		}

		pluginAssistant, ok := params.CurPlugin.(*AssistantPlugin)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, from, chatbot.ErrInvalidParamsInvalidCurPlugin.Error(), params.Msg)

			return false
		}

		file := params.Msg.GetFile()
		if file != nil {
			if file.FileType == chatbot.FileExcel {
				arr, err := chatbot.Xlsx2ArrayMap(file.Data)
				if err != nil {
					chatbot.SendTextMsg(params.ChatBot, from, err.Error(), params.Msg)

					return false
				}

				err = pluginAssistant.Mgr.Import(from.GetUserID(), arr)
				if err != nil {
					chatbot.SendTextMsg(params.ChatBot, from, err.Error(), params.Msg)

					return false
				}

				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "OK, I get it.", params.Msg)

				return true
			}
		}
	}

	return false
}

// Parse - parse command line
func (cmd *cmdImpNotes) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	file := params.Msg.GetFile()
	if file != nil {
		if file.FileType == chatbot.FileExcel {
			if len(params.LstStr) >= 1 {
				if params.LstStr[0] == "impnotes" {
					return chatbot.NewEmptyCommandLine("impnotes"), nil
				}
			}
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
