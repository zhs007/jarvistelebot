package chatbotdb

import (
	"encoding/base64"

	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb"
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
