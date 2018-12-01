package chatbot

import (
	"context"
	"sync"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/zhs007/jarvistelebot/chatbotdb"
)

// UserScripts - user's scripts
type UserScripts struct {
	Scripts []string
}

// UserScriptsMgr - user's scripts manager
type UserScriptsMgr struct {
	mapUser sync.Map
}

// Init - init
func (mgr *UserScriptsMgr) init(chatbot ChatBot) {
	chatbot.RegUserEventFunc(UserEventOnChgUserScript,
		func(ctx context.Context, chatbot ChatBot, eventid string, userID string) error {
			mgr.Load(chatbot.GetChatBotDB(), userID)

			return nil
		})
}

// Load - load user's scripts
func (mgr *UserScriptsMgr) Load(db *chatbotdb.ChatBotDB, userID string) (*UserScripts, error) {
	lst, err := db.GetUserScripts(userID, "")
	if err != nil {
		jarvisbase.Warn("UserScriptsMgr.Load:GetUserScripts", zap.Error(err))

		return nil, err
	}

	us := &UserScripts{}

	for _, v := range lst.Scripts {
		us.Scripts = append(us.Scripts, v.ScriptName)
	}

	mgr.mapUser.Store(userID, us)

	return us, nil
}

// Get - load user's scripts
func (mgr *UserScriptsMgr) Get(db *chatbotdb.ChatBotDB, userID string) (*UserScripts, error) {
	ret, ok := mgr.mapUser.Load(userID)
	if !ok {
		return mgr.Load(db, userID)
	}

	us, ok := ret.(*UserScripts)
	if !ok {
		return mgr.Load(db, userID)
	}

	return us, nil
}
