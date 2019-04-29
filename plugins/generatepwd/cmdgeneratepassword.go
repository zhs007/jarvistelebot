package plugingeneratepwd

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/generatepwd/proto"
)

const commandGeneratePassword = "generatepassword"

// cmdGeneratePassword - generatepassword
type cmdGeneratePassword struct {
}

// RunCommand - run command
func (cmd *cmdGeneratePassword) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		cmdGenPassword, ok := params.CommandLine.(*plugingeneratepwdpb.GeneratePassword)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidCommandLine.Error(), params.Msg)

			return false
		}

		// pluginGenPassword, ok := params.CurPlugin.(*generatePasswordPlugin)
		// if !ok {
		// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidPluginType.Error(), params.Msg)

		// 	return false
		// }

		pwd := genPassword(cmdGenPassword.Mode, int(cmdGenPassword.Length))

		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), pwd, params.Msg)

		return true
	}

	return false
}

// parse - parse command line
func (cmd *cmdGeneratePassword) parse(lstStr []string) (*plugingeneratepwdpb.GeneratePassword, error) {
	if len(lstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	if lstStr[0] != commandGeneratePassword {
		return nil, chatbot.ErrInvalidCommandLineCommand
	}

	flagset := pflag.NewFlagSet(commandGeneratePassword, pflag.ContinueOnError)

	var length = flagset.Int32P("length", "l", 16, "length")
	var mode = flagset.StringP("mode", "m", "normal", "mode")

	err := flagset.Parse(lstStr[1:])
	if err != nil {
		return nil, err
	}

	if *length > 0 && *mode != "" {
		uac := &plugingeneratepwdpb.GeneratePassword{
			Length: *length,
			Mode:   *mode,
		}

		return uac, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}

// Parse - parse command line
func (cmd *cmdGeneratePassword) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	return cmd.parse(params.LstStr)
}
