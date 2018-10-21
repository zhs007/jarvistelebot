package telebot

import "errors"

var (
	// ErrConfigFile - can't load config file
	ErrConfigFile = errors.New("can't load config file")
	// ErrInvalidConfigFile - invalid config file
	ErrInvalidConfigFile = errors.New("invalid config file")
	// ErrCfgTeleBotToken - invalid tele bot token
	ErrCfgTeleBotToken = errors.New("invalid tele bot token")
	// ErrCfgTeleBotType - invalid tele bot type
	ErrCfgTeleBotType = errors.New("invalid tele bot type")
	// ErrCfgWebHook - invalid webhook config
	ErrCfgWebHook = errors.New("invalid webhook config")
	// ErrCfgLogLevel - invalid log level
	ErrCfgLogLevel = errors.New("invalid log level")
	// ErrNewTeleBot - NewBotAPI err
	ErrNewTeleBot = errors.New("NewBotAPI err")
	// ErrInvalidUser - invalid teleUser
	ErrInvalidUser = errors.New("invalid teleUser")
)
