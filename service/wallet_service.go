package service

import (
	"context"
	"github.com/MCPutro/Go-MyWallet/entity/model"
)

type WalletService interface {
	AddWallet(ctx context.Context, newWallet *model.Wallet) (*model.Wallet, error)
}
