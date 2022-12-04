package service

import (
	"context"
	"github.com/MCPutro/Go-MyWallet/entity/web"
)

type ActivityService interface {
	GetActivityType(ctx context.Context) (*web.ResponseActivityType, error)
}
