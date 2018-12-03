package assistantdb

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
)

// inputTypeNote - Note
//		you can see assistantdb.graphql
var inputTypeNote = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "NoteInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"noteID": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"data": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
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

// inputTypeUserAssistantInfo - UserAssistantInfoInput
//		you can see coredb.graphql
var inputTypeUserAssistantInfo = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UserAssistantInfoInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"maxNoteID": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"keys": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

// inputTypeKeyInfo - KeyInfoInput
//		you can see assistantdb.graphql
var inputTypeKeyInfo = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "KeyInfoInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"noteIDs": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphqlext.Int64),
			},
		},
	},
)
