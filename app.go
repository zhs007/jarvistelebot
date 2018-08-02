package main

import (
	"github.com/zhs007/jarvistelebot/router"
)

// StartServ -
func StartServ(servaddr string) {
	r := router.Router
	router.SetRouter()
	r.Run(servaddr)
}
