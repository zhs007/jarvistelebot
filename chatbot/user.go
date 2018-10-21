package chatbot

// User - other user info
type User interface {
	// GetNickName - get nickname, composed of first name and last name
	GetNickName() string
	// GetUserID - get uid, uid is unique
	GetUserID() string
	// IsMaster - is master
	IsMaster() bool
}
