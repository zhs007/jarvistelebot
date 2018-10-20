package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhs007/jarvistelebot/base"
)

func main() {
	base.LoadConfig("./cfg/config.yaml")
	base.InitLogger()
	defer base.SyncLogger()

	tgbotapi.SetLogger(&telebotlog{})

	base.Info("tele bot start")
	startTeleBot()
}
