package service

import (
	"context"
	"firebase.google.com/go/v4/db"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/repository"
	"github.com/go-playground/validator/v10"
)

type walletServiceImpl struct {
	validate   *validator.Validate
	database   *db.Ref
	walletRepo repository.WalletRepository
}

func NewWalletService(validate *validator.Validate, database *db.Ref, walletRepo repository.WalletRepository) WalletService {
	return &walletServiceImpl{validate: validate, database: database, walletRepo: walletRepo}
}

func (w *walletServiceImpl) AddWallet(ctx context.Context, newWallet *model.Wallet) (*model.Wallet, error) {

	newWallet.IsActive = "Y"

	wallet, err := w.walletRepo.AddWallet(ctx, w.database, newWallet)

	if err != nil {
		return nil, err
	}

	return wallet, err
}
