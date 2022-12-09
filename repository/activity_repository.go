package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
)

type ActivityRepository interface {
	FindCompleteActivityByUID(ctx context.Context, tx *sql.Tx, userId string) ([]*web.Activity, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, userId string) (*[]model.Activity, error)
	FindActivityCategory(ctx context.Context, tx *sql.Tx) ([]*model.ActivityCategory, error)
	FindActivityCategoryById(ctx context.Context, tx *sql.Tx, categoryId uint) (*model.ActivityCategory, error)
	Save(ctx context.Context, tx *sql.Tx, activity *model.Activity) (*model.Activity, error)
}
