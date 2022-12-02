package service

import (
	"context"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
)

type UserService interface {
	Login(ctx context.Context, param string, password string) (*model.Users, error)
	Registration(ctx context.Context, userRegistration *web.UserRegistration) (*web.UserRegistrationResp, error)
}
