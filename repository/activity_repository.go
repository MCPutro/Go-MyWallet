package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/Go-MyWallet/entity/model"
)

type ActivityRepository interface {
	GetActivityTypes(ctx context.Context, tx *sql.Tx) (map[string]map[string]map[uint]string, error)
	GetActivityTypeById(ctx context.Context, tx *sql.Tx, categoryId uint) (*model.ActivityCategory, error)
	Save(ctx context.Context, tx *sql.Tx, activity *model.Activity) (*model.Activity, error)
}
