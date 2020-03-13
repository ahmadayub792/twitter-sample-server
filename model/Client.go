package model

import "github.com/jinzhu/gorm"

type Client struct {
	gorm.Model
	Name    string `sql:"type:VARCHAR(255);not null"`
	Users   []*ClientUser
	Targets []*Target `gorm:"many2many:client_targets"`
}
