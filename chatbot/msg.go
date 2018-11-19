package chatbot

import (
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

// MsgOption - option
type MsgOption struct {
	Text string
	ID   int
}

// Message - other user info
type Message interface {
	// GetFrom - get message sender
	GetFrom() User
	// GetTo - get user recive this msg
	GetTo() User
	// GetText - get message text
	GetText() string
	// GetTimeStamp - get timestamp
	GetTimeStamp() int64
	// GetChatID - get chatID
	GetChatID() string
	// GetMsgID - get message id
	GetMsgID() string
	// SetMsgID - set message id
	SetMsgID(msgid string)
	// SetChatID - set chat id
	SetChatID(chatid string)

	// AddOption - add option
	AddOption(text string) (int, error)
	// HasOptions - has any options
	HasOptions() bool
	// SelectOption - select option
	SelectOption(id int) error
	// GetSelected - get selected
	GetSelected() int
	// GetOption - get option
	GetOption(id int) string

	// ToProto - to proto message
	ToProto() *chatbotdbpb.Message
}

// BasicMessage - basic Message
type BasicMessage struct {
	Options    []*MsgOption
	IDSelected int
}

// AddOption - add option
func (msg *BasicMessage) AddOption(text string) (int, error) {
	if text == "" {
		return -1, ErrEmptyOption
	}

	for _, v := range msg.Options {
		if text == v.Text {
			return -1, ErrSameOption
		}
	}

	op := &MsgOption{
		Text: text,
		ID:   len(msg.Options) + 1,
	}

	msg.Options = append(msg.Options, op)

	return op.ID, nil
}

// HasOptions - has any options
func (msg *BasicMessage) HasOptions() bool {
	return len(msg.Options) > 0
}

// SelectOption - select option
func (msg *BasicMessage) SelectOption(id int) error {
	if msg.IDSelected > 0 {
		return ErrAlreadySelected
	}

	if id <= 0 || id > len(msg.Options) {
		return ErrInvalidOption
	}

	msg.IDSelected = id

	return nil
}

// GetSelected - get selected
func (msg *BasicMessage) GetSelected() int {
	return msg.IDSelected
}

// GetOption - get option
func (msg *BasicMessage) GetOption(id int) string {
	if id <= 0 || id > len(msg.Options) {
		return ""
	}

	return msg.Options[id-1].Text
}
