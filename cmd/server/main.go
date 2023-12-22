package main

import (
	"log"
	"rabi-salon/config"
	"rabi-salon/frameworks/database/gorm_adapter"
	"rabi-salon/frameworks/http/fiber_adapter"
	"time"
)

func main() {
	time.Local = time.UTC

	db := gorm_adapter.New(config.ProductionDatabase)

	if err := db.Start(); err != nil {
		panic(err)
	}

	httpServer := fiber_adapter.New(db)

	log.Fatal(httpServer.Start(config.Port))
}
