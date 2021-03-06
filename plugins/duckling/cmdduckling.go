package pluginduckling

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/duckling/proto"
)

const commandRequestDuckling = "duckling"

// cmdRequestDuckling - duckling
type cmdRequestDuckling struct {
}

// RunCommand - run command
func (cmd *cmdRequestDuckling) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		cmdRequestDuckling, ok := params.CommandLine.(*pluginducklingpb.RequestDuckling)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidCommandLine.Error(), params.Msg)

			return false
		}

		pluginDuckling, ok := params.CurPlugin.(*ducklingPlugin)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidPluginType.Error(), params.Msg)

			return false
		}

		reply, err := pluginDuckling.client.request(ctx, cmdRequestDuckling.Lang, cmdRequestDuckling.Text)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		str, err := chatbot.FormatJSON(reply)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
		} else {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), str, params.Msg)
		}

		return true
	}

	return false
}

// parse - parse command line
func (cmd *cmdRequestDuckling) parse(lstStr []string) (*pluginducklingpb.RequestDuckling, error) {
	if len(lstStr) <= 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	if lstStr[0] != commandRequestDuckling {
		return nil, chatbot.ErrInvalidCommandLineCommand
	}

	flagset := pflag.NewFlagSet(commandRequestDuckling, pflag.ContinueOnError)

	var lang = flagset.StringP("lang", "l", "", "language")
	var text = flagset.StringP("text", "t", "", "text")

	err := flagset.Parse(lstStr[1:])
	if err != nil {
		return nil, err
	}

	if *lang != "" && *text != "" {
		uac := &pluginducklingpb.RequestDuckling{
			Lang: *lang,
			Text: *text,
		}

		return uac, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}

// Parse - parse command line
func (cmd *cmdRequestDuckling) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	return cmd.parse(params.LstStr)
}
