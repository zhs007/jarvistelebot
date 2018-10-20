package main

import (
	"fmt"

	"github.com/zhs007/jarvistelebot/base"
)

// telebotlog - BotLogger
type telebotlog struct {
}

func (log *telebotlog) Println(v ...interface{}) {
	str := fmt.Sprintln(v...)
	base.Info(str)
}

func (log *telebotlog) Printf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	base.Info(str)
}
