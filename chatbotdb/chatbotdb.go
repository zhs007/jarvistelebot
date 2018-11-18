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

const queryGetMsg = `mutation Msg($chatID: Int64!) {
	msg(chatID: $chatID) {
		chatID
		from
		to
		text
		timeStamp
		msgID
		options
		selected
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

	jarvisbase.Info("chatbot.CoreDB.saveMsg",
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

	jarvisbase.Info("chatbot.CoreDB.GetMsg", jarvisbase.JSON("result", result))

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

	jarvisbase.Info("chatbot.CoreDB.updUser",
		jarvisbase.JSON("result", result))

	return nil
}

// import (
// 	"context"
// 	"path"

// 	"github.com/graphql-go/graphql"
// 	"go.uber.org/zap"

// 	"github.com/zhs007/ankadb"
// 	"github.com/zhs007/jarviscore/base"
// )

// // NewChatBotDB - new chatbot db
// func NewChatBotDB(dbpath string, httpAddr string, engine string) (*ankadb.AnkaDB, error) {
// 	cfg := ankadb.NewConfig()

// 	cfg.AddrHTTP = httpAddr
// 	cfg.PathDBRoot = dbpath
// 	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
// 		Name:   "chatbotdb",
// 		Engine: engine,
// 		PathDB: path.Join(dbpath, "chatbotdb"),
// 	})

// 	ankaDB, err := ankadb.NewAnkaDB(cfg, newDBLogic())
// 	if ankaDB == nil {
// 		jarvisbase.Error("NewChatBotDB", zap.Error(err))

// 		return nil, err
// 	}

// 	jarvisbase.Info("NewChatBotDB", zap.String("dbpath", dbpath),
// 		zap.String("httpAddr", httpAddr), zap.String("engine", engine))

// 	return ankaDB, err
// }

// // chatbotDB -
// type chatbotDB struct {
// 	schema graphql.Schema
// }

// // newDBLogic -
// func newDBLogic() ankadb.DBLogic {
// 	var schema, _ = graphql.NewSchema(
// 		graphql.SchemaConfig{
// 			Query:    typeQuery,
// 			Mutation: typeMutation,
// 			// Types:    curTypes,
// 		},
// 	)

// 	return &chatbotDB{
// 		schema: schema,
// 	}
// }

// // OnQuery -
// func (logic *chatbotDB) OnQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error) {
// 	result := graphql.Do(graphql.Params{
// 		Schema:         logic.schema,
// 		RequestString:  request,
// 		VariableValues: values,
// 		Context:        ctx,
// 	})
// 	// if len(result.Errors) > 0 {
// 	// 	fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
// 	// }

// 	return result, nil
// }

// // OnQueryStream -
// func (logic *chatbotDB) OnQueryStream(ctx context.Context, request string, values map[string]interface{}, funcOnQueryStream ankadb.FuncOnQueryStream) error {
// 	return nil
// }
