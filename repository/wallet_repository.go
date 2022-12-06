package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/Go-MyWallet/entity/model"
)

type WalletRepository interface {
	//firebase
	//AddWallet(ctx context.Context, database *db.Ref, newWallet *model.Wallet) (*model.Wallet, error)
	//GetWalletByUserId(ctx context.Context, database *db.Ref, uid string) (*[]model.Wallet, error)

	// postgreSQL
	Save(ctx context.Context, tx *sql.Tx, newWallet *model.Wallet) (*model.Wallet, error)
	Update(ctx context.Context, tx *sql.Tx, newWallet *model.Wallet) (*model.Wallet, error)
	AddAmount(ctx context.Context, tx *sql.Tx, walletId uint, uid string, amount int32) (uint32, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, uid string) (*[]model.Wallet, error)
	FindById(ctx context.Context, tx *sql.Tx, userid string, walletId uint32) (*model.Wallet, error)
	GetWalletType(ctx context.Context, tx *sql.Tx) (map[string]string, error)
}
