package plugingeneratepwd

import (
	"context"

	"github.com/golang/protobuf/proto"

	"github.com/zhs007/jarvistelebot/chatbot"
)

// PluginName - plugin name
const PluginName = "generatepassword"

// generatePasswordPlugin - generate password plugin
type generatePasswordPlugin struct {
	cmd *chatbot.CommandMap
}

// NewPlugin - new duckling plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - generatePasswordPlugin")

	// cfg := loadConfig(path.Join(cfgPath, "duckling.yaml"))
	// err := checkConfig(cfg)
	// if err != nil {
	// 	jarvisbase.Warn("pluginduckling.NewPlugin:checkConfig")

	// 	return nil, err
	// }

	cmd := chatbot.NewCommandMap()

	cmd.AddCommand("generatepassword", &cmdGeneratePassword{})

	// dbCrawler, err := newCrawlerDB(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)
	// if err != nil {
	// 	jarvisbase.Warn("plugincrawler.NewPlugin:newCrawlerDB")

	// 	return nil, err
	// }

	p := &generatePasswordPlugin{
		cmd: cmd,
	}

	return p, nil
}

// OnMessage - get message
func (p *generatePasswordPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
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
func (p *generatePasswordPlugin) GetPluginName() string {
	return PluginName
}

// OnStart - on start
func (p *generatePasswordPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *generatePasswordPlugin) GetPluginType() int {
	return chatbot.PluginTypeWritableCommand
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *generatePasswordPlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
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
