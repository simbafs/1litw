package bot

import (
	"1li/bot/msg"
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
)

func cmdStart(b *gotgbot.Bot, ctx *ext.Context) error {
	if u, err := user.Get(context.Background(), ctx.Message.From.Id); err == nil {
		ctx.EffectiveMessage.Reply(b, fmt.Sprintf(msg.WellcomeBack, u.Username), nil)
		return nil
	}

	if _, err := user.Add(context.Background(), ctx.Message.From.Username, ctx.Message.From.Id); err != nil {
		log.Printf("Error adding user: %v", err)
		ctx.EffectiveMessage.Reply(b, msg.ServerError, nil)
	}

	ctx.EffectiveMessage.Reply(b, fmt.Sprintf(msg.Wellcome, ctx.Message.From.Username), nil)

	return nil
}

func (bot *bot) cmdSync(b *gotgbot.Bot, ctx *ext.Context) error {
	if err := ssg.SyncFromDB(bot.w); err != nil {
		log.Printf("Error syncing from db: %v", err)

		ctx.EffectiveMessage.Reply(b, msg.ServerError, nil)
		return nil
	}
	ctx.EffectiveMessage.Reply(b, msg.Done, nil)
	return nil
}

func (bot *bot) shortURL(b *gotgbot.Bot, ctx *ext.Context) error {
	u, err := user.Get(context.Background(), ctx.Message.From.Id)
	if err != nil {
		ctx.EffectiveMessage.Reply(b, msg.Register, nil)
		return nil
	}

	if ok, err := user.GetPerm(context.Background(), ctx.EffectiveUser.Id, "create"); err != nil {
		log.Printf("Error getting permission: %v", err)
		ctx.EffectiveMessage.Reply(b, msg.ServerError, nil)
		return nil
	} else if !ok {
		ctx.EffectiveMessage.Reply(b, msg.PermissionDenied, nil)
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
			ctx.EffectiveMessage.Reply(b, msg.PermissionDeniedCustomCode, nil)
			return nil
		}
	}

	if !util.IsValidURL(target) {
		ctx.EffectiveMessage.Reply(b, msg.InvalidURL, nil)
		return nil
	}

	log.Printf("[%s] %s -> %s\n", ctx.Message.From.Username, code, target)

	// check is code exists
	if exists, err := record.Exists(context.Background(), code); err != nil {
		log.Printf("Error checking record: %v", err)
		ctx.EffectiveMessage.Reply(b, msg.ServerError, nil)
		return nil
	} else if exists {
		ctx.EffectiveMessage.Reply(b, msg.URLExist, nil)
		return nil
	}

	rec, err := record.Add(context.Background(), code, target, ctx.Message.From.Id)
	if err != nil {
		log.Printf("Error adding record: %v", err)

		ctx.EffectiveMessage.Reply(b, msg.ServerError, nil)
		return nil
	}

	if err := ssg.StaticGenOne(rec, bot.w.SetCode(rec.Code)); err != nil {
		log.Printf("Error generating static: %v", err)
	}

	ctx.EffectiveMessage.Reply(b, fmt.Sprintf(msg.ShortURL, bot.Config.Origin, code, target), nil)
	return nil
}
