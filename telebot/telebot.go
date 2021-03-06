package telebot

import (
	"context"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarvistelebot/chatbot"
	"go.uber.org/zap"
)

// startWithGetUpdate - start tele bot with getupdate mode
func startWithGetUpdate(ctx context.Context, cfg *Config, node jarviscore.JarvisNode) {
	bot, err := NewTeleChatBot(cfg)
	if err != nil {
		chatbot.Fatal("NewTeleChatBot", zap.Error(err))
	}

	// bot.Init(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)

	bot.Start(ctx, node)

	// bot, err := tgbotapi.NewBotAPI(token)
	// if err != nil {
	// 	chatbot.Fatal("tgbotapi.NewBotAPI", zap.Error(err))
	// }

	// bot.Debug = debugMode

	// chatbot.Info("Authorized on account " + bot.Self.UserName)

	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60

	// updates, err := bot.GetUpdatesChan(u)

	// for update := range updates {
	// 	if update.Message == nil { // ignore any non-Message Updates
	// 		continue
	// 	}

	// 	if IsMaster(update.Message.From.UserName) {
	// 		chatbot.Info("got master msg",
	// 			zap.String("username", update.Message.From.UserName),
	// 			zap.String("text", update.Message.Text))
	// 	} else {
	// 		chatbot.Info("got msg",
	// 			zap.String("username", update.Message.From.UserName),
	// 			zap.String("text", update.Message.Text))
	// 	}

	// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// 	msg.ReplyToMessageID = update.Message.MessageID

	// 	bot.Send(msg)
	// }
}

// StartTeleBot - start tele bot
func StartTeleBot(ctx context.Context, cfg *Config, node jarviscore.JarvisNode) {
	// cfg := GetConfig()

	if cfg.TeleBotType == TBTGetUpdate {
		chatbot.Info("tele bot type is " + TBTGetUpdate)

		startWithGetUpdate(ctx, cfg, node)
	}
}

// InitTeleBot - init telebot
func InitTeleBot(filename string) (*Config, error) {
	cfg, err := LoadConfig(filename)
	if err != nil {
		return nil, err
	}

	chatbot.InitLogger(cfg.lvl, cfg.LogPath == LOGPATHConsole, cfg.LogPath)

	tgbotapi.SetLogger(&telebotlog{})

	chatbot.Info("tele bot start")

	return cfg, nil
}

// ReleaseTeleBot - release telebot
func ReleaseTeleBot() {
	chatbot.SyncLogger()
}

// // IsMaster - is master?
// func IsMaster(uname string) bool {
// 	return cfg.TeleBotMaster == uname
// }
