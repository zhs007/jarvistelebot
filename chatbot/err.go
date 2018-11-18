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
	// ErrConfigFile - can't load config file
	ErrConfigFile = errors.New("can't load config file")
	// ErrInvalidConfigFile - invalid config file
	ErrInvalidConfigFile = errors.New("invalid config file")
	// ErrInvalidConfigCfgPath - invalid config cfgpath
	ErrInvalidConfigCfgPath = errors.New("invalid config cfgpath")
	// ErrInvalidConfigDownloadPath - invalid config downloadpath
	ErrInvalidConfigDownloadPath = errors.New("invalid config downloadpath")
	// ErrInvalidConfigAnkaDB - invalid config ankaDB
	ErrInvalidConfigAnkaDB = errors.New("invalid config ankaDB")
	// ErrSameOption - same option
	ErrSameOption = errors.New("same option")
	// ErrEmptyOption - empty option
	ErrEmptyOption = errors.New("empty option")
	// ErrAlreadySelected - already selected
	ErrAlreadySelected = errors.New("already selected")
	// ErrInvalidOption - invalid option
	ErrInvalidOption = errors.New("invalid option")
	// ErrInvalidMessageTo - invalid Message.To
	ErrInvalidMessageTo = errors.New("invalid Message.To")
)
