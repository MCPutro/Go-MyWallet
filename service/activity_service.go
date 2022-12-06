package service

import (
	"context"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
)

type ActivityService interface {
	GetActivityTypeById(ctx context.Context, categoryId uint) (*model.ActivityCategory, error)
	GetActivityType(ctx context.Context) (*web.ResponseActivityType, error)
	AddActivity(ctx context.Context, activity *model.Activity) (*web.ActivityResponse, error)
}
