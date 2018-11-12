package assistantdb

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb"
	pb "github.com/zhs007/jarvistelebot/assistantdb/proto"
)

var typeMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"updMsg": &graphql.Field{
			Type:        typeMessage,
			Description: "update message",
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

				curdb := anka.MgrDB.GetDB("assistantdb")
				if curdb == nil {
					return nil, ankadb.ErrCtxCurDB
				}

				msg := &pb.Message{}
				err := ankadb.GetMsgFromParam(params, "msg", msg)
				if err != nil {
					return nil, err
				}

				err = ankadb.PutMsg2DB(curdb, []byte(makeMessageKey(msg.MsgID)), msg)
				if err != nil {
					return nil, err
				}

				return msg, nil
			},
		},
		"updAssistantData": &graphql.Field{
			Type:        typeMessage,
			Description: "update AssistantData",
			Args: graphql.FieldConfigArgument{
				"dat": &graphql.ArgumentConfig{
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

				dat := &pb.AssistantData{}
				err := ankadb.GetMsgFromParam(params, "dat", dat)
				if err != nil {
					return nil, err
				}

				err = ankadb.PutMsg2DB(curdb, []byte(makeAssistantDataKey()), dat)
				if err != nil {
					return nil, err
				}

				return dat, nil
			},
		},
	},
})
