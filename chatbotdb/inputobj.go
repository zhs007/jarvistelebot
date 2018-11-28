package chatbotdb

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
			"chatID": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"from": &graphql.InputObjectFieldConfig{
				Type: typeUser,
			},
			"to": &graphql.InputObjectFieldConfig{
				Type: typeUser,
			},
			"text": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"timeStamp": &graphql.InputObjectFieldConfig{
				Type: graphqlext.Int64,
			},
			"msgID": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"options": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
			},
			"selected": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"file": &graphql.InputObjectFieldConfig{
				Type: typeFile,
			},
		},
	},
)

// inputTypeFile - File
//		you can see chatbotdb.graphql
var inputTypeFile = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "FileInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"filename": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"strData": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"fileType": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

// inputTypeUser - User
//		you can see chatbotdb.graphql
var inputTypeUser = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"nickName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"userID": &graphql.InputObjectFieldConfig{
				Type: graphql.ID,
			},
			"userName": &graphql.InputObjectFieldConfig{
				Type: graphql.ID,
			},
			"lastMsgID": &graphql.InputObjectFieldConfig{
				Type: graphqlext.Int64,
			},
			"scripts": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
			},
			"fileTemplates": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)
