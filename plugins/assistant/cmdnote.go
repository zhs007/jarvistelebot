package pluginassistant

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/assistant/proto"
)

// cmdNote - note
type cmdNote struct {
}

// RunCommand - run command
func (cmd *cmdNote) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
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

	if params.CommandLine != nil {
		notecmd, ok := params.CommandLine.(*pluginassistanepb.NoteCommand)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, from, chatbot.ErrInvalidCommandLine.Error(), params.Msg)

			return false
		}

		cn, err := pluginAssistant.Mgr.NewNote(from.GetUserID())
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, from, err.Error(), params.Msg)

			return false
		}

		if len(notecmd.Keys) > 0 {
			for _, v := range notecmd.Keys {
				cn.Keys = append(notecmd.Keys, v)
			}
		}

		params.MgrPlugins.SetCurPlugin(pluginAssistant)

		chatbot.SendTextMsg(params.ChatBot, from, "I get it, please tell me what this note is for recording.", params.Msg)
		chatbot.SendTextMsg(params.ChatBot, from, "You can input multiple segments for this note.", params.Msg)
		chatbot.SendTextMsg(params.ChatBot, from, "If you want to stop recording, you can send ``endnote``.", params.Msg)
	}

	return true
}

// Parse - parse command line
func (cmd *cmdNote) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) >= 1 {
		if params.LstStr[0] == "note" {
			if len(params.LstStr) >= 2 {
				flagset := pflag.NewFlagSet("note", pflag.ContinueOnError)

				var keys = flagset.StringSliceP("key", "k", []string{}, "you can set keywords")

				err := flagset.Parse(params.LstStr[1:])
				if err != nil {
					return nil, err
				}

				return &pluginassistanepb.NoteCommand{
					Keys: *keys,
				}, nil
			}

			return &pluginassistanepb.NoteCommand{}, nil
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
