package telegramctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Root -
func Root() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "login")
		// c.HTML(http.StatusOK, "login.html", gin.H{"user": "User"})
	}
}
