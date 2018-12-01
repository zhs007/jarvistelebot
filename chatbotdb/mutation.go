package chatbotdb

import (
	"encoding/base64"

	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb"
	pb "github.com/zhs007/jarvistelebot/chatbotdb/proto"
)

var typeMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"newMsg": &graphql.Field{
			Type:        typeMessage,
			Description: "new message",
			Args: graphql.FieldConfigArgument{
				"msg": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputTypeMessage),
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

				msg := &pb.Message{}
				err := ankadb.GetMsgFromParam(params, "msg", msg)
				if err != nil {
					return nil, err
				}

				if msg.File != nil && msg.File.StrData != "" {
					data, err := base64.StdEncoding.DecodeString(msg.File.StrData)
					if err != nil {
						return nil, err
					}

					msg.File.Data = data
					msg.File.StrData = ""
				}

				err = ankadb.PutMsg2DB(curdb, []byte(makeMessageKey(msg.ChatID)), msg)
				if err != nil {
					return nil, err
				}

				return msg, nil
			},
		},
		"updUser": &graphql.Field{
			Type:        typeUser,
			Description: "update user",
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputTypeUser),
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

				user := &pb.User{}
				err := ankadb.GetMsgFromParam(params, "user", user)
				if err != nil {
					return nil, err
				}

				// nickName := params.Args["nickName"].(string)
				// userID := params.Args["userID"].(string)
				// userName := params.Args["userName"].(string)
				// lastMsgID := params.Args["lastMsgID"].(int64)

				// user := &pb.User{
				// 	NickName:  nickName,
				// 	UserID:    userID,
				// 	UserName:  userName,
				// 	LastMsgID: lastMsgID,
				// }

				err = ankadb.PutMsg2DB(curdb, []byte(makeUserKey(user.UserID)), user)
				if err != nil {
					return nil, err
				}

				err = curdb.Put([]byte(makeUserNameKey(user.UserName)), []byte(user.UserID))
				if err != nil {
					return nil, err
				}

				return user, nil
			},
		},
		"updUserScript": &graphql.Field{
			Type:        typeUserScript,
			Description: "update user script",
			Args: graphql.FieldConfigArgument{
				"userID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"jarvisNodeName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"scriptName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"file": &graphql.ArgumentConfig{
					Type: inputTypeFile,
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
				jarvisNodeName := params.Args["jarvisNodeName"].(string)

				file := &pb.File{}
				err := ankadb.GetMsgFromParam(params, "file", file)
				if err != nil {
					return nil, err
				}

				if file.StrData != "" {
					file.Data = []byte(file.StrData)
					file.StrData = ""
				}

				userScript := &pb.UserScript{
					ScriptName:     scriptName,
					File:           file,
					JarvisNodeName: jarvisNodeName,
				}

				err = ankadb.PutMsg2DB(curdb, []byte(makeUserScriptKey(userID, scriptName)), userScript)
				if err != nil {
					return nil, err
				}

				// jarvisbase.Debug("updUserScript",
				// 	zap.String("key", makeUserScriptKey(userID, scriptName)),
				// 	jarvisbase.JSON("userScript", userScript))

				return userScript, nil
			},
		},
		"rmUserScript": &graphql.Field{
			Type:        graphql.String,
			Description: "remove user script",
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

				key := makeUserScriptKey(userID, scriptName)

				err := curdb.Delete([]byte(key))
				if err != nil {
					return nil, err
				}

				// jarvisbase.Debug("updUserScript",
				// 	zap.String("key", makeUserScriptKey(userID, scriptName)),
				// 	jarvisbase.JSON("userScript", userScript))

				return key, nil
			},
		},
	},
})
