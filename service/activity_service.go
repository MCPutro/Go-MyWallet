package service

import (
	"context"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
)

type ActivityService interface {
	GetActivityList(ctx context.Context, UID string) ([]*web.Activity, error)
	GetActivityCategoryById(ctx context.Context, categoryId uint) (*model.ActivityCategory, error)
	GetActivityCategory(ctx context.Context) (*web.ResponseActivityType, error)
	AddActivity(ctx context.Context, activity *model.Activity) (*web.NewActivityResponse, error)
}
