package main

import "github.com/mcyprian/gomsg/pkg/server"

func main() {
	config := loadConfig()
	server.Run(config.Port, config.ListLength)
}
