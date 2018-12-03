package assistant

import "errors"

var (
	// ErrNoUserAssistantInfo - no UserAssistantInfo
	ErrNoUserAssistantInfo = errors.New("no UserAssistantInfo")
	// ErrNoCurNote - no current node
	ErrNoCurNote = errors.New("no current node")
	// ErrInvalidCurNoteMode - invalid current node mode
	ErrInvalidCurNoteMode = errors.New("invalid current node mode")
)
