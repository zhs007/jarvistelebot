package assistantdb

import (
	"fmt"
)

const prefixKeyUserAssistantInfo = "uai:"
const prefixKeyNote = "note:"
const prefixKeyKeyInfo = "key:"

func makeNoteKey(userID string, noteID int64) string {
	return fmt.Sprintf("%v%v:%v", prefixKeyNote, userID, noteID)
}

func makeUserAssistantInfoKey(userID string) string {
	return fmt.Sprintf("%v%v", prefixKeyUserAssistantInfo, userID)
}

func makeKeyInfoKey(userID string, key string) string {
	return fmt.Sprintf("%v%v:%v", prefixKeyKeyInfo, userID, key)
}
