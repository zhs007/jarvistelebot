package chatbot

// import (
// 	"context"

// 	"github.com/zhs007/ankadb"
// 	"github.com/zhs007/jarvistelebot/chatbotdb"
// 	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
// )

// const querySaveMsg = `mutation NewMsg($chatID: ID!, $fromNickName: String!, $fromUserID: ID!, $text: String!, $timeStamp: Timestamp!) {
// 	newMsg(chatID: $chatID, fromNickName: $fromNickName, fromUserID: $fromUserID, text: $text, timeStamp: $timeStamp) {
// 		chatID
// 	}
// }`

// const queryUpdUser = `mutation UpdUser($nickName: String!, $userID: ID!, userName: ID!, $lastMsgID: Int64!) {
// 	updUser(nickName: $nickName, userID: $userID, userName: $userName, lastMsgID: $lastMsgID) {
// 		nickName
// 		userID
// 		userName
// 		lastMsgID
// 	}
// }`

// const queryGetMsg = `mutation Msg($chatID: Int64!) {
// 	msg(chatID: $chatID) {
// 		chatID
// 		from
// 		to
// 		text
// 		timeStamp
// 		msgID
// 		options
// 		selected
// 	}
// }`

// // CoreDB - chatbotdb
// type CoreDB struct {
// 	db *ankadb.AnkaDB
// }

// func newChatDB(cfg *Config) (*CoreDB, error) {
// 	db, err := chatbotdb.NewChatBotDB(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &CoreDB{
// 		db: db,
// 	}, nil
// }

// // SaveMsg - save message
// func (db *CoreDB) SaveMsg(msg Message) error {
// 	if db.db == nil {
// 		return ErrChatBotDBNil
// 	}

// 	params := make(map[string]interface{})
// 	params["chatID"] = msg.GetChatID()
// 	params["fromNickName"] = msg.GetFrom().GetNickName()
// 	params["fromUserID"] = msg.GetFrom().GetUserID()
// 	params["text"] = msg.GetText()
// 	params["timeStamp"] = msg.GetTimeStamp()

// 	result, err := db.db.LocalQuery(context.Background(), querySaveMsg, params)
// 	if err != nil {
// 		return err
// 	}

// 	Info("chatbot.CoreDB.saveMsg",
// 		JSON("result", result))

// 	return nil
// }

// // GetMsg - get message
// func (db *CoreDB) GetMsg(chatid string) (*chatbotdbpb.Message, error) {
// 	if db.db == nil {
// 		return nil, ErrChatBotDBNil
// 	}

// 	params := make(map[string]interface{})
// 	params["chatID"] = chatid

// 	result, err := db.db.LocalQuery(context.Background(), queryGetMsg, params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rmsg := &chatbotdb.ResultMsg{}
// 	err = ankadb.MakeObjFromResult(result, rmsg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	Info("chatbot.CoreDB.GetMsg", JSON("result", result))

// 	return &rmsg.Msg, nil
// }

// // UpdUser - update user
// func (db *CoreDB) UpdUser(user User) error {
// 	if db.db == nil {
// 		return ErrChatBotDBNil
// 	}

// 	params := make(map[string]interface{})
// 	params["nickName"] = user.GetNickName()
// 	params["userID"] = user.GetUserID()
// 	params["userName"] = user.GetUserName()
// 	params["lastMsgID"] = user.GetLastMsgID()

// 	result, err := db.db.LocalQuery(context.Background(), queryUpdUser, params)
// 	if err != nil {
// 		return err
// 	}

// 	Info("chatbot.CoreDB.updUser",
// 		JSON("result", result))

// 	return nil
// }
