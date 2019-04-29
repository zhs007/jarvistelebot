package main

import (
	"context"
	_ "net/http/pprof"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/jarvisnode"
	"github.com/zhs007/jarvistelebot/telebot"
	"go.uber.org/zap"
)

func main() {
	cfg, err := telebot.InitTeleBot("./cfg/config.yaml")
	if err != nil {
		chatbot.Error("InitTeleBot err.", zap.Error(err))

		return
	}

	defer telebot.ReleaseTeleBot()

	myni, err := jarvisnode.Init("./cfg/jarvisnode.yaml")
	if err != nil {
		chatbot.Error("jarvisnode.Init err.", zap.Error(err))

		return
	}

	// pprof
	jarviscore.InitPprof(myni)

	node, err := jarvisnode.NewNode(myni)
	if err != nil {
		chatbot.Error("jarvisnode.NewNode err.", zap.Error(err))

		return
	}

	go node.Start(context.Background())

	telebot.StartTeleBot(context.Background(), cfg, node)
}
