package assistant

import (
	"context"
	"sync"
	"time"

	"github.com/zhs007/jarvistelebot/assistantdb"

	"github.com/zhs007/jarvistelebot/assistantdb/proto"
)

const (
	// ModeInvalidType - invalid mode
	ModeInvalidType = 0
	// ModeInputData - input note data
	ModeInputData = 1
	// ModeInputKey - input note key
	ModeInputKey = 2
)

// UserAssistantInfo - user's assistant info
type UserAssistantInfo struct {
	MaxNoteID   int64    `json:"maxNoteID"`
	Keys        []string `json:"keys"`
	CurNote     *assistantdbpb.Note
	CurNoteMode int
}

// Mgr - assistant manager
type Mgr interface {
	// Start - start
	Start(ctx context.Context) error

	// NewNote - new Note
	NewNote(userID string) (*assistantdbpb.Note, error)
	// SaveCurNote - save current note
	SaveCurNote(userID string) (*assistantdbpb.Note, error)
	// AddCurNoteData - save current note
	AddCurNoteData(userID string, dat string) (*assistantdbpb.Note, error)
	// AddCurNoteKey - save current key
	AddCurNoteKey(userID string, key string) (*assistantdbpb.Note, error)
	// ChgCurNoteMode - change current note mode, mode likes ModeInputData or ModeInputKey
	ChgCurNoteMode(userID string, mode int) error
	// GetCurNoteMode - get current note mode
	GetCurNoteMode(userID string) int
	// GetUserAssistantInfo - get user assistant info
	GetUserAssistantInfo(userID string) *UserAssistantInfo
}

// assistantMgr - assistant manager
type assistantMgr struct {
	sync.Mutex

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

// GetUserAssistantInfo
func (mgr *assistantMgr) GetUserAssistantInfo(userID string) *UserAssistantInfo {
	uaii, ok := mgr.mapUser.Load(userID)
	if !ok {
		return nil
	}

	uai, ok := uaii.(*UserAssistantInfo)
	if !ok {
		return nil
	}

	return uai
}

// NewNote - new Note
func (mgr *assistantMgr) NewNote(userID string) (*assistantdbpb.Note, error) {
	uai := mgr.GetUserAssistantInfo(userID)
	if uai == nil {
		uai, err := mgr.loadUserAssistant(userID)
		if err != nil {
			return nil, err
		}

		mgr.mapUser.Store(userID, uai)

		return mgr.newNote(userID, uai)
	}

	// uaii, ok := mgr.mapUser.Load(userID)
	// if !ok {
	// 	uai, err := mgr.loadUserAssistant(userID)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	mgr.mapUser.Store(userID, uai)

	// 	return mgr.newNote(userID, uai)
	// }

	// uai, ok := uaii.(*UserAssistantInfo)
	// if !ok {
	// 	uai, err := mgr.loadUserAssistant(userID)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	mgr.mapUser.Store(userID, uai)

	// 	return mgr.newNote(userID, uai)
	// }

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

	uai.CurNote = note
	uai.CurNoteMode = ModeInputData

	return note, nil
}

// Start - start
func (mgr *assistantMgr) Start(ctx context.Context) error {
	return mgr.db.Start(ctx)
}

// SaveCurNote - save current note
func (mgr *assistantMgr) SaveCurNote(userID string) (*assistantdbpb.Note, error) {
	uai := mgr.GetUserAssistantInfo(userID)
	if uai == nil {
		return nil, ErrNoUserAssistantInfo
	}

	if uai.CurNote == nil {
		return nil, ErrNoCurNote
	}

	return mgr.db.UpdNote(userID, uai.CurNote)
}

// AddCurNoteData - save current note
func (mgr *assistantMgr) AddCurNoteData(userID string, dat string) (*assistantdbpb.Note, error) {
	uai := mgr.GetUserAssistantInfo(userID)
	if uai == nil {
		return nil, ErrNoUserAssistantInfo
	}

	if uai.CurNote == nil {
		return nil, ErrNoCurNote
	}

	uai.CurNote.Data = append(uai.CurNote.Data, dat)

	return nil, nil
}

// AddCurNoteKey - save current key
func (mgr *assistantMgr) AddCurNoteKey(userID string, key string) (*assistantdbpb.Note, error) {
	uai := mgr.GetUserAssistantInfo(userID)
	if uai == nil {
		return nil, ErrNoUserAssistantInfo
	}

	if uai.CurNote == nil {
		return nil, ErrNoCurNote
	}

	uai.CurNote.Keys = append(uai.CurNote.Keys, key)

	return nil, nil
}

// ChgCurNoteMode - change current note mode, mode likes ModeInputData or ModeInputKey
func (mgr *assistantMgr) ChgCurNoteMode(userID string, mode int) error {
	uai := mgr.GetUserAssistantInfo(userID)
	if uai == nil {
		return ErrNoUserAssistantInfo
	}

	if uai.CurNote == nil {
		return ErrNoCurNote
	}

	uai.CurNoteMode = mode

	return nil
}

// GetCurNoteMode - get current note mode
func (mgr *assistantMgr) GetCurNoteMode(userID string) int {
	uai := mgr.GetUserAssistantInfo(userID)
	if uai == nil {
		return ModeInvalidType
	}

	if uai.CurNote == nil {
		return ModeInvalidType
	}

	return uai.CurNoteMode
}
