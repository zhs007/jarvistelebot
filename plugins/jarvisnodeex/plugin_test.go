package pluginjarvisnodeex

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

// SetGroupID - set groupID
func (msg *testMessage) SetGroupID(groupID string) {

}

// GetGroupID - get groupID
func (msg *testMessage) GetGroupID() string {
	return ""
}

// InGroup - this message is from a group
func (msg *testMessage) InGroup() bool {
	return false
}

func Test_jarvisnodeexPlugin_IsMyMessage(t *testing.T) {
	chatbot.InitLogger(zapcore.InfoLevel, true, "./")

	p, err := NewPlugin("")
	if err != nil {
		t.Fatalf("jarvisnodeexPlugin NewPlugin Err")
	}

	arrOK := []string{
		"runscript -n jarvisroot",
		"runscript -n 123 -d abc.output -d abc1.output",
	}

	fd, err := chatbot.LoadScriptFile("../../test/test.sh")
	if err != nil {
		t.Fatalf("Test_jarvisnodeexPlugin_IsMyMessage %v", err)
	}

	for i := range arrOK {
		curmsg := &testMessage{
			strText: arrOK[i],
		}

		curmsg.SetFile(fd)

		params := &chatbot.MessageParams{
			ChatBot:    nil,
			MgrPlugins: nil,
			Msg:        curmsg,
			LstStr:     strings.Fields(curmsg.GetText()),
		}

		_, err := p.ParseMessage(params)
		if err != nil {
			t.Fatalf("Test_jarvisnodeexPlugin_IsMyMessage arrOK %v %v", arrOK[i], err)
		}
	}

	arrErr := []string{
		"runscript ",
		"runscript abc",
	}

	for i := range arrErr {
		curmsg := &testMessage{
			strText: arrErr[i],
		}

		curmsg.SetFile(fd)

		params := &chatbot.MessageParams{
			ChatBot:    nil,
			MgrPlugins: nil,
			Msg:        curmsg,
			LstStr:     strings.Fields(curmsg.GetText()),
		}

		_, err := p.ParseMessage(params)
		if err == nil {
			t.Fatalf("Test_jarvisnodeexPlugin_IsMyMessage arrOK %v", arrErr[i])
		}
	}

	t.Log("Test_jarvisnodeexPlugin_IsMyMessage OK")
}
