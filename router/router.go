package router

import (
	"github.com/zhs007/jarvistelebot/base"
	"github.com/zhs007/jarvistelebot/controller/telegram"

	"github.com/gin-gonic/gin"
)

// Router -
var Router *gin.Engine

func init() {
	Router = gin.Default()
}

// SetRouter -
func SetRouter() {
	// tpath := base.BuildResPath("./views")
	// base.Debug(tpath)

	// Router.Static("/js", base.BuildResPath("./publish/js"))
	// Router.Static("/css", base.BuildResPath("./publish/css"))

	// Router.LoadHTMLGlob(base.BuildResPath("./views") + "/*.html")

	config := base.GetConfig()

	Router.GET("/", telegramctrl.Index())

	Router.GET("/telegram", telegramctrl.Root())
	Router.POST("/telegram", telegramctrl.Root())

	Router.GET("/telegram/"+config.TelegramBotToken, telegramctrl.Root())
	Router.POST("/telegram/"+config.TelegramBotToken, telegramctrl.Root())
}
