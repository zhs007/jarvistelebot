package plugindtdata

import "errors"

var (
	// ErrConfigNoDTDataServAddr - no dtdataservaddr in dtdata.yaml
	ErrConfigNoDTDataServAddr = errors.New("no dtdataservaddr in dtdata.yaml")
	// ErrNoConfig - no dtdata.yaml
	ErrNoConfig = errors.New("no dtdata.yaml")
)
