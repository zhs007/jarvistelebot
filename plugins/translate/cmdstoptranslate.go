package plugintranslate

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/translate/proto"
)

// cmdStopTranslate - stoptranslate
type cmdStopTranslate struct {
}

// RunCommand - run command
func (cmd *cmdStopTranslate) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		_, ok := params.CommandLine.(*plugintranslatepb.StopTranslateCommand)
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

		pluginTranslate.translateParams = nil

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdStopTranslate) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	uac, err := parseStopTranslateCmd(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if uac != nil {
		return uac, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}

func parseStopTranslateCmd(lststr []string) (*plugintranslatepb.StopTranslateCommand, error) {
	// flagset := pflag.NewFlagSet("starttranslate", pflag.ContinueOnError)

	// var srclang = flagset.StringP("srclang", "s", "", "source language")
	// var destlang = flagset.StringP("destlang", "d", "", "destination language")
	// var platform = flagset.StringP("platform", "p", "google", "platform")

	// err := flagset.Parse(lststr)
	// if err != nil {
	// 	return nil, err
	// }

	return &plugintranslatepb.StopTranslateCommand{}, nil

	// if *srclang != "" && *destlang != "" {
	// 	uac := &plugintranslatepb.StopTranslateCommand{
	// 		Platform: *platform,
	// 		SrcLang:  *srclang,
	// 		DestLang: *destlang,
	// 	}

	// 	return uac, nil
	// }

	// return nil, nil
}
