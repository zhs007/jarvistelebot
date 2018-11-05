package chatbotdb

const prefixKeyMessage = "msg:"

func makeMessageKey(chatID string) string {
	return prefixKeyMessage + chatID
}
