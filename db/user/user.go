package user

import (
	"1li/db"
	"1li/ent"
	"1li/ent/user"
	"context"
)

func Add(ctx context.Context, username string, tgid int) (*ent.User, error) {
	return db.Client.User.Create().
		SetTgid(tgid).
		SetUsername(username).
		Save(ctx)
}

func Get(ctx context.Context, tgid int) (*ent.User, error) {
	return db.Client.User.Query().Where(user.Tgid(tgid)).Only(context.Background())
}
