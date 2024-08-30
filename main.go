package main

import (
	"1li/bot"
	"1li/config"
	"1li/db"
	"1li/writer"
	"log"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("failed initializing database: %v", err)
	}

	cfg := config.FromEnv()

	local := writer.Local{
		Base: cfg.Base,
	}

	go log.Fatal(local.ListenAndServe(cfg.Addr))

	log.Fatal(bot.Run(cfg, local))
}
