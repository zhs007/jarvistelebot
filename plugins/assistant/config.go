package pluginassistant

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// config - config
type config struct {
	AnkaDB struct {
		DBPath   string
		Engine   string
		HTTPAddr string
	}
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
