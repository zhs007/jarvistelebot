package pluginduckling

import "errors"

var (
	// ErrConfigNoDucklingServAddr - no ducklingservaddr in duckling.yaml
	ErrConfigNoDucklingServAddr = errors.New("no ducklingservaddr in duckling.yaml")
	// ErrNoDucklingConfig - no duckling.yaml
	ErrNoDucklingConfig = errors.New("no duckling.yaml")
)
