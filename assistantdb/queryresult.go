package assistantdb

import (
	pb "github.com/zhs007/jarvistelebot/assistantdb/proto"
)

// ResultUserAssistantInfo -
type ResultUserAssistantInfo struct {
	UserAssistantInfo struct {
		MaxNoteID int64    `json:"maxNoteID"`
		Keys      []string `json:"keys"`
	} `json:"userAssistantInfo"`
}

// ResultUpdNote -
type ResultUpdNote struct {
	UpdNote struct {
		NoteID     int64    `json:"noteID"`
		Data       []string `json:"data"`
		Keys       []string `json:"keys"`
		CreateTime int64    `json:"createTime"`
		UpdateTime int64    `json:"updateTime"`
	} `json:"updNote"`
}

// ResultUpdUserAssistantInfo -
type ResultUpdUserAssistantInfo struct {
	UpdUserAssistantInfo struct {
		MaxNoteID int64    `json:"maxNoteID"`
		Keys      []string `json:"keys"`
	} `json:"updUserAssistantInfo"`
}

// ResultNote -
type ResultNote struct {
	Note struct {
		NoteID     int64    `json:"noteID"`
		Data       []string `json:"data"`
		Keys       []string `json:"keys"`
		CreateTime int64    `json:"createTime"`
		UpdateTime int64    `json:"updateTime"`
	} `json:"note"`
}

// ResultUpdKeyInfo -
type ResultUpdKeyInfo struct {
	UpdKeyInfo struct {
		NoteIDs []int64 `json:"noteIDs"`
	} `json:"updKeyInfo"`
}

// ResultKeyInfo -
type ResultKeyInfo struct {
	KeyInfo struct {
		NoteIDs []int64 `json:"noteIDs"`
	} `json:"keyInfo"`
}

// ResultUserAssistantInfo2UserAssistantInfo - ResultUserAssistantInfo -> UserAssistantInfo
func ResultUserAssistantInfo2UserAssistantInfo(result *ResultUserAssistantInfo) *pb.UserAssistantInfo {
	dat := &pb.UserAssistantInfo{
		MaxNoteID: result.UserAssistantInfo.MaxNoteID,
		Keys:      result.UserAssistantInfo.Keys,
	}

	return dat
}

// ResultUpdUserAssistantInfo2UserAssistantInfo - ResultUpdUserAssistantInfo -> UserAssistantInfo
func ResultUpdUserAssistantInfo2UserAssistantInfo(result *ResultUpdUserAssistantInfo) *pb.UserAssistantInfo {
	dat := &pb.UserAssistantInfo{
		MaxNoteID: result.UpdUserAssistantInfo.MaxNoteID,
		Keys:      result.UpdUserAssistantInfo.Keys,
	}

	return dat
}

// ResultUpdNote2Note - ResultUpdNote -> Note
func ResultUpdNote2Note(result *ResultUpdNote) *pb.Note {
	note := &pb.Note{
		NoteID:     result.UpdNote.NoteID,
		Data:       result.UpdNote.Data,
		Keys:       result.UpdNote.Keys,
		CreateTime: result.UpdNote.CreateTime,
		UpdateTime: result.UpdNote.UpdateTime,
	}

	return note
}

// ResultNote2Note - ResultNote -> Message
func ResultNote2Note(result *ResultNote) *pb.Note {
	note := &pb.Note{
		NoteID:     result.Note.NoteID,
		Data:       result.Note.Data,
		Keys:       result.Note.Keys,
		CreateTime: result.Note.CreateTime,
		UpdateTime: result.Note.UpdateTime,
	}

	return note
}

// ResultUpdKeyInfo2KeyInfo - ResultUpdKeyInfo -> KeyInfo
func ResultUpdKeyInfo2KeyInfo(result *ResultUpdKeyInfo) *pb.KeyInfo {
	keyinfo := &pb.KeyInfo{
		NoteIDs: result.UpdKeyInfo.NoteIDs,
	}

	return keyinfo
}

// ResultKeyInfo2KeyInfo - ResultKeyInfo -> KeyInfo
func ResultKeyInfo2KeyInfo(result *ResultKeyInfo) *pb.KeyInfo {
	keyinfo := &pb.KeyInfo{
		NoteIDs: result.KeyInfo.NoteIDs,
	}

	return keyinfo
}
