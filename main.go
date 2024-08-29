package main

import (
	"1li/bot"
	"1li/db"
	"1li/writer"
	"log"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("failed initializing database: %v", err)
	}

	cfg := FromEnv()

	w := writer.NewGitHub(cfg.GitHubToken, cfg.User, cfg.Repo, cfg.Branch)

	log.Fatal(bot.Run(cfg.TgToken, w))
}
