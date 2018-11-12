package assistantdb

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
)

// inputTypeMessage - Message
//		you can see coredb.graphql
var inputTypeMessage = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "MessageInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"msgID": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"data": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"keys": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
			},
			"createTime": &graphql.InputObjectFieldConfig{
				Type: graphqlext.Int64,
			},
			"updateTime": &graphql.InputObjectFieldConfig{
				Type: graphqlext.Int64,
			},
		},
	},
)

// inputTypeAssistantData - AssistantData
//		you can see coredb.graphql
var inputTypeAssistantData = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "AssistantDataInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"maxMsgID": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"keys": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)
