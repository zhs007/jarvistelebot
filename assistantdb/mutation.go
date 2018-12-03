package assistantdb

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb"
	"github.com/zhs007/ankadb/graphqlext"
	pb "github.com/zhs007/jarvistelebot/assistantdb/proto"
)

var typeMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"updNote": &graphql.Field{
			Type:        typeNote,
			Description: "update note",
			Args: graphql.FieldConfigArgument{
				"userID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphqlext.Int64),
				},
				"note": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputTypeNote),
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

				note := &pb.Note{}
				err := ankadb.GetMsgFromParam(params, "note", note)
				if err != nil {
					return nil, err
				}

				err = ankadb.PutMsg2DB(curdb, []byte(makeNoteKey(userID, note.NoteID)), note)
				if err != nil {
					return nil, err
				}

				return note, nil
			},
		},
		"updUserAssistantInfo": &graphql.Field{
			Type:        typeUserAssistantInfo,
			Description: "update UserAssistantInfo",
			Args: graphql.FieldConfigArgument{
				"userID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphqlext.Int64),
				},
				"uai": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputTypeAssistantData),
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

				uai := &pb.UserAssistantInfo{}
				err := ankadb.GetMsgFromParam(params, "uai", uai)
				if err != nil {
					return nil, err
				}

				err = ankadb.PutMsg2DB(curdb, []byte(makeUserAssistantInfoKey(userID)), uai)
				if err != nil {
					return nil, err
				}

				return uai, nil
			},
		},
	},
})
