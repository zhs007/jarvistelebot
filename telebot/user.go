package telebot

// teleUser - telegram user
type teleUser struct {
	username string
	nickname string
	chatid   int64
}

func newUser(username string, chatid int64, nickname string) *teleUser {
	return &teleUser{
		username: username,
		nickname: nickname,
		chatid:   chatid,
	}
}

// GetNickName - get nickname, composed of first name and last name
func (user *teleUser) GetNickName() string {
	return user.username
}

// GetUserID - get uid, uid is unique
func (user *teleUser) GetUserID() string {
	return user.username
}

// IsMaster - is master
func (user *teleUser) IsMaster() bool {
	return cfg.TeleBotMaster == user.username
}