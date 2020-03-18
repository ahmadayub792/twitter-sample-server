package store

import (
	"github.com/ahmadayub792/twitter-sample-server/model"
	"github.com/jinzhu/gorm"
)

type UserStore struct {
	db *gorm.DB
}


func NewUserStore(db *gorm.DB) UserStorer {
	return &UserStore{db}
}

func (u *UserStore) FindOne(where ...interface{}) (*model.User, error) {
	user := model.User{}

	query := u.db.Model(model.User{}).First(&user, where...)
	if err := queryError(query); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserStore) Create(user *model.User) error {
	query := u.db.Model(model.User{}).Create(user)
	if err := queryError(query); err != nil {
		return err
	}
	return nil
}

func (u *UserStore) FindAll() ([]*model.User, error) {
	targets := []*model.User{}

	query := u.db.Model(model.User{}).Find(&targets)
	if err := queryError(query); err != nil {
		return nil, err
	}

	return targets, nil
}
