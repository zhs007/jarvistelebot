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
			"note": &graphql.Field{
				Type: typeNote,
				Args: graphql.FieldConfigArgument{
					"userID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
					"noteID": &graphql.ArgumentConfig{
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

					userID := params.Args["userID"].(string)
					noteID := params.Args["noteID"].(int64)

					note := &pb.Note{}
					err := ankadb.GetMsgFromDB(curdb, []byte(makeNoteKey(userID, noteID)), note)
					if err != nil {
						return nil, err
					}

					return note, nil
				},
			},
			"userAssistantInfo": &graphql.Field{
				Type: typeUserAssistantInfo,
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

					curdb := anka.MgrDB.GetDB("assistantdb")
					if curdb == nil {
						return nil, ankadb.ErrCtxCurDB
					}

					userID := params.Args["userID"].(string)

					haskey, err := curdb.Has([]byte(makeUserAssistantInfoKey(userID)))
					if err != nil {
						return nil, err
					}

					if !haskey {
						uai := &pb.UserAssistantInfo{
							MaxNoteID: 0,
						}

						return uai, nil
					}

					uai := &pb.UserAssistantInfo{}
					err = ankadb.GetMsgFromDB(curdb, []byte(makeUserAssistantInfoKey(userID)), uai)
					if err != nil {
						return nil, err
					}

					return uai, nil
				},
			},
		},
	},
)
