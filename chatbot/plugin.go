package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
)

// Plugin - chat bot plugin interface
type Plugin interface {
	// OnMessage - get message
	OnMessage(ctx context.Context, params *MessageParams) (bool, error)

	// ParseMessage - If this message is what I can process,
	//	it will return to the command line, otherwise it will return an error.
	ParseMessage(params *MessageParams) (proto.Message, error)

	// OnStart - on start
	OnStart(ctx context.Context) error

	// GetPluginType - get pluginType
	GetPluginType() int

	// GetPluginName - get plugin name
	GetPluginName() string
}
