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

					curkey := []byte(makeNoteKey(userID, noteID))

					note := &pb.Note{}

					has, err := curdb.Has(curkey)
					if err != nil {
						return nil, err
					}

					if has {
						err := ankadb.GetMsgFromDB(curdb, curkey, note)
						if err != nil {
							return nil, err
						}
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
			"keyInfo": &graphql.Field{
				Type: typeKeyInfo,
				Args: graphql.FieldConfigArgument{
					"userID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
					"key": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
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
					key := params.Args["key"].(string)

					dbkey := []byte(makeKeyInfoKey(userID, key))
					haskey, err := curdb.Has(dbkey)
					if err != nil {
						return nil, err
					}

					if !haskey {
						ki := &pb.KeyInfo{}

						return ki, nil
					}

					ki := &pb.KeyInfo{}
					err = ankadb.GetMsgFromDB(curdb, []byte(dbkey), ki)
					if err != nil {
						return nil, err
					}

					return ki, nil
				},
			},
		},
	},
)
