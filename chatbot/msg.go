package chatbot

// Message - other user info
type Message interface {
	// GetFrom - get message sender
	GetFrom() User
	// GetTo - get user recive this msg
	GetTo() User
	// GetText - get message text
	GetText() string
}
