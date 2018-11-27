package pluginassistant

import (
	"context"
	"fmt"
	"path"
	"strconv"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/zhs007/jarvistelebot/assistantdb"
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
	db *assistantdb.AssistantDB
}

// NewPlugin - reg assistant plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - assistantPlugin")

	cfg := loadConfig(path.Join(cfgPath, "assistant.yaml"))

	db, err := assistantdb.NewAssistantDB(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)
	if err != nil {
		return nil, err
	}

	return &assistantPlugin{
		db: db,
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
		ip := p.parseInput(params)

		if ip != nil {
			str := fmt.Sprintf("%+v", ip)
			jarvisbase.Debug("assistantPlugin.OnMessage:parseInput", zap.String("ret", str))

			if ip.isSave {
				msg, err := p.db.NewMsg(ip.dat, ip.keys)
				if err != nil {
					jarvisbase.Warn("assistantPlugin.OnMessage:NewMsg", zap.Error(err))

					return false, err
				}

				chatbot.SendTextMsg(params.ChatBot, from, fmt.Sprintf("ok. current msg is %+v", msg))
				// params.ChatBot.SendMsg(from, fmt.Sprintf("ok. current msg is %+v", msg))

				return true, nil
			}

			msg, err := p.db.GetMsg(ip.msgid)
			if err != nil {
				jarvisbase.Warn("assistantPlugin.OnMessage:GetMsg", zap.Error(err))

				return false, err
			}

			chatbot.SendTextMsg(params.ChatBot, from, fmt.Sprintf("ok. msg is %+v", msg))
			// params.ChatBot.SendMsg(from, fmt.Sprintf("ok. msg is %+v", msg))

			return true, nil

		}
	} else {
		chatbot.SendTextMsg(params.ChatBot, from, "sorry, you are not my master.")
		// params.ChatBot.SendMsg(from, "sorry, you are not my master.")
	}

	return false, nil
}

// GetComeInCode - if return is empty string, it means not comein
func (p *assistantPlugin) GetComeInCode() string {
	return PluginName
}

// IsMyMessage
func (p *assistantPlugin) IsMyMessage(params *chatbot.MessageParams) bool {
	if len(params.LstStr) > 1 && params.LstStr[0] == ">>" {
		for i := 2; i < len(params.LstStr)-1; i++ {
			if params.LstStr[i] == ">" {
				return true
			}
		}
	}

	if len(params.LstStr) == 3 && params.LstStr[0] == "<<" && params.LstStr[1] == "@" {
		_, err := strconv.ParseInt(params.LstStr[2], 10, 64)
		if err == nil {
			return true
		}
	}

	return false
}

// parseInput
func (p *assistantPlugin) parseInput(params *chatbot.MessageParams) *inputParams {
	if len(params.LstStr) > 1 && params.LstStr[0] == ">>" {
		ip := &inputParams{
			isSave: true,
		}

		ck := ""

		for i := 1; i < len(params.LstStr)-1; i++ {
			if params.LstStr[i] == ">" {
				if ck == "" {
					return nil
				}

				ip.keys = append(ip.keys, ck)

				ip.dat = ""
				for j := i + 1; j < len(params.LstStr); j++ {
					if ip.dat == "" {
						ip.dat += params.LstStr[j]
					} else {
						ip.dat += " " + params.LstStr[j]
					}

				}

				return nil
			}

			if ck == "" {
				ck += params.LstStr[i]
			} else {
				ck += " " + params.LstStr[i]
			}
		}

		return ip
	}

	if len(params.LstStr) == 3 && params.LstStr[0] == "<<" && params.LstStr[1] == "@" {
		msgid, err := strconv.ParseInt(params.LstStr[2], 10, 64)
		if err != nil {
			return nil
		}

		return &inputParams{
			isSave: false,
			msgid:  msgid,
		}
	}

	return nil
}

// OnStart - on start
func (p *assistantPlugin) OnStart(ctx context.Context) error {
	return p.db.Start(ctx)
}

// GetPluginType - get pluginType
func (p *assistantPlugin) GetPluginType() int {
	return chatbot.PluginTypeCommand
}
