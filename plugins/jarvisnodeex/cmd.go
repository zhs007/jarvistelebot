package pluginjarvisnodeex

// // cmdScripts - scripts
// func cmdScripts(ctx context.Context, params *chatbot.MessageParams) bool {
// 	files, _ := filepath.Glob(path.Join(params.ChatBot.GetConfig().DownloadPath, "scripts", "*.sh"))

// 	strret, err := chatbot.FormatJSONObj(files)
// 	if err != nil {
// 		str := fmt.Sprintf("%+v", files)
// 		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), str)
// 	} else {
// 		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret)
// 	}

// 	return true
// }
