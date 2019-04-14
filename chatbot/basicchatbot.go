package chatbot

import (
	"context"

	"github.com/zhs007/jarviscore/base"
	"github.com/zhs007/jarviscore/coredb"
	"go.uber.org/zap"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/basedef"
	"github.com/zhs007/jarvistelebot/chatbotdb"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

// BasicChatBot - base chatbot
type BasicChatBot struct {
	DB                   *chatbotdb.ChatBotDB
	Node                 jarviscore.JarvisNode
	MgrPlugins           PluginsMgr
	Config               *Config
	MgrUser              UserMgr
	mgrMsgCallback       *msgCallbackMgr
	mgrJsrvisMsgCallback *jarvisMsgCallbackMgr
	mgrEvent             *eventMgr
	mgrUserScripts       *UserScriptsMgr
	mgrFileTemplates     *FileTemplatesMgr
	mgrTimer             *TimerMgr
	// mgrUser              UserMgr
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
	base.mgrMsgCallback = newMsgCallbackMgr()
	base.mgrJsrvisMsgCallback = newJarvisMsgCallbackMgr()
	base.MgrUser = NewBasicUserMgr()
	base.mgrUserScripts = &UserScriptsMgr{}
	base.mgrFileTemplates = &FileTemplatesMgr{}
	base.mgrTimer = newTimerMgr()

	base.mgrEvent = newEventMgr()
	base.RegEventFunc(EventOnStarted, onEventStarted)

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
func (base *BasicChatBot) GetJarvisNodeCoreDB() *coredb.CoreDB {
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

// NewUserFromProto - new user from proto
func (base *BasicChatBot) NewUserFromProto(user *chatbotdbpb.User) User {
	u := base.MgrUser.GetUser(user.UserID)
	if u != nil {
		return u
	}

	return NewBasicUser(user.UserName, user.UserID, user.NickName, user.LastMsgID)
}

// GetUser - get user with userid
func (base *BasicChatBot) GetUser(userid string) (User, error) {
	u, err := base.DB.GetUser(userid)
	if err != nil {
		return nil, err
	}

	return base.NewUserFromProto(u), nil
}

// GetUserWithUserName - get user with user name
func (base *BasicChatBot) GetUserWithUserName(username string) (User, error) {
	u, err := base.DB.GetUserWithUserName(username)
	if err != nil {
		return nil, err
	}

	return base.NewUserFromProto(u), nil
}

// AddMsgCallback - add msgCallback
func (base *BasicChatBot) AddMsgCallback(msg Message, callback FuncMsgCallback) error {
	return base.mgrMsgCallback.addCallback(msg, callback)
}

// ProcMsgCallback - proc msgCallback
func (base *BasicChatBot) ProcMsgCallback(ctx context.Context, msg Message, id int) error {
	return base.mgrMsgCallback.procCallback(ctx, msg, id)
}

// DelMsgCallback - del msgCallback
func (base *BasicChatBot) DelMsgCallback(chatid string) error {
	return base.mgrMsgCallback.delCallback(chatid)
}

// AddJarvisMsgCallback - add jarvisMsgCallback
func (base *BasicChatBot) AddJarvisMsgCallback(destAddr string, ctrlid int64, callback FuncJarvisMsgCallback) error {
	return base.mgrJsrvisMsgCallback.addCallback(destAddr, ctrlid, callback)
}

// ProcJarvisMsgCallback - proc jarvisMsgCallback
func (base *BasicChatBot) ProcJarvisMsgCallback(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
	return base.mgrJsrvisMsgCallback.procCallback(ctx, msg)
}

// DelJarvisMsgCallback - del jarvisMsgCallback
func (base *BasicChatBot) DelJarvisMsgCallback(destAddr string, ctrlid int64) error {
	return base.mgrJsrvisMsgCallback.delCallback(destAddr, ctrlid)
}

// GetVersion - get version
func (base *BasicChatBot) GetVersion() string {
	return basedef.VERSION
}

// // NewEventMgr - new EventMgr
// func (base *BasicChatBot) NewEventMgr(chatbot ChatBot) {
// 	base.mgrEvent = newEventMgr(chatbot)

// 	base.RegEventFunc(EventOnStarted, onEventStarted)
// }

// RegEventFunc - reg event
func (base *BasicChatBot) RegEventFunc(eventid string, eventfunc FuncEvent) error {
	return base.mgrEvent.regEventFunc(eventid, eventfunc)
}

// OnEvent - on event
func (base *BasicChatBot) OnEvent(ctx context.Context, chatbot ChatBot, eventid string) error {
	return base.mgrEvent.onEvent(ctx, chatbot, eventid)
}

// RegUserEventFunc - reg event
func (base *BasicChatBot) RegUserEventFunc(eventid string, eventfunc FuncUserEvent) error {
	return base.mgrEvent.regUserEventFunc(eventid, eventfunc)
}

// OnUserEvent - on event
func (base *BasicChatBot) OnUserEvent(ctx context.Context, chatbot ChatBot, eventid string, userID string) error {
	return base.mgrEvent.onUserEvent(ctx, chatbot, eventid, userID)
}

// GetMaster - get master
func (base *BasicChatBot) GetMaster() User {
	if base.MgrUser.GetMasterUserID() != "" {
		user, err := base.DB.GetUser(base.MgrUser.GetMasterUserID())
		if err != nil {
			jarvisbase.Warn("BasicChatBot:GetMaster:GetUser", zap.Error(err))

			return nil
		}

		return base.NewUserFromProto(user)
	} else if base.MgrUser.GetMasterUserName() != "" {
		user, err := base.DB.GetUserWithUserName(base.MgrUser.GetMasterUserName())
		if err != nil {
			jarvisbase.Warn("BasicChatBot:GetMaster:GetUserWithUserName", zap.Error(err))

			return nil
		}

		return base.NewUserFromProto(user)
	}

	return nil
}

// SetMaster - set master, you can only set userid or username
func (base *BasicChatBot) SetMaster(userid string, username string) {
	base.MgrUser.SetMaster(userid, username)
}

// GetPluginsMgr - get plugins manager
func (base *BasicChatBot) GetPluginsMgr() PluginsMgr {
	return base.MgrPlugins
}

// GetUserScriptsMgr - get user scripts manager
func (base *BasicChatBot) GetUserScriptsMgr() *UserScriptsMgr {
	return base.mgrUserScripts
}

// GetFileTemplatesMgr - get user file template manager
func (base *BasicChatBot) GetFileTemplatesMgr() *FileTemplatesMgr {
	return base.mgrFileTemplates
}

// OnTimer - ontimer
func (base *BasicChatBot) OnTimer(ctx context.Context) {
	base.mgrTimer.OnTimer(ctx)
}

// AddTimer - add timer
func (base *BasicChatBot) AddTimer(timer int, times int, info string, onTimer FuncOnTimer) int {
	return base.mgrTimer.AddTimer(timer, times, info, onTimer)
}

// DeleteTimer - delete timer
func (base *BasicChatBot) DeleteTimer(timerid int) {
	base.mgrTimer.DeleteTimer(timerid)
}
