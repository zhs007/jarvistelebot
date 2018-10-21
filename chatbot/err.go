package chatbot

import "errors"

var (
	// ErrRepeatPlugins - the plugins repeat
	ErrRepeatPlugins = errors.New("the plugins repeat")
	// ErrPluginsEmpty - no plugins
	ErrPluginsEmpty = errors.New("no plugins")
	// ErrRepeatUserID - the userid repeat
	ErrRepeatUserID = errors.New("the userid repeat")
	// ErrMsgNoFrom - msg no from
	ErrMsgNoFrom = errors.New("msg no from")
)
