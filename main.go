package main

import (
	"1li/bot"
	"1li/config"
	"1li/db"
	"1li/writer"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("1li bot started")
	if err := db.InitDB(); err != nil {
		log.Fatalf("failed initializing database: %v", err)
	}

	cfg := config.FromEnv()

	local := writer.Local{
		Base: cfg.Base,
	}

	const maxRetries = 10
	const retryDelay = 10 * time.Second

	go func() {
		for i := 0; i < maxRetries; i++ {
			err := local.ListenAndServe(cfg.Addr)
			if err == nil {
				return
			}

			log.Printf("Server error: %v", err)

			if i < maxRetries-1 {
				log.Printf("Retrying in %v...", retryDelay)
				time.Sleep(retryDelay)
			} else {
				log.Fatalf("Failed to start server after %v retries", maxRetries)
			}
		}
	}()

	log.Fatal(bot.Run(cfg, local))
}
