package chatbot

import (
	"context"
)

// FuncMsgCallback - func(ctx, msg, id) error
type FuncMsgCallback func(ctx context.Context, msg Message, id int) error

type msgCallback struct {
	msg      Message
	callback FuncMsgCallback
}

type msgCallbackMgr struct {
	mapCallback map[string]*msgCallback
}

// newMsgCallbackMgr - new msgCallbackMgr
func newMsgCallbackMgr() *msgCallbackMgr {
	return &msgCallbackMgr{
		mapCallback: make(map[string]*msgCallback),
	}
}

// addMsgCallback - add msgCallback
func (mgr *msgCallbackMgr) addMsgCallback(msg Message, callback FuncMsgCallback) error {
	to := msg.GetTo()
	if to == nil {
		return ErrInvalidMessageTo
	}

	chatid := MakeChatID(to.GetUserID(), msg.GetMsgID())

	_, ok := mgr.mapCallback[chatid]
	if ok {
		return ErrSameMsgCallback
	}

	mgr.mapCallback[chatid] = &msgCallback{
		msg:      msg,
		callback: callback,
	}

	return nil
}

// procMsgCallback - proc msgCallback
func (mgr *msgCallbackMgr) procMsgCallback(ctx context.Context, msg Message, id int) error {
	cb, ok := mgr.mapCallback[msg.GetChatID()]
	if !ok {
		return ErrNoMsgCallback
	}

	return cb.callback(ctx, msg, id)
}

// procMsgCallback - proc msgCallback
func (mgr *msgCallbackMgr) delMsgCallback(chatid string) error {
	_, ok := mgr.mapCallback[chatid]
	if !ok {
		return ErrNoMsgCallback
	}

	delete(mgr.mapCallback, chatid)

	return nil
}
