package telebot

import (
	"fmt"

	"github.com/zhs007/jarvistelebot/chatbot"
)

// telebotlog - BotLogger
type telebotlog struct {
}

func (log *telebotlog) Println(v ...interface{}) {
	str := fmt.Sprintln(v...)
	chatbot.Info(str)
}

func (log *telebotlog) Printf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	chatbot.Info(str)
}
