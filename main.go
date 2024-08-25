package main

import (
	"1li/db"
	"log"
	"os"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("failed initializing database: %v", err)
	}

	if err := initServer("static"); err != nil {
		log.Fatalf("failed initializing server: %v", err)
	}

	go func() {
		log.Fatal(fileServer(":3000", os.DirFS("static")))
	}()

	log.Fatal(tgBot("7199207337:AAG_X5KQrXUkqtw_IeLtTe6CUAF2ZTNkLzA"))
}
