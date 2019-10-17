package server

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

// Message represents stored messages
type Message struct {
	Msg string `json:"msg" binding:"required"`
	Ts  int64  `json:"ts" binding:"required"`
}

// healthz is a health check endpoint
func healthz(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// addMessage stores received message to the datastore
func addMessage(c *gin.Context) {
	var json Message
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "nok"})
		return
	}
	db.Insert(json)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// getAllMessages returns last messages sorted by timestamp
func getAllMessages(c *gin.Context) {
	data := db.GetAll()
	sort.SliceStable(data, func(i, j int) bool {
		if data[i] == nil {
			return false
		}
		if data[j] == nil {
			return true
		}
		return data[i].(Message).Ts < data[j].(Message).Ts
	})
	c.JSON(http.StatusOK, gin.H{"data": data})
}
