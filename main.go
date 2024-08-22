package main

import (
	"1li/ent"
	"context"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"

	// _ "modernc.org/sqlite"

	_ "github.com/mattn/go-sqlite3"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func tgBot(token string, client *ent.Client) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		reply := func(text string) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}

		if update.Message.IsCommand() && update.Message.Command() == "start" {
			// create a new user
			ctx := context.Background()
			_, err := client.User.Create().
				SetTgid(update.Message.From.ID).
				SetUsername(update.Message.From.UserName).
				Save(ctx)
			if err != nil {
				log.Printf("Error adding user: %v", err)

				reply("Error adding user, try /start again later")
			}

			reply("Welcome to 1li! " + update.Message.From.UserName)

			continue
		}

		part := strings.SplitN(update.Message.Text, " ", 2)
		target := ""
		code := ""
		if len(part) == 1 {
			code = nonConflictCode(6)
			target = part[0]
		} else {
			code = part[0]
			target = part[1]
		}

		log.Printf("[%s] %s -> %s\n", update.Message.From.UserName, code, target)

		ctx := context.Background()
		rec, err := addRecord(client, ctx, code, target, update.Message.From.ID)
		if err != nil {
			log.Printf("Error adding record: %v", err)

			reply("Error adding record")

			continue
		}

		if err := StaticGenOne(rec); err != nil {
			log.Printf("Error generating static: %v", err)
		}

		reply(fmt.Sprintf("Add a record %s -> %s", code, target))
	}

	return nil
}

// fileServer serves files from the provided fs.FS.
func fileServer(addr string, files fs.FS) error {
	http.Handle("/", http.FileServer(http.FS(files)))
	return http.ListenAndServe(addr, nil)
}

func main() {
	client, err := ent.Open("sqlite3", "file:./ent.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Fatal(tgBot("7199207337:AAG_X5KQrXUkqtw_IeLtTe6CUAF2ZTNkLzA", client))
}
