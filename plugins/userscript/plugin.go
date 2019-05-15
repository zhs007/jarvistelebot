package pluginuserscript

import (
	"context"
	"fmt"

	"github.com/zhs007/jarviscore"
	jarvisbase "github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/golang/protobuf/proto"
	jarviscorepb "github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
	pluginuserscriptpb "github.com/zhs007/jarvistelebot/plugins/userscript/proto"
)

// PluginName - plugin name
const PluginName = "userscript"

// userscriptPlugin - userscript plugin
type userscriptPlugin struct {
}

// NewPlugin - new normal plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - userscriptPlugin")

	return &userscriptPlugin{}, nil
}

// OnMessage - get message
func (p *userscriptPlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if params.CommandLine != nil {
		rscmd, ok := params.CommandLine.(*pluginuserscriptpb.RunScriptCommand)
		if !ok {
			return false, chatbot.ErrInvalidCommandLine
		}

		us, err := params.ChatBot.GetChatBotDB().GetUserScript(from.GetUserID(), rscmd.ScriptName)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false, err
		}

		chatbot.SendTextMsg(params.ChatBot, from,
			"I will execute the script "+rscmd.ScriptName+" for "+us.JarvisNodeName, params.Msg)

		sf := &jarviscorepb.FileData{
			Filename: us.File.Filename,
			File:     us.File.Data,
		}
		ci, err := jarviscore.BuildCtrlInfoForScriptFile2(sf, nil, rscmd.ScriptName)
		// ci, err := jarviscore.BuildCtrlInfoForScriptFile(1, us.File.Filename, us.File.Data, "")
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)
			// jarvisbase.Warn("userscriptPlugin.OnMessage", zap.Error(err))

			return false, err
		}

		curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(us.JarvisNodeName)
		if curnode == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrNoJarvisNode.Error(), params.Msg)

			return false, chatbot.ErrNoJarvisNode
		}

		isrecv := false
		params.ChatBot.GetJarvisNode().RequestCtrl(ctx, curnode.Addr, ci,
			func(ctx context.Context, jarvisnode jarviscore.JarvisNode,
				lstResult []*jarviscore.JarvisMsgInfo) error {

				if !isrecv && len(lstResult) > 0 {
					if lstResult[len(lstResult)-1].JarvisResultType == jarviscore.JarvisResultTypeSend {

						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							jarviscore.AppendString("I send script ", rscmd.ScriptName, " to ", us.JarvisNodeName), params.Msg)

					} else if lstResult[len(lstResult)-1].JarvisResultType == jarviscore.JarvisResultTypeRemoved {

						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							jarviscore.AppendString(us.JarvisNodeName, " maybe restarted, you can resend the ", rscmd.ScriptName), params.Msg)

					} else if lstResult[len(lstResult)-1].Err != nil {
						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							lstResult[len(lstResult)-1].Err.Error(), params.Msg)
					} else if lstResult[len(lstResult)-1].Msg != nil {
						cm := lstResult[len(lstResult)-1].Msg

						if cm.MsgType == jarviscorepb.MSGTYPE_REPLY2 &&
							cm.ReplyType == jarviscorepb.REPLYTYPE_ERROR {

							chatbot.SendTextMsg(params.ChatBot,
								params.Msg.GetFrom(),
								fmt.Sprintf("%v reply err %v.",
									curnode.Name, cm.Err),
								params.Msg)

						} else if cm.MsgType == jarviscorepb.MSGTYPE_REPLY2 &&
							cm.ReplyType == jarviscorepb.REPLYTYPE_END {

							chatbot.SendTextMsg(params.ChatBot, from, "It's done.", params.Msg)

						} else if cm.MsgType == jarviscorepb.MSGTYPE_REPLY2 &&
							cm.ReplyType == jarviscorepb.REPLYTYPE_IGOTIT {

							chatbot.SendTextMsg(params.ChatBot,
								params.Msg.GetFrom(),
								fmt.Sprintf("%v has received the request (%v).",
									us.JarvisNodeName, rscmd.ScriptName),
								params.Msg)

							// params.ChatBot.AddJarvisMsgCallback(curnode.Addr, cm.ReplyMsgID,
							// 	func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
							// 		cr := msg.GetCtrlResult()
							// 		if cr == nil {
							// 			msgstr := fmt.Sprintf("%v", msg)
							// 			jarvisbase.Warn("userscriptPlugin.AddJarvisMsgCallback", zap.String("msg", msgstr))

							// 			chatbot.SendTextMsg(params.ChatBot, from, msgstr, params.Msg)

							// 			return nil
							// 		}

							// 		if cr.CtrlResult != "" {
							// 			chatbot.SendTextMsg(params.ChatBot, from, cr.CtrlResult, params.Msg)
							// 		}

							// 		chatbot.SendTextMsg(params.ChatBot, from, "It's done.", params.Msg)

							// 		return nil
							// 	})
						} else if cm.MsgType == jarviscorepb.MSGTYPE_REPLY_CTRL_RESULT {
							cr := cm.GetCtrlResult()
							if cr == nil {
								msgstr := fmt.Sprintf("%v", cm)
								jarvisbase.Warn("userscriptPlugin.AddJarvisMsgCallback", zap.String("msg", msgstr))

								chatbot.SendTextMsg(params.ChatBot, from, msgstr, params.Msg)

								return nil
							}

							if cr.CtrlResult != "" {
								chatbot.SendTextMsg(params.ChatBot, from, cr.CtrlResult, params.Msg)
							}

							if cr.ErrInfo != "" {
								chatbot.SendTextMsg(params.ChatBot, from, cr.ErrInfo, params.Msg)
							}

							// chatbot.SendTextMsg(params.ChatBot, from, "It's done.", params.Msg)
						}
					}
				}

				return nil
			})
	}

	return true, nil
}

// GetPluginName - get plugin name
func (p *userscriptPlugin) GetPluginName() string {
	return PluginName
}

// OnStart - on start
func (p *userscriptPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *userscriptPlugin) GetPluginType() int {
	return chatbot.PluginTypeNormal
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *userscriptPlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	if params.Msg.GetFrom() != nil && len(params.LstStr) == 1 {
		mgrUserScripts := params.ChatBot.GetUserScriptsMgr()
		us, err := mgrUserScripts.Get(params.ChatBot.GetChatBotDB(), params.Msg.GetFrom().GetUserID())
		if err != nil {
			jarvisbase.Warn("userscriptPlugin.ParseMessage:mgrUserScripts.Get", zap.Error(err))

			return nil, chatbot.ErrMsgNotMine
		}

		for _, v := range us.Scripts {
			if params.LstStr[0] == v {
				return &pluginuserscriptpb.RunScriptCommand{
					ScriptName: v,
				}, nil
			}
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
