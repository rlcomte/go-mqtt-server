package main

import (
	"fmt"

	service "github.com/mdzio/go-mqtt/service"
)

func main() {
	// Create a new server
	svr := &service.Server{
		KeepAlive:        300,               // seconds
		ConnectTimeout:   2,                 // seconds
		SessionsProvider: "mem",             // keeps sessions in memory
		Authenticator:    "mockSuccess",     // always succeed
		TopicsProvider:   "mem",             // keeps topic subscriptions in memory
	}

	// Listen and serve connections at localhost:18883
	fmt.Printf("Listening and serving connnections at localhost:18883\n")
	err := svr.ListenAndServe("tcp://:18883")

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}