package chatbot

import (
	"encoding/json"
	"time"

	"github.com/zhs007/jarviscore/base"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"go.uber.org/zap"
)

// FormatJSON - format JSON string
func FormatJSON(str string) (string, error) {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(str), &mapResult)
	if err != nil {
		return "", err
	}

	jsonStr, err := json.MarshalIndent(mapResult, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonStr), nil
}

// FormatJSONObj - format JSON string
func FormatJSONObj(obj interface{}) (string, error) {
	jsonStr, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonStr), nil
}

// SendTextMsg - sendmsg
func SendTextMsg(bot ChatBot, user User, text string) error {
	msg := bot.NewMsg("", "", nil, user, text, time.Now().Unix())

	_, err := bot.SendMsg(msg)

	return err
}

// SendMsgWithOptions - send msg with options
func SendMsgWithOptions(bot ChatBot, user User, text string, options []string, callback FuncMsgCallback) error {
	msg := bot.NewMsg("", "", nil, user, text, time.Now().Unix())

	for _, v := range options {
		msg.AddOption(v)
	}

	nmsg, err := bot.SendMsg(msg)
	if err != nil {
		return err
	}

	nmsg.SetChatID(MakeChatID(user.GetUserID(), nmsg.GetMsgID()))
	err = bot.SaveMsg(nmsg)
	if err != nil {
		jarvisbase.Warn("SendMsgWithOptions:SaveMsg", zap.Error(err))

		return err
	}

	err = bot.AddMsgCallback(nmsg, callback)
	if err != nil {
		jarvisbase.Warn("SendMsgWithOptions:AddMsgCallback", zap.Error(err))

		return err
	}

	return nil
}

// MakeChatID - make chatid
func MakeChatID(userid string, msgid string) string {
	return userid + ":" + msgid
}

// SendFileMsg - sendmsg
func SendFileMsg(bot ChatBot, user User, fd *chatbotdbpb.File) error {
	msg := bot.NewMsg("", "", nil, user, "", time.Now().Unix())
	msg.SetFile(fd)

	_, err := bot.SendMsg(msg)

	return err
}
