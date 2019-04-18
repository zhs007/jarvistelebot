package plugindtdata

import "errors"

var (
	// ErrConfigNoDTDataServAddr - no dtdataservaddr in crawler.yaml
	ErrConfigNoDTDataServAddr = errors.New("no crawlerservaddr in crawler.yaml")
	// ErrNoConfig - no dtdata.yaml
	ErrNoConfig = errors.New("no dtdata.yaml")
)
