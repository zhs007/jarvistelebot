package chatbot

// Plugins - chat bot plugins interface
type Plugins interface {
	// OnMessage - get message
	OnMessage(chatbot ChatBot, msg Message) (bool, error)
}
