package perms

// TODO: 
// 1. 拆分 response 和他們動作 ex 設定鍵盤
// 2. 除存狀態

import (
	"fmt"
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/conversation"
)

const (
	USER  = "user"
	PERM  = "perm"
	VALUE = "value"
)

var clearKeyboard = &gotgbot.SendMessageOpts{
	ReplyMarkup: &gotgbot.ReplyKeyboardRemove{
		RemoveKeyboard: true,
	},
}

var CMD = handlers.NewConversation(
	[]ext.Handler{handlers.NewCommand("perms", user)},
	map[string][]ext.Handler{
		USER:  {handlers.NewCommand("perms", user)},
		PERM:  {handlers.NewMessage(isUserShared, perm)},
		VALUE: {handlers.NewCallback(isPerms, value)},
	},
	&handlers.ConversationOpts{
		Exits:        []ext.Handler{handlers.NewCommand("exit", exit)},
		StateStorage: conversation.NewInMemoryStorage(conversation.KeyStrategySenderAndChat),
		AllowReEntry: true,
	},
)

// isuserShared checks if the message is a shared user request.
func isUserShared(msg *gotgbot.Message) bool {
	return msg.UsersShared != nil && msg.UsersShared.RequestId == msg.Chat.Id
}

// user is a response to the state USER.
func user(b *gotgbot.Bot, ctx *ext.Context) error {
	keyboard := gotgbot.ReplyKeyboardMarkup{
		Keyboard: [][]gotgbot.KeyboardButton{
			{gotgbot.KeyboardButton{
				Text: "選擇目標使用者",
				RequestUsers: &gotgbot.KeyboardButtonRequestUsers{
					RequestId: ctx.EffectiveChat.Id,
				},
			}},
		},
		ResizeKeyboard: true,
	}

	msgOpt := gotgbot.SendMessageOpts{
		ReplyMarkup: &keyboard,
	}

	ctx.EffectiveMessage.Reply(b, "請選擇目標使用者", &msgOpt)

	return handlers.NextConversationState(PERM)
}

func perm(b *gotgbot.Bot, ctx *ext.Context) error {
	keyboard := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				gotgbot.InlineKeyboardButton{
					Text:         "自訂短網址",
					CallbackData: "customCode",
				}, gotgbot.InlineKeyboardButton{
					Text:         "管理員",
					CallbackData: "admin",
				},
			}, {
				gotgbot.InlineKeyboardButton{
					Text:         "查看所有人的短資訊",
					CallbackData: "readAll",
				},
			},
		},
	}

	msgOpt := gotgbot.SendMessageOpts{
		ReplyMarkup: &keyboard,
	}

	// TODO: delete this message or clear keyboard without showing a message
	_, err := ctx.EffectiveChat.SendMessage(b, "claer keyboard", clearKeyboard)
	if err != nil {
		return fmt.Errorf("Fail to remove keyboard: %w", err)
	}

	ctx.EffectiveMessage.Reply(b, "請選擇權限", &msgOpt)

	return handlers.NextConversationState(VALUE)
}

// isPerms check if the callback query can be handled by this handler.
func isPerms(cq *gotgbot.CallbackQuery) bool {
	return cq.Data == "customCode" || cq.Data == "admin" || cq.Data == "readAll"
}

func value(b *gotgbot.Bot, ctx *ext.Context) error {
	ctx.EffectiveChat.SendMessage(b, "已設定"+ctx.CallbackQuery.Data, clearKeyboard)
	return handlers.NextConversationState(PERM)
}

func exit(b *gotgbot.Bot, ctx *ext.Context) error {
	log.Println("Exit")
	_, err := ctx.EffectiveMessage.Reply(b, "已退出設定", clearKeyboard)
	if err != nil {
		return fmt.Errorf("Fail to send message: %w", err)
	}
	return handlers.EndConversation()
}
