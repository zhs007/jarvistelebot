package pluginjarvisnode

import (
	"strings"
	"testing"

	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

// testMessage - test Message
type testMessage struct {
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

func Test_IsMyMessage(t *testing.T) {
	p := newPlugin()

	arrOK := []string{"> 123", "> haha", "   > haha    ", "> >haha"}
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
			t.Fatalf("Test_IsMyMessage Err")
		}
	}

	arrErr := []string{"123", ">> haha", ">   ", ">haha"}
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
			t.Fatalf("Test_IsMyMessage Err")
		}
	}

	t.Log("Test_IsMyMessage OK")
}
