package user

import (
	"1li/db"
	"1li/ent"
	"1li/ent/user"
	"context"
)

var admin = map[int]bool{901756183: true}

func Add(ctx context.Context, username string, tgid int) (*ent.User, error) {
	u, err := db.Client.User.Create().
		SetTgid(tgid).
		SetUsername(username).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if isAdmin, ok := admin[tgid]; ok && isAdmin {
		return Op(ctx, username)
	}

	return u, nil
}

func Get(ctx context.Context, tgid int) (*ent.User, error) {
	return db.Client.User.Query().Where(user.Tgid(tgid)).Only(context.Background())
}

func GetByUsername(ctx context.Context, username string) (*ent.User, error) {
	return db.Client.User.Query().Where(user.Username(username)).Only(context.Background())
}

func IsAdmin(ctx context.Context, tgid int) bool {
	u, err := Get(ctx, tgid)
	if err != nil {
		return false
	}

	return u.Admin
}

func SetAdmin(ctx context.Context, username string, admin bool) error {
	u, err := GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	_, err = db.Client.User.UpdateOneID(u.ID).SetAdmin(admin).Save(ctx)
	return err
}

func CanCustomCode(ctx context.Context, tgid int) bool {
	u, err := Get(ctx, tgid)
	if err != nil {
		return false
	}

	return u.CustomCode
}

func SetCustomCode(ctx context.Context, username string, customCode bool) error {
	u, err := GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	_, err = db.Client.User.UpdateOneID(u.ID).SetCustomCode(customCode).Save(ctx)
	return err
}

func CanReadAll(ctx context.Context, tgid int) bool {
	u, err := Get(ctx, tgid)
	if err != nil {
		return false
	}
	return u.ReadAll
}

func SetReadAll(ctx context.Context, username string, readAll bool) error {
	u, err := GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	_, err = db.Client.User.UpdateOneID(u.ID).SetReadAll(readAll).Save(ctx)
	return err
}

func Op(ctx context.Context, username string) (*ent.User, error) {
	u, err := GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return u.Update().
		SetAdmin(true).
		SetCustomCode(true).
		SetReadAll(true).
		Save(ctx)
}

func Deop(ctx context.Context, username string) (*ent.User, error) {
	u, err := GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return u.Update().
		SetAdmin(false).
		SetCustomCode(false).
		SetReadAll(false).
		Save(ctx)
}
