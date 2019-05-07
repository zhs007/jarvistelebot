package plugindtdata2

import (
	"context"
	"path"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarviscore/base"

	"github.com/zhs007/jarvistelebot/chatbot"
)

// PluginName - plugin name
const PluginName = "dtdata2"

// dtdata2Plugin - dtdata2 plugin
type dtdata2Plugin struct {
	cmd *chatbot.CommandMap
	cfg *config
}

// NewPlugin - new dtdata2 plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - dtdata2Plugin")

	cfg := loadConfig(path.Join(cfgPath, "dtdata2.yaml"))
	err := checkConfig(cfg)
	if err != nil {
		jarvisbase.Warn("plugindtdata2.NewPlugin:checkConfig")

		return nil, err
	}

	cmd := chatbot.NewCommandMap()

	cmd.AddCommand(CommandGameDayReport, &cmdGameDayReport{})

	p := &dtdata2Plugin{
		cmd: cmd,
		cfg: cfg,
	}

	return p, nil
}

// OnMessage - get message
func (p *dtdata2Plugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
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
func (p *dtdata2Plugin) GetPluginName() string {
	return PluginName
}

// OnStart - on start
func (p *dtdata2Plugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *dtdata2Plugin) GetPluginType() int {
	return chatbot.PluginTypeWritableCommand
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *dtdata2Plugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
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
