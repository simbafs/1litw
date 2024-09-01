package perms

// TODO:
// 1. 拆分 response 和他們動作 ex 設定鍵盤
// 2. 除存狀態

import (
	"1li/bot/msg"
	"1li/db/user"
	"1li/ent"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/conversation"
)

const (
	ASK_USER  = "askUser"
	ASK_PERM  = "askPerm"
	ASK_VALUE = "askValue"
	SET_PERM  = "setPerm"
)

var clearKeyboard = &gotgbot.SendMessageOpts{
	ReplyMarkup: &gotgbot.ReplyKeyboardRemove{
		RemoveKeyboard: true,
	},
}

var c = newClient()

var CMD = handlers.NewConversation(
	[]ext.Handler{handlers.NewCommand("perms", c.resUser)},
	map[string][]ext.Handler{
		ASK_USER:  {handlers.NewCommand("perms", c.resUser)},
		ASK_PERM:  {handlers.NewMessage(isUserShared, c.resPerm)},
		ASK_VALUE: {handlers.NewCallback(isQueryPerm, c.resValue)},
		SET_PERM:  {handlers.NewCallback(isQueryValue, c.resSetPerm)},
	},
	&handlers.ConversationOpts{
		Exits: []ext.Handler{
			handlers.NewCommand("exit", c.resExit),
			handlers.NewCallback(isQueryCancel, c.resExit),
		},
		StateStorage: conversation.NewInMemoryStorage(conversation.KeyStrategySenderAndChat),
		AllowReEntry: true,
	},
)

// state: ask user

func (c *client) resUser(b *gotgbot.Bot, ctx *ext.Context) error {
	if ok, err := user.GetPerm(context.Background(), ctx.EffectiveUser.Id, "superAdmin"); ent.IsNotFound(err) {
		ctx.EffectiveChat.SendMessage(b, msg.Register, nil)
		return c.resExit(b, ctx)
	} else if err != nil {
		ctx.EffectiveChat.SendMessage(b, msg.ServerError, nil)
		return err
	} else if !ok {
		ctx.EffectiveChat.SendMessage(b, msg.PermissionDenied, nil)
		return c.resExit(b, ctx)
	}

	if err := askUser(b, ctx); err != nil {
		return fmt.Errorf("Fail to ask user: %w", err)
	}

	return handlers.NextConversationState(ASK_PERM)
}

// state: ask perm

// isuserShared checks if the message is a shared user request.
func isUserShared(msg *gotgbot.Message) bool {
	return msg.UsersShared != nil && msg.UsersShared.RequestId == msg.Chat.Id
}

func (c *client) resPerm(b *gotgbot.Bot, ctx *ext.Context) error {
	if _, err := user.GetPerm(context.Background(), ctx.Message.UsersShared.Users[0].UserId, "superAdmin"); ent.IsNotFound(err) {
		ctx.EffectiveChat.SendMessage(b, msg.UserNotExist, nil)
		return c.resExit(b, ctx)
	} else if err != nil {
		ctx.EffectiveChat.SendMessage(b, msg.ServerError, nil)
		return err
	}

	c.set(ctx, data{
		Username: ctx.Message.UsersShared.Users[0].Username,
		UserId:   ctx.Message.UsersShared.Users[0].UserId,
	})
	if err := askPerm(b, ctx); err != nil {
		return fmt.Errorf("Fail to ask perm: %w", err)
	}

	return handlers.NextConversationState(ASK_VALUE)
}

// state: ask value

// isQueryPerm check if the callback query can be handled by this handler.
func isQueryPerm(cq *gotgbot.CallbackQuery) bool {
	return strings.HasPrefix(cq.Data, ASK_PERM)
}

func (c *client) resValue(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := b.AnswerCallbackQuery(ctx.CallbackQuery.Id, nil)
	if err != nil {
		return fmt.Errorf("Fail to answer callback query: %w", err)
	}

	perm, _ := strings.CutPrefix(ctx.CallbackQuery.Data, ASK_PERM)

	// set perm in client
	d, ok := c.get(ctx)
	if !ok {
		ctx.EffectiveChat.SendMessage(b, "還沒選取使用者", nil)
	}
	d.Perm = perm
	c.set(ctx, d)

	err = askValue(b, ctx, perm)
	if err != nil {
		return fmt.Errorf("Fail to ask value: %w", err)
	}

	return handlers.NextConversationState(SET_PERM)
}

// state: set perm

func isQueryValue(cq *gotgbot.CallbackQuery) bool {
	return strings.HasPrefix(cq.Data, ASK_VALUE)
}

func (c *client) resSetPerm(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := b.AnswerCallbackQuery(ctx.CallbackQuery.Id, nil)
	if err != nil {
		return fmt.Errorf("Fail to answer callback query: %w", err)
	}

	value := ctx.CallbackQuery.Data
	d, ok := c.get(ctx)
	if !ok {
		ctx.EffectiveChat.SendMessage(b, msg.UserNotSelected, nil)
	}
	value, _ = strings.CutPrefix(value, ASK_VALUE)
	d.Value = value == "true"

	_, err = user.Get(context.Background(), d.UserId)
	if ent.IsNotFound(err) {
		ctx.EffectiveChat.SendMessage(b, msg.UserNotExist, nil)
		return c.resExit(b, ctx)
	}

	if d.Perm == "all" {
		err = user.Op(context.Background(), d.UserId, d.Value)
	} else {
		err = user.SetPerm(context.Background(), d.UserId, d.Perm, d.Value)
	}
	if err != nil {
		return fmt.Errorf("Fail to set perm: %w", err)
	}

	log.Printf("set user permission: %s, perm: %s, value: %v\n", d.Username, d.Perm, d.Value)

	ctx.EffectiveChat.SendMessage(b, fmt.Sprintf(msg.PermSet, d.Perm, d.Value, d.Username), nil)

	// err = askPerm(b, ctx)
	// if err != nil {
	// 	return fmt.Errorf("Fail to ask perm: %w", err)
	// }
	// return handlers.NextConversationState(ASK_PERM)
	c.del(ctx)
	return handlers.EndConversation()
}

// state: exit

func isQueryCancel(cq *gotgbot.CallbackQuery) bool {
	return cq.Data == "cancel"
}

func (c *client) resExit(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.CallbackQuery != nil {
		_, err := b.AnswerCallbackQuery(ctx.CallbackQuery.Id, nil)
		if err != nil {
			return fmt.Errorf("Fail to answer callback query: %w", err)
		}
	}

	c.del(ctx)
	_, err := ctx.EffectiveMessage.Reply(b, "已退出設定", clearKeyboard)
	if err != nil {
		return fmt.Errorf("Fail to send message: %w", err)
	}
	return handlers.EndConversation()
}

func askUser(b *gotgbot.Bot, ctx *ext.Context) error {
	f := false
	keyboard := gotgbot.ReplyKeyboardMarkup{
		Keyboard: [][]gotgbot.KeyboardButton{
			{
				{
					Text: "選擇目標使用者",
					RequestUsers: &gotgbot.KeyboardButtonRequestUsers{
						RequestId:       ctx.EffectiveChat.Id,
						UserIsBot:       &f,
						RequestUsername: true,
					},
				},
			},
		},
		ResizeKeyboard: true,
	}

	msgOpt := gotgbot.SendMessageOpts{
		ReplyMarkup: &keyboard,
	}

	_, err := ctx.EffectiveChat.SendMessage(b, msg.UserNotSelected, &msgOpt)
	if err != nil {
		return fmt.Errorf("Fail to send message: %w", err)
	}

	return nil
}

func askPerm(b *gotgbot.Bot, ctx *ext.Context) error {
	keyboard := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			// TODO: rearrange buttons
			{
				{
					Text:         "自訂短網址",
					CallbackData: ASK_PERM + "customCode",
				}, {
					Text:         "管理員",
					CallbackData: ASK_PERM + "admin",
				},
			}, {
				{
					Text:         "查看所有資訊",
					CallbackData: ASK_PERM + "readAll",
				}, {
					Text:         "建立",
					CallbackData: ASK_PERM + "create",
				},
			}, {
				{
					Text:         "全部",
					CallbackData: ASK_PERM + "all",
				},
			}, {
				{
					Text:         "離開",
					CallbackData: "cancel",
				},
			},
		},
	}

	msgOpt := gotgbot.SendMessageOpts{
		ReplyMarkup: &keyboard,
	}

	// TODO: delete this message or clear keyboard without showing a message
	_, err := ctx.EffectiveChat.SendMessage(b, msg.KeyboardClear, clearKeyboard)
	if err != nil {
		return fmt.Errorf("Fail to remove keyboard: %w", err)
	}

	_, err = ctx.EffectiveChat.SendMessage(b, msg.PermNotSelected, &msgOpt)
	if err != nil {
		return fmt.Errorf("Fail to send message: %w", err)
	}

	return nil
}

func askValue(b *gotgbot.Bot, ctx *ext.Context, perm string) error {
	keyboard := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{
					Text:         "是",
					CallbackData: ASK_VALUE + "true",
				}, {
					Text:         "否",
					CallbackData: ASK_VALUE + "false",
				},
			}, {
				{
					Text:         "離開",
					CallbackData: "cancel",
				},
			},
		},
	}

	msgOpt := gotgbot.SendMessageOpts{
		ReplyMarkup: &keyboard,
	}

	_, err := ctx.EffectiveChat.SendMessage(b, fmt.Sprintf(msg.DoesSetPerm, perm), &msgOpt)
	if err != nil {
		return fmt.Errorf("Fail to send message: %w", err)
	}

	return nil
}
