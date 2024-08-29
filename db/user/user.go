package user

import (
	"1li/db"
	"1li/ent"
	"1li/ent/user"
	"1li/errorCollector"
	"context"
	"errors"
	"fmt"
)

var superAdmin = map[int64]bool{901756183: true}

// Add adds a user to the database.
func Add(ctx context.Context, username string, tgid int64) (*ent.User, error) {
	u, err := db.Client.User.Create().
		SetUserid(tgid).
		SetUsername(username).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if isAdmin, ok := superAdmin[tgid]; ok && isAdmin {
		Op(ctx, tgid, true)
	}

	return u, nil
}

// Get gets a user from the database by tgid.
func Get(ctx context.Context, tgid int64) (*ent.User, error) {
	return db.Client.User.Query().Where(user.Userid(tgid)).Only(context.Background())
}

// GetByUsername gets a user from the database by username.
func GetByUsername(ctx context.Context, username string) (*ent.User, error) {
	return db.Client.User.Query().Where(user.Username(username)).Only(context.Background())
}

var ErrUnknownPerm = errors.New("unknown permission")

func GetPerm(ctx context.Context, userid int64, perm string) (bool, error) {
	user, err := db.Client.User.Query().Where(user.Userid(userid)).Only(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get user: %w", err)
	}

	switch perm {
	case "superAdmin":
		return user.SuperAdmin, nil
	case "admin":
		return user.Admin, nil
	case "create":
		return user.Create, nil
	case "customCode":
		return user.CustomCode, nil
	default:
		return false, ErrUnknownPerm
	}
}

func Op(ctx context.Context, userid int64, value bool) error {
	err := errorCollector.New()
	err.Add(SetPerm(ctx, userid, "superAdmin", value))
	err.Add(SetPerm(ctx, userid, "admin", value))
	err.Add(SetPerm(ctx, userid, "create", value))
	err.Add(SetPerm(ctx, userid, "customCode", value))

	return err
}

func SetPerm(ctx context.Context, userid int64, perm string, value bool) error {
	// TODO: batch operation
	user, err := db.Client.User.Query().Where(user.Userid(userid)).Only(ctx)
	// TODO: process user not found here
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	switch perm {
	case "superAdmin":
		_, err = user.Update().SetSuperAdmin(value).Save(ctx)
	case "admin":
		_, err = user.Update().SetAdmin(value).Save(ctx)
	case "create":
		_, err = user.Update().SetCreate(value).Save(ctx)
	case "customCode":
		_, err = user.Update().SetCustomCode(value).Save(ctx)
	default:
		return ErrUnknownPerm
	}

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
