package main

import (
	"fmt"

	"github.com/mcyprian/gomsg/pkg/server"
)

func main() {
	config := loadConfig()
	router := server.NewRouter(config.ListLength)
	router.Run(fmt.Sprintf(":%v", config.Port))
}
