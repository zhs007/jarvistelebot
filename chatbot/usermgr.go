package chatbot

// UserMgr - other user info
type UserMgr interface {
	// GetUser - get user with userid
	GetUser(uid string) User
	// AddUser - add user
	AddUser(user User) error
	// IsMaster - is master
	IsMaster(user User) bool
	// SetMaster - set master, you can only set userid or username
	SetMaster(userid string, username string)
	// GetMasterUserID - get master userid
	GetMasterUserID() string
	// GetMasterUserName - get master username
	GetMasterUserName() string
}

// NewBasicUserMgr - new BasicUserMgr
func NewBasicUserMgr() *BasicUserMgr {
	return &BasicUserMgr{
		mapUser: make(map[string]User),
	}
}

// BasicUserMgr - default UserMgr
type BasicUserMgr struct {
	mapUser        map[string]User
	masterUserID   string
	masterUserName string
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

// GetMasterUserID - get master userid
func (mgr *BasicUserMgr) GetMasterUserID() string {
	return mgr.masterUserID
}

// GetMasterUserName - get master username
func (mgr *BasicUserMgr) GetMasterUserName() string {
	return mgr.masterUserName
}

// SetMaster - set master, you can only set userid or username
func (mgr *BasicUserMgr) SetMaster(userid string, username string) {
	mgr.masterUserID = userid
	mgr.masterUserName = username
}

// IsMaster - is master
func (mgr *BasicUserMgr) IsMaster(user User) bool {
	return user.GetUserName() == mgr.masterUserName || user.GetUserID() == mgr.masterUserID
}
