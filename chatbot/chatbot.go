package chatbot

import (
	"context"

	"github.com/zhs007/ankadb"
	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbotdb"
)

// ChatBot - chat bot interface
type ChatBot interface {
	// Init
	Init(cfgfilename string, mgr PluginsMgr) error
	// Start
	Start(ctx context.Context, node jarviscore.JarvisNode) error
	// SendMsg
	SendMsg(user User, text string) error
	// SaveMsg
	SaveMsg(msg Message) error
	// GetJarvisNodeCoreDB - get jarvis node coredb
	GetJarvisNodeCoreDB() *jarviscore.CoreDB
	// GetJarvisNode - get jarvis node
	GetJarvisNode() jarviscore.JarvisNode
	// GetConfig - get Config
	GetConfig() *Config

	// OnJarvisCtrlResult - event handle
	OnJarvisCtrlResult(ctx context.Context, msg *jarviscorepb.JarvisMsg) error
}

// BaseChatBot - base chatbot
type BaseChatBot struct {
	ChatBotDB  *ankadb.AnkaDB
	Node       jarviscore.JarvisNode
	MgrPlugins PluginsMgr
	cfg        *Config
}

const querySaveMsg = `mutation NewMsg($chatID: ID!, $fromNickName: String!, $fromUserID: ID!, $text: String!, $timeStamp: Timestamp!) {
	newMsg(chatID: $chatID, fromNickName: $fromNickName, fromUserID: $fromUserID, text: $text, timeStamp: $timeStamp) {
		chatID
	}
}`

// Init - init
func (base *BaseChatBot) Init(cfgfilename string, mgr PluginsMgr) error {
	cfg, err := LoadConfig(cfgfilename)
	if err != nil {
		return err
	}

	db, err := chatbotdb.NewChatBotDB(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)
	if err != nil {
		return err
	}

	base.ChatBotDB = db
	base.MgrPlugins = mgr
	base.cfg = cfg

	return nil
}

// SaveMsg - save message
func (base *BaseChatBot) SaveMsg(msg Message) error {
	if base.ChatBotDB == nil {
		return ErrChatBotDBNil
	}

	params := make(map[string]interface{})
	params["chatID"] = msg.GetChatID()
	params["fromNickName"] = msg.GetFrom().GetNickName()
	params["fromUserID"] = msg.GetFrom().GetUserID()
	params["text"] = msg.GetText()
	params["timeStamp"] = msg.GetTimeStamp()

	result, err := base.ChatBotDB.LocalQuery(context.Background(), querySaveMsg, params)
	if err != nil {
		return err
	}

	Info("BaseChatBot.SaveMsg",
		JSON("result", result))

	return nil
}

// Start - start chatbot
func (base *BaseChatBot) Start(ctx context.Context, node jarviscore.JarvisNode) error {
	base.Node = node

	go base.MgrPlugins.OnStart(ctx)

	return nil
}

// GetJarvisNodeCoreDB - get jarvis node coredb
func (base *BaseChatBot) GetJarvisNodeCoreDB() *jarviscore.CoreDB {
	return base.Node.GetCoreDB()
}

// GetJarvisNode - get jarvis node
func (base *BaseChatBot) GetJarvisNode() jarviscore.JarvisNode {
	return base.Node
}

// GetConfig - get Config
func (base *BaseChatBot) GetConfig() *Config {
	return base.cfg
}
