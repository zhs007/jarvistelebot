package chatbot

// UserMgr - other user info
type UserMgr interface {
	// GetUser - get user with userid
	GetUser(uid string) User
	// AddUser - add user
	AddUser(user User) error
	// IsMaster - is master
	IsMaster(user User) bool
}

// NewBasicUserMgr - new BasicUserMgr
func NewBasicUserMgr() *BasicUserMgr {
	return &BasicUserMgr{
		mapUser: make(map[string]User),
	}
}

// BasicUserMgr - default UserMgr
type BasicUserMgr struct {
	mapUser map[string]User
}

// GetUser - get user with userid
func (mgr *BasicUserMgr) GetUser(uid string) User {
	if user, ok := mgr.mapUser[uid]; ok {
		return user
	}

	return nil
}

// AddUser - add user
func (mgr *BasicUserMgr) AddUser(user User) error {
	if _, ok := mgr.mapUser[user.GetUserID()]; ok {
		return ErrRepeatUserID
	}

	mgr.mapUser[user.GetUserID()] = user

	return nil
}

// // IsMaster - is master
// func (mgr *BasicUserMgr) IsMaster(user User) bool {

// }
