package assistant

import (
	"context"
	"strconv"
	"strings"
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

// keyInfoMap - keyinfo map
type keyInfoMap struct {
	mapKeyInfo map[string]*assistantdbpb.KeyInfo
}

func newKeyInfoMap() *keyInfoMap {
	return &keyInfoMap{
		mapKeyInfo: make(map[string]*assistantdbpb.KeyInfo),
	}
}

func (m *keyInfoMap) addKey(key string, noteid int64) {
	v, ok := m.mapKeyInfo[key]
	if !ok {
		m.mapKeyInfo[key] = &assistantdbpb.KeyInfo{
			NoteIDs: []int64{noteid},
		}

		return
	}

	v.NoteIDs = append(v.NoteIDs, noteid)
}

// UserAssistantInfo - user's assistant info
type UserAssistantInfo struct {
	MaxNoteID   int64               `json:"maxNoteID"`
	Keys        []string            `json:"keys"`
	CurNote     *assistantdbpb.Note `json:"-"`
	CurNoteMode int                 `json:"-"`
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
	GetUserAssistantInfo(userID string) (*UserAssistantInfo, error)
	// GetNote - get note
	GetNote(userID string, noteID int64) (*assistantdbpb.Note, error)
	// RmNote - remove note
	RmNote(userID string, noteID int64) (string, error)
	// UpdNote - update user's note
	UpdNote(userID string, note *assistantdbpb.Note) (*assistantdbpb.Note, error)

	// RebuildKeys - rebuild note keywords
	RebuildKeys(userID string) (int64, int, error)
	// HasKeyword - has key
	HasKeyword(userID string, key string) bool
	// FindNoteWithKeyword - find note with keyword
	FindNoteWithKeyword(userID string, key string) ([]*assistantdbpb.Note, error)

	// Export - export all data
	Export(userID string) ([](map[string]interface{}), error)
	// Import - import all data
	Import(userID string, arr [](map[string]interface{})) error
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

// getUserAssistantInfo
func (mgr *assistantMgr) getUserAssistantInfo(userID string) *UserAssistantInfo {
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
	uai := mgr.getUserAssistantInfo(userID)
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

// insertKey - insert key to UserAssistantInfo
func (mgr *assistantMgr) insertKey(uai *UserAssistantInfo, key string) {
	for _, v := range uai.Keys {
		if v == key {
			return
		}
	}

	uai.Keys = append(uai.Keys, key)
}

// SaveCurNote - save current note
func (mgr *assistantMgr) SaveCurNote(userID string) (*assistantdbpb.Note, error) {
	uai := mgr.getUserAssistantInfo(userID)
	if uai == nil {
		return nil, ErrNoUserAssistantInfo
	}

	if uai.CurNote == nil {
		return nil, ErrNoCurNote
	}

	note, err := mgr.db.UpdNote(userID, uai.CurNote)
	if err != nil {
		return nil, err
	}

	for _, v := range note.Keys {
		mgr.insertKey(uai, v)
	}

	_, err = mgr.updUserAssistant(userID, uai)
	if err != nil {
		return nil, err
	}

	return note, nil
}

// AddCurNoteData - save current note
func (mgr *assistantMgr) AddCurNoteData(userID string, dat string) (*assistantdbpb.Note, error) {
	uai := mgr.getUserAssistantInfo(userID)
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
	uai := mgr.getUserAssistantInfo(userID)
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
	uai := mgr.getUserAssistantInfo(userID)
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
	uai := mgr.getUserAssistantInfo(userID)
	if uai == nil {
		return ModeInvalidType
	}

	if uai.CurNote == nil {
		return ModeInvalidType
	}

	return uai.CurNoteMode
}

// GetUserAssistantInfo - get user assistant info
func (mgr *assistantMgr) GetUserAssistantInfo(userID string) (*UserAssistantInfo, error) {
	uai := mgr.getUserAssistantInfo(userID)
	if uai == nil {
		uai, err := mgr.loadUserAssistant(userID)
		if err != nil {
			return nil, err
		}

		return uai, nil
	}

	return uai, nil
}

// GetNote - get note
func (mgr *assistantMgr) GetNote(userID string, noteID int64) (*assistantdbpb.Note, error) {
	return mgr.db.GetNote(userID, noteID)
}

// RebuildKeys - rebuild note keywords
func (mgr *assistantMgr) RebuildKeys(userID string) (int64, int, error) {
	mapKeyInfo := newKeyInfoMap()
	uai, err := mgr.GetUserAssistantInfo(userID)
	if err != nil {
		return 0, 0, err
	}

	var i int64
	for i = 1; i <= uai.MaxNoteID; i++ {
		note, e := mgr.GetNote(userID, i)
		if e != nil {
			err = e

			continue
		}
		if len(note.Data) > 0 {
			for _, v := range note.Keys {
				mapKeyInfo.addKey(v, i)
			}
		}
	}

	uai.Keys = nil
	keynums := 0
	for k, v := range mapKeyInfo.mapKeyInfo {
		_, e := mgr.db.UpdKeyInfo(userID, k, v)
		if e != nil {
			err = e
		}

		uai.Keys = append(uai.Keys, k)

		keynums++
	}

	return uai.MaxNoteID, keynums, err
}

// HasKeyword - has key
func (mgr *assistantMgr) HasKeyword(userID string, key string) bool {
	uai, err := mgr.GetUserAssistantInfo(userID)
	if err != nil {
		return false
	}

	for _, v := range uai.Keys {
		if v == key {
			return true
		}
	}

	return false
}

// FindNoteWithKeyword - find note with keyword
func (mgr *assistantMgr) FindNoteWithKeyword(userID string, key string) ([]*assistantdbpb.Note, error) {
	var lst []*assistantdbpb.Note

	ki, err := mgr.db.GetKeyInfo(userID, key)
	if err != nil {
		return nil, err
	}

	for _, v := range ki.NoteIDs {
		note, e := mgr.GetNote(userID, v)
		if e != nil {
			err = e

			continue
		}

		if len(note.Data) > 0 {
			lst = append(lst, note)
		}
	}

	return lst, nil
}

// Export - export all data
func (mgr *assistantMgr) Export(userID string) ([](map[string]interface{}), error) {
	var arr [](map[string]interface{})
	uai, err := mgr.GetUserAssistantInfo(userID)
	if err != nil {
		return nil, err
	}

	var i int64
	for i = 1; i <= uai.MaxNoteID; i++ {
		note, e := mgr.GetNote(userID, i)
		if e != nil {
			err = e

			continue
		}

		if len(note.Data) > 0 {
			co := make(map[string]interface{})

			for j, v := range note.Data {
				co["data"+strconv.Itoa(j)] = v
			}

			for j, v := range note.Keys {
				co["key"+strconv.Itoa(j)] = v
			}

			arr = append(arr, co)
		}
	}

	return arr, nil
}

// Import - import all data
func (mgr *assistantMgr) Import(userID string, arr [](map[string]interface{})) error {
	uai, err := mgr.GetUserAssistantInfo(userID)
	if err != nil {
		return err
	}

	var cn int64
	cn = int64(len(arr))
	if cn < uai.MaxNoteID {
		for i := cn + 1; i <= uai.MaxNoteID; i++ {
			mgr.RmNote(userID, i)
		}
	}

	uai.MaxNoteID = cn
	_, err = mgr.updUserAssistant(userID, uai)
	if err != nil {
		return err
	}

	for i, v := range arr {
		note := &assistantdbpb.Note{
			NoteID:     int64(i) + 1,
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
		}

		for k, ov := range v {
			if strings.Contains(k, "key") {
				note.Keys = append(note.Keys, ov.(string))
			} else if strings.Contains(k, "data") {
				note.Data = append(note.Data, ov.(string))
			}
		}

		note, err := mgr.db.UpdNote(userID, note)
		if err != nil {
			return err
		}
	}

	_, _, err = mgr.RebuildKeys(userID)
	if err != nil {
		return err
	}

	return nil
}

// RmNote - remove note
func (mgr *assistantMgr) RmNote(userID string, noteID int64) (string, error) {
	uai := mgr.getUserAssistantInfo(userID)
	if uai == nil {
		return "", ErrNoUserAssistantInfo
	}

	key, err := mgr.db.RmNote(userID, noteID)
	if err != nil {
		return "", err
	}

	return key, nil
}

// UpdNote - update user's note
func (mgr *assistantMgr) UpdNote(userID string, note *assistantdbpb.Note) (*assistantdbpb.Note, error) {
	uai := mgr.getUserAssistantInfo(userID)
	if uai == nil {
		return nil, ErrNoUserAssistantInfo
	}

	note, err := mgr.db.UpdNote(userID, note)
	if err != nil {
		return nil, err
	}

	for _, v := range note.Keys {
		mgr.insertKey(uai, v)
	}

	_, err = mgr.updUserAssistant(userID, uai)
	if err != nil {
		return nil, err
	}

	return note, nil
}
