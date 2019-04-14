package plugincrawler

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/crawler/proto"
)

// cmdGetArticles - getarticles
type cmdGetArticles struct {
}

// RunCommand - run command
func (cmd *cmdGetArticles) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		gacmd, ok := params.CommandLine.(*plugincrawlerpb.GetArticlesCommand)
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

		lst, err := pluginCrawler.client.getArticles(ctx, gacmd.Website)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
				err.Error(), params.Msg)

			return true
		}

		for _, v := range lst.Articles {
			if v.Summary != "" {
				chatbot.SendMarkdownMsg(params.ChatBot, params.Msg.GetFrom(),
					fmt.Sprintf("[*%v*](%v)\n%v", v.Title, v.Url, v.Summary), params.Msg)
			} else {
				chatbot.SendMarkdownMsg(params.ChatBot, params.Msg.GetFrom(),
					fmt.Sprintf("[*%v*](%v)", v.Title, v.Url), params.Msg)
			}
		}

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
func (cmd *cmdGetArticles) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("getarticles", pflag.ContinueOnError)

	var website = flagset.StringP("website", "w", "", "you want get this website.")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *website != "" {
		uac := &plugincrawlerpb.GetArticlesCommand{
			Website: *website,
		}

		return uac, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
