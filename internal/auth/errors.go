package auth

import "errors"

var (
	ErrUnauthorized    = errors.New("user is not authorized")
	ErrSignIncompleted = errors.New("failed to sign token")
)
