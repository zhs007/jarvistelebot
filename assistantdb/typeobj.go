package assistantdb

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
)

// typeMessage - Message
//		you can see assistantdb.graphql
var typeMessage = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Message",
		Fields: graphql.Fields{
			"msgID": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"data": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"keys": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"createTime": &graphql.Field{
				Type: graphqlext.Int64,
			},
			"updateTime": &graphql.Field{
				Type: graphqlext.Int64,
			},
		},
	},
)

// typeAssistantData - AssistantData
//		you can see assistantdb.graphql
var typeAssistantData = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AssistantData",
		Fields: graphql.Fields{
			"maxMsgID": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"keys": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)
