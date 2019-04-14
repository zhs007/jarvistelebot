package plugincrawler

import (
	"context"
	"encoding/json"

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

		lst, err := pluginCrawler.client.getArticles(ctx, gacmd.URL, gacmd.AttachJQuery)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
				err.Error(), params.Msg)
		}

		jret, err := json.Marshal(lst)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		strret, err := chatbot.FormatJSON(string(jret))
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), string(jret), params.Msg)
		} else {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret, params.Msg)
		}
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

	var url = flagset.StringP("url", "u", "", "you want export this url.")
	var jquery = flagset.BoolP("jquery", "j", false, "attach jquery")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *url != "" {
		uac := &plugincrawlerpb.GetArticlesCommand{
			URL:          *url,
			AttachJQuery: *jquery,
		}

		return uac, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
