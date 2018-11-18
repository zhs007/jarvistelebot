package chatbot

import (
	"context"
	"time"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbotdb"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

// ChatBot - chat bot interface
type ChatBot interface {
	// Init -
	Init(cfgfilename string, mgr PluginsMgr) error
	// Start -
	Start(ctx context.Context, node jarviscore.JarvisNode) error
	// SendMsg -
	SendMsg(msg Message) error
	// SaveMsg -
	SaveMsg(msg Message) error
	// NewMsg -
	NewMsg(chatid string, msgid string, from User, to User, text string, curtime int64) Message
	// NewMsgFromProto -
	NewMsgFromProto(msg *chatbotdbpb.Message) Message
	// GetMsg -
	GetMsg(chatid string) (Message, error)

	// GetJarvisNodeCoreDB - get jarvis node coredb
	GetJarvisNodeCoreDB() *jarviscore.CoreDB
	// GetJarvisNode - get jarvis node
	GetJarvisNode() jarviscore.JarvisNode

	// GetConfig - get Config
	GetConfig() *Config
	// GetChatBotDB - get ChatBotDB
	GetChatBotDB() *chatbotdb.ChatBotDB

	// IsMaster - is master
	IsMaster(user User) bool
	// GetUserMgr - get user manager
	GetUserMgr() UserMgr
	// NewUserFromProto - new user from proto
	NewUserFromProto(user *chatbotdbpb.User) User
	// GetUser -
	GetUser(userid string) (User, error)

	// OnJarvisCtrlResult - event handle
	OnJarvisCtrlResult(ctx context.Context, msg *jarviscorepb.JarvisMsg) error
}

// BasicChatBot - base chatbot
type BasicChatBot struct {
	DB         *chatbotdb.ChatBotDB
	Node       jarviscore.JarvisNode
	MgrPlugins PluginsMgr
	Config     *Config
	MgrUser    UserMgr
}

// NewBasicChatBot - new BasicChatBot
func NewBasicChatBot() *BasicChatBot {
	return &BasicChatBot{}
}

// Init - init
func (base *BasicChatBot) Init(cfgfilename string, mgr PluginsMgr) error {
	cfg, err := LoadConfig(cfgfilename)
	if err != nil {
		return err
	}

	db, err := chatbotdb.NewChatBotDB(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)
	if err != nil {
		return err
	}

	base.DB = db
	base.MgrPlugins = mgr
	base.Config = cfg

	return nil
}

// SaveMsg - save message
func (base *BasicChatBot) SaveMsg(msg Message) error {
	return base.DB.SaveMsg(msg.ToProto())
}

// Start - start chatbot
func (base *BasicChatBot) Start(ctx context.Context, node jarviscore.JarvisNode) error {
	base.Node = node

	go base.MgrPlugins.OnStart(ctx)

	return nil
}

// GetJarvisNodeCoreDB - get jarvis node coredb
func (base *BasicChatBot) GetJarvisNodeCoreDB() *jarviscore.CoreDB {
	return base.Node.GetCoreDB()
}

// GetJarvisNode - get jarvis node
func (base *BasicChatBot) GetJarvisNode() jarviscore.JarvisNode {
	return base.Node
}

// GetConfig - get Config
func (base *BasicChatBot) GetConfig() *Config {
	return base.Config
}

// IsMaster - is master
func (base *BasicChatBot) IsMaster(user User) bool {
	return base.MgrUser.IsMaster(user)
}

// GetUserMgr - get user manager
func (base *BasicChatBot) GetUserMgr() UserMgr {
	return base.MgrUser
}

// GetChatBotDB - get ChatBotDB
func (base *BasicChatBot) GetChatBotDB() *chatbotdb.ChatBotDB {
	return base.DB
}

// SendTextMsg - sendmsg
func SendTextMsg(bot ChatBot, user User, text string) error {
	msg := bot.NewMsg("", "", nil, user, text, time.Now().Unix())

	return bot.SendMsg(msg)
}

// NewUserFromProto - new user from proto
func (base *BasicChatBot) NewUserFromProto(user *chatbotdbpb.User) User {
	u := base.MgrUser.GetUser(user.UserID)
	if u != nil {
		return u
	}

	return NewBasicUser(user.UserName, user.UserID, user.NickName, user.LastMsgID)
}

// GetUser -
func (base *BasicChatBot) GetUser(userid string) (User, error) {
	u, err := base.DB.GetUser(userid)
	if err != nil {
		return nil, err
	}

	return base.NewUserFromProto(u), nil
}
