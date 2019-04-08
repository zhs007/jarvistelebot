package plugintranslate

import "sync"

type userinfo struct {
	srcLang     string
	destLang    string
	retranslate bool
}

type groupUserInfo struct {
	mapUser sync.Map
}

func (g *groupUserInfo) clearUser(uid string) {
	g.mapUser.Delete(uid)
}

func (g *groupUserInfo) setUser(uid string, srclang string, destlang string, retranslate bool) {
	ui := &userinfo{
		srcLang:     srclang,
		destLang:    destlang,
		retranslate: retranslate,
	}

	g.mapUser.Store(uid, ui)
}

func (g *groupUserInfo) getUser(uid string) *userinfo {
	v, ok := g.mapUser.Load(uid)
	if ok {
		ui, okt := v.(*userinfo)
		if !okt {
			return nil
		}

		return ui
	}

	return nil
}

type groupInfo struct {
	mapGroup sync.Map
}

func (g *groupInfo) setGroupUser(groupid string, uid string, srclang string, destlang string, retranslate bool) {
	gui := g.getGroup(groupid)
	if gui == nil {
		gui = &groupUserInfo{}
		gui.setUser(uid, srclang, destlang, retranslate)

		g.mapGroup.Store(groupid, gui)

		return
	}

	gui.setUser(uid, srclang, destlang, retranslate)
}

func (g *groupInfo) getGroup(groupid string) *groupUserInfo {
	v, ok := g.mapGroup.Load(groupid)
	if ok {
		gui, okt := v.(*groupUserInfo)
		if !okt {
			return nil
		}

		return gui
	}

	return nil
}

func (g *groupInfo) getGroupUserInfo(groupid string, uid string) *userinfo {
	v, ok := g.mapGroup.Load(groupid)
	if ok {
		gui, okt := v.(*groupUserInfo)
		if !okt {
			return nil
		}

		return gui.getUser(uid)
	}

	return nil
}

func (g *groupInfo) clearGroupUser(groupid string, uid string) {
	gui := g.getGroup(groupid)
	if gui == nil {
		return
	}

	gui.clearUser(uid)
}

func (g *groupInfo) clearGroup(groupid string) {
	g.mapGroup.Delete(groupid)
}
