package model

import (
	"database/sql/driver"
)

type ClientRole string

const (
	ClientRoleAdmin   ClientRole = "admin"
	ClientRoleManager ClientRole = "manager"
	ClientRoleAnalyst ClientRole = "analyst"
)

func (r *ClientRole) Scan(value interface{}) error {
	*r = ClientRole(value.([]byte))
	return nil
}

func (r ClientRole) Value() (driver.Value, error) {
	return string(r), nil
}
