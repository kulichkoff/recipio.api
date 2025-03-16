package db

import "errors"

var (
	ErrDBNotConnected = errors.New("DB not connected")
)
