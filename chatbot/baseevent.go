package chatbot

import (
	"context"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"
)

// onEventStarted - on started
func onEventStarted(ctx context.Context, chatbot ChatBot, eventid string) error {

	user := chatbot.GetMaster()
	if user != nil {
		err := SendTextMsg(chatbot, user, "Master, I am restarted.")
		if err != nil {
			jarvisbase.Warn("onEventStarted:SendTextMsg", zap.Error(err))
		}
	}

	chatbot.GetUserScriptsMgr().init(chatbot)

	return nil
}
