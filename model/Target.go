package model

import "github.com/jinzhu/gorm"

type Target struct {
	gorm.Model
	Type    string `gorm:"unique_index:idx_targets_type_handler" sql:"not null"`
	Handler string `gorm:"unique_index:idx_targets_type_handler" sql:"not null"`
}
