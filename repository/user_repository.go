package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/Go-MyWallet/entity/model"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, newUser model.Users) (*model.Users, error)
	FindAll(ctx context.Context, tx *sql.Tx) (*[]model.Users, error)
	FindByUsernameOrEmail(ctx context.Context, tx *sql.Tx, param string) (*model.Users, error)
	GetListAccount(ctx context.Context, tx *sql.Tx) (map[string]bool, error)
	//FindByEmail(ctx context.Context, tx *sql.Tx, email string) *model.Users

}
