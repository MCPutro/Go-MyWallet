package web

import "github.com/MCPutro/Go-MyWallet/entity/model"

type UserRegistration struct {
	Username string `validate:"alphanum,required,max=15,min=6"`
	FullName string `validate:"required,max=25,min=3"`
	Password string `validate:"required,max=100,min=3"`
	Email    string `validate:"required,email"`
	Imei     string
	DeviceId string
}

type UserResp struct {
	UserId         string                   `json:"UserId,omitempty" `
	Username       string                   `json:"Username,omitempty" `
	FullName       string                   `json:"FullName,omitempty" `
	Authentication model.UserAuthentication `json:"Authentication,omitempty" `
	Data           map[string]string        `json:"Data,omitempty" `
}

type UserLogin struct {
	Account  string `validate:"alphanum,required,max=15,min=6"`
	Password string `validate:"required,max=100,min=3"`
}
