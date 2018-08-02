package router

import (
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

	Router.GET("/telegram/612523255:AAHDo94TU3wZVfYwBwbZpYlUjeeSGyJ3o-c", telegramctrl.Root())
	Router.POST("/telegram/612523255:AAHDo94TU3wZVfYwBwbZpYlUjeeSGyJ3o-c", telegramctrl.Root())
}
