package plugindtdata2

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	dtdataclient "github.com/zhs007/dtdataserv/client"
	"github.com/zhs007/jarviscore"
	jarviscorepb "github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
	plugindtdata2pb "github.com/zhs007/jarvistelebot/plugins/dtdata2/proto"
)

// CommandGameDayReport - command
const CommandGameDayReport = "gamedayreport"

// cmdGameDayReport - gamedayreport
type cmdGameDayReport struct {
}

// RunCommand - run command
func (cmd *cmdGameDayReport) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		cmdGameDayReport, ok := params.CommandLine.(*plugindtdata2pb.GameDayReportCommand)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidCommandLine.Error(), params.Msg)

			return false
		}

		pluginDTData2, ok := params.CurPlugin.(*dtdata2Plugin)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidPluginType.Error(), params.Msg)

			return false
		}

		curnode := params.ChatBot.GetJarvisNode().FindNode(pluginDTData2.cfg.DTDataServAddr)
		if curnode == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrNoJarvisNode.Error(), params.Msg)

			return false
		}

		ci, err := dtdataclient.BuildCtrlInfoForGameDayReport(cmdGameDayReport.Env, cmdGameDayReport.DayTime,
			cmdGameDayReport.Currency, int(cmdGameDayReport.ScaleMoney))
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		params.ChatBot.GetJarvisNode().RequestCtrl(ctx, curnode.Addr, ci,
			func(ctx context.Context, jarvisnode jarviscore.JarvisNode,
				lstResult []*jarviscore.JarvisMsgInfo) error {

				if len(lstResult) > 0 {
					curret := lstResult[len(lstResult)-1]

					if curret.JarvisResultType == jarviscore.JarvisResultTypeSend {

						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							jarviscore.AppendString("I send the task to ", curnode.Name),
							params.Msg)

					} else if lstResult[len(lstResult)-1].JarvisResultType == jarviscore.JarvisResultTypeRemoved {

						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							jarviscore.AppendString(curnode.Name, " maybe restarted, you can restart the gamedayreport."), params.Msg)

					} else if curret.Err != nil {

						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							curret.Err.Error(), params.Msg)

					} else if curret.Msg != nil {
						cm := curret.Msg

						if cm.MsgType == jarviscorepb.MSGTYPE_REPLY2 &&
							cm.ReplyType == jarviscorepb.REPLYTYPE_ERROR {

							chatbot.SendTextMsg(params.ChatBot,
								params.Msg.GetFrom(),
								jarviscore.AppendString(curnode.Name, " reply err ", cm.Err, " ."),
								params.Msg)

						} else if cm.MsgType == jarviscorepb.MSGTYPE_REPLY2 &&
							cm.ReplyType == jarviscorepb.REPLYTYPE_END {

							chatbot.SendTextMsg(params.ChatBot, from, "It's done.", params.Msg)

						} else if cm.MsgType == jarviscorepb.MSGTYPE_REPLY2 &&
							cm.ReplyType == jarviscorepb.REPLYTYPE_ISME {

							chatbot.SendTextMsg(params.ChatBot,
								params.Msg.GetFrom(),
								jarviscore.AppendString(curnode.Name, " has received the request the task."),
								params.Msg)

						} else if cm.MsgType == jarviscorepb.MSGTYPE_REPLY_CTRL_RESULT {

							cr := cm.GetCtrlResult()
							if cr != nil {
								reply, err := dtdataclient.GetReplyFromCtrlResult(cr)
								if err != nil {
									chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
										curret.Err.Error(), params.Msg)

									return nil
								}

								chatbot.SendTextMsg(params.ChatBot, from,
									jarviscore.AppendString(pluginDTData2.cfg.URL, reply.Token),
									params.Msg)

								return nil
							}
						}
					}
				}

				return nil
			})

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdGameDayReport) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet(CommandGameDayReport, pflag.ContinueOnError)

	var env = flagset.StringP("env", "e", "", "Which environment data do you want to see?")
	var dayTime = flagset.StringP("daytime", "t", "", "What day do you want to see?")
	var currency = flagset.StringP("currency", "c", "", "Which currency do you want to see?")
	var scalemoney = flagset.Int32P("scalemoney", "s", -1, "scale money")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *env != "" && *dayTime != "" && *currency != "" {
		gdrc := &plugindtdata2pb.GameDayReportCommand{
			Env:        *env,
			DayTime:    *dayTime,
			Currency:   *currency,
			ScaleMoney: *scalemoney,
		}

		return gdrc, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
