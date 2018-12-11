package pluginassistant

import (
	"context"
	"path"

	"github.com/golang/protobuf/proto"

	"github.com/zhs007/jarvistelebot/assistant"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// PluginName - plugin name
const PluginName = "assistant"

// inputParams - parse input string to inputParams
type inputParams struct {
	isSave bool
	dat    string
	keys   []string
	msgid  int64
}

// AssistantPlugin - assistant plugin
type AssistantPlugin struct {
	Mgr assistant.Mgr
	cmd *chatbot.CommandMap

	// db *assistantdb.AssistantDB
}

// NewPlugin - reg assistant plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - AssistantPlugin")

	cfg := loadConfig(path.Join(cfgPath, "assistant.yaml"))

	mgr, err := assistant.NewAssistantMgr(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)
	if err != nil {
		return nil, err
	}

	cmd := chatbot.NewCommandMap()

	cmd.AddCommand("note", &cmdNote{})
	cmd.AddCommand("endnote", &cmdEndNote{})
	cmd.AddCommand("endkey", &cmdEndKey{})
	cmd.AddCommand("mynotes", &cmdMyNotes{})
	cmd.AddCommand("rebuildkeywords", &cmdRebuildKeywords{})

	return &AssistantPlugin{
		Mgr: mgr,
		cmd: cmd,
	}, nil
}

// // RegPlugin - reg assistant plugin
// func RegPlugin(cfgPath string, mgr chatbot.PluginsMgr) error {
// 	chatbot.Info("RegPlugin - assistantPlugin")

// 	cfg := loadConfig(path.Join(cfgPath, "assistant.yaml"))

// 	db, err := assistantdb.NewAssistantDB(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)
// 	if err != nil {
// 		return err
// 	}

// 	mgr.RegPlugin(&assistantPlugin{
// 		db: db,
// 	})

// 	return nil
// }

// OnMessage - get message
func (p *AssistantPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if params.ChatBot.IsMaster(from) {
		if len(params.LstStr) >= 1 {
			p.cmd.Run(ctx, params.LstStr[0], params)

			return true, nil
		}

		ct := p.Mgr.GetCurNoteMode(from.GetUserID())
		if ct == assistant.ModeInvalidType {
			chatbot.SendTextMsg(params.ChatBot, from, "Sorry, I found some problems, please restart.", params.Msg)

			return true, assistant.ErrInvalidCurNoteMode
		} else if ct == assistant.ModeInputData {
			p.Mgr.AddCurNoteData(from.GetUserID(), params.Msg.GetText())

			chatbot.SendTextMsg(params.ChatBot, from, "I recorded for note, and then?", params.Msg)
		} else if ct == assistant.ModeInputKey {
			p.Mgr.AddCurNoteKey(from.GetUserID(), params.Msg.GetText())

			chatbot.SendTextMsg(params.ChatBot, from, "I recorded key for note, and then?", params.Msg)
		}

		return true, nil
	}

	chatbot.SendTextMsg(params.ChatBot, from, "Sorry, you are not my master.", params.Msg)

	return false, nil
}

// GetPluginName - get plugin name
func (p *AssistantPlugin) GetPluginName() string {
	return PluginName
}

// OnStart - on start
func (p *AssistantPlugin) OnStart(ctx context.Context) error {
	return p.Mgr.Start(ctx)
}

// GetPluginType - get pluginType
func (p *AssistantPlugin) GetPluginType() int {
	return chatbot.PluginTypeWritableCommand
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *AssistantPlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) >= 1 {
		if p.cmd.HasCommand(params.LstStr[0]) {
			return p.cmd.ParseCommandLine(params.LstStr[0], params)
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
