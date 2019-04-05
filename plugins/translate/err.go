package plugintranslate

import "errors"

var (
	// ErrConfigNoTranslateServAddr - no translateservaddr in translate.yaml
	ErrConfigNoTranslateServAddr = errors.New("no crawlernodeaddr in translate.yaml")
	// ErrNoConfig - no config
	ErrNoConfig = errors.New("no config")
)
