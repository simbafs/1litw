package main

import (
	"1li/bot"
	"1li/db"
	"1li/fileserver"
	"log"
	"os"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("failed initializing database: %v", err)
	}

	if err := fileserver.Init("static"); err != nil {
		log.Fatalf("failed initializing server: %v", err)
	}

	go func() {
		log.Fatal(fileserver.ListenAndServe(":3000", os.DirFS("static")))
	}()

	log.Fatal(bot.Run("7199207337:AAG_X5KQrXUkqtw_IeLtTe6CUAF2ZTNkLzA"))
}
