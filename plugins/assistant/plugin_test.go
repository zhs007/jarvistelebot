package pluginassistant

import (
	"testing"

	jarvisbase "github.com/zhs007/jarviscore/base"

	"go.uber.org/zap/zapcore"

	"github.com/zhs007/jarvistelebot/chatbot"
	chatbotdbpb "github.com/zhs007/jarvistelebot/chatbotdb/proto"
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

func Test_assistantPlugin_IsMyMessage(t *testing.T) {
	subname := jarvisbase.BuildLogSubFilename("testnode", "0.1.0")
	jarvisbase.InitLogger(zapcore.InfoLevel, true, "./", subname)
	chatbot.InitLogger(zapcore.InfoLevel, true, "./")

	p, err := NewPlugin("../../test")
	if err != nil {
		t.Fatalf("Test_assistantPlugin_IsMyMessage NewPlugin Err")
	}

	arrOK := []string{
		"expnotes --graph",
		"expnotes -g",
		"expnotes -g true",
		"expnotes",
	}
	for i := range arrOK {
		curmsg := &testMessage{
			strText: arrOK[i],
		}

		params := &chatbot.MessageParams{
			ChatBot:    nil,
			MgrPlugins: nil,
			Msg:        curmsg,
			LstStr:     chatbot.SplitString(curmsg.GetText()),
		}

		_, err := p.ParseMessage(params)
		if err != nil {
			t.Fatalf("Test_assistantPlugin_IsMyMessage arrOK %v %v", arrOK[i], err)
		}
	}

	// arrErr := []string{"123", "haha", ">   ", ">haha", "help", "user -userid 123 -n 456"}
	// for i := range arrErr {
	// 	curmsg := &testMessage{
	// 		strText: arrErr[i],
	// 	}

	// 	params := &chatbot.MessageParams{
	// 		ChatBot:    nil,
	// 		MgrPlugins: nil,
	// 		Msg:        curmsg,
	// 		LstStr:     strings.Fields(curmsg.GetText()),
	// 	}

	// 	_, err := p.ParseMessage(params)
	// 	if err == nil {
	// 		t.Fatalf("Test_assistantPlugin_IsMyMessage arrOK %v", arrErr[i])
	// 	}
	// }

	t.Log("Test_assistantPlugin_IsMyMessage OK")
}
