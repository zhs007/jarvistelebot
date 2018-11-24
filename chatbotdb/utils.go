package chatbotdb

const prefixKeyMessage = "msg:"
const prefixKeyUser = "user:"
const prefixKeyUserName = "uname:"

func makeMessageKey(chatID string) string {
	return prefixKeyMessage + chatID
}

func makeUserKey(userID string) string {
	return prefixKeyUser + userID
}

func makeUserNameKey(userName string) string {
	return prefixKeyUserName + userName
}
