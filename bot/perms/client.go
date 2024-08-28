package perms

import (
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// TODO: rename thie

type data struct {
	Username string
	UserId   int64
	Perm     string
	Value    bool
}

type client struct {
	rwMux sync.RWMutex

	user map[int64]data
}

func newClient() *client {
	return &client{
		user: make(map[int64]data),
	}
}

func (c *client) get(ctx *ext.Context) (data, bool) {
	c.rwMux.RLock()
	defer c.rwMux.RUnlock()

	user, ok := c.user[ctx.EffectiveChat.Id]
	return user, ok
}

func (c *client) set(ctx *ext.Context, value data) {
	c.rwMux.Lock()
	defer c.rwMux.Unlock()

	c.user[ctx.EffectiveChat.Id] = value
}

func (c *client) del(ctx *ext.Context) {
	c.rwMux.Lock()
	defer c.rwMux.Unlock()

	delete(c.user, ctx.EffectiveChat.Id)
}
