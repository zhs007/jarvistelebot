package plugincore

// // cmdVersion - version
// func cmdVersion(ctx context.Context, params *chatbot.MessageParams) bool {
// 	chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), params.ChatBot.GetVersion())

// 	return true
// }

// // cmdUsers - users
// func cmdUsers(ctx context.Context, params *chatbot.MessageParams) bool {
// 	// coredb := params.ChatBot.GetJarvisNodeCoreDB()

// 	lst, err := params.ChatBot.GetChatBotDB().GetUsers(100)
// 	if err != nil {
// 		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())

// 		return true
// 	}

// 	strret, err := chatbot.FormatJSONObj(lst)
// 	if err != nil {
// 		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
// 	} else {
// 		chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret)
// 	}

// 	return true
// }

// // cmdUser - user
// func cmdUser(ctx context.Context, params *chatbot.MessageParams) bool {
// 	if len(params.LstStr) == 3 {
// 		lst, err := params.ChatBot.GetChatBotDB().GetUser(params.LstStr[2])
// 		if err != nil {
// 			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
// 		}

// 		strret, err := chatbot.FormatJSONObj(lst)
// 		if err != nil {
// 			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), err.Error())
// 		} else {
// 			chatbot.SendTextMsg(params.ChatBot, params.Msg.GetFrom(), strret)
// 		}

// 		return true
// 	}

// 	return false
// }
