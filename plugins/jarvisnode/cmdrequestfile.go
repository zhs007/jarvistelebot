package pluginjarvisnode

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/jarviscore"
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
	if params.CommandLine != nil {
		rfcmd, ok := params.CommandLine.(*pluginjarvisnodepb.RequestFileCommand)
		if !ok {
			return false
		}

		curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(rfcmd.JarvisNodeName)
		if curnode == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), "Sorry, I can't find this node.", params.Msg)

			return true
		}

		rf := &jarviscorepb.RequestFile{
			Filename: rfcmd.FileFullPath,
		}

		// params.ChatBot.GetJarvisNode().RequestFile(ctx, curnode.Addr, rf, nil)
		var buf bytes.Buffer

		params.ChatBot.GetJarvisNode().RequestFile(ctx, curnode.Addr, rf,
			func(ctx context.Context, jarvisnode jarviscore.JarvisNode, lstResult []*jarviscore.JarvisMsgInfo) error {

				// for ; lastresultindex < len(lstResult); lastresultindex++ {

				curmsg := lstResult[len(lstResult)-1].Msg
				if curmsg != nil {
					if lstResult[len(lstResult)-1].JarvisResultType == jarviscore.JarvisResultTypeSend {

						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							jarviscore.AppendString("I send request ", rfcmd.FileFullPath, " to ", curnode.Name), params.Msg)

					} else if curmsg.MsgType == jarviscorepb.MSGTYPE_REPLY2 {
						if curmsg.ReplyType == jarviscorepb.REPLYTYPE_IGOTIT {
							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
								fmt.Sprintf("%v has received the file request(%v).",
									rfcmd.JarvisNodeName, rfcmd.FileFullPath),
								params.Msg)
						} else if curmsg.ReplyType == jarviscorepb.REPLYTYPE_ERROR {
							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
								curmsg.Err,
								params.Msg)
						}

					} else if curmsg.MsgType == jarviscorepb.MSGTYPE_REPLY_REQUEST_FILE {
						isend, err := chatbot.ProcReplyRequestFile(curmsg, &buf)
						if err != nil {
							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
								err.Error(), params.Msg)

							return err
						}

						if isend {
							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
								fmt.Sprintf("The %v:%v received %v bytes, the file is received, I will send it to you.",
									rfcmd.JarvisNodeName, rfcmd.FileFullPath, buf.Len()),
								params.Msg)

							chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
								Filename: chatbot.GetFileNameFromFullPath(rfcmd.FileFullPath),
								Data:     buf.Bytes(),
							}, params.Msg)
						} else {
							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
								fmt.Sprintf("The %v:%v received %v bytes.",
									rfcmd.JarvisNodeName, rfcmd.FileFullPath, buf.Len()),
								params.Msg)
						}
					}
				}

				if lstResult[len(lstResult)-1].Err != nil {
					chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
						lstResult[len(lstResult)-1].Err.Error(), params.Msg)
				}
				// }

				return nil
			})

		// params.ChatBot.AddJarvisMsgCallback(curnode.Addr, 0, func(ctx context.Context, msg *jarviscorepb.JarvisMsg) error {
		// 	if msg.MsgType == jarviscorepb.MSGTYPE_REPLY_REQUEST_FILE {
		// 		fd := msg.GetFile()

		// 		chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
		// 			Filename: chatbot.GetFileNameFromFullPath(fd.Filename),
		// 			Data:     fd.File,
		// 		}, params.Msg)
		// 	}

		// 	return nil
		// })
	}

	return true
}

// Parse - parse command line
func (cmd *cmdRequestFile) ParseCommandLine(params *chatbot.MessageParams) (proto.Message, error) {
	if len(params.LstStr) < 1 {
		return nil, chatbot.ErrInvalidCommandLineItemNums
	}

	flagset := pflag.NewFlagSet("requestfile", pflag.ContinueOnError)

	var nodename = flagset.StringP("nodename", "n", "", "you want request file from this node")
	var fullpath = flagset.StringP("filepath", "f", "", "you want request this file")

	err := flagset.Parse(params.LstStr[1:])
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
