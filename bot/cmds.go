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
	if u, err := user.Get(context.Background(), ctx.Message.From.Id); err == nil {
		ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Welcome back %s", u.Username), nil)
		return nil
	}

	if _, err := user.Add(context.Background(), ctx.Message.From.Username, ctx.Message.From.Id); err != nil {
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

func shortURL(b *gotgbot.Bot, ctx *ext.Context) error {
	u, err := user.Get(context.Background(), ctx.Message.From.Id)
	if err != nil {
		ctx.EffectiveMessage.Reply(b, "Please /start first", nil)
		return nil
	}

	if ok, err := user.GetPerm(context.Background(), ctx.EffectiveUser.Id, "create"); err != nil {
		log.Printf("Error getting permission: %v", err)
		ctx.EffectiveMessage.Reply(b, "查詢權限時發生錯誤", nil)
		return nil
	} else if !ok {
		ctx.EffectiveMessage.Reply(b, "權限不足", nil)
		return nil
	}

	part := strings.SplitN(ctx.Message.Text, " ", 2)
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
			ctx.EffectiveMessage.Reply(b, "你沒有權限自訂短網址", nil)
			return nil
		}
	}

	if !util.IsValidURL(target) {
		ctx.EffectiveMessage.Reply(b, "Target is not a valid URL", nil)
		return nil
	}

	log.Printf("[%s] %s -> %s\n", ctx.Message.From.Username, code, target)

	// check is code exists
	if exists, err := record.Exists(context.Background(), code); err != nil {
		log.Printf("Error checking record: %v", err)
		ctx.EffectiveMessage.Reply(b, "Error checking record", nil)
		return nil
	} else if exists {
		ctx.EffectiveMessage.Reply(b, "Code already exists", nil)
		return nil
	}

	rec, err := record.Add(context.Background(), code, target, ctx.Message.From.Id)
	if err != nil {
		log.Printf("Error adding record: %v", err)

		ctx.EffectiveMessage.Reply(b, "Error adding record", nil)
		return nil
	}

	if err := ssg.StaticGenOne(rec); err != nil {
		log.Printf("Error generating static: %v", err)
	}

	ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Add a record %s -> %s", code, target), nil)
	return nil
}
