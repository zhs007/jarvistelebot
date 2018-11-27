package chatbotdb

const prefixKeyMessage = "msg:"
const prefixKeyUser = "user:"
const prefixKeyUserName = "uname:"
const prefixKeyUserScript = "userscript:"

func makeMessageKey(chatID string) string {
	return prefixKeyMessage + chatID
}

func makeUserKey(userID string) string {
	return prefixKeyUser + userID
}

func makeUserNameKey(userName string) string {
	return prefixKeyUserName + userName
}

func makeUserScriptKey(userID string, scriptName string) string {
	return prefixKeyUserScript + userID + ":" + scriptName
}
