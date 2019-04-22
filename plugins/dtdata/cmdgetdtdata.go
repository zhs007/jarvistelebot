package plugindtdata

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/dtdata/proto"
)

// cmdGetDTData - getdtdata
type cmdGetDTData struct {
}

// RunCommand - run command
func (cmd *cmdGetDTData) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		cmdGetDTData, ok := params.CommandLine.(*plugindtdatapb.GetDTDataCommand)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidCommandLine.Error(), params.Msg)

			return false
		}

		pluginDTData, ok := params.CurPlugin.(*dtdataPlugin)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidPluginType.Error(), params.Msg)

			return false
		}

		reply, err := pluginDTData.client.getDTData(ctx, cmdGetDTData.Mode, cmdGetDTData.StartTime, cmdGetDTData.EndTime)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		str, err := chatbot.FormatJSONObj(reply)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
		} else {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), str, params.Msg)
		}

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdGetDTData) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("getdtdata", pflag.ContinueOnError)

	var mode = flagset.StringP("mode", "m", "", "What data do you want?")
	var startTime = flagset.StringP("starttime", "s", "", "start time")
	var endTime = flagset.StringP("endtime", "e", "", "end time")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *mode != "" {
		uac := &plugindtdatapb.GetDTDataCommand{
			Mode:      *mode,
			StartTime: *startTime,
			EndTime:   *endTime,
		}

		return uac, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
