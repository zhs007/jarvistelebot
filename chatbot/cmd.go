package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
)

// Command - command
type Command interface {
	// // Parse - parse command line
	// Parse(lstcmdfield []string) (proto.Message, error)
	// RunCommand - run command
	RunCommand(ctx context.Context, params *MessageParams) bool
	// ParseCommandLine - parse command line
	ParseCommandLine(params *MessageParams) (proto.Message, error)
}
