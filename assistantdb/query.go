package assistantdb

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb"
	"github.com/zhs007/ankadb/graphqlext"
	pb "github.com/zhs007/jarvistelebot/assistantdb/proto"
)

var typeQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"msg": &graphql.Field{
				Type: typeMessage,
				Args: graphql.FieldConfigArgument{
					"msgID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphqlext.Int64),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					anka := ankadb.GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
					if anka == nil {
						return nil, ankadb.ErrCtxAnkaDB
					}

					curdb := anka.MgrDB.GetDB("assistantdb")
					if curdb == nil {
						return nil, ankadb.ErrCtxCurDB
					}

					msgID := params.Args["msgID"].(int64)

					msg := &pb.Message{}
					err := ankadb.GetMsgFromDB(curdb, []byte(makeMessageKey(msgID)), msg)
					if err != nil {
						return nil, err
					}

					return msg, nil
				},
			},
		},
	},
)
