package telebot

import (
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

// teleMsg - telegram msg
type teleMsg struct {
	chatbot.BasicMessage

	chatID    string
	msgID     string
	from      chatbot.User
	to        chatbot.User
	text      string
	timeStamp int64
	groupID   string
}

// func newMsg(msgID string, from chatbot.User, text string, date int) *teleMsg {
// 	return &teleMsg{
// 		chatID:    from.GetUserID() + ":" + msgID,
// 		msgID:     msgID,
// 		from:      from,
// 		text:      text,
// 		timeStamp: int64(date),
// 	}
// }

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
	pbmsg := msg.BasicMessage.ToProto()

	pbmsg.ChatID = msg.chatID
	pbmsg.MsgID = msg.msgID
	pbmsg.Text = msg.text
	pbmsg.TimeStamp = msg.timeStamp

	if msg.from != nil {
		pbmsg.From = msg.from.ToProto()
	}

	if msg.to != nil {
		pbmsg.To = msg.to.ToProto()
	}

	return pbmsg
}

// GetTimeStamp - get timestamp
func (msg *teleMsg) GetTimeStamp() int64 {
	return msg.timeStamp
}

// GetChatID - get chatID
func (msg *teleMsg) GetChatID() string {
	return msg.chatID
}

// GetMsgID - get message id
func (msg *teleMsg) GetMsgID() string {
	return msg.msgID
}

// SetMsgID - set message id
func (msg *teleMsg) SetMsgID(msgid string) {
	if msg.from == nil {
		msg.chatID = ":" + msgid
		msg.msgID = msgid

		return
	}
	msg.chatID = msg.from.GetUserID() + ":" + msgid
	msg.msgID = msgid
}

// SetChatID - set chat id
func (msg *teleMsg) SetChatID(chatid string) {
	msg.chatID = chatid
}

// SetText - set text
func (msg *teleMsg) SetText(text string) {
	msg.text = text
}

// SetGroupID - set groupID
func (msg *teleMsg) SetGroupID(groupID string) {
	msg.groupID = groupID
}

// GetGroupID - get groupID
func (msg *teleMsg) GetGroupID() string {
	return msg.groupID
}

// InGroup - this message is from a group
func (msg *teleMsg) InGroup() bool {
	return msg.groupID != ""
}
