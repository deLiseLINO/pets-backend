package models

import "time"

type OTP struct {
	Code           string
	Length         int
	NextSendTime   time.Time
	ExparationTime time.Time
}
