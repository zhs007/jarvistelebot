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

// addCallback - add msgCallback
func (mgr *msgCallbackMgr) addCallback(msg Message, callback FuncMsgCallback) error {
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

// procCallback - proc msgCallback
func (mgr *msgCallbackMgr) procCallback(ctx context.Context, msg Message, id int) error {
	cb, ok := mgr.mapCallback[msg.GetChatID()]
	if !ok {
		return ErrNoMsgCallback
	}

	return cb.callback(ctx, msg, id)
}

// delCallback - proc msgCallback
func (mgr *msgCallbackMgr) delCallback(chatid string) error {
	_, ok := mgr.mapCallback[chatid]
	if !ok {
		return ErrNoMsgCallback
	}

	delete(mgr.mapCallback, chatid)

	return nil
}
