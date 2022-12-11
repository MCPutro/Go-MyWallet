package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
)

type ActivityRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, actId uint8, userId string) (*model.Activity, error)
	DeleteById(ctx context.Context, tx *sql.Tx, actId uint8, userId string) error
	FindCompleteActivityByUID(ctx context.Context, tx *sql.Tx, userId string) ([]*web.Activity, error)
	//FindByUserIdxxx(ctx context.Context, tx *sql.Tx, userId string) (*[]model.Activity, error)
	FindActivityCategory(ctx context.Context, tx *sql.Tx) ([]*model.ActivityCategory, error)
	FindActivityCategoryById(ctx context.Context, tx *sql.Tx, categoryId uint) (*model.ActivityCategory, error)
	Save(ctx context.Context, tx *sql.Tx, activity *model.Activity) (*model.Activity, error)
}
