package main

import (
	"os"
	"strconv"

	"github.com/jinzhu/copier"
)

const (
	defaultPort       = 8080
	defaultListLength = 100
)

type config struct {
	ListLength int `json:"msg"`
	Port       int `json:"port"`
}

func loadConfig() *config {
	var err Error
	port, ok := os.LookupEnv("GOMSG_PORT")
	if !ok {
		port = defaultPort
	} else {
		port, err = strconv.ParseInt(port, 10, 32)
		if err != nil {
			port = defaultPort
		}
	}
	listLength, ok := os.LookupEnv("LIST_LENGTH")
	if ok {
		listLength, err = strconv.ParseInt(port, 10, 32)
		if err != nil {
			listLength = defaultListLength
		}
	} else {
		listLength = defaultListLength
	}
	return &config{Port: port, ListLength: listLength}
}
