package telebot

import (
	"context"
	"fmt"
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
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
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
// func (cb *teleChatBot) SendMsg(user chatbot.User, text string) error {
// 	// u, ok := (user).(*teleUser)
// 	// if !ok {
// 	// 	return ErrInvalidUser
// 	// }

// 	// chatid, err := strconv.ParseInt(user.GetUserID(), 10, 64)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// telemsg := tgbotapi.NewMessage(chatid, text)

// 	// var arr []tgbotapi.InlineKeyboardButton

// 	// arr = append(arr, tgbotapi.NewInlineKeyboardButtonData("yes", "yes, i am."))
// 	// arr = append(arr, tgbotapi.NewInlineKeyboardButtonData("no", "no, i am not"))

// 	// var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(arr)

// 	// telemsg.ReplyMarkup = numericKeyboard

// 	// cb.teleBotAPI.Send(telemsg)

// 	return nil
// }

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

// procMessageUser
func (cb *teleChatBot) procMessageUser(msg *tgbotapi.Message) (chatbot.User, error) {
	userid := strconv.Itoa(msg.From.ID)
	user := cb.MgrUser.GetUser(userid)
	if user == nil {
		user = chatbot.NewBasicUser(msg.From.UserName, userid,
			msg.From.FirstName+" "+msg.From.LastName, 0)

		cb.MgrUser.AddUser(user)
		cb.GetChatBotDB().UpdUser(user.ToProto())
	}

	return user, nil
}

// procCallbackQuery
func (cb *teleChatBot) procCallbackQuery(ctx context.Context, query *tgbotapi.CallbackQuery) error {
	if query.Message != nil {
		msgid := strconv.Itoa(query.Message.MessageID)

		user, err := cb.procMessageUser(query.Message)
		if err != nil {
			chatbot.Warn("teleChatBot.procCallbackQuery:procMessageUser", zap.Error(err))

			return err
		}

		msg, err := cb.GetMsg(makeChatID(user.GetUserID(), msgid))
		if err != nil {
			chatbot.Warn("teleChatBot.procCallbackQuery:GetMsg", zap.Error(err))

			return err
		}

		if msg.GetSelected() > 0 {
			chatbot.SendTextMsg(cb, user, "Sorry, you have made a choice.")

			return nil
		}

		id, err := strconv.Atoi(query.Data)
		if err != nil {
			chatbot.Warn("teleChatBot.procCallbackQuery:GetID", zap.Error(err))

			return err
		}

		err = cb.ProcMsgCallback(ctx, msg, id)
		if err != nil {
			chatbot.Warn("teleChatBot.procCallbackQuery:ProcMsgCallback", zap.Error(err))

			chatbot.SendTextMsg(cb, user, "Sorry, I found some problems, please start over.")

			return err
		}

		cb.DelMsgCallback(msg.GetChatID())
	}

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
		if update.CallbackQuery != nil {
			cb.procCallbackQuery(ctx, update.CallbackQuery)

			continue
		}

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		user, err := cb.procMessageUser(update.Message)
		if err != nil {
			chatbot.Warn("teleChatBot.Start:procMessageUser", zap.Error(err))

			continue
		}

		lastmsgid := int64(update.Message.MessageID)
		if lastmsgid <= user.GetLastMsgID() {
			continue
		}

		user.UpdLastMsgID(lastmsgid)
		cb.GetChatBotDB().UpdUser(user.ToProto())

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

		msgid := strconv.Itoa(update.Message.MessageID)
		msg := cb.NewMsg(makeChatID(user.GetUserID(), msgid), msgid, user, nil,
			update.Message.Text, int64(update.Message.Date))
		// msg := newMsg(strconv.Itoa(update.Message.MessageID),
		// 	user, update.Message.Text, update.Message.Date)

		err = cb.SaveMsg(msg)
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

	chatbot.SendTextMsg(cb, cb.scriptUser, cr.CtrlResult)
	// cb.SendMsg(cb.scriptUser, cr.CtrlResult)

	return nil
}

// NewMsg
func (cb *teleChatBot) NewMsg(chatid string, msgid string, from chatbot.User, to chatbot.User,
	text string, curtime int64) chatbot.Message {

	return &teleMsg{
		chatID:    chatid,
		msgID:     msgid,
		from:      from,
		to:        to,
		text:      text,
		timeStamp: int64(curtime),
	}
}

// SendMsg
func (cb *teleChatBot) SendMsg(msg chatbot.Message) (chatbot.Message, error) {
	tgmsg := msg.(*teleMsg)

	to := msg.GetTo()
	if to == nil {
		return nil, chatbot.ErrInvalidMessageTo
	}

	chatid, err := strconv.ParseInt(to.GetUserID(), 10, 64)
	if err != nil {
		return nil, err
	}

	telemsg := tgbotapi.NewMessage(chatid, msg.GetText())

	if msg.HasOptions() {
		var lst []tgbotapi.InlineKeyboardButton

		for _, v := range tgmsg.Options {
			lst = append(lst, tgbotapi.NewInlineKeyboardButtonData(v.Text, fmt.Sprintf("%v", v.ID)))
		}

		reply := tgbotapi.NewInlineKeyboardMarkup(lst)

		telemsg.ReplyMarkup = reply
	}

	destmsg, err := cb.teleBotAPI.Send(telemsg)
	if err != nil {
		return nil, err
	}

	msg.SetMsgID(fmt.Sprintf("%v", destmsg.MessageID))

	return msg, nil
}

// NewMsgFromProto
func (cb *teleChatBot) NewMsgFromProto(msg *chatbotdbpb.Message) chatbot.Message {

	cmsg := &teleMsg{
		chatID:    msg.GetChatID(),
		msgID:     msg.GetMsgID(),
		text:      msg.GetText(),
		timeStamp: msg.GetTimeStamp(),
	}

	if msg.GetFrom() != nil {
		cmsg.from = cb.NewUserFromProto(msg.GetFrom())
	}

	if msg.GetTo() != nil {
		cmsg.to = cb.NewUserFromProto(msg.GetTo())
	}

	for _, v := range msg.Options {
		cmsg.AddOption(v)
	}

	return cmsg
}

// GetMsg -
func (cb *teleChatBot) GetMsg(chatid string) (chatbot.Message, error) {
	msg, err := cb.DB.GetMsg(chatid)
	if err != nil {
		return nil, err
	}

	return cb.NewMsgFromProto(msg), nil
}
