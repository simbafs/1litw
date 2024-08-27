package bot

import (
	"1li/db/record"
	"1li/db/user"
	"1li/ssg"
	"1li/util"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func reply(message *tgbotapi.Message, bot *tgbotapi.BotAPI) func(text string) {
	return func(text string) {
		msg := tgbotapi.NewMessage(message.Chat.ID, text)
		msg.ReplyToMessageID = message.MessageID

		bot.Send(msg)
	}
}

func cmdStart(b *gotgbot.Bot, ctx *ext.Context) error {
	if u, err := user.Get(context.Background(), int(ctx.Message.From.Id)); err == nil {
		ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Welcome back %s", u.Username), nil)
		return nil
	}

	if _, err := user.Add(context.Background(), ctx.Message.From.Username, int(ctx.Message.From.Id)); err != nil {
		log.Printf("Error adding user: %v", err)
		ctx.EffectiveMessage.Reply(b, "Error adding user, try /start again later", nil)
	}

	ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Welcome to 1li! %s", ctx.Message.From.Username), nil)

	return nil
}

func cmdSync(b *gotgbot.Bot, ctx *ext.Context) error {
	if err := ssg.SyncFromDB(); err != nil {
		log.Printf("Error syncing from db: %v", err)

		ctx.EffectiveMessage.Reply(b, "Error syncing from db", nil)
		return nil
	}
	ctx.EffectiveMessage.Reply(b, "Sync from db done", nil)
	return nil
}

func CMDOp(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	reply := reply(update.Message, bot)
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

func CMDDeop(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	reply := reply(update.Message, bot)
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

func ShortURL(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	reply := reply(update.Message, bot)
	u, err := user.Get(context.Background(), update.Message.From.ID)
	if err != nil {
		reply("Please /start first")
		return
	}

	part := strings.SplitN(update.Message.Text, " ", 2)
	target := ""
	code := ""
	if len(part) == 1 {
		code = util.NonConflictCode(6)
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

	if !util.IsValidURL(target) {
		reply("Target is not a valid URL")
		return
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

	if err := ssg.StaticGenOne(rec); err != nil {
		log.Printf("Error generating static: %v", err)
	}

	reply(fmt.Sprintf("Add a record %s -> %s", code, target))
}
