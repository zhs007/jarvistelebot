package chatbot

// ChatBot - chat bot interface
type ChatBot interface {
	// Start
	Start() error
	// SendMsg
	SendMsg(user User, text string) error
}
