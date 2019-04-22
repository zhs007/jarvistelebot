package plugincrawler

import (
	"context"
	"fmt"

	"github.com/zhs007/jarviscore/base"

	"github.com/zhs007/jarviscore"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/crawler/proto"
)

// cmdSubscribeArticles - subscribe articles
type cmdSubscribeArticles struct {
}

// RunCommand - run command
func (cmd *cmdSubscribeArticles) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		sacmd, ok := params.CommandLine.(*plugincrawlerpb.SubscribeArticlesCommand)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidCommandLine.Error(), params.Msg)

			return false
		}

		pluginCrawler, ok := params.CurPlugin.(*crawlerPlugin)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
				chatbot.ErrInvalidParamsInvalidCurPlugin.Error(), params.Msg)

			return false
		}

		timerid := params.ChatBot.AddTimer(int(sacmd.Timer), -1,
			jarviscore.AppendString(params.Msg.GetFrom().GetUserID(), "cmdSubscribeArticles.timer"),
			func(ctx context.Context) error {
				jarvisbase.Info("cmdSubscribeArticles.timer")

				for _, website := range sacmd.Websites {
					lst, err := pluginCrawler.client.getArticles(ctx, website)
					if err != nil {
						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							err.Error(), params.Msg)

						return err
					}

					for _, v := range lst.Articles {
						isok, err := pluginCrawler.dbCrawler.addArticle(ctx,
							params.Msg.GetFrom().GetUserID(), website, v.Url)

						if err != nil {
							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
								err.Error(), params.Msg)
						}

						if isok {
							if v.Lang == "en" {
								zhtitle, err := pluginCrawler.client.translate(ctx, v.Title, "en", "zh-CN")
								if err != nil {
									chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
										err.Error(), params.Msg)
								}

								if v.Summary != "" {
									zhsummary, err := pluginCrawler.client.translate(ctx, v.Summary, "en", "zh-CN")
									if err != nil {
										chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
											err.Error(), params.Msg)
									}

									if v.Premium {
										chatbot.SendMarkdownMsg(params.ChatBot, params.Msg.GetFrom(),
											fmt.Sprintf("[%v](%v) Premium\n%v\n%v\n%v",
												v.Title, v.Url, zhtitle, v.Summary, zhsummary), params.Msg)
									} else {
										chatbot.SendMarkdownMsg(params.ChatBot, params.Msg.GetFrom(),
											fmt.Sprintf("[%v](%v)\n%v\n%v\n%v",
												v.Title, v.Url, zhtitle, v.Summary, zhsummary), params.Msg)
									}
								} else {
									if v.Premium {
										chatbot.SendMarkdownMsg(params.ChatBot, params.Msg.GetFrom(),
											fmt.Sprintf("[%v](%v) Premium\n%v", v.Title, v.Url, zhtitle), params.Msg)
									} else {
										chatbot.SendMarkdownMsg(params.ChatBot, params.Msg.GetFrom(),
											fmt.Sprintf("[%v](%v)\n%v", v.Title, v.Url, zhtitle), params.Msg)
									}
								}
							} else {
								if v.Summary != "" {
									if v.Premium {
										chatbot.SendMarkdownMsg(params.ChatBot, params.Msg.GetFrom(),
											fmt.Sprintf("[%v](%v) Premium\n%v", v.Title, v.Url, v.Summary), params.Msg)
									} else {
										chatbot.SendMarkdownMsg(params.ChatBot, params.Msg.GetFrom(),
											fmt.Sprintf("[%v](%v)\n%v", v.Title, v.Url, v.Summary), params.Msg)
									}
								} else {
									if v.Premium {
										chatbot.SendMarkdownMsg(params.ChatBot, params.Msg.GetFrom(),
											fmt.Sprintf("[%v](%v) Premium", v.Title, v.Url), params.Msg)
									} else {
										chatbot.SendMarkdownMsg(params.ChatBot, params.Msg.GetFrom(),
											fmt.Sprintf("[%v](%v)", v.Title, v.Url), params.Msg)
									}
								}
							}
						}
					}
				}
				return nil
			})

		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
			fmt.Sprintf("It's done, the timerid is %v", timerid), params.Msg)

		// jret, err := json.Marshal(lst)
		// if err != nil {
		// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

		// 	return false
		// }

		// strret, err := chatbot.FormatJSON(string(jret))
		// if err != nil {
		// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), string(jret), params.Msg)
		// } else {
		// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret, params.Msg)
		// }
		// return runExportArticle(ctx, params, eacmd)
	}

	return false
}

// Parse - parse command line
func (cmd *cmdSubscribeArticles) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("subscribearticles", pflag.ContinueOnError)

	var websites = flagset.StringSliceP("website", "w", []string{}, "you want to subscribe to this website.")
	var timer = flagset.Int32P("timer", "t", 5*60, "timer")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if len(*websites) != 0 && *timer > 0 {
		uac := &plugincrawlerpb.SubscribeArticlesCommand{
			Websites: *websites,
			Timer:    *timer,
		}

		return uac, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
