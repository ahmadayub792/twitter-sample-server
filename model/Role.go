package model

import (
	"database/sql/driver"
)

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleClient Role = "client"
)

func (r *Role) Scan(value interface{}) error {
	*r = Role(value.([]byte))
	return nil
}

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}
