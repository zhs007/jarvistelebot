package telebot

import (
	"github.com/zhs007/jarvistelebot/chatbot"
)

// teleUserMgr - tele user manager
type teleUserMgr struct {
	*chatbot.BasicUserMgr

	masterUserName string
}

// newTeleUserMgr - new default user mgr
func newTeleUserMgr(master string) chatbot.UserMgr {
	return &teleUserMgr{
		BasicUserMgr:   chatbot.NewBasicUserMgr(),
		masterUserName: master,
	}
}

// IsMaster - is master
func (mgr *teleUserMgr) IsMaster(user chatbot.User) bool {
	return user.GetUserName() == mgr.masterUserName
}
