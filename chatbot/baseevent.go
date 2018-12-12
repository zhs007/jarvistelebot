package chatbot

import (
	"context"

	"github.com/zhs007/jarvistelebot/basedef"
)

// onEventStarted - on started
func onEventStarted(ctx context.Context, chatbot ChatBot, eventid string) error {

	user := chatbot.GetMaster()
	if user != nil {
		SendTextMsg(chatbot, user, "Master, I am restarted.", nil)
		SendTextMsg(chatbot, user, "My version is "+basedef.VERSION, nil)
	}

	chatbot.GetUserScriptsMgr().init(chatbot)
	chatbot.GetFileTemplatesMgr().init(chatbot)

	return nil
}
