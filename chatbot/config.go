package chatbot

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config - config
type Config struct {
	CfgPath      string
	DownloadPath string

	AnkaDB struct {
		DBPath   string
		Engine   string
		HTTPAddr string
	}
}

// LoadConfig - load config
func LoadConfig(filename string) (*Config, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, ErrConfigFile
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, ErrConfigFile
	}

	cfg := &Config{}

	err = yaml.Unmarshal(fd, &cfg)
	if err != nil {
		return nil, ErrInvalidConfigFile
	}

	err = checkConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// checkConfig - check config file
func checkConfig(cfg *Config) error {
	if cfg.CfgPath == "" {
		return ErrInvalidConfigCfgPath
	}

	if cfg.DownloadPath == "" {
		return ErrInvalidConfigDownloadPath
	}

	if cfg.AnkaDB.DBPath == "" || cfg.AnkaDB.Engine != "leveldb" {
		return ErrInvalidConfigAnkaDB
	}

	return nil
}
