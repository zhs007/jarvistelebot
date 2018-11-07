package chatbot

import "context"

// Plugin - chat bot plugin interface
type Plugin interface {
	// OnMessage - get message
	OnMessage(ctx context.Context, params *MessageParams) (bool, error)
	// GetComeInCode - if return is empty string, it means not comein
	GetComeInCode() string
	// IsMyMessage
	IsMyMessage(params *MessageParams) bool
}
