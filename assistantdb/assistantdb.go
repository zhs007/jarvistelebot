package assistantdb

import (
	"context"
	"path"
	"time"

	"github.com/zhs007/ankadb"
	"github.com/zhs007/jarviscore/base"
	pb "github.com/zhs007/jarvistelebot/assistantdb/proto"
	"go.uber.org/zap"
)

const queryAssistantData = `{
	assistantData {
		maxMsgID
		keys
	}
}`

const queryUpdMsg = `mutation UpdMsg($msg: MessageInput!) {
	updMsg(msg: $msg) {
		msgID
		data
		keys
		createTime
		updateTime
	}
}`

const queryUpdAssistantData = `mutation UpdAssistantData($dat: AssistantDataInput!) {
	updAssistantData(dat: $dat) {
		maxMsgID
		keys
	}
}`

// AssistantDB -
type AssistantDB struct {
	ankaDB *ankadb.AnkaDB
	dat    *pb.AssistantData
}

// NewAssistantDB - new assistant db
func NewAssistantDB(dbpath string, httpAddr string, engine string) (*AssistantDB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   "assistantdb",
		Engine: engine,
		PathDB: path.Join(dbpath, "assistantdb"),
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, newDBLogic())
	if ankaDB == nil {
		jarvisbase.Error("NewAssistantDB", zap.Error(err))

		return nil, err
	}

	jarvisbase.Info("NewAssistantDB", zap.String("dbpath", dbpath),
		zap.String("httpAddr", httpAddr), zap.String("engine", engine))

	db := &AssistantDB{
		ankaDB: ankaDB,
	}

	err = db.loadAssistantDB()
	if err != nil {
		jarvisbase.Error("NewAssistantDB", zap.Error(err))

		return nil, err
	}

	return db, err
}

func (db *AssistantDB) loadAssistantDB() error {
	result, err := db.ankaDB.LocalQuery(context.Background(), queryAssistantData, nil)
	if err != nil {
		return err
	}

	jarvisbase.Info("AssistantDB.loadAssistantDB",
		jarvisbase.JSON("result", result))

	rad := &ResultAssistantData{}
	err = ankadb.MakeObjFromResult(result, rad)
	if err != nil {
		return err
	}

	db.dat = ResultAssistantData2AssistantData(rad)

	return nil
}

// newMsgID - new MsgID
func (db *AssistantDB) newMsgID() int64 {
	if db.dat == nil {
		jarvisbase.Error("AssistantDB.newMsgID", zap.Error(ErrNoAssistantData))

		return -1
	}

	db.dat.MaxMsgID++

	return db.dat.MaxMsgID
}

// updMsg - update msg to db
func (db *AssistantDB) updMsg(msg *pb.Message) error {
	params := make(map[string]interface{})

	err := ankadb.MakeParamsFromMsg(params, "msg", msg)
	if err != nil {
		return err
	}

	result, err := db.ankaDB.LocalQuery(context.Background(), queryUpdMsg, params)
	if err != nil {
		return err
	}

	rum := &ResultUpdMsg{}
	err = ankadb.MakeObjFromResult(result, rum)
	if err != nil {
		return err
	}

	return nil
}

// updAssistantData - update AssistantData to db
func (db *AssistantDB) updAssistantData(dat *pb.AssistantData) error {
	params := make(map[string]interface{})

	err := ankadb.MakeParamsFromMsg(params, "dat", dat)
	if err != nil {
		return err
	}

	result, err := db.ankaDB.LocalQuery(context.Background(), queryUpdAssistantData, params)
	if err != nil {
		return err
	}

	rum := &ResultUpdMsg{}
	err = ankadb.MakeObjFromResult(result, rum)
	if err != nil {
		return err
	}

	return nil
}

// NewMsg - new Message
func (db *AssistantDB) NewMsg(dat string, keys []string) error {
	msg := &pb.Message{
		MsgID:      db.newMsgID(),
		Data:       dat,
		Keys:       keys,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	err := db.updMsg(msg)
	if err != nil {
		return err
	}

	err = db.updAssistantData(db.dat)
	if err != nil {
		return err
	}

	return nil
}

// Start - start
func (db *AssistantDB) Start(ctx context.Context) error {
	return db.ankaDB.Start(ctx)
}
