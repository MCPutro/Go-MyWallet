package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/query"
)

type walletRepositoryImpl struct{}

func (w *walletRepositoryImpl) AddWallet(ctx context.Context, tx *sql.Tx, newWallet *model.Wallet) (*model.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w *walletRepositoryImpl) GetWalletByUserId(ctx context.Context, tx *sql.Tx, uid string) (*[]model.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w *walletRepositoryImpl) GetWalletType(ctx context.Context, tx *sql.Tx) (map[string]string, error) {
	querySQL := query.GetWalletType

	rows, err := tx.QueryContext(ctx, querySQL)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	resp := make(map[string]string)

	var itemCode, itemName string

	for rows.Next() {
		err := rows.Scan(&itemCode, &itemName)
		if err != nil {
			return nil, err
		}

		resp[itemCode] = itemName

	}

	return resp, nil
}

func NewWalletRepository() WalletRepository {
	return &walletRepositoryImpl{}
}
