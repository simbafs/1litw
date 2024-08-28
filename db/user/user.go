package user

import (
	"1li/db"
	"1li/ent"
	"1li/ent/user"
	"context"
	"errors"
	"fmt"
)

var admin = map[int64]bool{901756183: true}

// Add adds a user to the database.
func Add(ctx context.Context, username string, tgid int64) (*ent.User, error) {
	u, err := db.Client.User.Create().
		SetTgid(tgid).
		SetUsername(username).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if isAdmin, ok := admin[tgid]; ok && isAdmin {
		SetPerm(context.Background(), tgid, "admin", true)
		SetPerm(context.Background(), tgid, "readAll", true)
		SetPerm(context.Background(), tgid, "customCode", true)
	}

	return u, nil
}

// Get gets a user from the database by tgid.
func Get(ctx context.Context, tgid int64) (*ent.User, error) {
	return db.Client.User.Query().Where(user.Tgid(tgid)).Only(context.Background())
}

// GetByUsername gets a user from the database by username.
func GetByUsername(ctx context.Context, username string) (*ent.User, error) {
	return db.Client.User.Query().Where(user.Username(username)).Only(context.Background())
}

var ErrUnknownPerm = errors.New("unknown permission")

func GetPerm(ctx context.Context, userid int64, perm string) (bool, error) {
	user, err := db.Client.User.Query().Where(user.Tgid(userid)).Only(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get user: %w", err)
	}

	switch perm {
	case "customCode":
		return user.CustomCode, nil
	case "admin":
		return user.Admin, nil
	case "readAll":
		return user.ReadAll, nil
	default:
		return false, ErrUnknownPerm
	}
}

func SetPerm(ctx context.Context, userid int64, perm string, value bool) error {
	// TODO: batch operation
	user, err := db.Client.User.Query().Where(user.Tgid(userid)).Only(ctx)
	// TODO: process user not found here
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	switch perm {
	case "customCode":
		_, err = user.Update().SetCustomCode(value).Save(ctx)
	case "admin":
		_, err = user.Update().SetAdmin(value).Save(ctx)
	case "readAll":
		_, err = user.Update().SetReadAll(value).Save(ctx)
	default:
		return ErrUnknownPerm
	}

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
