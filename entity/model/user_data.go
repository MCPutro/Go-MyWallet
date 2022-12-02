package model

import "time"

type UserData struct {
	UserId      string
	DataKey     string `validate:"required,max=15"`
	DataValue   string `validate:"required,max=15"`
	CreatedDate time.Time
}
