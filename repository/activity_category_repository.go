package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/Go-MyWallet/entity/model"
)

type ActivityCategoryRepository interface {
	FindActivityCategory(ctx context.Context, tx *sql.Tx) ([]*model.ActivityCategory, error)
	FindActivityCategoryById(ctx context.Context, tx *sql.Tx, categoryId uint) (*model.ActivityCategory, error)
}
