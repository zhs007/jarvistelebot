package telebot

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarviscore/base"
	"github.com/zhs007/jarviscore/proto"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/assistant"
	"github.com/zhs007/jarvistelebot/plugins/core"
	"github.com/zhs007/jarvistelebot/plugins/filetemplate"
	"github.com/zhs007/jarvistelebot/plugins/filetransfer"
	"github.com/zhs007/jarvistelebot/plugins/jarvisnode"
	"github.com/zhs007/jarvistelebot/plugins/jarvisnodeex"
	"github.com/zhs007/jarvistelebot/plugins/normal"
	"github.com/zhs007/jarvistelebot/plugins/timestamp"
	"github.com/zhs007/jarvistelebot/plugins/usermgr"
	"github.com/zhs007/jarvistelebot/plugins/userscript"
	"github.com/zhs007/jarvistelebot/plugins/xlsx2json"

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
	mgrPlugins.RegPlugin(pluginassistant.PluginName, pluginassistant.NewPlugin)
	mgrPlugins.RegPlugin(pluginjarvisnode.PluginName, pluginjarvisnode.NewPlugin)
	mgrPlugins.RegPlugin(plugintimestamp.PluginName, plugintimestamp.NewPlugin)
	mgrPlugins.RegPlugin(pluginxlsx2json.PluginName, pluginxlsx2json.NewPlugin)
	mgrPlugins.RegPlugin(pluginfiletransfer.PluginName, pluginfiletransfer.NewPlugin)
	mgrPlugins.RegPlugin(pluginnormal.PluginName, pluginnormal.NewPlugin)
	mgrPlugins.RegPlugin(plugincore.PluginName, plugincore.NewPlugin)
	mgrPlugins.RegPlugin(pluginjarvisnodeex.PluginName, pluginjarvisnodeex.NewPlugin)
	mgrPlugins.RegPlugin(pluginusermgr.PluginName, pluginusermgr.NewPlugin)
	mgrPlugins.RegPlugin(pluginuserscript.PluginName, pluginuserscript.NewPlugin)
	mgrPlugins.RegPlugin(pluginfiletemplate.PluginName, pluginfiletemplate.NewPlugin)

	mgrPlugins.SetDefaultPlugin(cfg.DefaultPlugin)

	for _, v := range cfg.Plugins {
		mgrPlugins.NewPlugin(v)
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

	mgrPlugins := chatbot.NewPluginsMgr(cfg.CfgPath)

	regPlugins(cfg, mgrPlugins)

	tcb := &teleChatBot{
		BasicChatBot: chatbot.NewBasicChatBot(),
		teleBotAPI:   bot,
		// mgrUser:    chatbot.NewUserMgr(),
		// MgrPlugins: mgrPlugins,
	}

	// tcb.MgrUser = newTeleUserMgr(cfg.TeleBotMaster)

	tcb.Init(path.Join(cfg.CfgPath, "chatbot.yaml"), mgrPlugins)

	tcb.SetMaster("", cfg.TeleBotMaster)
	// tcb.NewEventMgr(tcb)

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

func (cb *teleChatBot) procDocumentWithMsg(msg chatbot.Message, doc *tgbotapi.Document) error {
	jarvisbase.Debug("teleChatBot.procDocument")

	file, err := cb.teleBotAPI.GetFile(tgbotapi.FileConfig{
		FileID: doc.FileID,
	})
	if err != nil {
		return err
	}

	// localfn := path.Join(cb.GetConfig().DownloadPath, doc.FileName)
	// if doc.MimeType == "text/x-script.sh" {
	// 	localfn = path.Join(cb.GetConfig().DownloadPath, "scripts", doc.FileName)
	// }

	// jarvisbase.Info("teleChatBot.procDocument",
	// 	jarvisbase.JSON("file", file))

	url := file.Link(cb.teleBotAPI.Token)

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	// f, err := os.Create(localfn)
	// if err != nil {
	// 	return err
	// }
	// io.Copy(f, res.Body)

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	fileobj := &chatbotdbpb.File{
		Filename: doc.FileName,
		Data:     buf.Bytes(),
		FileType: doc.MimeType,
	}

	msg.SetFile(fileobj)

	// dat, err := ioutil.ReadFile(localfn)
	// if err != nil {
	// 	jarvisbase.Warn("load script file", zap.Error(err))

	// 	return err
	// }

	// ci, err := jarviscore.BuildCtrlInfoForScriptFile(1, doc.FileName, dat, "")

	// cb.Node.RequestCtrl(ctx, "1NutSP6ypvLtHpqHaxtjJMmEUbMfLUdp9a", ci)

	return nil
}

func (cb *teleChatBot) procPhotoWithMsg(msg chatbot.Message, photo *tgbotapi.PhotoSize) error {
	jarvisbase.Debug("teleChatBot.procPhotoWithMsg")

	file, err := cb.teleBotAPI.GetFile(tgbotapi.FileConfig{
		FileID: photo.FileID,
	})
	if err != nil {
		return err
	}

	url := file.Link(cb.teleBotAPI.Token)

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	fileobj := &chatbotdbpb.File{
		Filename: photo.FileID,
		Data:     buf.Bytes(),
		FileType: chatbot.FileTypePhoto,
	}

	msg.SetFile(fileobj)

	return nil
}

// procMessageUser
func (cb *teleChatBot) procMessageUser(user *tgbotapi.User) (chatbot.User, error) {
	userid := strconv.Itoa(user.ID)
	curuser := cb.MgrUser.GetUser(userid)
	if curuser == nil {
		curuser = chatbot.NewBasicUser(user.UserName, userid,
			user.FirstName+" "+user.LastName, 0)

		cb.MgrUser.AddUser(curuser)
		cb.GetChatBotDB().UpdUser(curuser.ToProto())
	}

	return curuser, nil
}

// procCallbackQuery
func (cb *teleChatBot) procCallbackQuery(ctx context.Context, query *tgbotapi.CallbackQuery) error {
	if query.Message != nil {
		msgid := strconv.Itoa(query.Message.MessageID)

		user, err := cb.procMessageUser(query.From)
		if err != nil {
			chatbot.Warn("teleChatBot.procCallbackQuery:procMessageUser", zap.Error(err))

			return err
		}

		msg, err := cb.GetMsg(chatbot.MakeChatID(user.GetUserID(), msgid))
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

		err = msg.SelectOption(id)
		if err != nil {
			chatbot.Warn("teleChatBot.procCallbackQuery:SelectOption", zap.Error(err))

			return err
		}

		cb.DelMsgCallback(msg.GetChatID())

		err = cb.SaveMsg(msg)
		if err != nil {
			chatbot.Warn("teleChatBot.procCallbackQuery:SaveMsg", zap.Error(err))

			return err
		}

		editText := tgbotapi.NewEditMessageText(
			query.Message.Chat.ID,
			query.Message.MessageID,
			query.Message.Text+fmt.Sprintf(" (you choice %v)", msg.GetOption(id)),
		)

		cb.teleBotAPI.Send(editText)

		configAlert := tgbotapi.NewCallback(query.ID, "")
		cb.teleBotAPI.AnswerCallbackQuery(configAlert)
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

	node.RegMsgEventFunc(jarviscore.EventOnReplyRequestFile,
		func(curctx context.Context, jarvisnode jarviscore.JarvisNode, msg *jarviscorepb.JarvisMsg) error {
			return cb.OnJarvisCtrlResult(curctx, msg)
		})

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := cb.teleBotAPI.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	cb.OnEvent(ctx, cb, chatbot.EventOnStarted)

	for update := range updates {
		if update.CallbackQuery != nil {
			cb.procCallbackQuery(ctx, update.CallbackQuery)

			continue
		}

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		user, err := cb.procMessageUser(update.Message.From)
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

		if update.Message.Text == "" && update.Message.Document == nil && update.Message.Photo == nil {
			continue
		}

		msgid := strconv.Itoa(update.Message.MessageID)
		msg := cb.NewMsg(chatbot.MakeChatID(user.GetUserID(), msgid), msgid, user, nil,
			update.Message.Text, int64(update.Message.Date))

		if update.Message.Document != nil {
			err = cb.procDocumentWithMsg(msg, update.Message.Document)
			if err != nil {
				chatbot.Warn("teleChatBot.Start:procDocumentWithMsg", zap.Error(err))
			}

			msg.SetText(update.Message.Caption)
		} else if update.Message.Photo != nil && len(*update.Message.Photo) > 0 {
			if len(*update.Message.Photo) == 2 {
				err = cb.procPhotoWithMsg(msg, &(*update.Message.Photo)[1])
				if err != nil {
					chatbot.Warn("teleChatBot.Start:procPhotoWithMsg", zap.Error(err))
				}

				msg.SetText(update.Message.Caption)
			}
		}

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

	err := cb.ProcJarvisMsgCallback(ctx, msg)
	if err != nil {
		jarvisbase.Warn("teleChatBot.OnJarvisCtrlResult:ProcJarvisMsgCallback", zap.Error(err))

		return err
	}

	err = cb.DelJarvisMsgCallback(msg.SrcAddr, 0)
	if err != nil {
		jarvisbase.Warn("teleChatBot.OnJarvisCtrlResult:DelJarvisMsgCallback", zap.Error(err))

		return err
	}

	// cr := msg.GetCtrlResult()

	// chatbot.SendTextMsg(cb, cb.scriptUser, cr.CtrlResult)
	// cb.SendMsg(cb.scriptUser, cr.CtrlResult)

	return nil
}

// NewMsg
func (cb *teleChatBot) NewMsg(chatid string, msgid string, from chatbot.User, to chatbot.User,
	text string, curtime int64) chatbot.Message {

	msg := &teleMsg{
		chatID:    chatid,
		msgID:     msgid,
		from:      from,
		to:        to,
		text:      text,
		timeStamp: int64(curtime),
	}

	return msg
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

	fd := msg.GetFile()
	if fd != nil {
		fb := tgbotapi.FileBytes{
			Name:  fd.Filename,
			Bytes: fd.Data,
		}

		jarvisbase.Debug("teleChatBot.SendMsg:file",
			zap.String("filename", fb.Name),
			zap.Int("datalen", len(fb.Bytes)))

		telemsg := tgbotapi.NewDocumentUpload(chatid, fb)

		destmsg, err := cb.teleBotAPI.Send(telemsg)
		if err != nil {
			jarvisbase.Warn("teleChatBot.SendMsg:sendfile", zap.Error(err))

			return nil, err
		}

		msg.SetMsgID(fmt.Sprintf("%v", destmsg.MessageID))

		return msg, nil
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

	if msg.GetFile() != nil {
		cmsg.SetFile(msg.GetFile())
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

	cmsg.SelectOption(int(msg.GetSelected()))

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
