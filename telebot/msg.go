package telebot

import "github.com/zhs007/jarvistelebot/chatbot"

// teleMsg - telegram msg
type teleMsg struct {
	from chatbot.User
	to   chatbot.User
	text string
}

func newMsg(from chatbot.User, text string) *teleMsg {
	return &teleMsg{
		from: from,
		text: text,
	}
}

// GetFrom - get message sender
func (msg *teleMsg) GetFrom() chatbot.User {
	return msg.from
}

// GetTo - get user recive this msg
func (msg *teleMsg) GetTo() chatbot.User {
	return msg.to
}

// GetText - get message text
func (msg *teleMsg) GetText() string {
	return msg.text
}
