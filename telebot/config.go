package telebot

import (
	"io/ioutil"
	"os"

	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"
)

const (
	// TBTGetUpdate - getupdate
	TBTGetUpdate = "getupdate"
	// TBTWebHook - webhook
	TBTWebHook = "webhook"
	// LOGPATHConsole - console
	LOGPATHConsole = "console"
	// LLDebug - debug
	LLDebug = "debug"
	// LLInfo - info
	LLInfo = "info"
	// LLWarn - warn
	LLWarn = "warn"
	// LLError - error
	LLError = "error"
)

// Config - config
type Config struct {
	TeleBotToken  string
	TeleBotMaster string
	TeleBotType   string
	WebHookURL    string
	WebKookAddr   string
	LogPath       string
	LogLevel      string
	DebugMode     bool
	CfgPath       string
	Plugins       []string
	DefaultPlugin string

	lvl zapcore.Level
}

// var cfg Config
// var onceCfg sync.Once

// LoadConfig - load config
func LoadConfig(filename string) (*Config, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(fd, cfg)
	if err != nil {
		return nil, err
	}

	err = checkConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// checkConfig - check config file
func checkConfig(cfg *Config) error {
	if cfg.TeleBotToken == "" {
		return ErrCfgTeleBotToken
	}

	if cfg.TeleBotType != TBTGetUpdate && cfg.TeleBotType != TBTWebHook {
		return ErrCfgTeleBotType
	}

	if cfg.TeleBotType == TBTWebHook {
		if cfg.WebHookURL == "" || cfg.WebKookAddr == "" {
			return ErrCfgWebHook
		}
	}

	if cfg.LogLevel == LLDebug {
		cfg.lvl = zapcore.DebugLevel
	} else if cfg.LogLevel == LLInfo {
		cfg.lvl = zapcore.InfoLevel
	} else if cfg.LogLevel == LLWarn {
		cfg.lvl = zapcore.WarnLevel
	} else if cfg.LogLevel == LLError {
		cfg.lvl = zapcore.ErrorLevel
	} else {
		return ErrCfgLogLevel
	}

	return nil
}

// // LoadConfig - load config
// func LoadConfig(filename string) (err error) {
// 	onceCfg.Do(func() {
// 		err = loadConfig(filename)
// 	})

// 	return
// }

// // GetConfig - get Config
// func GetConfig() *Config {
// 	return &cfg
// }
