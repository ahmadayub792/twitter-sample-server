package store

import (
	"github.com/ahmadayub792/twitter-sample-server/model"
	"github.com/jinzhu/gorm"
)

type ClientStore struct {
	db *gorm.DB
}

var _ ClientStorer = &ClientStore{}

func NewClientStore(db *gorm.DB) ClientStorer {
	return &ClientStore{db}
}

func (c *ClientStore) Create(client *model.Client) error {
	query := c.db.Model(model.Client{}).Create(client)
	if err := queryError(query); err != nil {
		return err
	}
	return nil
}

func (c *ClientStore) FindOne(where ...interface{}) (*model.Client, error) {
	client := model.Client{}

	query := c.db.Model(model.Client{}).First(&client, where...)
	if err := queryError(query); err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *ClientStore) AddUser(client model.Client, user model.User, role model.ClientRole) error {
	query := c.db.Model(model.ClientUser{}).Create(&model.ClientUser{
		Client: &client,
		User:   &user,
		Role:   role,
	})

	return queryError(query)
}
