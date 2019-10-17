package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mcyprian/gomsg/pkg/datastore"
)

var db datastore.MsgDatastore

func Run() {
	r := gin.Default()
	if db == nil {
		db = datastore.NewMsgMemoryBuffer(100)
	}
	r.GET("/healthz", healthz)
	r.POST("/message", addMessage)
	r.GET("/messages", getAllMessages)

	r.Run()
}
