package assistantdb

import (
	"fmt"
)

const prefixUserAssistantInfo = "uai:"
const prefixKeyNote = "note:"

func makeNoteKey(userID string, noteID int64) string {
	return fmt.Sprintf("%v%v:%v", prefixKeyNote, userID, noteID)
}

func makeUserAssistantInfoKey(userID string) string {
	return fmt.Sprintf("%v%v", prefixKeyNote, userID)
}
