package main

import (
	"os"
	"strconv"
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
	port := defaultPort
	listLength := defaultListLength
	portStr, ok := os.LookupEnv("GOMSG_PORT")
	if ok {
		convertedPort, err := strconv.ParseInt(portStr, 10, 32)
		if err == nil {
			port = int(convertedPort)
		}
	}
	listLengthStr, ok := os.LookupEnv("LIST_LENGTH")
	if ok {
		convertedListLength, err := strconv.ParseInt(listLengthStr, 10, 32)
		if err == nil {
			listLength = int(convertedListLength)
		}
	}
	return &config{Port: port, ListLength: listLength}
}
