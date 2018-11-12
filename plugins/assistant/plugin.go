package pluginassistant

import (
	"context"
	"fmt"
	"path"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/zhs007/jarvistelebot/assistantdb"
	"github.com/zhs007/jarvistelebot/chatbot"
)

// assistantPlugin - assistant plugin
type assistantPlugin struct {
	db *assistantdb.AssistantDB
}

// RegPlugin - reg assistant plugin
func RegPlugin(cfgPath string, mgr chatbot.PluginsMgr) error {
	chatbot.Info("RegPlugin - assistantPlugin")

	cfg := loadConfig(path.Join(cfgPath, "assistant.yaml"))

	db, err := assistantdb.NewAssistantDB(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)
	if err != nil {
		return err
	}

	mgr.RegPlugin(&assistantPlugin{
		db: db,
	})

	return nil
}

// OnMessage - get message
func (p *assistantPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if from.IsMaster() {
		dat, keys := p.parseInput(params)

		str := fmt.Sprintf("%v%v", dat, keys)
		jarvisbase.Debug("assistantPlugin.OnMessage:parseInput", zap.String("ret", str))

		if len(dat) > 0 {
			msg, err := p.db.NewMsg(dat, keys)
			if err != nil {
				jarvisbase.Warn("assistantPlugin.OnMessage:NewMsg", zap.Error(err))

				return false, err
			}

			params.ChatBot.SendMsg(from, fmt.Sprintf("ok. current msgID is %+v", msg))

			return true, nil
		}
	} else {
		params.ChatBot.SendMsg(from, "sorry, you are not my master.")
	}

	return false, nil
}

// GetComeInCode - if return is empty string, it means not comein
func (p *assistantPlugin) GetComeInCode() string {
	return "assistant"
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

	return false
}

// parseInput
func (p *assistantPlugin) parseInput(params *chatbot.MessageParams) (dat string, keys []string) {
	ck := ""

	if len(params.LstStr) > 1 && params.LstStr[0] == ">>" {
		for i := 1; i < len(params.LstStr)-1; i++ {
			if params.LstStr[i] == ">" {
				if ck == "" {
					return
				}

				keys = append(keys, ck)

				dat = ""
				for j := i + 1; j < len(params.LstStr); j++ {
					if dat == "" {
						dat += params.LstStr[j]
					} else {
						dat += " " + params.LstStr[j]
					}

				}

				return
			}

			if ck == "" {
				ck += params.LstStr[i]
			} else {
				ck += " " + params.LstStr[i]
			}
		}
	}

	return
}

// OnStart - on start
func (p *assistantPlugin) OnStart(ctx context.Context) error {
	return p.db.Start(ctx)
}
