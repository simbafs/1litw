package bot

import (
	"1li/bot/perms"
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

// Run starts the bot.
func Run(token string) error {
	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {
		return err
	}

	me, err := bot.GetMe(nil)
	if err != nil {
		return err
	}
	log.Printf("Authorized on account %s", me.Username)

	dispatcher := ext.NewDispatcher(nil) // TODO: error handler

	dispatcher.AddHandler(handlers.NewCommand("start", cmdStart))
	dispatcher.AddHandler(handlers.NewCommand("sync", cmdSync))
	dispatcher.AddHandler(perms.CMD)
	dispatcher.AddHandler(handlers.NewMessage(message.All, shortURL))

	updater := ext.NewUpdater(dispatcher, nil)

	err = updater.StartPolling(bot, nil)
	if err != nil {
		return err
	}

	log.Println("Bot is running")
	updater.Idle()
	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60
	//
	// updates, err := bot.GetUpdatesChan(u)
	//
	// for update := range updates {
	// 	if update.Message == nil { // ignore any non-Message Updates
	// 		continue
	// 	}
	//
	// 	if update.Message.IsCommand() {
	// 		switch update.Message.Command() {
	// 		case "start":
	// 			CMDStart(update, bot)
	// 		case "sync":
	// 			CMDSync(update, bot)
	// 		case "op":
	// 			CMDOp(update, bot)
	// 		case "deop":
	// 			CMDDeop(update, bot)
	// 		default:
	// 			reply(update.Message, bot)("I don't know that command")
	// 		}
	// 	} else {
	// 		ShortURL(update, bot)
	// 	}
	// }

	return nil
}
