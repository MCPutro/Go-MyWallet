package model

import "time"

type Users struct {
	UserId         string
	AccountId      string
	Username       string
	FullName       string
	CreatedDate    time.Time
	Status         string
	Authentication UserAuthentication
	Data           map[string]string
}
