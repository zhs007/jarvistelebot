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
					Type: graphql.NewNonNull(graphql.ID),
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
		"rmNote": &graphql.Field{
			Type:        graphql.String,
			Description: "remove note",
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

				key := makeNoteKey(userID, noteID)

				err := curdb.Delete([]byte(key))
				if err != nil {
					return nil, err
				}

				return key, nil
			},
		},
		"updUserAssistantInfo": &graphql.Field{
			Type:        typeUserAssistantInfo,
			Description: "update UserAssistantInfo",
			Args: graphql.FieldConfigArgument{
				"userID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"uai": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputTypeUserAssistantInfo),
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
		"updKeyInfo": &graphql.Field{
			Type:        typeKeyInfo,
			Description: "update keyinfo",
			Args: graphql.FieldConfigArgument{
				"userID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"key": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"keyinfo": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputTypeKeyInfo),
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

				keyinfo := &pb.KeyInfo{}
				err := ankadb.GetMsgFromParam(params, "keyinfo", keyinfo)
				if err != nil {
					return nil, err
				}

				err = ankadb.PutMsg2DB(curdb, []byte(makeKeyInfoKey(userID, key)), keyinfo)
				if err != nil {
					return nil, err
				}

				return keyinfo, nil
			},
		},
	},
})
