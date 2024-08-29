package bot

import (
	"1li/bot/perms"
	"1li/writer"
	"log"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

type fs struct {
	w writer.Writer
}

// Run starts the bot.
func Run(token string, w writer.Writer) error {
	fs := &fs{w: w}

	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		return err
	}

	me, err := b.GetMe(nil)
	if err != nil {
		return err
	}
	log.Printf("Authorized on account %s", me.Username)

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Printf("An error occured: %v", err)
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	}) // TODO: error handler

	dispatcher.AddHandler(handlers.NewCommand("start", cmdStart))
	dispatcher.AddHandler(handlers.NewCommand("sync", fs.cmdSync))
	dispatcher.AddHandler(perms.CMD)
	dispatcher.AddHandler(handlers.NewMessage(message.All, fs.shortURL))

	updater := ext.NewUpdater(dispatcher, nil)

	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		return err
	}

	log.Println("Bot is running")
	updater.Idle()

	return nil
}
