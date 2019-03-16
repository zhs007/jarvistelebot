package chatbot

import (
	"context"

	"github.com/zhs007/jarviscore/coredb"

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
	SendMsg(msg Message) (Message, error)
	// SaveMsg -
	SaveMsg(msg Message) error
	// NewMsg -
	NewMsg(chatid string, msgid string, from User, to User, text string, curtime int64) Message
	// NewMsgFromProto -
	NewMsgFromProto(msg *chatbotdbpb.Message) Message
	// GetMsg -
	GetMsg(chatid string) (Message, error)
	// GetPluginsMgr - get plugins manager
	GetPluginsMgr() PluginsMgr
	// GetUserScriptsMgr - get user scripts manager
	GetUserScriptsMgr() *UserScriptsMgr
	// GetFileTemplatesMgr - get user file template manager
	GetFileTemplatesMgr() *FileTemplatesMgr

	// GetJarvisNodeCoreDB - get jarvis node coredb
	GetJarvisNodeCoreDB() *coredb.CoreDB
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
	// GetUser - get user with userid
	GetUser(userid string) (User, error)
	// GetUserWithUserName - get user with user name
	GetUserWithUserName(username string) (User, error)
	// GetMaster - get master
	GetMaster() User
	// SetMaster - set master, you can only set userid or username
	SetMaster(userid string, username string)

	// OnJarvisCtrlResult - event handle
	OnJarvisCtrlResult(ctx context.Context, msg *jarviscorepb.JarvisMsg) error

	// AddMsgCallback - add msgCallback
	AddMsgCallback(msg Message, callback FuncMsgCallback) error
	// ProcMsgCallback - proc msgCallback
	ProcMsgCallback(ctx context.Context, msg Message, id int) error
	// DelMsgCallback - del msgCallback
	DelMsgCallback(chatid string) error

	// AddJarvisMsgCallback - add jarvisMsgCallback
	AddJarvisMsgCallback(destAddr string, ctrlid int64, callback FuncJarvisMsgCallback) error
	// ProcJarvisMsgCallback - proc jarvisMsgCallback
	ProcJarvisMsgCallback(ctx context.Context, msg *jarviscorepb.JarvisMsg) error
	// DelJarvisMsgCallback - del jarvisMsgCallback
	DelJarvisMsgCallback(destAddr string, ctrlid int64) error

	// GetVersion - get version
	GetVersion() string

	// // NewEventMgr - new EventMgr
	// //	Because chatbot is an interface that is not fully implemented,
	// //	it can only be implemented like this.
	// NewEventMgr(chatbot ChatBot)
	// RegEventFunc - reg event
	RegEventFunc(eventid string, eventfunc FuncEvent) error
	// OnEvent - on event
	OnEvent(ctx context.Context, chatbot ChatBot, eventid string) error
	// RegUserEventFunc - reg event
	RegUserEventFunc(eventid string, eventfunc FuncUserEvent) error
	// OnUserEvent - on event
	OnUserEvent(ctx context.Context, chatbot ChatBot, eventid string, userID string) error
}
