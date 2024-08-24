package main

import (
	"1li/ent"
	"context"
	"log"
	"os"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	RegisterSqlite3Driver()

	client, err := ent.Open("sqlite3", "./1litw.sqlite?_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := initServer("static"); err != nil {
		log.Fatalf("failed initializing server: %v", err)
	}

	go func() {
		log.Fatal(fileServer(":3000", os.DirFS("static")))
	}()

	log.Fatal(tgBot("7199207337:AAG_X5KQrXUkqtw_IeLtTe6CUAF2ZTNkLzA", client))
}
