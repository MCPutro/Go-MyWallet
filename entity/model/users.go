package model

import "time"

type Users struct {
	UserId         string
	Username       string
	FirstName      string
	LastName       string
	CreatedDate    time.Time
	Status         string
	Authentication UserAuthentication
	Data           []UserData
}
