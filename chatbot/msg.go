package chatbot

import (
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

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
	// ToProto - to proto message
	ToProto() *chatbotdbpb.Message
}
