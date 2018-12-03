package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
)

// MessageParams - Message func params
type MessageParams struct {
	ChatBot     ChatBot
	MgrPlugins  PluginsMgr
	Msg         Message
	LstStr      []string
	CommandLine proto.Message
	CurPlugin   Plugin
}

// // FuncCommand - func ([]string)
// type FuncCommand func(context.Context, *MessageParams) bool

// CommandMap - command list
type CommandMap struct {
	mapCmd map[string]Command
}

// NewCommandMap - new CommandMap
func NewCommandMap() *CommandMap {
	return &CommandMap{
		mapCmd: make(map[string]Command),
	}
}

// AddCommand - add command with cmd
func (m *CommandMap) AddCommand(cmd string, c Command) {
	m.mapCmd[cmd] = c
}

// ParseCommandLine - run func with cmd
func (m *CommandMap) ParseCommandLine(cmd string, params *MessageParams) (proto.Message, error) {
	c, ok := m.mapCmd[cmd]
	if ok {
		return c.ParseCommandLine(params)
	}

	return nil, ErrNoCommand
}

// Run - run func with cmd
func (m *CommandMap) Run(ctx context.Context, cmd string, params *MessageParams) bool {
	c, ok := m.mapCmd[cmd]
	if ok {
		return c.RunCommand(ctx, params)
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
