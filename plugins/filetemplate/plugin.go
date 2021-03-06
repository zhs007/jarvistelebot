package pluginfiletemplate

import (
	"bytes"
	"context"
	"fmt"

	"github.com/zhs007/jarviscore"
	jarvisbase "github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/golang/protobuf/proto"
	jarviscorepb "github.com/zhs007/jarviscore/proto"
	"github.com/zhs007/jarvistelebot/chatbot"
	chatbotdbpb "github.com/zhs007/jarvistelebot/chatbotdb/proto"
	pluginfiletemplatepb "github.com/zhs007/jarvistelebot/plugins/filetemplate/proto"
)

// PluginName - plugin name
const PluginName = "filetemplate"

// filetemplatePlugin - filetemplate plugin
type filetemplatePlugin struct {
}

// NewPlugin - new normal plugin
func NewPlugin(cfgPath string) (chatbot.Plugin, error) {
	chatbot.Info("NewPlugin - filetemplatePlugin")

	return &filetemplatePlugin{}, nil
}

// OnMessage - get message
func (p *filetemplatePlugin) OnMessage(ctx context.Context, params *chatbot.MessageParams) (bool, error) {
	from := params.Msg.GetFrom()
	if from == nil {
		return false, chatbot.ErrMsgNoFrom
	}

	if params.CommandLine != nil {
		ftcmd, ok := params.CommandLine.(*pluginfiletemplatepb.FileTemplateCommand)
		if !ok {
			return false, chatbot.ErrInvalidCommandLine
		}

		ft, err := params.ChatBot.GetChatBotDB().GetFileTemplate(from.GetUserID(), ftcmd.FileTemplateName)
		if err != nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

			return false, err
		}

		file := params.Msg.GetFile()
		if file != nil {
			chatbot.SendTextMsg(params.ChatBot, from,
				fmt.Sprintf("I will send %v to %v.", ft.FullPath, ft.JarvisNodeName), params.Msg)

			curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(ft.JarvisNodeName)
			if curnode == nil {
				chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrNoJarvisNode.Error(), params.Msg)

				return false, chatbot.ErrNoJarvisNode
			}

			fd := &jarviscorepb.FileData{
				File:     file.Data,
				Filename: ft.FullPath,
			}

			sendfilelastresultindex := 0

			params.ChatBot.GetJarvisNode().SendFile2(ctx, curnode.Addr, fd,
				func(ctx context.Context, jarvisnode jarviscore.JarvisNode, lstResult []*jarviscore.JarvisMsgInfo) error {

					if len(lstResult) > 0 {
						if lstResult[len(lstResult)-1].JarvisResultType == jarviscore.JarvisResultTypeSend {
							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
								jarviscore.AppendString("I send ", ft.FullPath, " to ", ft.JarvisNodeName), params.Msg)
						} else if lstResult[len(lstResult)-1].JarvisResultType == jarviscore.JarvisResultTypeRemoved {

							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
								jarviscore.AppendString(ft.JarvisNodeName, " maybe restarted, you can resend the file."), params.Msg)
						}
					}

					for ; sendfilelastresultindex < len(lstResult); sendfilelastresultindex++ {
						curmsg := lstResult[sendfilelastresultindex].Msg
						if curmsg != nil {
							if curmsg.MsgType == jarviscorepb.MSGTYPE_REPLY2 {
								if curmsg.ReplyType == jarviscorepb.REPLYTYPE_IGOTIT {
									chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
										fmt.Sprintf("%v has received the file (%v).",
											ft.JarvisNodeName, ft.FullPath),
										params.Msg)
								} else if curmsg.ReplyType == jarviscorepb.REPLYTYPE_ERROR {
									chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
										curmsg.Err,
										params.Msg)
								}
							} else if curmsg.MsgType == jarviscorepb.MSGTYPE_REPLY2 {
								if curmsg.ReplyType == jarviscorepb.REPLYTYPE_END {
									chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
										fmt.Sprintf("%v:%v is sent and the verification is successful.",
											ft.JarvisNodeName, ft.FullPath),
										params.Msg)
								}
							}
						}
					}

					return nil
				})

			if ft.SubfilesPath != "" {
				chatbot.SendTextMsg(params.ChatBot, from,
					fmt.Sprintf("I processing subfiles."), params.Msg)

				err = procSubfiles(params.ChatBot.GetChatBotDB(), from.GetUserID(), ft.JarvisNodeName, file.Data, ft.SubfilesPath)
				if err != nil {
					chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

					return false, err
				}

				chatbot.SendTextMsg(params.ChatBot, from,
					fmt.Sprintf("I reload your filetemplates."), params.Msg)

				mgrFileTemplates := params.ChatBot.GetFileTemplatesMgr()
				_, err := mgrFileTemplates.Load(params.ChatBot.GetChatBotDB(), from.GetUserID())
				if err != nil {
					chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error(), params.Msg)

					return false, err
				}
			}

			return true, nil
		}

		chatbot.SendTextMsg(params.ChatBot, from,
			fmt.Sprintf("I will get %v:%v.", ft.JarvisNodeName, ft.FullPath), params.Msg)

		curnode := params.ChatBot.GetJarvisNode().FindNodeWithName(ft.JarvisNodeName)
		if curnode == nil {
			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), chatbot.ErrNoJarvisNode.Error(), params.Msg)

			return false, chatbot.ErrNoJarvisNode
		}

		rf := &jarviscorepb.RequestFile{
			Filename: ft.FullPath,
		}

		// lastresultindex := 0
		// currecvlen := int64(0)
		var buf bytes.Buffer

		params.ChatBot.GetJarvisNode().RequestFile(ctx, curnode.Addr, rf,
			func(ctx context.Context, jarvisnode jarviscore.JarvisNode, lstResult []*jarviscore.JarvisMsgInfo) error {

				if len(lstResult) > 0 {
					if lstResult[len(lstResult)-1].JarvisResultType == jarviscore.JarvisResultTypeSend {
						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							jarviscore.AppendString("I send request ", ft.FullPath, " from ", ft.JarvisNodeName), params.Msg)
					} else if lstResult[len(lstResult)-1].JarvisResultType == jarviscore.JarvisResultTypeRemoved {

						chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
							jarviscore.AppendString(ft.JarvisNodeName, " maybe restarted, you can retry request the file."), params.Msg)
					}
				}

				curmsg := lstResult[len(lstResult)-1].Msg
				if curmsg != nil {
					if curmsg.MsgType == jarviscorepb.MSGTYPE_REPLY2 {
						if curmsg.ReplyType == jarviscorepb.REPLYTYPE_IGOTIT {
							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
								fmt.Sprintf("%v has received the file request(%v).",
									ft.JarvisNodeName, ft.FullPath),
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
									ft.JarvisNodeName, ft.FullPath, buf.Len()),
								params.Msg)

							chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
								Filename: ft.FileTemplateName,
								Data:     buf.Bytes(),
							}, params.Msg)
						} else {
							chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
								fmt.Sprintf("The %v:%v received %v bytes.",
									ft.JarvisNodeName, ft.FullPath, buf.Len()),
								params.Msg)
						}

						// curfi := curmsg.GetFile()
						// if curfi == nil {
						// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
						// 		jarviscore.ErrNoFileData.Error(), params.Msg)

						// 	return jarviscore.ErrNoFileData
						// }

						// if curfi.TotalLength == curfi.Length {

						// 	chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
						// 		Filename: ft.FileTemplateName,
						// 		Data:     curfi.File,
						// 	})

						// } else {
						// 	currecvlen = currecvlen + int64(len(curmsg.GetFile().File))

						// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
						// 		fmt.Sprintf("The %v:%v length is %v, I received %v bytes.",
						// 			ft.JarvisNodeName, ft.FullPath, curfi.TotalLength, currecvlen),
						// 		params.Msg)

						// 	if curfi.FileMD5String != "" {
						// 		var b bytes.Buffer

						// 		for j := 0; j < len(lstResult); j++ {
						// 			if lstResult[j].Msg != nil &&
						// 				lstResult[j].Msg.MsgType == jarviscorepb.MSGTYPE_REPLY_REQUEST_FILE &&
						// 				lstResult[j].Msg.GetFile() != nil {

						// 				n, err := b.Write(lstResult[j].Msg.GetFile().File)
						// 				if err != nil {
						// 					chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
						// 						fmt.Sprintf("WriteError: %v", err.Error()),
						// 						params.Msg)

						// 					continue
						// 				}

						// 				if n != len(lstResult[j].Msg.GetFile().File) {
						// 					chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
						// 						fmt.Sprintf("WriteError: invalid length"),
						// 						params.Msg)

						// 					continue
						// 				}
						// 			}
						// 		}

						// 		strmd5 := jarviscore.GetMD5String(b.Bytes())
						// 		if strmd5 != curfi.FileMD5String {
						// 			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
						// 				fmt.Sprintf("The file length is %v(%v), the MD5 is %v(%v).",
						// 					curfi.TotalLength, b.Len(), curfi.FileMD5String, strmd5),
						// 				params.Msg)

						// 			continue
						// 		}

						// 		chatbot.SendFileMsg(params.ChatBot, params.Msg.GetFrom(), &chatbotdbpb.File{
						// 			Filename: ft.FileTemplateName,
						// 			Data:     b.Bytes(),
						// 		})
						// 	}
						// }
					}
				}

				if lstResult[len(lstResult)-1].Err != nil {
					chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(),
						lstResult[len(lstResult)-1].Err.Error(), params.Msg)
				}
				// }

				return nil
			})
	}

	return true, nil
}

// GetPluginName - get plugin name
func (p *filetemplatePlugin) GetPluginName() string {
	return PluginName
}

// OnStart - on start
func (p *filetemplatePlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginType - get pluginType
func (p *filetemplatePlugin) GetPluginType() int {
	return chatbot.PluginTypeNormal
}

// ParseMessage - If this message is what I can process,
//	it will return to the command line, otherwise it will return an error.
func (p *filetemplatePlugin) ParseMessage(params *chatbot.MessageParams) (proto.Message, error) {
	file := params.Msg.GetFile()
	if file != nil {
		if params.Msg.GetText() == "" {
			mgrFileTemplates := params.ChatBot.GetFileTemplatesMgr()
			ft, err := mgrFileTemplates.Get(params.ChatBot.GetChatBotDB(), params.Msg.GetFrom().GetUserID())
			if err != nil {
				jarvisbase.Warn("filetemplatePlugin.ParseMessage:mgrFileTemplates.Get", zap.Error(err))

				return nil, chatbot.ErrMsgNotMine
			}

			for _, v := range ft.Templates {
				if file.Filename == v {
					return &pluginfiletemplatepb.FileTemplateCommand{
						FileTemplateName: v,
					}, nil
				}
			}
		}

		return nil, chatbot.ErrMsgNotMine
	}

	if params.Msg.GetFrom() != nil && len(params.LstStr) == 1 {
		mgrFileTemplates := params.ChatBot.GetFileTemplatesMgr()
		ft, err := mgrFileTemplates.Get(params.ChatBot.GetChatBotDB(), params.Msg.GetFrom().GetUserID())
		if err != nil {
			jarvisbase.Warn("filetemplatePlugin.ParseMessage:mgrFileTemplates.Get", zap.Error(err))

			return nil, chatbot.ErrMsgNotMine
		}

		for _, v := range ft.Templates {
			if params.LstStr[0] == v {
				return &pluginfiletemplatepb.FileTemplateCommand{
					FileTemplateName: v,
				}, nil
			}
		}
	}

	return nil, chatbot.ErrMsgNotMine
}
