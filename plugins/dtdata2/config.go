package plugindtdata2

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// config - config
type config struct {
	DTDataServAddr string
	URL            string
}

// LoadConfig - load config
func loadConfig(filename string) *config {
	fi, err := os.Open(filename)
	if err != nil {
		return nil
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil
	}

	cfg := &config{}

	err = yaml.Unmarshal(fd, cfg)
	if err != nil {
		return nil
	}

	return cfg
}

// checkConfig -
func checkConfig(cfg *config) error {
	if cfg == nil {
		return ErrNoConfig
	}

	if cfg.DTDataServAddr == "" {
		return ErrConfigNoDTDataServAddr
	}

	if cfg.URL == "" {
		return ErrConfigNoURL
	}

	return nil
}
