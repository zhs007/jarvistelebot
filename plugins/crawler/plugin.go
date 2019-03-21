package plugincrawler

import (
	"context"
	"path"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarviscore/base"
	"github.com/zhs007/jarviscore/proto"
	"go.uber.org/zap"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// PluginName - plugin name
const PluginName = "crawler"

// crawlerPlugin - crawler plugin
type crawlerPlugin struct {
	cmd *chatbot.CommandMap
	cfg *config
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
		cmd: cmd,
		cfg: cfg,
	}

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

	file := params.Msg.GetFile()
	if file != nil {
		if file.FileType == chatbot.FileTypeShellScript {
			sf := &jarviscorepb.FileData{
				Filename: file.Filename,
				File:     file.Data,
			}
			ci, err := jarviscore.BuildCtrlInfoForScriptFile2(1, sf, nil)
			if err != nil {
				jarvisbase.Warn("jarvisnodeexPlugin.OnMessage", zap.Error(err))

				return false, err
			}

			curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(params.Msg.GetText())
			if curnode == nil {
				return false, nil
			}

			params.ChatBot.GetJarvisNode().RequestCtrl(ctx, curnode.Addr, ci, nil)

			params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
				cr := msg.GetCtrlResult()

				chatbot.SendTextMsg(params.ChatBot, from, cr.CtrlResult, params.Msg)

				return nil
			})

			return true, nil
		}

		return false, nil
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
