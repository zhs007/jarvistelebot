package chatbot

import (
	"context"

	"github.com/zhs007/ankadb"
	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarvistelebot/chatbotdb"
)

// ChatBot - chat bot interface
type ChatBot interface {
	// Init
	Init(dbpath string, httpAddr string, engine string) error
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
}

// BaseChatBot - base chatbot
type BaseChatBot struct {
	ChatBotDB  *ankadb.AnkaDB
	Node       jarviscore.JarvisNode
	mgrPlugins PluginsMgr
}

const querySaveMsg = `mutation NewMsg($chatID: ID!, $fromNickName: String!, $fromUserID: ID!, $text: String!, $timeStamp: Timestamp!) {
	newMsg(chatID: $chatID, fromNickName: $fromNickName, fromUserID: $fromUserID, text: $text, timeStamp: $timeStamp) {
		chatID
	}
}`

// Init - init
func (base *BaseChatBot) Init(dbpath string, httpAddr string, engine string) error {
	db, err := chatbotdb.NewChatBotDB(dbpath, httpAddr, engine)
	if err != nil {
		return err
	}

	base.ChatBotDB = db

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

	go base.mgrPlugins.OnStart(ctx)

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
