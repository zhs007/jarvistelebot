package main

import (
	"github.com/zhs007/jarvistelebot/telebot"
)

func main() {
	err := telebot.InitTeleBot("./cfg/config.yaml")
	if err != nil {
		telebot.Error("InitTeleBot err.")
	}

	defer telebot.ReleaseTeleBot()

	telebot.StartTeleBot()
}
