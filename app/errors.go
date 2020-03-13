package app

import "errors"

var (
	ErrAuthFailed = errors.New("authentication failed")
	ErrIncorrectPassword = errors.New("incorrect password")
)
