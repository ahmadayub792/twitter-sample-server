package store

import "github.com/jinzhu/gorm"

func queryError(query *gorm.DB) error {
	if err := query.Error; err != nil{
		if query.RecordNotFound() {
			return ErrNotFound
		}

		return err
	}

	return nil
}