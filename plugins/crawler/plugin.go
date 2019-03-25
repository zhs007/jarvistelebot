package plugincrawler

import (
	"context"
	"path"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarviscore/base"

	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/crawler/proto"
)

// PluginName - plugin name
const PluginName = "crawler"

// crawlerPlugin - crawler plugin
type crawlerPlugin struct {
	cmd       *chatbot.CommandMap
	cfg       *config
	urlParser *URLParser
}

// NewPlugin - new jarvisnode plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - crawlerPlugin")

	cfg := loadConfig(path.Join(cfgPath, "crawler.yaml"))
	err := checkConfig(cfg)
	if err != nil {
		jarvisbase.Warn("plugincrawler.NewPlugin:checkConfig")

		return nil, err
	}

	cmd := chatbot.NewCommandMap()

	cmd.AddCommand("exparticle", &cmdExpArticle{})
	cmd.AddCommand("updcrawler", &cmdUpdCrawler{})

	p := &crawlerPlugin{
		cmd:       cmd,
		cfg:       cfg,
		urlParser: &URLParser{},
	}

	p.urlParser.Reg(articleSMZDM, parseArticleSMZDM)

	return p, nil
}

// OnMessage - get message
func (p *crawlerPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if !params.ChatBot.IsMaster(from) {
		return false, nil
	}

	ret := p.urlParser.ParseURL(params.LstStr[0])
	if ret != nil {
		if ret.URLType == "article" {
			return runExportArticle(ctx, params, &plugincrawlerpb.ExpArticleCommand{
				URL: ret.URL,
				PDF: ret.PDF,
			}), nil
		}
	}

	if len(params.LstStr) >= 1 {
		p.cmd.Run(ctx, params.LstStr[0], params)

		return true, nil
	}

	return false, nil
}

// GetPluginName - get plugin name
func (p *crawlerPlugin) GetPluginName() string {
	return PluginName
}

// OnStart - on start
func (p *crawlerPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *crawlerPlugin) GetPluginType() int {
	return chatbot.PluginTypeWritableCommand
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *crawlerPlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) >= 1 {
		if p.cmd.HasCommand(params.LstStr[0]) {
			return p.cmd.ParseCommandLine(params.LstStr[0], params)
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
