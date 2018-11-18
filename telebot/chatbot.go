package telebot

import (
	"context"
	"io"
	"io/ioutil"
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
	*chatbot.BasicChatBot

	teleBotAPI *tgbotapi.BotAPI
	cfg        *Config
	// mgrUser    chatbot.UserMgr

	scriptUser chatbot.User
	// mgrPlugins chatbot.PluginsMgr
}

func regPlugins(cfg *Config, mgrPlugins chatbot.PluginsMgr) {
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
func NewTeleChatBot(cfg *Config) (chatbot.ChatBot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.TeleBotToken)
	if err != nil {
		chatbot.Fatal("tgbotapi.NewBotAPI", zap.Error(err))
	}

	bot.Debug = cfg.DebugMode

	chatbot.Info("Authorized on account " + bot.Self.UserName)

	mgrPlugins := chatbot.NewPluginsMgr()

	regPlugins(cfg, mgrPlugins)

	tcb := &teleChatBot{
		BasicChatBot: chatbot.NewBasicChatBot(),
		teleBotAPI:   bot,
		// mgrUser:    chatbot.NewUserMgr(),
		// MgrPlugins: mgrPlugins,
	}

	tcb.MgrUser = newTeleUserMgr(cfg.TeleBotMaster)

	tcb.Init(path.Join(cfg.CfgPath, "chatbot.yaml"), mgrPlugins)

	return tcb, nil
}

// SendMsg -
func (cb *teleChatBot) SendMsg(user chatbot.User, text string) error {
	// u, ok := (user).(*teleUser)
	// if !ok {
	// 	return ErrInvalidUser
	// }

	chatid, err := strconv.ParseInt(user.GetUserID(), 10, 64)
	if err != nil {
		return err
	}

	telemsg := tgbotapi.NewMessage(chatid, text)

	var arr []tgbotapi.InlineKeyboardButton

	arr = append(arr, tgbotapi.NewInlineKeyboardButtonData("yes", "yes, i am."))
	arr = append(arr, tgbotapi.NewInlineKeyboardButtonData("no", "no, i am not"))

	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(arr)

	telemsg.ReplyMarkup = numericKeyboard

	cb.teleBotAPI.Send(telemsg)

	return nil
}

func (cb *teleChatBot) procDocument(ctx context.Context, node jarviscore.JarvisNode, doc *tgbotapi.Document) error {
	jarvisbase.Debug("teleChatBot.procDocument")

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

	dat, err := ioutil.ReadFile(localfn)
	if err != nil {
		jarvisbase.Warn("load script file", zap.Error(err))

		return err
	}

	ci, err := jarviscore.BuildCtrlInfoForScriptFile(1, doc.FileName, dat, "")

	cb.Node.RequestCtrl(ctx, "1NutSP6ypvLtHpqHaxtjJMmEUbMfLUdp9a", ci)

	return nil
}

// Start
func (cb *teleChatBot) Start(ctx context.Context, node jarviscore.JarvisNode) error {
	cb.BasicChatBot.Start(ctx, node)

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

		user := cb.MgrUser.GetUser(update.Message.From.UserName)
		if user == nil {
			userid := strconv.Itoa(update.Message.From.ID)

			user = chatbot.NewBasicUser(update.Message.From.UserName, userid,
				update.Message.From.FirstName+" "+update.Message.From.LastName, int64(update.Message.MessageID))

			cb.MgrUser.AddUser(user)
			cb.GetChatBotDB().UpdUser(user)
		} else {
			lastmsgid := int64(update.Message.MessageID)

			if lastmsgid <= user.GetLastMsgID() {
				continue
			}
			user.UpdLastMsgID(lastmsgid)
			cb.GetChatBotDB().UpdUser(user)
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
