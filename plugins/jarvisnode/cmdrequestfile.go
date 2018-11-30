package pluginjarvisnode

import (
	"context"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
	"github.com/zhs007/jarvistelebot/chatbotdb/proto"
	"github.com/zhs007/jarvistelebot/plugins/jarvisnode/proto"
)

// cmdRequestFile - request file
type cmdRequestFile struct {
}

// RunCommand - run command
func (cmd *cmdRequestFile) RunCommand(ctx context.Context, params *chatbot.MessageParams) bool {
	fn := strings.Join(params.LstStr[2:], " ")

	arr := strings.Split(fn, ":")
	if len(arr) < 2 {
		return false
	}

	curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(arr[0])
	if curnode == nil {
		return false
	}

	rf := &jarviscorepb.RequestFile{
		Filename: strings.Join(arr[1:], ":"),
	}

	params.ChatBot.GetJarvisNode().RequestFile(ctx, curnode.Addr, rf)

	params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
		if msg.MsgType == jarviscorepb.MSGTYPE_REPLY_REQUEST_FILE {
			fd := msg.GetFile()

			chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
				Filename: chatbot.GetFileNameFromFullPath(fd.Filename),
				Data:     fd.File,
			})
		}

		return nil
	})

	// chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), params.ChatBot.GetVersion())

	return true
}

// Parse - parse command line
func (cmd *cmdRequestFile) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 2 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("requestfile", pflag.ContinueOnError)

	var nodename = flagset.StringP("nodename", "n", "", "you want request file from this node")
	var fullpath = flagset.StringP("filepath", "f", "", "you want request this file")

	err := flagset.Parse(params.LstStr[2:])
	if err != nil {
		return nil, err
	}

	if *nodename == "" || *fullpath == "" {
		args := flagset.Args()
		if len(args) == 1 {
			arr := strings.Split(args[0], ":")
			if len(arr) == 2 {
				return &pluginjarvisnodepb.RequestFileCommand{
					JarvisNodeName: arr[0],
					FileFullPath:   arr[1],
				}, nil
			}
		} else if len(args) == 2 {
			return &pluginjarvisnodepb.RequestFileCommand{
				JarvisNodeName: args[0],
				FileFullPath:   args[1],
			}, nil
		}

		return nil, chatbot.ErrInvalidCommandLine
	}

	return &pluginjarvisnodepb.RequestFileCommand{
		JarvisNodeName: *nodename,
		FileFullPath:   *fullpath,
	}, nil
}
