package app

import (
	"github.com/ahmadayub792/twitter-sample-server/model"
	"github.com/ahmadayub792/twitter-sample-server/store"
)

type App struct {
	UserStore   store.UserStorer
	ClientStore store.ClientStorer
	TargetStore store.TargetStorer

	PasswordHasher Hasher

	User *model.User

	TokenSecret []byte
}

func (a *App) setUserSession(user *model.User) {
	a.User = user
}

func (a *App) getUserSession() *model.User {
	return a.User
}
