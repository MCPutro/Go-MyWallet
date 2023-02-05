package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
)

type ActivityRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, actId uint32, userId string) (*model.Activity, error)
	DeleteById(ctx context.Context, tx *sql.Tx, actId uint32, userId string) error
	Save(ctx context.Context, tx *sql.Tx, activity *model.Activity) (*model.Activity, error)
	FindDetailActivityByUID(ctx context.Context, tx *sql.Tx, userId string) ([]*web.Activity, error)
}
