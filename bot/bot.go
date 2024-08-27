package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Run starts the bot.
func Run(token string) error {
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

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				CMDStart(update, bot)
			case "sync":
				CMDSync(update, bot)
			case "op":
				CMDOp(update, bot)
			case "deop":
				CMDDeop(update, bot)
			default:
				reply(update.Message, bot)("I don't know that command")
			}
		} else {
			ShortURL(update, bot)
		}
	}

	return nil
}
