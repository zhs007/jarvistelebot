package plugincrawler

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarviscore"
	jarvisbase "github.com/zhs007/jarviscore/base"
	jarviscorepb "github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
	plugincrawlerpb "github.com/zhs007/jarvistelebot/plugins/crawler/proto"
	"go.uber.org/zap"
)

// cmdUpdCrawler - updcrawler
type cmdUpdCrawler struct {
}

// RunCommand - run command
func (cmd *cmdUpdCrawler) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		_, ok := params.CommandLine.(*plugincrawlerpb.UpdCrawlerCommand)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidCommandLine.Error(), params.Msg)

			return false
		}

		if params.CurPlugin == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidParamsNoCurPlugin.Error(), params.Msg)

			return false
		}

		pluginCrawler, ok := params.CurPlugin.(*crawlerPlugin)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidParamsInvalidCurPlugin.Error(), params.Msg)

			return false
		}

		curnode := params.ChatBot.GetJarvisNode().FindNode(pluginCrawler.cfg.CrawlerNodeAddr)
		if curnode == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "Sorry, I can't find crawler node.", params.Msg)

			return false
		}

		chatbot.SendTextMsg(params.ChatBot, from,
			"I will execute the script updcrawler for "+curnode.Name, params.Msg)

		buf, err := updateCrawler(
			&UpdateCrawlerParam{CrawlerPath: pluginCrawler.cfg.CrawlerPath},
			pluginCrawler.cfg.UpdateScript)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		sf := &jarviscorepb.FileData{
			Filename: "updcrawler",
			File:     buf,
		}

		ci, err := jarviscore.BuildCtrlInfoForScriptFile2(sf, nil, "updcrawler")
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		isrecv := false
		params.ChatBot.GetJarvisNode().RequestCtrl(ctx, curnode.Addr, ci,
			func(ctx context.Context, jarvisnode jarviscore.JarvisNode,
				lstResult []*jarviscore.JarvisMsgInfo) error {

				if !isrecv && len(lstResult) > 0 {
					if lstResult[len(lstResult)-1].JarvisResultType == jarviscore.JarvisResultTypeSend {

						chatbot.SendTextMsg(params.ChatBot, from,
							"I send the script updcrawler for "+curnode.Name, params.Msg)

					} else if lstResult[len(lstResult)-1].JarvisResultType == jarviscore.JarvisResultTypeRemoved {

						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							jarviscore.AppendString(curnode.Name, " maybe restarted, you can restart the crawler."), params.Msg)

					} else if lstResult[len(lstResult)-1].Err != nil {

						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							lstResult[len(lstResult)-1].Err.Error(), params.Msg)

					} else if lstResult[len(lstResult)-1].Msg != nil {

						cm := lstResult[len(lstResult)-1].Msg
						if cm.MsgType == jarviscorepb.MSGTYPE_REPLY2 &&
							cm.ReplyType == jarviscorepb.REPLYTYPE_ERROR {

							chatbot.SendTextMsg(params.ChatBot,
								params.Msg.GetFrom(),
								fmt.Sprintf("%v reply err %v.",
									curnode.Name, cm.Err),
								params.Msg)

						} else if cm.MsgType == jarviscorepb.MSGTYPE_REPLY2 &&
							cm.ReplyType == jarviscorepb.REPLYTYPE_ISME {

							chatbot.SendTextMsg(params.ChatBot,
								params.Msg.GetFrom(),
								fmt.Sprintf("%v has received the request (updcrawler).",
									curnode.Name),
								params.Msg)

							params.ChatBot.AddJarvisMsgCallback(curnode.Addr, cm.ReplyMsgID,
								func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
									cr := msg.GetCtrlResult()
									if cr == nil {
										msgstr := fmt.Sprintf("%v", msg)
										jarvisbase.Warn("userscriptPlugin.AddJarvisMsgCallback", zap.String("msg", msgstr))

										chatbot.SendTextMsg(params.ChatBot, from, msgstr, params.Msg)

										return nil
									}

									if cr.CtrlResult != "" {
										chatbot.SendTextMsg(params.ChatBot, from, cr.CtrlResult, params.Msg)
									}

									if cr.ErrInfo != "" {
										chatbot.SendTextMsg(params.ChatBot, from, cr.ErrInfo, params.Msg)
									}

									chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "It's done.", params.Msg)

									return nil
								})
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
func (cmd *cmdUpdCrawler) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	return &plugincrawlerpb.UpdCrawlerCommand{}, nil
}
