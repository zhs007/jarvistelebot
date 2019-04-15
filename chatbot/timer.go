package chatbot

import (
	"context"
	"sync"
	"time"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"
)

// FuncOnTimer - func onTimer(ctx) error
type FuncOnTimer func(ctx context.Context) error

// Timer - timer
type Timer struct {
	timerid  int
	timer    int
	lasttime int64
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
		timerid:  mgr.curTimerid,
		timer:    timer,
		times:    times,
		info:     info,
		onTimer:  onTimer,
		lasttime: time.Now().Unix(),
	}

	mgr.mapTimer.Store(t.timerid, t)

	return t.timerid
}

// DeleteTimer - delete timer
func (mgr *TimerMgr) DeleteTimer(timerid int) {
	mgr.mapTimer.Delete(timerid)
}

// OnTimer - on timer
func (mgr *TimerMgr) OnTimer(ctx context.Context) {
	ct := time.Now().Unix()

	// jarvisbase.Info("TimerMgr.OnTimer",
	// 	zap.Int64("curtime", ct))

	mgr.mapTimer.Range(func(key, val interface{}) bool {
		t, typeok := val.(*Timer)
		if typeok {
			// jarvisbase.Info("TimerMgr.OnTimer:range",
			// 	zap.Int64("lasttime", t.lasttime),
			// 	zap.Int("timer", t.timer))

			if ct-t.lasttime >= int64(t.timer) {
				err := t.onTimer(ctx)
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

				t.lasttime += int64(t.timer)
			}
		}

		return true
	})
}
