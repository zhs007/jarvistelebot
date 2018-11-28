package chatbot

import "context"

// MessageParams - Message func params
type MessageParams struct {
	ChatBot    ChatBot
	MgrPlugins PluginsMgr
	Msg        Message
	LstStr     []string
}

// FuncCommand - func ([]string)
type FuncCommand func(context.Context, *MessageParams) bool

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

// Run - run func with cmd
func (m *CommandMap) Run(ctx context.Context, cmd string, params *MessageParams) bool {
	f, ok := m.mapCmd[cmd]
	if ok {
		return f(ctx, params)
	}

	return false
}

// HasCommand - has command
func (m *CommandMap) HasCommand(cmd string) bool {
	_, ok := m.mapCmd[cmd]
	if ok {
		return true
	}

	return false
}
