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

// assistantPlugin - assistant plugin
type assistantPlugin struct {
	mgr assistant.Mgr
	cmd *chatbot.CommandMap

	// db *assistantdb.AssistantDB
}

// NewPlugin - reg assistant plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - assistantPlugin")

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

	return &assistantPlugin{
		mgr: mgr,
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
func (p *assistantPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if params.ChatBot.IsMaster(from) {
		if len(params.LstStr) >= 2 && params.LstStr[0] == ">>" {
			p.cmd.Run(ctx, params.LstStr[1], params)

			return true, nil
		}

		ct := p.mgr.GetCurNoteMode(from.GetUserID())
		if ct == assistant.ModeInvalidType {
			chatbot.SendTextMsg(params.ChatBot, from, "Sorry, I found some problems, please restart.")

			return true, assistant.ErrInvalidCurNoteMode
		} else if ct == assistant.ModeInputData {
			p.mgr.AddCurNoteData(from.GetUserID(), params.Msg.GetText())

			chatbot.SendTextMsg(params.ChatBot, from, "I recorded for note, and then?")
		} else if ct == assistant.ModeInputKey {
			p.mgr.AddCurNoteKey(from.GetUserID(), params.Msg.GetText())

			chatbot.SendTextMsg(params.ChatBot, from, "I recorded key for note, and then?")
		}

		return true, nil

		// if params.CommandLine != nil {
		// 	notecmd, ok := params.CommandLine.(*pluginassistanepb.NoteCommand)
		// 	if !ok {
		// 		return false, chatbot.ErrInvalidCommandLine
		// 	}

		// 	cn, err := p.mgr.NewNote(from.GetUserID())
		// 	if err != nil {
		// 		chatbot.SendTextMsg(params.ChatBot, from, err.Error())

		// 		return false, chatbot.ErrInvalidCommandLine
		// 	}

		// 	if len(notecmd.Keys) > 0 {
		// 		for _, v := range notecmd.Keys {
		// 			cn.Keys = append(notecmd.Keys, v)
		// 		}
		// 	}

		// 	chatbot.SendTextMsg(params.ChatBot, from, "I get it, please tell me what to record.")
		// } else {

		// }
		// // ip := p.parseInput(params)

		// if ip != nil {
		// 	// str := fmt.Sprintf("%+v", ip)
		// 	// jarvisbase.Debug("assistantPlugin.OnMessage:parseInput", zap.String("ret", str))

		// 	if ip.isSave {
		// 		msg, err := p.db.NewMsg(ip.dat, ip.keys)
		// 		if err != nil {
		// 			jarvisbase.Warn("assistantPlugin.OnMessage:NewMsg", zap.Error(err))

		// 			return false, err
		// 		}

		// 		chatbot.SendTextMsg(params.ChatBot, from, fmt.Sprintf("ok. current msg is %+v", msg))
		// 		// params.ChatBot.SendMsg(from, fmt.Sprintf("ok. current msg is %+v", msg))

		// 		return true, nil
		// 	}

		// 	msg, err := p.db.GetMsg(ip.msgid)
		// 	if err != nil {
		// 		jarvisbase.Warn("assistantPlugin.OnMessage:GetMsg", zap.Error(err))

		// 		return false, err
		// 	}

		// 	chatbot.SendTextMsg(params.ChatBot, from, fmt.Sprintf("ok. msg is %+v", msg))
		// 	// params.ChatBot.SendMsg(from, fmt.Sprintf("ok. msg is %+v", msg))

		// 	return true, nil

		// }
	}

	chatbot.SendTextMsg(params.ChatBot, from, "Sorry, you are not my master.")
	// params.ChatBot.SendMsg(from, "sorry, you are not my master.")

	return false, nil
}

// GetPluginName - get plugin name
func (p *assistantPlugin) GetPluginName() string {
	return PluginName
}

// // IsMyMessage
// func (p *assistantPlugin) IsMyMessage(params *chatbot.MessageParams) bool {
// 	// if len(params.LstStr) > 1 && params.LstStr[0] == ">>" {
// 	// 	for i := 2; i < len(params.LstStr)-1; i++ {
// 	// 		if params.LstStr[i] == ">" {
// 	// 			return true
// 	// 		}
// 	// 	}
// 	// }

// 	// if len(params.LstStr) == 3 && params.LstStr[0] == "<<" && params.LstStr[1] == "@" {
// 	// 	_, err := strconv.ParseInt(params.LstStr[2], 10, 64)
// 	// 	if err == nil {
// 	// 		return true
// 	// 	}
// 	// }

// 	return false
// }

// // parseInput
// func (p *assistantPlugin) parseInput(params *chatbot.MessageParams) *inputParams {
// 	if len(params.LstStr) > 1 && params.LstStr[0] == ">>" {
// 		ip := &inputParams{
// 			isSave: true,
// 		}

// 		ck := ""

// 		for i := 1; i < len(params.LstStr)-1; i++ {
// 			if params.LstStr[i] == ">" {
// 				if ck == "" {
// 					return nil
// 				}

// 				ip.keys = append(ip.keys, ck)

// 				ip.dat = ""
// 				for j := i + 1; j < len(params.LstStr); j++ {
// 					if ip.dat == "" {
// 						ip.dat += params.LstStr[j]
// 					} else {
// 						ip.dat += " " + params.LstStr[j]
// 					}

// 				}

// 				return nil
// 			}

// 			if ck == "" {
// 				ck += params.LstStr[i]
// 			} else {
// 				ck += " " + params.LstStr[i]
// 			}
// 		}

// 		return ip
// 	}

// 	if len(params.LstStr) == 3 && params.LstStr[0] == "<<" && params.LstStr[1] == "@" {
// 		msgid, err := strconv.ParseInt(params.LstStr[2], 10, 64)
// 		if err != nil {
// 			return nil
// 		}

// 		return &inputParams{
// 			isSave: false,
// 			msgid:  msgid,
// 		}
// 	}

// 	return nil
// }

// OnStart - on start
func (p *assistantPlugin) OnStart(ctx context.Context) error {
	return p.mgr.Start(ctx)
}

// GetPluginType - get pluginType
func (p *assistantPlugin) GetPluginType() int {
	return chatbot.PluginTypeWritableCommand
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *assistantPlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) >= 2 && params.LstStr[0] == ">>" {
		if p.cmd.HasCommand(params.LstStr[1]) {
			return p.cmd.ParseCommandLine(params.LstStr[1], params)
		}
	}

	// if len(params.LstStr) >= 1 && params.LstStr[0] == ">>" {
	// 	if params.LstStr[1] == "note" {
	// 		if len(params.LstStr) >= 2 {
	// 			flagset := pflag.NewFlagSet("note", pflag.ContinueOnError)

	// 			var keys = flagset.StringSliceP("key", "k", []string{}, "you can set keywords")

	// 			err := flagset.Parse(params.LstStr[2:])
	// 			if err != nil {
	// 				return nil, err
	// 			}

	// 			return &pluginassistanepb.NoteCommand{
	// 				Keys: *keys,
	// 			}, nil
	// 		}

	// 		return &pluginassistanepb.NoteCommand{}, nil
	// 	}
	// }

	return nil, chatbot.ErrMsgNotMine
}
