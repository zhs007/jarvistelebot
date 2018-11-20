package chatbot

import (
	"context"
	"fmt"

	"github.com/zhs007/jarviscore/proto"
)

// FuncJarvisMsgCallback - func(ctx, msg) error
type FuncJarvisMsgCallback func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error

type jarvisMsgCallback struct {
	destAddr string
	ctrlid   int64
	callback FuncJarvisMsgCallback
}

type jarvisMsgCallbackMgr struct {
	mapJarvisMsgCallback map[string]*jarvisMsgCallback
}

// newJarvisMsgCallbackMgr - new jarvisMsgCallbackMgr
func newJarvisMsgCallbackMgr() *jarvisMsgCallbackMgr {
	return &jarvisMsgCallbackMgr{
		mapJarvisMsgCallback: make(map[string]*jarvisMsgCallback),
	}
}

// MakeJarvisMsgCallbackID - make callbackid with jarvisMsgCallback
func MakeJarvisMsgCallbackID(destAddr string, ctrlid int64) string {
	return fmt.Sprintf("%v", destAddr)
	// return fmt.Sprintf("%v:%v", destAddr, ctrlid)
}

// addCallback - add jarvisMsgCallback
func (mgr *jarvisMsgCallbackMgr) addCallback(destAddr string, ctrlid int64, callback FuncJarvisMsgCallback) error {
	callbackid := MakeJarvisMsgCallbackID(destAddr, ctrlid)

	_, ok := mgr.mapJarvisMsgCallback[callbackid]
	if ok {
		return ErrSameJarvisMsgCallback
	}

	mgr.mapJarvisMsgCallback[callbackid] = &jarvisMsgCallback{
		destAddr: destAddr,
		ctrlid:   ctrlid,
		callback: callback,
	}

	return nil
}

// procCallback - proc msgCallback
func (mgr *jarvisMsgCallbackMgr) procCallback(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
	callbackid := MakeJarvisMsgCallbackID(msg.SrcAddr, 0)

	cb, ok := mgr.mapJarvisMsgCallback[callbackid]
	if !ok {
		return ErrNoJarvisMsgCallback
	}

	return cb.callback(ctx, msg)
}

// delCallback - proc msgCallback
func (mgr *jarvisMsgCallbackMgr) delCallback(destAddr string, ctrlid int64) error {
	callbackid := MakeJarvisMsgCallbackID(destAddr, ctrlid)

	_, ok := mgr.mapJarvisMsgCallback[callbackid]
	if !ok {
		return ErrNoJarvisMsgCallback
	}

	delete(mgr.mapJarvisMsgCallback, callbackid)

	return nil
}
