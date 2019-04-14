package plugincrawler

import (
	"io/ioutil"
	"os"

	"github.com/zhs007/jarviscore"
	"gopkg.in/yaml.v2"
)

// config - config
type config struct {
	CrawlerServAddr  string
	CrawlerNodeAddr  string
	CrawlerPath      string
	UpdateScript     string
	ExpArticleScript string
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

	if cfg.CrawlerServAddr == "" {
		return ErrConfigNoCrawlerServAddr
	}

	if cfg.CrawlerNodeAddr == "" {
		return ErrConfigNoCrawlerNodeAddr
	}

	if !jarviscore.IsValidNodeAddr(cfg.CrawlerNodeAddr) {
		return ErrConfigInvalidCrawlerNodeAddr
	}

	if cfg.CrawlerPath == "" {
		return ErrConfigNoCrawlerPath
	}

	if cfg.UpdateScript == "" {
		return ErrConfigNoUpdateScript
	}

	if cfg.ExpArticleScript == "" {
		return ErrConfigNoExpArticleScript
	}

	return nil
}
