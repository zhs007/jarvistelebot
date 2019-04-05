package plugintranslate

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// config - config
type config struct {
	TranslateServAddr string
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

	if cfg.TranslateServAddr == "" {
		return ErrConfigNoTranslateServAddr
	}

	return nil
}
