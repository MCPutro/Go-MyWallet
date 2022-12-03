package model

import "time"

type Users struct {
	UserId         string             `json:"UserId,omitempty"`
	Username       string             `json:"Username,omitempty"`
	FirstName      string             `json:"FirstName,omitempty"`
	LastName       string             `json:"LastName,omitempty"`
	CreatedDate    time.Time          `json:"CreatedDate,omitempty"`
	Status         string             `json:"Status,omitempty"`
	Authentication UserAuthentication `json:"Authentication,omitempty"`
	Data           map[string]string  `json:"Data,omitempty"`
}
