package store

import (
	"errors"

	"github.com/ahmadayub792/twitter-sample-server/model"
)

var (
	ErrNotFound = errors.New("not found")
)

type UserStorer interface {
	Create(*model.User) error
	FindOne(where ...interface{}) (*model.User, error)
}

type ClientStorer interface {
	Create(*model.Client) error
	FindOne(where ...interface{}) (*model.Client, error)
	AddUser(model.Client, model.User, model.ClientRole) error
}

type TargetStorer interface {
	Create(*model.Target) error
	FindOne(where ...interface{}) (*model.Target, error)
	FindAll() ([]*model.Target, error)
}
