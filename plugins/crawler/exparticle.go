package plugincrawler

import (
	"bytes"
	"context"
	"fmt"
	"path"

	"github.com/zhs007/jarviscore"
	jarvisbase "github.com/zhs007/jarviscore/base"
	jarviscorepb "github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
	chatbotdbpb "github.com/zhs007/jarvistelebot/chatbotdb/proto"
	plugincrawlerpb "github.com/zhs007/jarvistelebot/plugins/crawler/proto"
	"go.uber.org/zap"
)

// RunCommand - run command
func runExportArticle(ctx context.Context, params *chatbot.MessageParams, eacmd *plugincrawlerpb.ExpArticleCommand) bool {

	if params.CommandLine != nil {
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

		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
			"I will execute the script exparticle for "+curnode.Name, params.Msg)

		buf, err := expArticle(&ExpArticleParam{
			CrawlerPath: pluginCrawler.cfg.CrawlerPath,
			URL:         eacmd.URL,
			PDF:         eacmd.PDF,
		}, pluginCrawler.cfg.ExpArticleScript)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		sf := &jarviscorepb.FileData{
			Filename: "exparticle",
			File:     buf,
		}

		ci, err := jarviscore.BuildCtrlInfoForScriptFile3(sf, []string{
			path.Join(pluginCrawler.cfg.CrawlerPath, "./output/", eacmd.PDF),
		}, "exparticle")
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		var filebuf bytes.Buffer

		isrecv := false
		params.ChatBot.GetJarvisNode().RequestCtrl(ctx, curnode.Addr, ci,
			func(ctx context.Context, jarvisnode jarviscore.JarvisNode,
				lstResult []*jarviscore.JarvisMsgInfo) error {

				if !isrecv && len(lstResult) > 0 {
					if lstResult[len(lstResult)-1].JarvisResultType == jarviscore.JarvisResultTypeSend {

						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							"I send the script exparticle for "+curnode.Name, params.Msg)

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
								fmt.Sprintf("%v has received the request (exparticle).",
									curnode.Name),
								params.Msg)

							params.ChatBot.AddJarvisMsgCallback(curnode.Addr, cm.ReplyMsgID,
								func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
									cr := msg.GetCtrlResult()
									if cr == nil {
										msgstr := fmt.Sprintf("%v", msg)
										jarvisbase.Warn("userscriptPlugin.AddJarvisMsgCallback", zap.String("msg", msgstr))

										chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), msgstr, params.Msg)

										return nil
									}

									if cr.CtrlResult != "" {
										chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), cr.CtrlResult, params.Msg)
									}

									if cr.ErrInfo != "" {
										chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), cr.ErrInfo, params.Msg)
									}

									// chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "It's done.", params.Msg)

									return nil
								})
						} else if cm.MsgType == jarviscorepb.MSGTYPE_REPLY_REQUEST_FILE {
							isend, err := chatbot.ProcReplyRequestFile(cm, &filebuf)
							if err != nil {
								chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
									err.Error(), params.Msg)

								return err
							}

							if isend {
								chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
									fmt.Sprintf("The %v received %v bytes, the file is received, I will send it to you.",
										eacmd.PDF, filebuf.Len()),
									params.Msg)

								chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
									Filename: eacmd.PDF,
									Data:     filebuf.Bytes(),
								}, params.Msg)

								filebuf.Reset()
							} else {
								chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
									fmt.Sprintf("The %v received %v bytes.",
										eacmd.PDF, filebuf.Len()),
									params.Msg)
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
