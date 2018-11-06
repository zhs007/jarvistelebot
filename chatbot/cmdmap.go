package chatbot

// MessageParams - Message func params
type MessageParams struct {
	ChatBot    ChatBot
	MgrPlugins PluginsMgr
	Msg        Message
	LstStr     []string
}

// FuncCommand - func ([]string)
type FuncCommand func(*MessageParams) bool

// CommandMap - command list
type CommandMap struct {
	mapCmd map[string]FuncCommand
}

// NewCommandMap - new CommandMap
func NewCommandMap() *CommandMap {
	return &CommandMap{
		mapCmd: make(map[string]FuncCommand),
	}
}

// RegFunc - reg func with cmd
func (m *CommandMap) RegFunc(cmd string, f FuncCommand) {
	m.mapCmd[cmd] = f
}

// Run - reg func with cmd
func (m *CommandMap) Run(cmd string, params *MessageParams) bool {
	f, ok := m.mapCmd[cmd]
	if ok {
		return f(params)
	}

	return false
}
