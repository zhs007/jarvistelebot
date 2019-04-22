package pluginduckling

import (
	"context"
	"path"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarviscore/base"

	"github.com/zhs007/jarvistelebot/chatbot"
)

// PluginName - plugin name
const PluginName = "duckling"

// ducklingPlugin - dtdata plugin
type ducklingPlugin struct {
	cmd    *chatbot.CommandMap
	cfg    *config
	client *ducklingClient
}

// NewPlugin - new duckling plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - ducklingPlugin")

	cfg := loadConfig(path.Join(cfgPath, "duckling.yaml"))
	err := checkConfig(cfg)
	if err != nil {
		jarvisbase.Warn("pluginduckling.NewPlugin:checkConfig")

		return nil, err
	}

	cmd := chatbot.NewCommandMap()

	cmd.AddCommand("duckling", &cmdRequestDuckling{})

	// dbCrawler, err := newCrawlerDB(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)
	// if err != nil {
	// 	jarvisbase.Warn("plugincrawler.NewPlugin:newCrawlerDB")

	// 	return nil, err
	// }

	p := &ducklingPlugin{
		cmd:    cmd,
		cfg:    cfg,
		client: newDucklingClient(cfg),
	}

	return p, nil
}

// OnMessage - get message
func (p *ducklingPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if !params.ChatBot.IsMaster(from) {
		return false, nil
	}

	if params.CommandLine != nil {
		p.cmd.Run(ctx, params.LstStr[0], params)

		return true, nil
	}

	// if len(params.LstStr) >= 1 {
	// 	ret := p.urlParser.ParseURL(params.LstStr[0])
	// 	if ret != nil {
	// 		if ret.URLType == "article" {
	// 			return runExportArticle(ctx, params, &plugincrawlerpb.ExpArticleCommand{
	// 				URL: ret.URL,
	// 				PDF: ret.PDF,
	// 			}), nil
	// 		}
	// 	}

	// 	p.cmd.Run(ctx, params.LstStr[0], params)

	// 	return true, nil
	// }

	return false, nil
}

// GetPluginName - get plugin name
func (p *ducklingPlugin) GetPluginName() string {
	return PluginName
}

// OnStart - on start
func (p *ducklingPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *ducklingPlugin) GetPluginType() int {
	return chatbot.PluginTypeWritableCommand
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *ducklingPlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) > 1 {
		// ret := p.urlParser.ParseURL(params.LstStr[0])
		// if ret != nil {
		// 	return &plugincrawlerpb.URLCommand{}, nil
		// }

		if p.cmd.HasCommand(params.LstStr[0]) {
			return p.cmd.ParseCommandLine(params.LstStr[0], params)
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
