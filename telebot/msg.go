package telebot

import (
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

// teleMsg - telegram msg
type teleMsg struct {
	chatID    string
	from      chatbot.User
	to        chatbot.User
	text      string
	timeStamp int64
}

func newMsg(chatID string, from chatbot.User, text string, date int) *teleMsg {
	return &teleMsg{
		chatID:    chatID,
		from:      from,
		text:      text,
		timeStamp: int64(date),
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

// ToProto - ToProto - to proto message
func (msg *teleMsg) ToProto() *chatbotdbpb.Message {
	return &chatbotdbpb.Message{
		From:      msg.from.ToProto(),
		To:        msg.to.ToProto(),
		Text:      msg.text,
		TimeStamp: msg.timeStamp,
	}
}

// GetTimeStamp - get timestamp
func (msg *teleMsg) GetTimeStamp() int64 {
	return msg.timeStamp
}

// GetChatID - get chatID
func (msg *teleMsg) GetChatID() string {
	return msg.chatID
}
