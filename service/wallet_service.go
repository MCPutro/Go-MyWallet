package service

import (
	"context"
	"github.com/MCPutro/Go-MyWallet/entity/model"
)

type WalletService interface {
	AddWallet(ctx context.Context, newWallet *model.Wallet) (*model.Wallet, error)
	//UpdateWallet(ctx context.Context, wallet *model.Wallet) (*model.Wallet, error)
	GetWalletByUserId(ctx context.Context, UID string) ([]*model.Wallet, error)
	GetWalletById(ctx context.Context, userid string, walletId uint32) (*model.Wallet, error)
	GetWalletType(ctx context.Context) (*[]model.WalletType, error)
	DeleteWallet(ctx context.Context, userid string, walletId uint32) error
}
