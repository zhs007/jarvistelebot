package plugincore

import (
	"strings"
	"testing"

	"go.uber.org/zap/zapcore"

	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

// testMessage - test Message
type testMessage struct {
	chatbot.BasicMessage

	strText string
}

// GetFrom - get message sender
func (msg *testMessage) GetFrom() chatbot.User {
	return nil
}

// GetTo - get user recive this msg
func (msg *testMessage) GetTo() chatbot.User {
	return nil
}

// GetText - get message text
func (msg *testMessage) GetText() string {
	return msg.strText
}

// GetTimeStamp - get timestamp
func (msg *testMessage) GetTimeStamp() int64 {
	return 0
}

// GetChatID - get chatID
func (msg *testMessage) GetChatID() string {
	return ""
}

// ToProto - to proto message
func (msg *testMessage) ToProto() *chatbotdbpb.Message {
	return nil
}

// GetMsgID - get message id
func (msg *testMessage) GetMsgID() string {
	return ""
}

// SetMsgID - set message id
func (msg *testMessage) SetMsgID(msgid string) {
}

// SetChatID - set chat id
func (msg *testMessage) SetChatID(chatid string) {

}

// SetText - set text
func (msg *testMessage) SetText(text string) {
	// msg.text = text
}

// GetOption - get option
func (msg *testMessage) GetOption(id int) string {
	return ""
}

func Test_IsMyMessage(t *testing.T) {
	chatbot.InitLogger(zapcore.InfoLevel, true, "./")

	p, err := NewPlugin("")
	if err != nil {
		t.Fatalf("Test_IsMyMessage NewPlugin Err")
	}

	arrOK := []string{"> version", "   > version    ", "> version"}
	for i := range arrOK {
		curmsg := &testMessage{
			strText: arrOK[i],
		}

		params := &chatbot.MessageParams{
			ChatBot:    nil,
			MgrPlugins: nil,
			Msg:        curmsg,
			LstStr:     strings.Fields(curmsg.GetText()),
		}

		if !p.IsMyMessage(params) {
			t.Fatalf("Test_IsMyMessage arrOK %v", arrOK[i])
		}
	}

	arrErr := []string{"123", ">> haha", ">   ", ">haha", "> help"}
	for i := range arrErr {
		curmsg := &testMessage{
			strText: arrErr[i],
		}

		params := &chatbot.MessageParams{
			ChatBot:    nil,
			MgrPlugins: nil,
			Msg:        curmsg,
			LstStr:     strings.Fields(curmsg.GetText()),
		}

		if p.IsMyMessage(params) {
			t.Fatalf("Test_IsMyMessage arrErr %v", arrErr[i])
		}
	}

	t.Log("Test_IsMyMessage OK")
}