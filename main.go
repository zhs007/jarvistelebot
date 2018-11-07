package main

import (
	"context"

	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/jarvisnode"
	"github.com/zhs007/jarvistelebot/telebot"
	"go.uber.org/zap"
)

func main() {
	err := telebot.InitTeleBot("./cfg/config.yaml")
	if err != nil {
		chatbot.Error("InitTeleBot err.", zap.Error(err))
	}

	defer telebot.ReleaseTeleBot()

	myni, err := jarvisnode.Init("./cfg/jarvisnode.yaml")
	if err != nil {
		chatbot.Error("jarvisnode.Init err.", zap.Error(err))
	}

	node := jarvisnode.NewNode(myni)
	go node.Start(context.Background())

	telebot.StartTeleBot(context.Background(), node)
}
