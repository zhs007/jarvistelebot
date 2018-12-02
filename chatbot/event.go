package chatbot

import (
	"context"
)

var (
	// EventOnStarted - started
	EventOnStarted = "started"

	// UserEventOnChgUserScript - chguserscript
	UserEventOnChgUserScript = "chguserscript"
	// UserEventOnChgUserFileTemplate - chguserfiletemplate
	UserEventOnChgUserFileTemplate = "chguserfiletemplate"
)

// FuncEvent - func event
type FuncEvent func(ctx context.Context, chatbot ChatBot, eventid string) error

// FuncUserEvent - user func event
type FuncUserEvent func(ctx context.Context, chatbot ChatBot, eventid string, userID string) error

// eventMgr event manager
type eventMgr struct {
	mapEvent     map[string]([]FuncEvent)
	mapUserEvent map[string]([]FuncUserEvent)
	// chatbot      ChatBot
}

func newEventMgr() *eventMgr {
	mgr := &eventMgr{
		mapEvent:     make(map[string]([]FuncEvent)),
		mapUserEvent: make(map[string]([]FuncUserEvent)),
		// chatbot:      chatbot,
	}

	return mgr
}

func (mgr *eventMgr) checkEventID(eventid string) bool {
	return eventid == EventOnStarted
}

func (mgr *eventMgr) checkUserEventID(eventid string) bool {
	return eventid == UserEventOnChgUserScript
}

func (mgr *eventMgr) regEventFunc(eventid string, eventfunc FuncEvent) error {
	if !mgr.checkEventID(eventid) {
		return ErrInvalidEventID
	}

	mgr.mapEvent[eventid] = append(mgr.mapEvent[eventid], eventfunc)

	return nil
}

func (mgr *eventMgr) regUserEventFunc(eventid string, eventfunc FuncUserEvent) error {
	if !mgr.checkUserEventID(eventid) {
		return ErrInvalidEventID
	}

	mgr.mapUserEvent[eventid] = append(mgr.mapUserEvent[eventid], eventfunc)

	return nil
}

func (mgr *eventMgr) onEvent(ctx context.Context, chatbot ChatBot, eventid string) error {
	lst, ok := mgr.mapEvent[eventid]
	if !ok {
		return nil
	}

	for i := range lst {
		lst[i](ctx, chatbot, eventid)
	}

	return nil
}

func (mgr *eventMgr) onUserEvent(ctx context.Context, chatbot ChatBot, eventid string, userID string) error {
	lst, ok := mgr.mapUserEvent[eventid]
	if !ok {
		return nil
	}

	for i := range lst {
		lst[i](ctx, chatbot, eventid, userID)
	}

	return nil
}
