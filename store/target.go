package store

import (
	"github.com/ahmadayub792/twitter-sample-server/model"
	"github.com/jinzhu/gorm"
)

type TargetStore struct {
	db *gorm.DB
}

var _ TargetStorer = &TargetStore{}

func NewTargetStore(db *gorm.DB) TargetStorer {
	return &TargetStore{db}
}

func (t *TargetStore) Create(target *model.Target) error {
	query := t.db.Model(model.Target{}).Create(target)
	if err := queryError(query); err != nil {
		return err
	}
	return nil
}

func (t *TargetStore) FindOne(where ...interface{}) (*model.Target, error) {
	target := model.Target{}

	query := t.db.Model(model.Target{}).First(&target, where...)
	if err := queryError(query); err != nil {
		return nil, err
	}
	return &target, nil
}

func (t *TargetStore) FindAll() ([]*model.Target, error) {
	targets := []*model.Target{}

	query := t.db.Model(model.Target{}).Find(&targets)
	if err := queryError(query); err != nil {
		return nil, err
	}
	return targets, nil
}
