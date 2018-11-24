package chatbot

import (
	"context"
)

var (
	// EventOnStarted - started
	EventOnStarted = "started"
)

// FuncEvent - func event
type FuncEvent func(ctx context.Context, eventid string, chatbot ChatBot) error

// eventMgr event manager
type eventMgr struct {
	mapEvent map[string]([]FuncEvent)
	chatbot  ChatBot
}

func newEventMgr(chatbot ChatBot) *eventMgr {
	mgr := &eventMgr{
		mapEvent: make(map[string]([]FuncEvent)),
		chatbot:  chatbot,
	}

	return mgr
}

func (mgr *eventMgr) checkEventID(eventid string) bool {
	return eventid == EventOnStarted
}

func (mgr *eventMgr) regEventFunc(eventid string, eventfunc FuncEvent) error {
	if !mgr.checkEventID(eventid) {
		return ErrInvalidEventID
	}

	mgr.mapEvent[eventid] = append(mgr.mapEvent[eventid], eventfunc)

	return nil
}

func (mgr *eventMgr) onEvent(ctx context.Context, eventid string) error {
	lst, ok := mgr.mapEvent[eventid]
	if !ok {
		return nil
	}

	for i := range lst {
		lst[i](ctx, eventid, mgr.chatbot)
	}

	return nil
}
