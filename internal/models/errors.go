package models

import (
	"errors"
)

var (
	ErrMismattchCode = errors.New("mismatch code")
	ErrOTPNotFound   = errors.New("otp not found")
	ErrUserNotFound  = errors.New("user not found")
)
