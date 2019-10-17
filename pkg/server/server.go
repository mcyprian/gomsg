package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mcyprian/gomsg/pkg/datastore"
)

var db datastore.MsgDatastore

func Run(port int, listLenght int) {
	r := gin.Default()
	if db == nil {
		db = datastore.NewMsgMemoryBuffer(listLenght)
	}
	r.GET("/healthz", healthz)
	r.POST("/message", addMessage)
	r.GET("/messages", getAllMessages)

	r.Run(fmt.Sprintf(":%v", port))
}
