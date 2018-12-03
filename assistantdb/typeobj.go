package assistantdb

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
)

// typeNote - Note
//		you can see assistantdb.graphql
var typeNote = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Note",
		Fields: graphql.Fields{
			"noteID": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"data": &graphql.Field{
				Type: graphql.NewList(graphql.String),
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

// typeUserAssistantInfo - UserAssistantInfo
//		you can see assistantdb.graphql
var typeUserAssistantInfo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserAssistantInfo",
		Fields: graphql.Fields{
			"maxNoteID": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"keys": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

// typeKeyInfo - KeyInfo
//		you can see assistantdb.graphql
var typeKeyInfo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "KeyInfo",
		Fields: graphql.Fields{
			"noteIDs": &graphql.Field{
				Type: graphql.NewList(graphqlext.Int64),
			},
		},
	},
)
