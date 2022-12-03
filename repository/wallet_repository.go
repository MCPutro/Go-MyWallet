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

	//postgres
	AddWallet(ctx context.Context, tx *sql.Tx, newWallet *model.Wallet) (*model.Wallet, error)
	GetWalletByUserId(ctx context.Context, tx *sql.Tx, uid string) (*[]model.Wallet, error)
	GetWalletType(ctx context.Context, tx *sql.Tx) (map[string]string, error)
}
