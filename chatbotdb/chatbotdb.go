package chatbotdb

import (
	"context"
	"encoding/base64"
	"path"

	"github.com/zhs007/ankadb"
	"github.com/zhs007/jarviscore/base"
	pb "github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"go.uber.org/zap"
)

const querySaveMsg = `mutation NewMsg($msg: MessageInput!) {
	newMsg(msg: $msg) {
		chatID
	}
}`

const querySaveUserScript = `mutation UpdUserScript($userID: ID!, $jarvisNodeName: ID!, $scriptName: ID!, $file: FileInput!) {
	updUserScript(userID: $userID, jarvisNodeName: $jarvisNodeName, scriptName: $scriptName, file: $file) {
		scriptName
	}
}`

const queryUpdUser = `mutation UpdUser($user: UserInput!) {
	updUser(user: $user) {
		nickName
		userID
		userName
		lastMsgID
	}
}`

const queryRmScript = `mutation RemoveUserScript($userID: ID!, $scriptName: ID!) {
	removeUserScript(userID: $userID, scriptName: $scriptName)
}`

const queryGetMsg = `query Msg($chatID: ID!) {
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
		file {
			filename
			strData
			fileType
		}
	}
}`

const queryGetUser = `query User($userID: ID!) {
	user(userID: $userID) {
		nickName
		userID
		userName
		lastMsgID
	}
}`

const queryGetUserWithUserName = `query UserWithUserName($userName: ID!) {
	userWithUserName(userName: $userName) {
		nickName
		userID
		userName
		lastMsgID
	}
}`

const queryGetUserScript = `query UserScript($userID: ID!, $scriptName: ID!) {
	userScript(userID: $userID, scriptName: $scriptName) {
		scriptName
		jarvisNodeName
		file {
			filename
			strData
			fileType
		}
	}
}`

const queryGetUsers = `query Users($snapshotID: Int64!, $beginIndex: Int!, $nums: Int!) {
	users(snapshotID: $snapshotID, beginIndex: $beginIndex, nums: $nums) {
		snapshotID, endIndex, maxIndex, 
		users {
			nickName
			userID
			userName
			lastMsgID
		}
	}
}`

const queryGetUserScripts = `query UserScripts($userID: String!, $jarvisNodeName: String!) {
	userScripts(userID: $userID, jarvisNodeName: $jarvisNodeName) {
		scripts {
			scriptName
		}
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

	if msg.File != nil && msg.File.Data != nil {
		msg.File.StrData = base64.StdEncoding.EncodeToString(msg.File.Data)
	}

	params := make(map[string]interface{})
	err := ankadb.MakeParamsFromMsg(params, "msg", msg)
	if err != nil {
		return err
	}

	// params["chatID"] = msg.GetChatID()
	// params["fromNickName"] = msg.GetFrom().GetNickName()
	// params["fromUserID"] = msg.GetFrom().GetUserID()
	// params["text"] = msg.GetText()
	// params["timeStamp"] = msg.GetTimeStamp()

	result, err := db.db.LocalQuery(context.Background(), querySaveMsg, params)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.SaveMsg:LocalQuery", zap.Error(err))

		return err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.SaveMsg:GetResultError", zap.Error(err))

		return err
	}

	jarvisbase.Debug("ChatBotDB.saveMsg",
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
		jarvisbase.Warn("ChatBotDB.GetMsg:LocalQuery", zap.Error(err))

		return nil, err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetMsg:GetResultError", zap.Error(err))

		return nil, err
	}

	jarvisbase.Debug("ChatBotDB.GetMsg", jarvisbase.JSON("result", result))

	rmsg := &ResultMsg{}
	err = ankadb.MakeObjFromResult(result, rmsg)
	if err != nil {
		return nil, err
	}

	pbmsg, err := ResultMsg2Msg(rmsg)
	if err != nil {
		return nil, err
	}

	return pbmsg, nil
}

// UpdUser - update user
func (db *ChatBotDB) UpdUser(user *pb.User) error {
	if db.db == nil {
		return ErrChatBotDBNil
	}

	params := make(map[string]interface{})
	err := ankadb.MakeParamsFromMsg(params, "user", user)
	if err != nil {
		return err
	}

	// params := make(map[string]interface{})
	// params["nickName"] = user.GetNickName()
	// params["userID"] = user.GetUserID()
	// params["userName"] = user.GetUserName()
	// params["lastMsgID"] = user.GetLastMsgID()

	result, err := db.db.LocalQuery(context.Background(), queryUpdUser, params)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.UpdUser:LocalQuery", zap.Error(err))

		return err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.UpdUser:GetResultError", zap.Error(err))

		return err
	}

	jarvisbase.Debug("ChatBotDB.updUser",
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
		jarvisbase.Warn("ChatBotDB.GetUser:LocalQuery", zap.Error(err))

		return nil, err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUser:GetResultError", zap.Error(err))

		return nil, err
	}

	ruser := &ResultUser{}
	err = ankadb.MakeObjFromResult(result, ruser)
	if err != nil {
		return nil, err
	}

	// jarvisbase.Debug("ChatBotDB.GetUser", jarvisbase.JSON("result", result))

	return ResultUser2User(ruser)
}

// GetUserWithUserName - get user with username
func (db *ChatBotDB) GetUserWithUserName(username string) (*pb.User, error) {
	if db.db == nil {
		return nil, ErrChatBotDBNil
	}

	params := make(map[string]interface{})
	params["userName"] = username

	result, err := db.db.LocalQuery(context.Background(), queryGetUserWithUserName, params)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUserWithUserName:LocalQuery", zap.Error(err))

		return nil, err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUserWithUserName:GetResultError", zap.Error(err))

		return nil, err
	}

	ruser := &ResultUserWithUserName{}
	err = ankadb.MakeObjFromResult(result, ruser)
	if err != nil {
		return nil, err
	}

	// jarvisbase.Debug("ChatBotDB.GetUserWithUserName", jarvisbase.JSON("result", result))

	return ResultUserWithUserName2User(ruser)
}

// SaveUserScript - save user script
func (db *ChatBotDB) SaveUserScript(userID string, userScript *pb.UserScript) error {
	if db.db == nil {
		return ErrChatBotDBNil
	}

	if userScript.File != nil && userScript.File.Data != nil {
		userScript.File.StrData = string(userScript.File.Data)
	}

	params := make(map[string]interface{})

	fd := *userScript.File
	fd.Data = nil
	err := ankadb.MakeParamsFromMsg(params, "file", &fd)
	if err != nil {
		return err
	}

	params["userID"] = userID
	params["scriptName"] = userScript.ScriptName
	params["jarvisNodeName"] = userScript.JarvisNodeName

	result, err := db.db.LocalQuery(context.Background(), querySaveUserScript, params)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.SaveUserScript:LocalQuery", zap.Error(err))

		return err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.SaveUserScript:GetResultError", zap.Error(err))

		return err
	}

	// jarvisbase.Debug("ChatBotDB.SaveUserScript",
	// 	jarvisbase.JSON("result", result))

	return nil
}

// GetUserScript - get user script
func (db *ChatBotDB) GetUserScript(userID string, scriptName string) (*pb.UserScript, error) {
	if db.db == nil {
		return nil, ErrChatBotDBNil
	}

	params := make(map[string]interface{})
	params["userID"] = userID
	params["scriptName"] = scriptName

	result, err := db.db.LocalQuery(context.Background(), queryGetUserScript, params)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUserScript:LocalQuery", zap.Error(err))

		return nil, err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUserScript:GetResultError", zap.Error(err))

		return nil, err
	}

	// jarvisbase.Debug("ChatBotDB.GetUserScript", jarvisbase.JSON("result", result))

	rus := &ResultUserScript{}
	err = ankadb.MakeObjFromResult(result, rus)
	if err != nil {
		return nil, err
	}

	userScript, err := ResultUserScript2UserScript(rus)
	if err != nil {
		return nil, err
	}

	return userScript, nil
}

// GetUsers - get users
func (db *ChatBotDB) GetUsers(nums int) (*pb.UserList, error) {
	if db.db == nil {
		return nil, ErrChatBotDBNil
	}

	params := make(map[string]interface{})
	params["snapshotID"] = int64(0)
	params["beginIndex"] = 0
	params["nums"] = nums

	result, err := db.db.LocalQuery(context.Background(), queryGetUsers, params)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUsers:LocalQuery", zap.Error(err))

		return nil, err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUsers:GetResultError", zap.Error(err))

		return nil, err
	}

	// s, err := json.Marshal(result)
	// if err != nil {
	// 	jarvisbase.Error("CoreDB.GetUsers", zap.Error(err))

	// 	return nil, err
	// }

	// jarvisbase.Debug("ChatBotDB.GetUsers:Marshal", jarvisbase.JSON("result", result))

	us := &ResultUsers{}
	err = ankadb.MakeObjFromResult(result, us)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUsers:MakeObjFromResult", zap.Error(err))

		return nil, err
	}

	lst, err := ResultUsers2UserList(us)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUsers:ResultUsers2UserList", zap.Error(err))

		return nil, err
	}

	return lst, nil

	// if db.db == nil {
	// 	return nil, ErrChatBotDBNil
	// }

	// params := make(map[string]interface{})
	// params["snapshotID"] = int64(0)
	// params["beginIndex"] = 0
	// params["nums"] = nums

	// jarvisbase.Debug("GetUsers",
	// 	zap.String("query string", queryUsers))

	// result, err := db.db.LocalQuery(context.Background(), queryUsers, params)
	// if err != nil {
	// 	jarvisbase.Warn("ChatBotDB.GetUsers:LocalQuery", zap.Error(err))

	// 	return nil, err
	// }

	// s, err := json.Marshal(result)
	// if err != nil {
	// 	jarvisbase.Error("CoreDB.GetNodes", zap.Error(err))

	// 	return nil, err
	// }

	// jarvisbase.Debug("GetUsers",
	// 	jarvisbase.JSON("resultstr", string(s)))

	// us := &ResultUsers{}
	// err = ankadb.MakeObjFromResult(result, us)
	// if err != nil {
	// 	return nil, err
	// }

	// jarvisbase.Debug("GetUsers",
	// 	jarvisbase.JSON("result", us))

	// lst, err := ResultUsers2UserList(us)
	// if err != nil {
	// 	return nil, err
	// }

	// return lst, nil
}

// GetUserScripts - get user scripts
func (db *ChatBotDB) GetUserScripts(userID string, jarvisNodeName string) (*pb.UserScriptList, error) {
	if db.db == nil {
		return nil, ErrChatBotDBNil
	}

	params := make(map[string]interface{})
	params["userID"] = userID
	params["jarvisNodeName"] = jarvisNodeName

	result, err := db.db.LocalQuery(context.Background(), queryGetUserScripts, params)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUserScripts:LocalQuery", zap.Error(err))

		return nil, err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUserScripts:GetResultError", zap.Error(err))

		return nil, err
	}

	// if result.HasErrors() {
	// 	var errstr string

	// 	for i, v := range result.Errors {
	// 		str := fmt.Sprintf("Error-%v: %v", (i + 1), v.Error())
	// 		if i > 0 {
	// 			errstr = errstr + " " + str
	// 		} else {
	// 			errstr = str
	// 		}
	// 	}

	// 	return nil, errors.New(errstr)
	// }

	// s, err := json.Marshal(result)
	// if err != nil {
	// 	jarvisbase.Error("CoreDB.GetUsers", zap.Error(err))

	// 	return nil, err
	// }

	// jarvisbase.Debug("ChatBotDB.GetUserScripts", jarvisbase.JSON("result", result))

	us := &ResultUserScripts{}
	err = ankadb.MakeObjFromResult(result, us)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUserScripts:MakeObjFromResult", zap.Error(err))

		return nil, err
	}

	lst, err := ResultUserScripts2UserScriptList(us)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUserScripts:ResultUserScripts2UserScriptList", zap.Error(err))

		return nil, err
	}

	return lst, nil

	// if db.db == nil {
	// 	return nil, ErrChatBotDBNil
	// }

	// params := make(map[string]interface{})
	// params["snapshotID"] = int64(0)
	// params["beginIndex"] = 0
	// params["nums"] = nums

	// jarvisbase.Debug("GetUsers",
	// 	zap.String("query string", queryUsers))

	// result, err := db.db.LocalQuery(context.Background(), queryUsers, params)
	// if err != nil {
	// 	jarvisbase.Warn("ChatBotDB.GetUsers:LocalQuery", zap.Error(err))

	// 	return nil, err
	// }

	// s, err := json.Marshal(result)
	// if err != nil {
	// 	jarvisbase.Error("CoreDB.GetNodes", zap.Error(err))

	// 	return nil, err
	// }

	// jarvisbase.Debug("GetUsers",
	// 	jarvisbase.JSON("resultstr", string(s)))

	// us := &ResultUsers{}
	// err = ankadb.MakeObjFromResult(result, us)
	// if err != nil {
	// 	return nil, err
	// }

	// jarvisbase.Debug("GetUsers",
	// 	jarvisbase.JSON("result", us))

	// lst, err := ResultUsers2UserList(us)
	// if err != nil {
	// 	return nil, err
	// }

	// return lst, nil
}

// RemoveUserScripts - remove user scripts
func (db *ChatBotDB) RemoveUserScripts(userID string, scriptName string) error {
	if db.db == nil {
		return ErrChatBotDBNil
	}

	params := make(map[string]interface{})
	params["userID"] = userID
	params["scriptName"] = scriptName

	result, err := db.db.LocalQuery(context.Background(), queryRmScript, params)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.RemoveUserScripts:LocalQuery", zap.Error(err))

		return err
	}

	err = ankadb.GetResultError(result)
	if err != nil {
		jarvisbase.Warn("ChatBotDB.GetUserScripts:GetResultError", zap.Error(err))

		return err
	}

	return nil

	// if result.HasErrors() {
	// 	var errstr string

	// 	for i, v := range result.Errors {
	// 		str := fmt.Sprintf("Error-%v: %v", (i + 1), v.Error())
	// 		if i > 0 {
	// 			errstr = errstr + " " + str
	// 		} else {
	// 			errstr = str
	// 		}
	// 	}

	// 	return nil, errors.New(errstr)
	// }

	// s, err := json.Marshal(result)
	// if err != nil {
	// 	jarvisbase.Error("CoreDB.GetUsers", zap.Error(err))

	// 	return nil, err
	// }

	// jarvisbase.Debug("ChatBotDB.GetUserScripts", jarvisbase.JSON("result", result))

	// us := &ResultUserScripts{}
	// err = ankadb.MakeObjFromResult(result, us)
	// if err != nil {
	// 	jarvisbase.Warn("ChatBotDB.GetUserScripts:MakeObjFromResult", zap.Error(err))

	// 	return nil, err
	// }

	// lst, err := ResultUserScripts2UserScriptList(us)
	// if err != nil {
	// 	jarvisbase.Warn("ChatBotDB.GetUserScripts:ResultUserScripts2UserScriptList", zap.Error(err))

	// 	return nil, err
	// }

	// return lst, nil

	// if db.db == nil {
	// 	return nil, ErrChatBotDBNil
	// }

	// params := make(map[string]interface{})
	// params["snapshotID"] = int64(0)
	// params["beginIndex"] = 0
	// params["nums"] = nums

	// jarvisbase.Debug("GetUsers",
	// 	zap.String("query string", queryUsers))

	// result, err := db.db.LocalQuery(context.Background(), queryUsers, params)
	// if err != nil {
	// 	jarvisbase.Warn("ChatBotDB.GetUsers:LocalQuery", zap.Error(err))

	// 	return nil, err
	// }

	// s, err := json.Marshal(result)
	// if err != nil {
	// 	jarvisbase.Error("CoreDB.GetNodes", zap.Error(err))

	// 	return nil, err
	// }

	// jarvisbase.Debug("GetUsers",
	// 	jarvisbase.JSON("resultstr", string(s)))

	// us := &ResultUsers{}
	// err = ankadb.MakeObjFromResult(result, us)
	// if err != nil {
	// 	return nil, err
	// }

	// jarvisbase.Debug("GetUsers",
	// 	jarvisbase.JSON("result", us))

	// lst, err := ResultUsers2UserList(us)
	// if err != nil {
	// 	return nil, err
	// }

	// return lst, nil
}
