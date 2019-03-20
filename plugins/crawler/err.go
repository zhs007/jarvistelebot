package plugincrawler

import "errors"

var (
	// ErrConfigNoCrawlerNodeAddr - no crawlernodeaddr in crawler.yaml
	ErrConfigNoCrawlerNodeAddr = errors.New("no crawlernodeaddr in crawler.yaml")
	// ErrConfigInvalidCrawlerNodeAddr - invalid crawlernodeaddr in crawler.yaml
	ErrConfigInvalidCrawlerNodeAddr = errors.New("invalid crawlernodeaddr in crawler.yaml")
	// ErrConfigNoCrawlerPath - no crawlerpath in crawler.yaml
	ErrConfigNoCrawlerPath = errors.New("no crawlerpath in crawler.yaml")
	// ErrNoConfig - no config
	ErrNoConfig = errors.New("no config")
	// ErrConfigNoUpdateScript - no updatescript in crawler.yaml
	ErrConfigNoUpdateScript = errors.New("no updatescript in crawler.yaml")
	// ErrConfigNoExpArticleScript - no exparticlescript in crawler.yaml
	ErrConfigNoExpArticleScript = errors.New("no exparticlescript in crawler.yaml")
)
