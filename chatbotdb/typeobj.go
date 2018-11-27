package chatbotdb

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
)

// typeUser - User
//		you can see chatbotdb.graphql
var typeUser = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"nickName": &graphql.Field{
				Type: graphql.String,
			},
			"userID": &graphql.Field{
				Type: graphql.ID,
			},
			"userName": &graphql.Field{
				Type: graphql.ID,
			},
			"lastMsgID": &graphql.Field{
				Type: graphqlext.Int64,
			},
		},
	},
)

// typeFile - File
//		you can see chatbotdb.graphql
var typeFile = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "File",
		Fields: graphql.Fields{
			"filename": &graphql.Field{
				Type: graphql.String,
			},
			"strData": &graphql.Field{
				Type: graphql.String,
			},
			"fileType": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// typeMessage - Message
//		you can see chatbotdb.graphql
var typeMessage = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Message",
		Fields: graphql.Fields{
			"chatID": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"from": &graphql.Field{
				Type: typeUser,
			},
			"to": &graphql.Field{
				Type: typeUser,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"timeStamp": &graphql.Field{
				Type: graphqlext.Int64,
			},
			"msgID": &graphql.Field{
				Type: graphql.String,
			},
			"options": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"selected": &graphql.Field{
				Type: graphql.Int,
			},
			"file": &graphql.Field{
				Type: typeFile,
			},
		},
	},
)

// typeUserScript - UserScript
//		you can see chatbotdb.graphql
var typeUserScript = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserScript",
		Fields: graphql.Fields{
			"scriptName": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"file": &graphql.Field{
				Type: typeFile,
			},
		},
	},
)
