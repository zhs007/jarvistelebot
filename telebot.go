package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhs007/jarvistelebot/base"
	"go.uber.org/zap"
)

// startWithGetUpdate - start tele bot with getupdate mode
func startWithGetUpdate(token string, debugMode bool) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		base.Fatal("tgbotapi.NewBotAPI", zap.Error(base.ErrNewTeleBot))
	}

	bot.Debug = debugMode

	base.Info("Authorized on account " + bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		base.Info("got msg",
			zap.String("username", update.Message.From.UserName),
			zap.String("text", update.Message.Text))

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

// startTeleBot - start tele bot
func startTeleBot() {
	cfg := base.GetConfig()

	if cfg.TeleBotType == base.TBTGetUpdate {
		base.Info("tele bot type is " + base.TBTGetUpdate)

		startWithGetUpdate(cfg.TeleBotToken, cfg.DebugMode)
	}
}
