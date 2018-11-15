package telebot

import (
	"context"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarviscore/base"
	"github.com/zhs007/jarviscore/proto"

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

	scriptUser chatbot.User
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

	tcb.Init(path.Join(cfg.CfgPath, "chatbot.yaml"), mgrPlugins)

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

func (cb *teleChatBot) procDocument(ctx context.Context, node jarviscore.JarvisNode, doc *tgbotapi.Document) error {
	file, err := cb.teleBotAPI.GetFile(tgbotapi.FileConfig{
		FileID: doc.FileID,
	})
	if err != nil {
		return err
	}

	localfn := path.Join(cb.GetConfig().DownloadPath, doc.FileName)
	if doc.MimeType == "text/x-script.sh" {
		localfn = path.Join(cb.GetConfig().DownloadPath, "scripts", doc.FileName)
	}

	jarvisbase.Info("teleChatBot.procDocument",
		jarvisbase.JSON("file", file))

	url := file.Link(cb.teleBotAPI.Token)

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	f, err := os.Create(localfn)
	if err != nil {
		return err
	}
	io.Copy(f, res.Body)

	// dat, err := ioutil.ReadFile(localfn)
	// if err != nil {
	// 	jarvisbase.Warn("load script file", zap.Error(err))

	// 	return err
	// }

	// ci, err := jarviscore.BuildCtrlInfoForScriptFile(1, doc.FileName, dat, "")

	// cb.Node.SendCtrl(ctx, "1NutSP6ypvLtHpqHaxtjJMmEUbMfLUdp9a", ci)

	return nil
}

// Start
func (cb *teleChatBot) Start(ctx context.Context, node jarviscore.JarvisNode) error {
	cb.BaseChatBot.Start(ctx, node)

	node.RegMsgEventFunc(jarviscore.EventOnCtrlResult,
		func(curctx context.Context, jarvisnode jarviscore.JarvisNode, msg *jarviscorepb.JarvisMsg) error {
			return cb.OnJarvisCtrlResult(curctx, msg)
		})

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

		if update.Message.Document != nil {
			err = cb.procDocument(ctx, node, update.Message.Document)
			if err != nil {
				chatbot.Warn("teleChatBot.Start:procDocument", zap.Error(err))
			}

			cb.scriptUser = user

			continue
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

// OnJarvisCtrlResult - event handle
func (cb *teleChatBot) OnJarvisCtrlResult(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
	cr := msg.GetCtrlResult()
	cb.SendMsg(cb.scriptUser, cr.CtrlResult)

	return nil
}
