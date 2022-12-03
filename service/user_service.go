package service

import (
	"context"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
)

type UserService interface {
	FindAll(ctx context.Context) (*[]model.Users, error)
	Login(ctx context.Context, param string, password string) (*model.Users, error)
	Registration(ctx context.Context, userRegistration *web.UserRegistration) (*web.UserRegistrationResp, error)
}
