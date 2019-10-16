package server

import "github.com/gin-gonic/gin"

func Run() {
	r := gin.Default()
	r.GET("/healthz", healthz)
	r.Run()
}
