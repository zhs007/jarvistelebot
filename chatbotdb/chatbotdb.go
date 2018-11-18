package chatbotdb

import (
	"context"
	"path"

	"github.com/zhs007/ankadb"
	"github.com/zhs007/jarviscore/base"
	pb "github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"go.uber.org/zap"
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

const queryGetMsg = `mutation Msg($chatID: String!) {
	msg(chatID: $chatID) {
		chatID
		from {		
			nickName
			userID
			userName
			lastMsgID
		}
		to {
			nickName
			userID
			userName
			lastMsgID
		}
		text
		timeStamp
		msgID
		options
		selected
	}
}`

const queryGetUser = `mutation User($uerID: String!) {
	user(uerID: $uerID) {
		nickName
		userID
		userName
		lastMsgID
	}
}`

// ChatBotDB - chatbotdb
type ChatBotDB struct {
	db *ankadb.AnkaDB
}

// NewChatBotDB - new ChatBotDB
func NewChatBotDB(dbpath string, httpAddr string, engine string) (*ChatBotDB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   "chatbotdb",
		Engine: engine,
		PathDB: path.Join(dbpath, "chatbotdb"),
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, newDBLogic())
	if ankaDB == nil {
		jarvisbase.Error("NewChatBotDB", zap.Error(err))

		return nil, err
	}

	jarvisbase.Info("NewChatBotDB", zap.String("dbpath", dbpath),
		zap.String("httpAddr", httpAddr), zap.String("engine", engine))

	// return ankaDB, err

	// db, err := chatbotdb.NewChatBotDB(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)
	// if err != nil {
	// 	return nil, err
	// }

	return &ChatBotDB{
		db: ankaDB,
	}, nil
}

// SaveMsg - save message
func (db *ChatBotDB) SaveMsg(msg *pb.Message) error {
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

	jarvisbase.Info("ChatBotDB.saveMsg",
		jarvisbase.JSON("result", result))

	return nil
}

// GetMsg - get message
func (db *ChatBotDB) GetMsg(chatid string) (*pb.Message, error) {
	if db.db == nil {
		return nil, ErrChatBotDBNil
	}

	params := make(map[string]interface{})
	params["chatID"] = chatid

	result, err := db.db.LocalQuery(context.Background(), queryGetMsg, params)
	if err != nil {
		return nil, err
	}

	rmsg := &ResultMsg{}
	err = ankadb.MakeObjFromResult(result, rmsg)
	if err != nil {
		return nil, err
	}

	jarvisbase.Info("ChatBotDB.GetMsg", jarvisbase.JSON("result", result))

	return &rmsg.Msg, nil
}

// UpdUser - update user
func (db *ChatBotDB) UpdUser(user *pb.User) error {
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

	jarvisbase.Info("ChatBotDB.updUser",
		jarvisbase.JSON("result", result))

	return nil
}

// GetUser - get user
func (db *ChatBotDB) GetUser(userid string) (*pb.User, error) {
	if db.db == nil {
		return nil, ErrChatBotDBNil
	}

	params := make(map[string]interface{})
	params["userID"] = userid

	result, err := db.db.LocalQuery(context.Background(), queryGetUser, params)
	if err != nil {
		return nil, err
	}

	ruser := &ResultUser{}
	err = ankadb.MakeObjFromResult(result, ruser)
	if err != nil {
		return nil, err
	}

	jarvisbase.Info("ChatBotDB.GetUser", jarvisbase.JSON("result", result))

	return &ruser.User, nil
}
