package plugintranslate

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/translate/proto"
)

// cmdTranslate - translate
type cmdTranslate struct {
}

// RunCommand - run command
func (cmd *cmdTranslate) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		eacmd, ok := params.CommandLine.(*plugintranslatepb.TranslateCommand)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
				chatbot.ErrInvalidCommandLine.Error(), params.Msg)

			return false
		}

		pluginTranslate, ok := params.CurPlugin.(*translatePlugin)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
				chatbot.ErrInvalidParamsInvalidCurPlugin.Error(), params.Msg)

			return false
		}

		if eacmd.Run {
			pluginTranslate.translateParams = eacmd
		} else {
			pluginTranslate.translateParams = nil
		}

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdTranslate) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("runscript", pflag.ContinueOnError)

	var srclang = flagset.StringP("srclang", "s", "", "source language")
	var destlang = flagset.StringP("destlang", "d", "", "destination language")
	var platform = flagset.StringP("platform", "p", "", "platform")
	var run = flagset.BoolP("run", "r", true, "run")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *srclang != "" && *destlang != "" {
		uac := &plugintranslatepb.TranslateCommand{
			Platform: *platform,
			SrcLang:  *srclang,
			DestLang: *destlang,
			Run:      *run,
		}

		return uac, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
