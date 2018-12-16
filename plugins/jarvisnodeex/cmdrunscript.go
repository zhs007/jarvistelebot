package pluginjarvisnodeex

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/jarvisnodeex/proto"
)

// cmdRunScript - runscript
type cmdRunScript struct {
}

// RunCommand - run command
func (cmd *cmdRunScript) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	from := params.Msg.GetFrom()
	if from == nil {
		return false
	}

	if params.CommandLine != nil {
		rscmd, ok := params.CommandLine.(*pluginjarvisnodeexpb.RunScriptCommand)
		if !ok {
			return false
		}

		curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(rscmd.JarvisNodeName)
		if curnode == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "Sorry, I can't find this node.", params.Msg)

			return true
		}

		if params.CurPlugin == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidParamsNoCurPlugin.Error(), params.Msg)

			return false
		}

		pluginJarvisnodeex, ok := params.CurPlugin.(*jarvisnodeexPlugin)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidParamsInvalidCurPlugin.Error(), params.Msg)

			return false
		}

		pd, ok := from.GetPluginData(PluginName)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrNoUserPluginData.Error(), params.Msg)

			return false
		}

		jnpd, ok := pd.(*pluginjarvisnodeexpb.PluginData)
		if !ok {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidUserPluginDataType.Error(), params.Msg)

			jnpd = &pluginjarvisnodeexpb.PluginData{}
			from.StorePluginData(PluginName, jnpd)

			return false
		}

		jnpd.RunScript = rscmd
		from.StorePluginData(PluginName, jnpd)

		params.MgrPlugins.SetCurPlugin(pluginJarvisnodeex)

		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "I get it, please send me some files to run script.", params.Msg)
		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "If you want to start run script, you can send ``start``.", params.Msg)

		// sf := &jarviscorepb.FileData{
		// 	Filename: rscmd.ScriptFile.Filename,
		// 	File:     rscmd.ScriptFile.Data,
		// }
		// ci, err := jarviscore.BuildCtrlInfoForScriptFile2(1, sf, nil)
		// if err != nil {
		// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

		// 	return false
		// }

		// params.ChatBot.GetJarvisNode().RequestCtrl(ctx, curnode.Addr, ci)

		// params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
		// 	cr := msg.GetCtrlResult()

		// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), cr.CtrlResult, params.Msg)

		// 	return nil
		// })

		return true

		// rf := &jarviscorepb.RequestFile{
		// 	Filename: rfcmd.FileFullPath,
		// }

		// params.ChatBot.GetJarvisNode().RequestFile(ctx, curnode.Addr, rf)

		// params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
		// 	if msg.MsgType == jarviscorepb.MSGTYPE_REPLY_REQUEST_FILE {
		// 		fd := msg.GetFile()

		// 		chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
		// 			Filename: chatbot.GetFileNameFromFullPath(fd.Filename),
		// 			Data:     fd.File,
		// 		})
		// 	}

		// 	return nil
		// })
	}

	// fn := strings.Join(params.LstStr[2:], " ")

	// arr := strings.Split(fn, ":")
	// if len(arr) < 2 {
	// 	return false
	// }

	// curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(arr[0])
	// if curnode == nil {
	// 	return false
	// }

	// rf := &jarviscorepb.RequestFile{
	// 	Filename: strings.Join(arr[1:], ":"),
	// }

	// params.ChatBot.GetJarvisNode().RequestFile(ctx, curnode.Addr, rf)

	// params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
	// 	if msg.MsgType == jarviscorepb.MSGTYPE_REPLY_REQUEST_FILE {
	// 		fd := msg.GetFile()

	// 		chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
	// 			Filename: chatbot.GetFileNameFromFullPath(fd.Filename),
	// 			Data:     fd.File,
	// 		})
	// 	}

	// 	return nil
	// })

	// // chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), params.ChatBot.GetVersion())

	return false
}

// Parse - parse command line
func (cmd *cmdRunScript) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("runscript", pflag.ContinueOnError)

	var nodename = flagset.StringP("nodename", "n", "", "you want run script with this node")
	var deststr = flagset.StringArrayP("destfiles", "d", nil, "you want get this files when script finished.")

	err := flagset.Parse(params.LstStr[1:])
	if err != nil {
		return nil, err
	}

	if *nodename != "" {
		rsc := &pluginjarvisnodeexpb.RunScriptCommand{
			JarvisNodeName: *nodename,
		}

		if deststr != nil {
			rsc.DestGlobPatterns = *deststr
		}

		file := params.Msg.GetFile()
		if file != nil {
			if file.FileType == chatbot.FileTypeShellScript {
				rsc.ScriptFile = &pluginjarvisnodeexpb.FileData{
					Filename: file.Filename,
					FileType: file.FileType,
					Data:     file.Data,
				}
			} else {
				return nil, chatbot.ErrOnlyScriptFile
			}
		}

		return rsc, nil
	}

	return nil, chatbot.ErrInvalidCommandLine
}
