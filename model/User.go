package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `sql:"unique;not null"`
	Password string `sql:"not null"`
	FullName string `sql:"type:VARCHAR(255);not null"`
	Role     Role   `sql:"type:role;not null"`
	Clients  []*ClientUser
}
