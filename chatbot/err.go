package chatbot

import "errors"

var (
	// ErrRepeatPlugins - The plugins repeat
	ErrRepeatPlugins = errors.New("The plugins repeat")
	// ErrPluginsEmpty - No plugins
	ErrPluginsEmpty = errors.New("No plugins")
	// ErrRepeatUserID - The userid repeat
	ErrRepeatUserID = errors.New("The userid repeat")
	// ErrMsgNoFrom - Message no from
	ErrMsgNoFrom = errors.New("Message no from")
	// ErrConfigFile - Can't load config file
	ErrConfigFile = errors.New("Can't load config file")
	// ErrInvalidConfigFile - Invalid config file
	ErrInvalidConfigFile = errors.New("Invalid config file")
	// ErrInvalidConfigCfgPath - Invalid config cfgpath
	ErrInvalidConfigCfgPath = errors.New("Invalid config cfgpath")
	// ErrInvalidConfigDownloadPath - Invalid config downloadpath
	ErrInvalidConfigDownloadPath = errors.New("Invalid config downloadpath")
	// ErrInvalidConfigAnkaDB - Invalid config ankaDB
	ErrInvalidConfigAnkaDB = errors.New("Invalid config ankaDB")
	// ErrSameOption - Same option
	ErrSameOption = errors.New("Same option")
	// ErrEmptyOption - Empty option
	ErrEmptyOption = errors.New("Empty option")
	// ErrAlreadySelected - Already selected
	ErrAlreadySelected = errors.New("Already selected")
	// ErrInvalidOption - Invalid option
	ErrInvalidOption = errors.New("Invalid option")
	// ErrInvalidMessageTo - Invalid Message.To
	ErrInvalidMessageTo = errors.New("Invalid Message.To")
	// ErrSameMsgCallback - Same msgcallback
	ErrSameMsgCallback = errors.New("Same msgcallback")
	// ErrNoMsgCallback - Can't find msgcallback
	ErrNoMsgCallback = errors.New("Can't find msgcallback")
	// ErrInvalidPluginType - Invalid plugin type
	ErrInvalidPluginType = errors.New("Invalid plugin type")
	// ErrSameJarvisMsgCallback - Same jarvismsgcallback
	ErrSameJarvisMsgCallback = errors.New("Same jarvismsgcallback")
	// ErrNoJarvisMsgCallback - Can't find jarvismsgcallback
	ErrNoJarvisMsgCallback = errors.New("Can't find jarvismsgcallback")
	// ErrEmptyMsgText - Empty message text
	ErrEmptyMsgText = errors.New("Empty message text")
	// ErrInvalidEventID - Invalid eventid
	ErrInvalidEventID = errors.New("Invalid eventid")
	// ErrSamePluginName - Same plugin name
	ErrSamePluginName = errors.New("Same plugin name")
	// ErrNoPluginName - No plugin name
	ErrNoPluginName = errors.New("No plugin name")
	// ErrNoCommand - No command
	ErrNoCommand = errors.New("No command")
	// ErrInvalidCommandLineItemNums - Invalid command line item nums
	ErrInvalidCommandLineItemNums = errors.New("Invalid command line item nums")
	// ErrMsgNotMine - This message should not be processed by me
	ErrMsgNotMine = errors.New("This message should not be processed by me")
	// ErrInvalidCommandLine - Invalid command line
	ErrInvalidCommandLine = errors.New("Invalid command line")
	// ErrNoJarvisNode - Can't find jarvisnode
	ErrNoJarvisNode = errors.New("Can't find jarvisnode")
	// ErrInvalidParamsNoCurPlugin - Invalid params: no current plugin
	ErrInvalidParamsNoCurPlugin = errors.New("Invalid params: no current plugin")
	// ErrInvalidParamsInvalidCurPlugin - Invalid params: invalid current plugin
	ErrInvalidParamsInvalidCurPlugin = errors.New("Invalid params: invalid current plugin")
	// ErrEmptyArray - Empty array
	ErrEmptyArray = errors.New("Empty array")
)
