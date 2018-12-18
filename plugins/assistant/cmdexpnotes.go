package pluginassistant

import (
	"context"
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/assistant/proto"
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

	if params.CommandLine == nil {
		chatbot.SendTextMsg(params.ChatBot, from, chatbot.ErrInvalidCommandLine.Error(), params.Msg)

		return false
	}

	encmd, ok := params.CommandLine.(*pluginassistanepb.ExpNotesCommand)
	if !ok {
		chatbot.SendTextMsg(params.ChatBot, from, chatbot.ErrInvalidCommandLine.Error(), params.Msg)

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

	if encmd.Graph {
		nkg, err := pluginAssistant.Mgr.ExportGraph(from.GetUserID())
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		jsonstr, err := json.Marshal(nkg)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
			Filename: "notes.json",
			Data:     jsonstr,
		})

		return true
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

	return true
}

// Parse - parse command line
func (cmd *cmdExpNotes) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) >= 1 {

		flagset := pflag.NewFlagSet("expnotes", pflag.ContinueOnError)

		var graph = flagset.BoolP("graph", "g", false, "use graph mode")

		err := flagset.Parse(params.LstStr[1:])
		if err != nil {
			return nil, err
		}

		return &pluginassistanepb.ExpNotesCommand{
			Graph: *graph,
		}, nil
	}

	return nil, chatbot.ErrMsgNotMine
}
