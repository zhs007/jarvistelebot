package plugindtdata2

import "errors"

var (
	// ErrConfigNoDTDataServAddr - no dtdataservaddr in dtdata2.yaml
	ErrConfigNoDTDataServAddr = errors.New("no dtdataservaddr in dtdata2.yaml")
	// ErrNoConfig - no dtdata.yaml
	ErrNoConfig = errors.New("no dtdata2.yaml")
)
