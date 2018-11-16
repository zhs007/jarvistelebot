package chatbot

import (
	"context"

	"github.com/zhs007/ankadb"
	"github.com/zhs007/jarvistelebot/chatbotdb"
)

const querySaveMsg = `mutation NewMsg($chatID: ID!, $fromNickName: String!, $fromUserID: ID!, $text: String!, $timeStamp: Timestamp!) {
	newMsg(chatID: $chatID, fromNickName: $fromNickName, fromUserID: $fromUserID, text: $text, timeStamp: $timeStamp) {
		chatID
	}
}`

const queryUpdUser = `mutation UpdUser($nickName: String!, $userID: ID!, userName: ID!, $lastMsgID: Int64!) {
	updUser(nickName: $nickName, userID: $userID, userName: $userName, lastMsgID: $lastMsgID) {
		nickName
		userID
		userName
		lastMsgID
	}
}`

// chatBotDB - chatbotdb
type chatBotDB struct {
	db *ankadb.AnkaDB
}

func newChatDB(cfg *Config) (*chatBotDB, error) {
	db, err := chatbotdb.NewChatBotDB(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)
	if err != nil {
		return nil, err
	}

	return &chatBotDB{
		db: db,
	}, nil
}

// saveMsg - save message
func (db *chatBotDB) saveMsg(msg Message) error {
	if db.db == nil {
		return ErrChatBotDBNil
	}

	params := make(map[string]interface{})
	params["chatID"] = msg.GetChatID()
	params["fromNickName"] = msg.GetFrom().GetNickName()
	params["fromUserID"] = msg.GetFrom().GetUserID()
	params["text"] = msg.GetText()
	params["timeStamp"] = msg.GetTimeStamp()

	result, err := db.db.LocalQuery(context.Background(), querySaveMsg, params)
	if err != nil {
		return err
	}

	Info("chatBotDB.saveMsg",
		JSON("result", result))

	return nil
}

// updUser - update user
func (db *chatBotDB) updUser(user User) error {
	if db.db == nil {
		return ErrChatBotDBNil
	}

	params := make(map[string]interface{})
	params["nickName"] = user.GetNickName()
	params["userID"] = user.GetUserID()
	params["userName"] = user.GetUserName()
	params["lastMsgID"] = user.GetLastMsgID()

	result, err := db.db.LocalQuery(context.Background(), queryUpdUser, params)
	if err != nil {
		return err
	}

	Info("chatBotDB.updUser",
		JSON("result", result))

	return nil
}
