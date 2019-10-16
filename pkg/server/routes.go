package server

import "github.com/gin-gonic/gin"

// healthz is a health check endpoint
func healthz(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
