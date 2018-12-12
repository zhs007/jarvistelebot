package chatbot

import (
	"context"
	"fmt"
	"sync"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/zhs007/jarvistelebot/chatbotdb"
)

// FileTemplates - file template
type FileTemplates struct {
	Templates []string
}

// FileTemplatesMgr - user's file template manager
type FileTemplatesMgr struct {
	mapUser sync.Map
}

// Init - init
func (mgr *FileTemplatesMgr) init(chatbot ChatBot) {
	chatbot.RegUserEventFunc(UserEventOnChgUserFileTemplate,
		func(ctx context.Context, chatbot ChatBot, eventid string, userID string) error {
			jarvisbase.Debug("UserEventOnChgUserFileTemplate")

			mgr.Load(chatbot.GetChatBotDB(), userID)

			return nil
		})
}

// Load - load user's scripts
func (mgr *FileTemplatesMgr) Load(db *chatbotdb.ChatBotDB, userID string) (*FileTemplates, error) {
	lst, err := db.GetFileTemplates(userID, "")
	if err != nil {
		jarvisbase.Warn("FileTemplatesMgr.Load:GetFileTemplates", zap.Error(err))

		return nil, err
	}

	ft := &FileTemplates{}

	for _, v := range lst.Templates {
		ft.Templates = append(ft.Templates, v.FileTemplateName)
	}

	mgr.mapUser.Store(userID, ft)

	jarvisbase.Debug("FileTemplatesMgr.Load", zap.String("ft", fmt.Sprintf("%v", ft)))

	return ft, nil
}

// Get - load user's scripts
func (mgr *FileTemplatesMgr) Get(db *chatbotdb.ChatBotDB, userID string) (*FileTemplates, error) {
	ret, ok := mgr.mapUser.Load(userID)
	if !ok {
		return mgr.Load(db, userID)
	}

	ft, ok := ret.(*FileTemplates)
	if !ok {
		return mgr.Load(db, userID)
	}

	return ft, nil
}
