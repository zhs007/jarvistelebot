package chatbot

import (
	"sync"
	"time"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"
)

// FuncOnTimer - func onTimer() error
type FuncOnTimer func() error

// Timer - timer
type Timer struct {
	timerid  int
	timer    int
	lasttime int
	times    int
	info     string
	onTimer  FuncOnTimer
}

// TimerMgr - timer manager
type TimerMgr struct {
	curTimerid int
	mapTimer   sync.Map
}

// newTimerMgr - new TimerMgr
func newTimerMgr() *TimerMgr {
	return &TimerMgr{
		curTimerid: 1,
	}
}

// AddTimer - add timer
func (mgr *TimerMgr) AddTimer(timer int, times int, info string, onTimer FuncOnTimer) int {
	mgr.curTimerid++

	t := &Timer{
		timerid: mgr.curTimerid,
		timer:   timer,
		times:   times,
		info:    info,
		onTimer: onTimer,
	}

	mgr.mapTimer.Store(t.timerid, t)

	return t.timerid
}

// DeleteTimer - delete timer
func (mgr *TimerMgr) DeleteTimer(timerid int) {
	mgr.mapTimer.Delete(timerid)
}

// OnTimer - on timer
func (mgr *TimerMgr) OnTimer() {
	ct := time.Now().Second()

	mgr.mapTimer.Range(func(key, val interface{}) bool {
		t, typeok := val.(*Timer)
		if typeok {
			if ct-t.lasttime >= t.timer {
				err := t.onTimer()
				if err != nil {
					jarvisbase.Warn("TimerMgr.OnTimer:onTimer",
						zap.Error(err),
						zap.String("timer.info", t.info))
				}

				if t.times > 0 {
					t.times--

					if t.times <= 0 {
						mgr.DeleteTimer(t.timerid)

						return true
					}
				}

				t.lasttime += t.timer
			}
		}

		return true
	})
}
