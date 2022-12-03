package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/query"
)

type walletRepositoryImpl struct{}

func (w *walletRepositoryImpl) AddWallet(ctx context.Context, tx *sql.Tx, newWallet *model.Wallet) (*model.Wallet, error) {
	queryInsert := "INSERT INTO public.wallets (user_id, wallet_name, type) VALUES ($1, $2, $3);"

	result, err := tx.ExecContext(ctx, queryInsert, newWallet.UserId, newWallet.Name, newWallet.Type)
	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	return newWallet, nil
}

func (w *walletRepositoryImpl) GetWalletByUserId(ctx context.Context, tx *sql.Tx, uid string) (*[]model.Wallet, error) {
	querySQL := query.GetWalletByUserId

	rows, err := tx.QueryContext(ctx, querySQL, uid)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var walletList []model.Wallet
	var tWallet model.Wallet

	for rows.Next() {
		rows.Scan(&tWallet.UserId, &tWallet.WalletId, &tWallet.Name, &tWallet.Type)

		walletList = append(walletList, tWallet)
	}

	if len(walletList) > 0 {
		return &walletList, nil
	}

	return nil, errors.New("wallet list is empty")
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
