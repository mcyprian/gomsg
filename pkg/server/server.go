package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mcyprian/gomsg/pkg/datastore"
)

var db datastore.MsgDatastore

// NewRouter set up the router for gomsg server
func NewRouter(listLenght int) *gin.Engine {
	router := gin.Default()
	if db == nil {
		db = datastore.NewMsgMemoryBuffer(listLenght)
	}
	router.GET("/healthz", healthz)
	router.POST("/message", addMessage)
	router.GET("/messages", getAllMessages)

	return router
}
