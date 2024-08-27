package main

import (
	"1li/db/record"
	"1li/db/user"

	// "1li/ent/user"
	"context"
	"fmt"
	"log"
	"strings"

	// _ "github.com/mattn/go-sqlite3"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func CMDStart(update tgbotapi.Update, reply func(text string)) {
	// check if user exists
	if u, err := user.Get(context.Background(), update.Message.From.ID); err == nil {
		reply("Welcome back " + u.Username)
		return
	}

	// create a new user
	if _, err := user.Add(context.Background(), update.Message.From.UserName, update.Message.From.ID); err != nil {
		log.Printf("Error adding user: %v", err)
		reply("Error adding user, try /start again later")
	}

	reply("Welcome to 1li! " + update.Message.From.UserName)

	return
}

func CMDSync(update tgbotapi.Update, reply func(text string)) {
	if err := SyncFromDB(); err != nil {
		log.Printf("Error syncing from db: %v", err)

		reply("Error syncing from db")
	}
	reply("Sync from db done")
}

func CMDOp(update tgbotapi.Update, reply func(text string)) {
	if !user.IsAdmin(context.Background(), update.Message.From.ID) {
		reply("You are not an admin")
		return
	}

	part := strings.SplitN(update.Message.Text, " ", 2)
	if len(part) != 2 {
		reply("Usage: /op <username>")
		return
	}

	if _, err := user.Op(context.Background(), part[1]); err != nil {
		log.Printf("Error setting admin: %v", err)
		reply("Error setting admin")
		return
	}

	reply("Set admin for " + part[1])
}

func CMDDeop(update tgbotapi.Update, reply func(text string)) {
	if !user.IsAdmin(context.Background(), update.Message.From.ID) {
		reply("You are not an admin")
		return
	}

	part := strings.SplitN(update.Message.Text, " ", 2)
	if len(part) != 2 {
		reply("Usage: /deop <username>")
		return
	}

	if _, err := user.Deop(context.Background(), part[1]); err != nil {
		log.Printf("Error setting admin: %v", err)
		reply("Error setting admin")
		return
	}

	reply("Deop for " + part[1])
}

func ShortURL(update tgbotapi.Update, reply func(text string)) {
	u, err := user.Get(context.Background(), update.Message.From.ID)
	if err != nil {
		reply("Please /start first")
		return
	}

	part := strings.SplitN(update.Message.Text, " ", 2)
	target := ""
	code := ""
	if len(part) == 1 {
		code = nonConflictCode(6)
		target = part[0]
	} else {
		if u.CustomCode {
			code = part[0]
			target = part[1]
		} else {
			reply("You are not allowed to use custom code")
			return
		}
	}

	log.Printf("[%s] %s -> %s\n", update.Message.From.UserName, code, target)

	// check is code exists
	if exists, err := record.Exists(context.Background(), code); err != nil {
		log.Printf("Error checking record: %v", err)
		reply("Error checking record")
		return
	} else if exists {
		reply("Code already exists")
		return
	}

	rec, err := record.Add(context.Background(), code, target, update.Message.From.ID)
	if err != nil {
		log.Printf("Error adding record: %v", err)

		reply("Error adding record")
		return
	}

	if err := StaticGenOne(rec); err != nil {
		log.Printf("Error generating static: %v", err)
	}

	reply(fmt.Sprintf("Add a record %s -> %s", code, target))
}

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
				CMDStart(update, reply)
			case "sync":
				CMDSync(update, reply)
			case "op":
				CMDOp(update, reply)
			case "deop":
				CMDDeop(update, reply)
			default:
				reply("I don't know that command")
			}
		} else {
			ShortURL(update, reply)
		}
	}

	return nil
}
