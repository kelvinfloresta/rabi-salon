package main

import (
	"log"
	"rabi-salon/config"
	"rabi-salon/fixtures"
	"rabi-salon/frameworks/http/fiber_adapter"
	"time"
)

func main() {
	time.Local = time.UTC
	server := fiber_adapter.New(fixtures.TestDatabase)

	if err := fixtures.TestDatabase.Start(); err != nil {
		panic(err)
	}

	log.Fatal(server.Start(config.TestPort))
}
