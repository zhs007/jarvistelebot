package chatbotdb

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
)

// typeUser - User
//		you can see charbotdb.graphql
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

// typeMessage - Message
//		you can see charbotdb.graphql
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
		},
	},
)
