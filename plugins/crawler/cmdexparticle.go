package plugincrawler

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/crawler/proto"
)

// cmdExpArticle - exparticle
type cmdExpArticle struct {
}

// RunCommand - run command
func (cmd *cmdExpArticle) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		eacmd, ok := params.CommandLine.(*plugincrawlerpb.ExpArticleCommand)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidCommandLine.Error(), params.Msg)

			return false
		}

		return runExportArticle(ctx, params, eacmd)
	}

	return false
}

// Parse - parse command line
func (cmd *cmdExpArticle) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("runscript", pflag.ContinueOnError)

	var url = flagset.StringP("url", "u", "", "you want export this url.")
	var pdf = flagset.StringP("pdf", "p", "", "you want get this pdf file.")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *url != "" && *pdf != "" {
		uac := &plugincrawlerpb.ExpArticleCommand{
			URL: *url,
			PDF: *pdf,
		}

		return uac, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
