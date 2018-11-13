package telebot

import (
	"context"
	"strconv"

	"github.com/zhs007/jarviscore/base"

	"github.com/zhs007/jarviscore"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/assistant"
	"github.com/zhs007/jarvistelebot/plugins/jarvisnode"
	"github.com/zhs007/jarvistelebot/plugins/normal"
	"github.com/zhs007/jarvistelebot/plugins/timestamp"
	"go.uber.org/zap"
)

// teleChatBot - tele chat bot
type teleChatBot struct {
	chatbot.BaseChatBot

	teleBotAPI *tgbotapi.BotAPI
	mgrUser    chatbot.UserMgr
	// mgrPlugins chatbot.PluginsMgr
}

func regPlugins(mgrPlugins chatbot.PluginsMgr) {
	err := pluginassistant.RegPlugin(cfg.CfgPath, mgrPlugins)
	if err != nil {
		jarvisbase.Warn("telbot.regPlugins:pluginassistant.RegPlugin", zap.Error(err))
	}

	err = pluginjarvisnode.RegPlugin(cfg.CfgPath, mgrPlugins)
	if err != nil {
		jarvisbase.Warn("telbot.regPlugins:pluginjarvisnode.RegPlugin", zap.Error(err))
	}

	err = plugintimestamp.RegPlugin(cfg.CfgPath, mgrPlugins)
	if err != nil {
		jarvisbase.Warn("telbot.regPlugins:plugintimestamp.RegPlugin", zap.Error(err))
	}

	err = pluginnormal.RegPlugin(cfg.CfgPath, mgrPlugins)
	if err != nil {
		jarvisbase.Warn("telbot.regPlugins:pluginnormal.RegPlugin", zap.Error(err))
	}
}

// NewTeleChatBot - new tele chat bot
func NewTeleChatBot(token string, debugMode bool) (chatbot.ChatBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		chatbot.Fatal("tgbotapi.NewBotAPI", zap.Error(err))
	}

	bot.Debug = debugMode

	chatbot.Info("Authorized on account " + bot.Self.UserName)

	mgrPlugins := chatbot.NewPluginsMgr()

	regPlugins(mgrPlugins)

	tcb := &teleChatBot{
		teleBotAPI: bot,
		mgrUser:    chatbot.NewUserMgr(),
		// MgrPlugins: mgrPlugins,
	}

	tcb.Init(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine, mgrPlugins)

	return tcb, nil
}

// SendMsg -
func (cb *teleChatBot) SendMsg(user chatbot.User, text string) error {
	u, ok := (user).(*teleUser)
	if !ok {
		return ErrInvalidUser
	}

	telemsg := tgbotapi.NewMessage(u.chatid, text)
	cb.teleBotAPI.Send(telemsg)

	return nil
}

// Start
func (cb *teleChatBot) Start(ctx context.Context, node jarviscore.JarvisNode) error {
	cb.BaseChatBot.Start(ctx, node)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := cb.teleBotAPI.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		user := cb.mgrUser.GetUser(update.Message.From.UserName)
		if user == nil {
			user = newUser(update.Message.From.UserName, int64(update.Message.From.ID),
				update.Message.From.FirstName+" "+update.Message.From.LastName)

			cb.mgrUser.AddUser(user)
		}

		if update.Message.Text == "" {
			continue
		}

		msg := newMsg(user.GetUserID()+":"+strconv.Itoa(update.Message.MessageID),
			user, update.Message.Text, update.Message.Date)

		err := cb.SaveMsg(msg)
		if err != nil {
			chatbot.Warn("teleChatBot.Start", zap.Error(err))
		}

		curctx, cancel := context.WithCancel(ctx)
		defer cancel()

		err = cb.MgrPlugins.OnMessage(curctx, cb, msg)
		if err != nil {
			chatbot.Error("mgrPlugins.OnMessage", zap.Error(err))
		}

		// if IsMaster(update.Message.From.UserName) {
		// 	chatbot.Info("got master msg",
		// 		zap.String("username", update.Message.From.UserName),
		// 		zap.String("text", update.Message.Text))
		// } else {
		// 	chatbot.Info("got msg",
		// 		zap.String("username", update.Message.From.UserName),
		// 		zap.String("text", update.Message.Text))
		// }

		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// msg.ReplyToMessageID = update.Message.MessageID

		// cb.SendMsg()
		// bot.Send(msg)
	}

	return nil
}

// // GetPluginsMgr - get PluginsMgr
// func (cb *teleChatBot) GetPluginsMgr() chatbot.PluginsMgr {
// 	return cb.mgrPlugins
// }
