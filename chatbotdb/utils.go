package chatbotdb

const prefixKeyMessage = "msg:"
const prefixKeyUser = "user:"

func makeMessageKey(chatID string) string {
	return prefixKeyMessage + chatID
}

func makeUserKey(userID string) string {
	return prefixKeyUser + userID
}
