package plugincrawler

import "errors"

var (
	// ErrConfigNoCrawlerServAddr - no crawlerservaddr in crawler.yaml
	ErrConfigNoCrawlerServAddr = errors.New("no crawlerservaddr in crawler.yaml")
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
	// ErrInvalidURLParserType - invalid URLParser type
	ErrInvalidURLParserType = errors.New("invalid URLParser type")
	// ErrDuplicateURLParserType - duplicate URLParser type
	ErrDuplicateURLParserType = errors.New("duplicate URLParser type")
)
