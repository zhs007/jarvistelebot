package chatbot

import (
	"sync"

	"github.com/golang/protobuf/proto"

	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

// User - other user info
type User interface {
	// GetNickName - get nickname, composed of first name and last name
	GetNickName() string
	// GetUserName - get username
	GetUserName() string
	// GetUserID - get uid, uid is unique
	GetUserID() string
	// // IsMaster - is master
	// IsMaster() bool
	// ToProto - to proto user
	ToProto() *chatbotdbpb.User
	// GetLastMsgID - get lastMsgID
	GetLastMsgID() int64
	// UpdLastMsgID - update lastMsgID
	UpdLastMsgID(lastmsgid int64)

	// SetUserName - set username
	SetUserName(uname string)

	// StorePluginData -
	StorePluginData(pluginName string, data proto.Message)
	// GetPluginData -
	GetPluginData(pluginName string) (proto.Message, bool)
}

// BasicUser - basic User
type BasicUser struct {
	User          *chatbotdbpb.User
	mapPluginData sync.Map
}

// NewBasicUser - new BasicUser
func NewBasicUser(username string, userid string, nickname string, lastMsgID int64) *BasicUser {
	return &BasicUser{
		User: &chatbotdbpb.User{
			NickName:  nickname,
			UserID:    userid,
			UserName:  username,
			LastMsgID: lastMsgID,
		},
	}
}

// ToProto - to proto user
func (bu *BasicUser) ToProto() *chatbotdbpb.User {
	return bu.User
}

// GetNickName - get nickname, composed of first name and last name
func (bu *BasicUser) GetNickName() string {
	return bu.User.NickName
}

// GetUserID - get uid, uid is unique
func (bu *BasicUser) GetUserID() string {
	return bu.User.UserID
}

// GetLastMsgID - get lastMsgID
func (bu *BasicUser) GetLastMsgID() int64 {
	return bu.User.LastMsgID
}

// UpdLastMsgID - update lastMsgID
func (bu *BasicUser) UpdLastMsgID(lastmsgid int64) {
	bu.User.LastMsgID = lastmsgid
}

// GetUserName - get username
func (bu *BasicUser) GetUserName() string {
	return bu.User.UserName
}

// StorePluginData -
func (bu *BasicUser) StorePluginData(pluginName string, data proto.Message) {
	bu.mapPluginData.Store(pluginName, data)
}

// GetPluginData -
func (bu *BasicUser) GetPluginData(pluginName string) (proto.Message, bool) {
	v, ok := bu.mapPluginData.Load(pluginName)
	if ok {
		return v.(proto.Message), true
	}

	return nil, false
}

// SetUserName - set username
func (bu *BasicUser) SetUserName(uname string) {
	bu.User.UserName = uname
}
