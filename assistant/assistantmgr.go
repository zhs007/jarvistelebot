package assistant

import (
	"context"
	"sync"
	"time"

	"github.com/zhs007/jarvistelebot/assistantdb"

	"github.com/zhs007/jarvistelebot/assistantdb/proto"
)

// UserAssistantInfo - user's assistant info
type UserAssistantInfo struct {
	MaxNoteID int64
	Keys      []string
	CurNode   *assistantdbpb.Note
}

// Mgr - assistant manager
type Mgr interface {
	// Start - start
	Start(ctx context.Context) error
	// NewNote - new Note
	NewNote(userID string) (*assistantdbpb.Note, error)
}

// assistantMgr - assistant manager
type assistantMgr struct {
	sync.Locker

	mapUser sync.Map
	db      *assistantdb.AssistantDB
}

// NewAssistantMgr - new AssistantMgr
func NewAssistantMgr(dbpath string, httpAddr string, engine string) (Mgr, error) {
	db, err := assistantdb.NewAssistantDB(dbpath, httpAddr, engine)
	if err != nil {
		return nil, err
	}

	return &assistantMgr{
		db: db,
	}, nil
}

// loadUserAssistant - load userAssistant from db
func (mgr *assistantMgr) loadUserAssistant(userID string) (*UserAssistantInfo, error) {
	pbuai, err := mgr.db.LoadUserAssistantInfo(userID)
	if err != nil {
		return nil, err
	}

	uai := &UserAssistantInfo{
		MaxNoteID: pbuai.MaxNoteID,
		Keys:      pbuai.Keys,
	}

	return uai, nil
}

// updUserAssistant - update userAssistant from db
func (mgr *assistantMgr) updUserAssistant(userID string, uai *UserAssistantInfo) (*UserAssistantInfo, error) {
	pbuai := &assistantdbpb.UserAssistantInfo{
		MaxNoteID: uai.MaxNoteID,
		Keys:      uai.Keys,
	}

	pbuai, err := mgr.db.UpdUserAssistantInfo(userID, pbuai)
	if err != nil {
		return nil, err
	}

	uai.MaxNoteID = pbuai.MaxNoteID
	uai.Keys = pbuai.Keys

	return uai, nil
}

// NewNote - new Note
func (mgr *assistantMgr) NewNote(userID string) (*assistantdbpb.Note, error) {
	uaii, ok := mgr.mapUser.Load(userID)
	if !ok {
		uai, err := mgr.loadUserAssistant(userID)
		if err != nil {
			return nil, err
		}

		mgr.mapUser.Store(userID, uai)

		return mgr.newNote(userID, uai)
	}

	uai, ok := uaii.(*UserAssistantInfo)
	if !ok {
		uai, err := mgr.loadUserAssistant(userID)
		if err != nil {
			return nil, err
		}

		mgr.mapUser.Store(userID, uai)

		return mgr.newNote(userID, uai)
	}

	return mgr.newNote(userID, uai)
}

// newNote - new Note
func (mgr *assistantMgr) newNote(userID string, uai *UserAssistantInfo) (*assistantdbpb.Note, error) {
	mgr.Lock()
	defer mgr.Unlock()

	uai.MaxNoteID = uai.MaxNoteID + 1
	curNoteID := uai.MaxNoteID
	_, err := mgr.updUserAssistant(userID, uai)
	if err != nil {
		return nil, err
	}

	note := &assistantdbpb.Note{
		NoteID:     curNoteID,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	uai.CurNode = note

	return note, nil
}

// Start - start
func (mgr *assistantMgr) Start(ctx context.Context) error {
	return mgr.db.Start(ctx)
}
