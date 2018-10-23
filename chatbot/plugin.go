package chatbot

// Plugin - chat bot plugin interface
type Plugin interface {
	// OnMessage - get message
	OnMessage(chatbot ChatBot, mgr PluginsMgr, msg Message) (bool, error)
	// GetComeInCode - if return is empty string, it means not comein
	GetComeInCode() string
}
