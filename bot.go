package main

import (
	"1li/db"
	"1li/ent/user"
	"context"
	"fmt"
	"log"
	"strings"

	// _ "github.com/mattn/go-sqlite3"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func tgBot(token string) error {
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

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				// check if user exists
				u, err := db.Client.User.Query().Where(user.Tgid(update.Message.From.ID)).Only(context.Background())
				if err == nil {
					reply("Welcome back " + u.Username)
					continue
				}

				// create a new user
				_, err = db.Client.User.Create().
					SetTgid(update.Message.From.ID).
					SetUsername(update.Message.From.UserName).
					Save(context.Background())
				if err != nil {
					log.Printf("Error adding user: %v", err)

					reply("Error adding user, try /start again later")
				}

				reply("Welcome to 1li! " + update.Message.From.UserName)
			case "sync":
				if err := SyncFromDB(); err != nil {
					log.Printf("Error syncing from db: %v", err)

					reply("Error syncing from db")
				}
				reply("Sync from db done")
			}
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

		rec, err := db.AddUser(context.Background(), code, target, update.Message.From.ID)
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
