package assistantdb

import (
	"fmt"
)

const keyAssistantData = "assistantdata"
const prefixKeyMessage = "msg:"

func makeMessageKey(msgID int64) string {
	return fmt.Sprintf("%v%v", prefixKeyMessage, msgID)
}

func makeAssistantDataKey() string {
	return keyAssistantData
}
