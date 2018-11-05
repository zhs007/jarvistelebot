package chatbotdb

import (
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
				"chatID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"fromNickName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"fromUserID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"text": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"timeStamp": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphqlext.Timestamp),
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
				fromNickName := params.Args["fromNickName"].(string)
				fromUserID := params.Args["fromUserID"].(string)
				text := params.Args["text"].(string)
				timeStamp := params.Args["timeStamp"].(int64)

				msg := &pb.Message{
					ChatID: chatID,
					From: &pb.User{
						NickName: fromNickName,
						UserID:   fromUserID,
					},
					Text:      text,
					TimeStamp: timeStamp,
				}

				err := ankadb.PutMsg2DB(curdb, []byte(makeMessageKey(chatID)), msg)
				if err != nil {
					return nil, err
				}

				return msg, nil
			},
		},
	},
})
