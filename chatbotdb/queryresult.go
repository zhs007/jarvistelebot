package chatbotdb

import (
	pb "github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

// ResultMsg -
type ResultMsg struct {
	// Msg pb.Message `json:"msg"`

	Msg struct {
		ChatID string `json:"chatID"`

		From struct {
			NickName  string `json:"nickName"`
			UserID    string `json:"userID"`
			UserName  string `json:"userName"`
			LastMsgID int64  `json:"lastMsgID"`
		} `json:"from"`

		To struct {
			NickName  string `json:"nickName"`
			UserID    string `json:"userID"`
			UserName  string `json:"userName"`
			LastMsgID int64  `json:"lastMsgID"`
		} `json:"to"`

		Text      string   `json:"text"`
		TimeStamp int64    `json:"timeStamp"`
		MsgID     string   `json:"msgID"`
		Options   []string `json:"options"`
		Selected  int      `json:"selected"`
	} `json:"msg"`
}

// ResultUser -
type ResultUser struct {
	User pb.User `json:"user"`
}

// ResultMsg2Msg - ResultMsg -> Message
func ResultMsg2Msg(result *ResultMsg) *pb.Message {
	msg := &pb.Message{
		ChatID:    result.Msg.ChatID,
		Text:      result.Msg.Text,
		TimeStamp: result.Msg.TimeStamp,
		MsgID:     result.Msg.MsgID,
		Selected:  int32(result.Msg.Selected),
	}

	if result.Msg.From.UserID != "" {
		msg.From = &pb.User{
			NickName:  result.Msg.From.NickName,
			UserID:    result.Msg.From.UserID,
			UserName:  result.Msg.From.UserName,
			LastMsgID: result.Msg.From.LastMsgID,
		}
	}

	if result.Msg.To.UserID != "" {
		msg.To = &pb.User{
			NickName:  result.Msg.To.NickName,
			UserID:    result.Msg.To.UserID,
			UserName:  result.Msg.To.UserName,
			LastMsgID: result.Msg.To.LastMsgID,
		}
	}

	for _, v := range result.Msg.Options {
		msg.Options = append(msg.Options, v)
	}

	return msg
}
