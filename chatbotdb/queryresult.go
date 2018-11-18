package chatbotdb

import (
	pb "github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

// ResultMsg -
type ResultMsg struct {
	Msg pb.Message `json:"msg"`

	// Msg struct {
	// 	ChatID string `json:"chatID"`

	// 	From struct {
	// 		NickName  string `json:"nickName"`
	// 		UserID    string `json:"userID"`
	// 		UserName  string `json:"userName"`
	// 		LastMsgID int64  `json:"lastMsgID"`
	// 	} `json:"from"`

	// 	To struct {
	// 		NickName  string `json:"nickName"`
	// 		UserID    string `json:"userID"`
	// 		UserName  string `json:"userName"`
	// 		LastMsgID int64  `json:"lastMsgID"`
	// 	} `json:"to"`

	// 	Text      string   `json:"text"`
	// 	TimeStamp int64    `json:"timeStamp"`
	// 	MsgID     string   `json:"msgID"`
	// 	Options   []string `json:"options"`
	// 	Selected  int      `json:"selected"`
	// } `json:"msg"`
}

// ResultUser -
type ResultUser struct {
	User pb.User `json:"user"`
}
