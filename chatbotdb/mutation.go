package chatbotdb

import (
	"encoding/base64"

	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb"
	"github.com/zhs007/ankadb/graphqlext"
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
				"nickName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"userID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"userName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"lastMsgID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphqlext.Int64),
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

				nickName := params.Args["nickName"].(string)
				userID := params.Args["userID"].(string)
				userName := params.Args["userName"].(string)
				lastMsgID := params.Args["lastMsgID"].(int64)

				user := &pb.User{
					NickName:  nickName,
					UserID:    userID,
					UserName:  userName,
					LastMsgID: lastMsgID,
				}

				err := ankadb.PutMsg2DB(curdb, []byte(makeUserKey(userID)), user)
				if err != nil {
					return nil, err
				}

				return user, nil
			},
		},
	},
})
