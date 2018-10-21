package chatbot

// UserMgr - other user info
type UserMgr interface {
	// GetUser - get user with userid
	GetUser(uid string) User
	// AddUser - add user
	AddUser(user User) error
}

// NewUserMgr - new default user mgr
func NewUserMgr() UserMgr {
	return &userMgr{
		mapUser: make(map[string]User),
	}
}

// userMgr - default UserMgr
type userMgr struct {
	mapUser map[string]User
}

// GetUser - get user with userid
func (mgr *userMgr) GetUser(uid string) User {
	if user, ok := mgr.mapUser[uid]; ok {
		return user
	}

	return nil
}

// AddUser - add user
func (mgr *userMgr) AddUser(user User) error {
	if _, ok := mgr.mapUser[user.GetUserID()]; ok {
		return ErrRepeatUserID
	}

	mgr.mapUser[user.GetUserID()] = user

	return nil
}
