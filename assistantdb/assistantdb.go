package assistantdb

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb"
	"github.com/zhs007/jarviscore/base"
	pb "github.com/zhs007/jarvistelebot/assistantdb/proto"
	"go.uber.org/zap"
)

const queryUserAssistantInfo = `query UserAssistantInfo($userID: ID!) {
	userAssistantInfo(userID: $userID) {
		maxNoteID
		keys
	}
}`

const queryGetNote = `query Note($userID: ID!, $noteID: Int64!) {
	note(userID: $userID, noteID: $noteID) {
		noteID
		data
		keys
		createTime
		updateTime
	}
}`

const queryUpdNote = `mutation UpdNote($userID: ID!, $note: NoteInput!) {
	updNote(userID: $userID, note: $note) {
		noteID
		data
		keys
		createTime
		updateTime
	}
}`

const queryRmNote = `mutation RmNote($userID: ID!, $noteID: Int64!) {
	rmNote(userID: $userID, noteID: $noteID)
}`

const queryUpdKeyInfo = `mutation UpdKeyInfo($userID: ID!, $key: String!, $keyinfo: KeyInfoInput!) {
	updKeyInfo(userID: $userID, key: $key, keyinfo: $keyinfo) {
		noteIDs
	}
}`

const queryGetKeyInfo = `query KeyInfo($userID: ID!, $key: String!) {
	keyInfo(userID: $userID, key: $key) {
		noteIDs
	}
}`

const queryUpdAssistantData = `mutation UpdUserAssistantInfo($userID: ID!, $uai: UserAssistantInfoInput!) {
	updUserAssistantInfo(userID: $userID, uai: $uai) {
		maxNoteID
		keys
	}
}`

// AssistantDB -
type AssistantDB struct {
	ankaDB ankadb.AnkaDB
	// dat    *pb.UserAssistantInfo
}

// NewAssistantDB - new assistant db
func NewAssistantDB(dbpath string, httpAddr string, engine string) (*AssistantDB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   "assistantdb",
		Engine: engine,
		PathDB: "assistantdb",
	})

	dblogic, err := ankadb.NewBaseDBLogic(graphql.SchemaConfig{
		Query:    typeQuery,
		Mutation: typeMutation,
	})
	if err != nil {
		jarvisbase.Error("newdb", zap.Error(err))

		return nil, err
	}

	ankaDB, err := ankadb.NewAnkaDB(cfg, dblogic)
	if ankaDB == nil {
		jarvisbase.Error("NewAssistantDB", zap.Error(err))

		return nil, err
	}

	jarvisbase.Info("NewAssistantDB", zap.String("dbpath", dbpath),
		zap.String("httpAddr", httpAddr), zap.String("engine", engine))

	db := &AssistantDB{
		ankaDB: ankaDB,
	}

	// err = db.loadAssistantDB()
	// if err != nil {
	// 	jarvisbase.Error("NewAssistantDB", zap.Error(err))

	// 	return nil, err
	// }

	return db, err
}

// LoadUserAssistantInfo - load UserAssistantInfo
func (db *AssistantDB) LoadUserAssistantInfo(userID string) (*pb.UserAssistantInfo, error) {
	params := make(map[string]interface{})
	params["userID"] = userID

	result, err := db.ankaDB.Query(context.Background(), queryUserAssistantInfo, params)
	if err != nil {
		return nil, err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		return nil, err
	}

	// jarvisbase.Info("AssistantDB.LoadUserAssistantInfo",
	// 	jarvisbase.JSON("result", result))

	ruai := &ResultUserAssistantInfo{}
	err = ankadb.MakeObjFromResult(result, ruai)
	if err != nil {
		return nil, err
	}

	return ResultUserAssistantInfo2UserAssistantInfo(ruai), nil
}

// // newNoteID - new NoteID
// func (db *AssistantDB) newNoteID() int64 {
// 	if db.dat == nil {
// 		jarvisbase.Error("AssistantDB.newNoteID", zap.Error(ErrNoAssistantData))

// 		return -1
// 	}

// 	db.dat.MaxNoteID++

// 	return db.dat.MaxNoteID
// }

// UpdNote - update note to db
func (db *AssistantDB) UpdNote(userID string, note *pb.Note) (*pb.Note, error) {
	params := make(map[string]interface{})

	params["userID"] = userID

	err := ankadb.MakeParamsFromMsg(params, "note", note)
	if err != nil {
		return nil, err
	}

	result, err := db.ankaDB.Query(context.Background(), queryUpdNote, params)
	if err != nil {
		return nil, err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		return nil, err
	}

	jarvisbase.Info("AssistantDB.UpdNote",
		jarvisbase.JSON("result", result))

	run := &ResultUpdNote{}
	err = ankadb.MakeObjFromResult(result, run)
	if err != nil {
		return nil, err
	}

	return ResultUpdNote2Note(run), nil
}

// RmNote - remove note to db
func (db *AssistantDB) RmNote(userID string, noteID int64) (string, error) {
	params := make(map[string]interface{})

	params["userID"] = userID
	params["noteID"] = noteID

	result, err := db.ankaDB.Query(context.Background(), queryRmNote, params)
	if err != nil {
		return "", err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		return "", err
	}

	jarvisbase.Info("AssistantDB.RmNote",
		jarvisbase.JSON("result", result))

	run := &ResultRmNote{}
	err = ankadb.MakeObjFromResult(result, run)
	if err != nil {
		return "", err
	}

	return ResultRmNote2String(run), nil
}

// UpdUserAssistantInfo - update AssistantData to db
func (db *AssistantDB) UpdUserAssistantInfo(userID string, uai *pb.UserAssistantInfo) (*pb.UserAssistantInfo, error) {
	params := make(map[string]interface{})

	params["userID"] = userID

	err := ankadb.MakeParamsFromMsg(params, "uai", uai)
	if err != nil {
		return nil, err
	}

	result, err := db.ankaDB.Query(context.Background(), queryUpdAssistantData, params)
	if err != nil {
		return nil, err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		return nil, err
	}

	jarvisbase.Info("AssistantDB.UpdUserAssistantInfo",
		jarvisbase.JSON("result", result))

	ruuai := &ResultUpdUserAssistantInfo{}
	err = ankadb.MakeObjFromResult(result, ruuai)
	if err != nil {
		return nil, err
	}

	return ResultUpdUserAssistantInfo2UserAssistantInfo(ruuai), nil
}

// GetNote - get note from db
func (db *AssistantDB) GetNote(userID string, noteID int64) (*pb.Note, error) {
	params := make(map[string]interface{})

	params["userID"] = userID
	params["noteID"] = noteID

	result, err := db.ankaDB.Query(context.Background(), queryGetNote, params)
	if err != nil {
		jarvisbase.Warn("AssistantDB.GetNote", zap.Error(err), jarvisbase.JSON("params", params))

		return nil, err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		jarvisbase.Warn("AssistantDB.GetNote", zap.Error(err), jarvisbase.JSON("params", params))

		return nil, err
	}

	// jarvisbase.Info("AssistantDB.GetNote",
	// 	jarvisbase.JSON("result", result))

	rn := &ResultNote{}
	err = ankadb.MakeObjFromResult(result, rn)
	if err != nil {
		jarvisbase.Warn("AssistantDB.GetNote", zap.Error(err), jarvisbase.JSON("params", params))

		return nil, err
	}

	return ResultNote2Note(rn), nil
}

// // UpdNote - update note
// func (db *AssistantDB) UpdNote(userID string, note *pb.Note) (*pb.Note, error) {
// 	// msg := &pb.Note{
// 	// 	NoteID:     db.newNoteID(),
// 	// 	Data:       dat,
// 	// 	Keys:       keys,
// 	// 	CreateTime: time.Now().Unix(),
// 	// 	UpdateTime: time.Now().Unix(),
// 	// }

// 	// if msg.MsgID < 0 {
// 	// 	return nil, ErrNoAssistantData
// 	// }

// 	err := db.updMsg(msg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = db.updAssistantData(db.dat)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return msg, nil
// }

// Start - start
func (db *AssistantDB) Start(ctx context.Context) error {
	return db.ankaDB.Start(ctx)
}

// UpdKeyInfo - update keyinfo to db
func (db *AssistantDB) UpdKeyInfo(userID string, key string, keyinfo *pb.KeyInfo) (*pb.KeyInfo, error) {
	params := make(map[string]interface{})

	params["userID"] = userID
	params["key"] = key

	err := ankadb.MakeParamsFromMsg(params, "keyinfo", keyinfo)
	if err != nil {
		jarvisbase.Warn("AssistantDB.UpdKeyInfo", zap.Error(err), jarvisbase.JSON("params", params))

		return nil, err
	}

	result, err := db.ankaDB.Query(context.Background(), queryUpdKeyInfo, params)
	if err != nil {
		jarvisbase.Warn("AssistantDB.UpdKeyInfo", zap.Error(err), jarvisbase.JSON("params", params))

		return nil, err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		jarvisbase.Warn("AssistantDB.UpdKeyInfo", zap.Error(err), jarvisbase.JSON("params", params))

		return nil, err
	}

	// jarvisbase.Info("AssistantDB.UpdKeyInfo",
	// 	jarvisbase.JSON("result", result))

	ruki := &ResultUpdKeyInfo{}
	err = ankadb.MakeObjFromResult(result, ruki)
	if err != nil {
		jarvisbase.Warn("AssistantDB.UpdKeyInfo", zap.Error(err), jarvisbase.JSON("params", params))

		return nil, err
	}

	return ResultUpdKeyInfo2KeyInfo(ruki), nil
}

// GetKeyInfo - get keyinfo
func (db *AssistantDB) GetKeyInfo(userID string, key string) (*pb.KeyInfo, error) {
	params := make(map[string]interface{})

	params["userID"] = userID
	params["key"] = key

	result, err := db.ankaDB.Query(context.Background(), queryGetKeyInfo, params)
	if err != nil {
		return nil, err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		return nil, err
	}

	jarvisbase.Info("AssistantDB.GetKeyInfo",
		jarvisbase.JSON("result", result))

	rki := &ResultKeyInfo{}
	err = ankadb.MakeObjFromResult(result, rki)
	if err != nil {
		return nil, err
	}

	return ResultKeyInfo2KeyInfo(rki), nil
}
