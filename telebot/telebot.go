package telebot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhs007/jarvistelebot/chatbot"
	"go.uber.org/zap"
)

// startWithGetUpdate - start tele bot with getupdate mode
func startWithGetUpdate(token string, debugMode bool) {
	bot, err := NewTeleChatBot(token, debugMode)
	if err != nil {
		chatbot.Fatal("NewTeleChatBot", zap.Error(err))
	}

	bot.Start()

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
func StartTeleBot() {
	cfg := GetConfig()

	if cfg.TeleBotType == TBTGetUpdate {
		chatbot.Info("tele bot type is " + TBTGetUpdate)

		startWithGetUpdate(cfg.TeleBotToken, cfg.DebugMode)
	}
}

// InitTeleBot - init telebot
func InitTeleBot(filename string) error {
	err := LoadConfig(filename)
	if err != nil {
		return err
	}

	chatbot.InitLogger(cfg.lvl, cfg.LogPath == LOGPATHConsole, cfg.LogPath)

	tgbotapi.SetLogger(&telebotlog{})

	chatbot.Info("tele bot start")

	return nil
}

// ReleaseTeleBot - release telebot
func ReleaseTeleBot() {
	chatbot.SyncLogger()
}

// // IsMaster - is master?
// func IsMaster(uname string) bool {
// 	return cfg.TeleBotMaster == uname
// }
