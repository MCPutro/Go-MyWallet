package service

import (
	"context"
	"database/sql"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/MCPutro/Go-MyWallet/repository"
	"github.com/go-playground/validator/v10"
)

type walletServiceImpl struct {
	validate   *validator.Validate
	walletRepo repository.WalletRepository
	db         *sql.DB //postgresql
	//database   *db.Ref //firebase

}

func (w *walletServiceImpl) GetWalletType(ctx context.Context) (*[]model.WalletType, error) {
	//create db transaction
	conn, err := w.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.CommitOrRollback(err, beginTx)
		helper.ConnClose(conn)
	}()
	if err != nil {
		return nil, err
	}

	walletType, err := w.walletRepo.GetWalletType(ctx, beginTx)
	if err != nil {
		return nil, err
	}

	var wt []model.WalletType = nil

	for key, value := range walletType {
		wt = append(wt, model.WalletType{
			WalletCode: key,
			WalletName: value,
		})
	}

	return &wt, nil

}

func (w *walletServiceImpl) GetWalletByUserId(ctx context.Context, UID string) ([]*model.Wallet, error) {
	//create db transaction
	conn, err := w.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.CommitOrRollback(err, beginTx)
		helper.ConnClose(conn)
	}()
	if err != nil {
		return nil, err
	}

	walletsByUserId, err := w.walletRepo.FindByUserId(ctx, beginTx, UID)
	if err != nil {
		return nil, err
	}

	return walletsByUserId, nil
}

func (w *walletServiceImpl) GetWalletById(ctx context.Context, userid string, walletId uint32) (*model.Wallet, error) {

	//create db transaction
	conn, err := w.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.CommitOrRollback(err, beginTx)
		helper.ConnClose(conn)
	}()
	if err != nil {
		return nil, err
	}

	wallet, err := w.walletRepo.FindById(ctx, beginTx, userid, walletId)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

//func (w *walletServiceImpl) UpdateWallet(ctx context.Context, wallet *model.Wallet) (*model.Wallet, error) {
//	//validation data
//	err := w.validate.Struct(wallet)
//	if err != nil {
//		return nil, err
//	}
//	//newWallet.IsActive = "Y"
//
//	//create db transaction
//	conn, err := w.db.Conn(ctx)
//	beginTx, err := conn.BeginTx(ctx, nil)
//	defer func() {
//		helper.CommitOrRollback(err, beginTx)
//		helper.ConnClose(conn)
//	}()
//	if err != nil {
//		return nil, err
//	}
//
//	updated, err := w.walletRepo.Update(ctx, beginTx, wallet)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return updated, nil
//}

func (w *walletServiceImpl) AddWallet(ctx context.Context, newWallet *model.Wallet) (*model.Wallet, error) {

	//validation data
	err2 := w.validate.Struct(newWallet)
	if err2 != nil {
		return nil, err2
	}
	newWallet.IsActive = "Y"

	//create db transaction
	conn, err := w.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.CommitOrRollback(err, beginTx)
		helper.ConnClose(conn)
	}()
	if err != nil {
		return nil, err
	}

	//check wallet id is exists or not
	existing, err := w.walletRepo.FindById(ctx, beginTx, newWallet.UserId, newWallet.WalletId)
	if err != nil {
		return nil, err
	}

	//wallet already exist, update data
	var wallet *model.Wallet
	if existing != nil {
		wallet, err = w.walletRepo.Update(ctx, beginTx, newWallet)
	} else {
		wallet, err = w.walletRepo.Save(ctx, beginTx, newWallet)
	}

	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (w *walletServiceImpl) DeleteWallet(ctx context.Context, userid string, walletId uint32) error {
	//create db transaction
	conn, err := w.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.CommitOrRollback(err, beginTx)
		helper.ConnClose(conn)
	}()
	if err != nil {
		return err
	}

	return w.walletRepo.DeleteById(ctx, beginTx, userid, walletId)
}

func NewWalletService(validate *validator.Validate, database *sql.DB, walletRepo repository.WalletRepository) WalletService {
	return &walletServiceImpl{validate: validate, db: database, walletRepo: walletRepo}
}
