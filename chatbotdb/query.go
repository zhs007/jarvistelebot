package chatbotdb

import (
	"encoding/base64"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb"
	"github.com/zhs007/ankadb/graphqlext"
	"github.com/zhs007/ankadb/proto"
	pb "github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

var typeQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"msg": &graphql.Field{
				Type: typeMessage,
				Args: graphql.FieldConfigArgument{
					"chatID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					anka := ankadb.GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
					if anka == nil {
						return nil, ankadb.ErrCtxAnkaDB
					}

					curdb := anka.MgrDB.GetDB("chatbotdb")
					if curdb == nil {
						return nil, ankadb.ErrCtxCurDB
					}

					chatID := params.Args["chatID"].(string)

					msg := &pb.Message{}
					err := ankadb.GetMsgFromDB(curdb, []byte(makeMessageKey(chatID)), msg)
					if err != nil {
						return nil, err
					}

					if msg.File != nil && msg.File.Data != nil {
						msg.File.StrData = base64.StdEncoding.EncodeToString(msg.File.Data)
						// msg.File.Data = nil
					}

					return msg, nil
				},
			},
			"user": &graphql.Field{
				Type: typeUser,
				Args: graphql.FieldConfigArgument{
					"userID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					anka := ankadb.GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
					if anka == nil {
						return nil, ankadb.ErrCtxAnkaDB
					}

					curdb := anka.MgrDB.GetDB("chatbotdb")
					if curdb == nil {
						return nil, ankadb.ErrCtxCurDB
					}

					userID := params.Args["userID"].(string)

					user := &pb.User{}
					err := ankadb.GetMsgFromDB(curdb, []byte(makeUserKey(userID)), user)
					if err != nil {
						return nil, err
					}

					return user, nil
				},
			},
			"users": &graphql.Field{
				Type: typeUserList,
				Args: graphql.FieldConfigArgument{
					"snapshotID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphqlext.Int64),
					},
					"beginIndex": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"nums": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					jarvisbase.Debug("query users")

					anka := ankadb.GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
					if anka == nil {
						return nil, ankadb.ErrCtxAnkaDB
					}

					curdb := anka.MgrDB.GetDB("chatbotdb")
					if curdb == nil {
						return nil, ankadb.ErrCtxCurDB
					}

					mgrSnapshot := anka.MgrDB.GetMgrSnapshot("chatbotdb")
					if mgrSnapshot == nil {
						return nil, ankadb.ErrCtxSnapshotMgr
					}

					curit := curdb.NewIteratorWithPrefix([]byte(prefixKeyUser))
					jarvisbase.Debug("curdb.NewIteratorWithPrefix")
					for curit.Next() {
						key := curit.Key()
						jarvisbase.Debug("curdb.NewIteratorWithPrefix", zap.String("key", string(key)))
					}
					curit.Release()
					err := curit.Error()
					if err != nil {
						jarvisbase.Debug("curdb.NewIteratorWithPrefix", zap.Error(err))

						return nil, err
					}

					snapshotID := params.Args["snapshotID"].(int64)
					beginIndex := params.Args["beginIndex"].(int)
					nums := params.Args["nums"].(int)
					if beginIndex < 0 || nums <= 0 {
						return nil, ankadb.ErrQuertParams
					}

					lstUser := &pb.UserList{}
					var pSnapshot *ankadbpb.Snapshot

					if snapshotID > 0 {
						pSnapshot = mgrSnapshot.Get(snapshotID)
					} else {
						var err error
						pSnapshot, err = mgrSnapshot.NewSnapshot([]byte(prefixKeyUser))
						if err != nil {
							return nil, ankadb.ErrCtxSnapshotMgr
						}
					}

					lstUser.SnapshotID = pSnapshot.SnapshotID
					lstUser.MaxIndex = int32(len(pSnapshot.Keys))

					curi := beginIndex
					for ; curi < len(pSnapshot.Keys) && len(lstUser.Users) < nums; curi++ {
						cui := &pb.User{}
						err := ankadb.GetMsgFromDB(curdb, []byte(pSnapshot.Keys[curi]), cui)
						if err == nil {
							lstUser.Users = append(lstUser.Users, cui)
						}
					}

					lstUser.EndIndex = int32(curi)

					return lstUser, nil
				},
			},
			"userWithUserName": &graphql.Field{
				Type: typeUser,
				Args: graphql.FieldConfigArgument{
					"userName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					anka := ankadb.GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
					if anka == nil {
						return nil, ankadb.ErrCtxAnkaDB
					}

					curdb := anka.MgrDB.GetDB("chatbotdb")
					if curdb == nil {
						return nil, ankadb.ErrCtxCurDB
					}

					userName := params.Args["userName"].(string)
					uid, err := curdb.Get([]byte(makeUserNameKey(userName)))
					if err != nil {
						return nil, err
					}

					user := &pb.User{}
					err = ankadb.GetMsgFromDB(curdb, []byte(makeUserKey(string(uid))), user)
					if err != nil {
						return nil, err
					}

					return user, nil
				},
			},
			"userScript": &graphql.Field{
				Type: typeUserScript,
				Args: graphql.FieldConfigArgument{
					"userID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
					"scriptName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					anka := ankadb.GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
					if anka == nil {
						return nil, ankadb.ErrCtxAnkaDB
					}

					curdb := anka.MgrDB.GetDB("chatbotdb")
					if curdb == nil {
						return nil, ankadb.ErrCtxCurDB
					}

					userID := params.Args["userID"].(string)
					scriptName := params.Args["scriptName"].(string)

					userScript := &pb.UserScript{}
					err := ankadb.GetMsgFromDB(curdb, []byte(makeUserScriptKey(userID, scriptName)), userScript)
					if err != nil {
						return nil, err
					}

					if userScript.File != nil && userScript.File.Data != nil {
						userScript.File.StrData = string(userScript.File.Data)
					}

					return userScript, nil
				},
			},
		},
	},
)
