package repository

import (
	"context"
	"database/sql"
)

type ActivityRepository interface {
	GetActivityTypes(ctx context.Context, tx *sql.Tx) (map[string]map[string]map[uint]string, error)
}
