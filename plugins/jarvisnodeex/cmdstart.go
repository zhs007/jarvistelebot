package pluginjarvisnodeex

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/plugins/jarvisnodeex/proto"
)

// cmdStart - start
type cmdStart struct {
}

// RunCommand - run command
func (cmd *cmdStart) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
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

		_, ok = params.CurPlugin.(*jarvisnodeexPlugin)
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

			return false
		}

		if jnpd.RunScript == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrInvalidUserPluginData.Error(), params.Msg)

			return false
		}

		var arr []*jarviscorepb.FileData
		for _, v := range jnpd.RunScript.Files {
			fd := &jarviscorepb.FileData{
				Filename: v.Filename,
				File:     v.Data,
				DestPath: v.FullPath,
			}

			arr = append(arr, fd)
		}

		sf := &jarviscorepb.FileData{
			Filename: rscmd.ScriptFile.Filename,
			File:     rscmd.ScriptFile.Data,
		}
		ci, err := jarviscore.BuildCtrlInfoForScriptFile2(1, sf, arr)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false
		}

		params.ChatBot.GetJarvisNode().RequestCtrl(ctx, curnode.Addr, ci,
			func(ctx context.Context, jarvisnode jarviscore.JarvisNode, request *jarviscorepb.JarvisMsg,
				reply *jarviscorepb.JarvisMsg) (bool, error) {

				return true, nil
			})

		params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
			cr := msg.GetCtrlResult()

			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), cr.CtrlResult, params.Msg)

			return nil
		})

		params.MgrPlugins.SetCurPlugin(nil)

		jnpd.RunScript = nil
		from.StorePluginData(PluginName, jnpd)

		return true
	}

	return false
}

// Parse - parse command line
func (cmd *cmdStart) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	from := params.Msg.GetFrom()
	if from != nil {
		pd, ok := from.GetPluginData(PluginName)
		if !ok {
			return nil, chatbot.ErrNoUserPluginData
		}

		jnpd, ok := pd.(*pluginjarvisnodeexpb.PluginData)
		if !ok {
			return nil, chatbot.ErrInvalidUserPluginDataType
		}

		if jnpd.RunScript == nil {
			return nil, chatbot.ErrInvalidUserPluginData
		}

		if len(params.LstStr) != 1 {
			return nil, chatbot.ErrInvalidCommandLineItemNums
		}

		if params.LstStr[0] == "start" {
			return jnpd.RunScript, nil
		}
	}

	// flagset := pflag.NewFlagSet("start", pflag.ContinueOnError)

	// err := flagset.Parse(params.LstStr[1:])
	// if err != nil {
	// 	return nil, err
	// }

	// if *nodename != "" {
	// 	rsc := &pluginjarvisnodeexpb.RunScriptCommand{
	// 		JarvisNodeName: *nodename,
	// 	}

	// 	if deststr != nil {
	// 		rsc.DestGlobPatterns = *deststr
	// 	}

	// 	file := params.Msg.GetFile()
	// 	if file != nil {
	// 		if file.FileType == chatbot.FileTypeShellScript {
	// 			rsc.ScriptFile = &pluginjarvisnodeexpb.FileData{
	// 				Filename: file.Filename,
	// 				FileType: file.FileType,
	// 				Data:     file.Data,
	// 			}
	// 		} else {
	// 			return nil, chatbot.ErrOnlyScriptFile
	// 		}
	// 	}

	// 	return rsc, nil
	// }

	return nil, chatbot.ErrInvalidCommandLine
}
