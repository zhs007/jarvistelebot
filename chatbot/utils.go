package chatbot

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarviscore/base"
	"github.com/zhs007/jarvistelebot/basedef"
	"github.com/zhs007/jarvistelebot/chatbot/proto"
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
func SendTextMsg(bot ChatBot, user User, text string, srcmsg Message) error {
	// jarvisbase.Debug("SendTextMsg", zap.String("text", text))

	if len(text) >= basedef.MaxTextMessageSize {
		return SendFileMsg(bot, user, &chatbotdbpb.File{
			Filename: GetMD5String([]byte(text)) + ".txt",
			Data:     []byte(text),
		})
	}

	msg := bot.NewMsg("", "", nil, user, text, time.Now().Unix())
	if srcmsg != nil && srcmsg.InGroup() {
		// jarvisbase.Debug("SendTextMsg", zap.String("groupid", srcmsg.GetGroupID()))

		msg.SetGroupID(srcmsg.GetGroupID())
	}

	_, err := bot.SendMsg(msg)
	if err != nil {
		jarvisbase.Warn("SendTextMsg", zap.Error(err))
	}

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
	if err != nil {
		jarvisbase.Warn("SendFileMsg", zap.Error(err))
	}

	return err
}

// GetMD5String - md5 buf and return string
func GetMD5String(buf []byte) string {
	return fmt.Sprintf("%x", md5.Sum(buf))
}

// GetFileNameFromFullPath - get filename from fullpathfilname
func GetFileNameFromFullPath(fullname string) string {
	arr := strings.Split(fullname, "/")
	if len(arr) <= 1 {
		return fullname
	}

	return arr[len(arr)-1]
}

// GetFileNameFromFullPathNoExt - get filename from fullpathfilname
func GetFileNameFromFullPathNoExt(fullname string) string {
	fn := GetFileNameFromFullPath(fullname)

	arr := strings.Split(fn, ".")
	if len(arr) <= 1 {
		return fn
	}

	return strings.Join(arr[:len(arr)-1], ".")
}

// NewEmptyCommandLine - new EmptyMessage
func NewEmptyCommandLine(cmd string) proto.Message {
	return &chatbotpb.EmptyCommand{
		Cmd: cmd,
	}
}
