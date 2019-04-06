package plugintranslate

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/translate/proto"
)

// cmdStartTranslate - translate
type cmdStartTranslate struct {
}

// RunCommand - run command
func (cmd *cmdStartTranslate) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		cmd, ok := params.CommandLine.(*plugintranslatepb.StartTranslateCommand)
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

		if params.Msg.InGroup() && cmd.Username != "" {
			user, err := params.ChatBot.GetChatBotDB().GetUserWithUserName(cmd.Username)
			if err != nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
					err.Error(), params.Msg)

				return false
			}

			pluginTranslate.mapGroupInfo.setGroupUser(
				params.Msg.GetGroupID(),
				user.UserID,
				cmd.SrcLang,
				cmd.DestLang,
			)

			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
				fmt.Sprintf("I got it, I will translate %v to %v for %v.%v(%v-%v)",
					cmd.SrcLang, cmd.DestLang, params.Msg.GetGroupID(), cmd.Username, user.UserID, params.Msg.GetChatID()), params.Msg)
		} else {
			pluginTranslate.translateParams = cmd

			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
				fmt.Sprintf("I got it, I will translate %v to %v for you",
					cmd.SrcLang, cmd.DestLang), params.Msg)
		}

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdStartTranslate) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	uac, err := parseStartTranslateCmd(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if uac != nil {
		return uac, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}

func parseStartTranslateCmd(lststr []string) (*plugintranslatepb.StartTranslateCommand, error) {
	flagset := pflag.NewFlagSet("starttranslate", pflag.ContinueOnError)

	var srclang = flagset.StringP("srclang", "s", "", "source language")
	var destlang = flagset.StringP("destlang", "d", "", "destination language")
	var platform = flagset.StringP("platform", "p", "google", "platform")
	var username = flagset.StringP("username", "u", "", "username")

	err := flagset.Parse(lststr)
	if err != nil {
		return nil, err
	}

	if *srclang != "" && *destlang != "" {
		uac := &plugintranslatepb.StartTranslateCommand{
			Platform: *platform,
			SrcLang:  *srclang,
			DestLang: *destlang,
			Username: *username,
		}

		return uac, nil
	}

	return nil, nil
}
