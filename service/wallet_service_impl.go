package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/MCPutro/Go-MyWallet/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type walletServiceImpl struct {
	validate   *validator.Validate
	walletRepo repository.WalletRepository
	db         *sql.DB
	log        *logrus.Logger
}

func (w *walletServiceImpl) UpdateWallet(ctx context.Context, wallet *model.Wallet) (*model.Wallet, error) {

	/* validation data */
	if err := w.validate.Struct(wallet); err != nil {
		return nil, err
	}

	/* create db transaction */
	conn, err := w.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.Close(err, beginTx, conn)
	}()
	if err != nil {
		return nil, err
	}

	/* check wallet id is exists or not */
	existing, err := w.walletRepo.FindById(ctx, beginTx, wallet.UserId, wallet.WalletId)
	if err != nil {
		return nil, err
	}

	if existing == nil {
		return nil, errors.New("wallet not found")
	}

	return existing, nil
}

func (w *walletServiceImpl) GetWalletType(ctx context.Context) ([]*model.WalletType, error) {
	/* create db transaction */
	conn, err := w.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.Close(err, beginTx, conn)
	}()
	if err != nil {
		return nil, err
	}

	/* call repo func */
	walletType, err := w.walletRepo.GetWalletType(ctx, beginTx)
	if err != nil {
		return nil, err
	}

	/* variable for response */
	var wt []*model.WalletType = nil

	/* fetch data */
	for key, value := range walletType {
		wt = append(wt, &model.WalletType{
			WalletCode: key,
			WalletName: value,
		})
	}

	return wt, nil
}

func (w *walletServiceImpl) GetWalletByUserId(ctx context.Context, UID string) ([]*model.Wallet, error) {
	/* create db transaction */
	conn, err := w.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.Close(err, beginTx, conn)
	}()
	if err != nil {
		return nil, err
	}

	/*get user id*/
	id, _ := helper.GetUserIdAndAccountId(UID)

	/* call repo func */
	walletsByUserId, err := w.walletRepo.FindByUserId(ctx, beginTx, id)
	if err != nil {
		return nil, err
	}

	return walletsByUserId, nil
}

func (w *walletServiceImpl) GetWalletById(ctx context.Context, userid string, walletId uint32) (*model.Wallet, error) {
	/*logging start*/
	w.log.WithFields(logrus.Fields{
		"state": "START", "payload": fmt.Sprintf("userId : %s ; walletId : %d", userid, walletId), //fmt.Sprintf("%+v", newWallet),
	}).Infoln(ctx.Value(fiber.HeaderXRequestID).(string))

	/* create db transaction */
	conn, err := w.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.Close(err, beginTx, conn)
	}()
	if err != nil {
		return nil, err
	}

	/*get user id*/
	id, _ := helper.GetUserIdAndAccountId(userid)

	/*logging out*/
	w.log.WithFields(logrus.Fields{
		"state": "OUT", "payload": fmt.Sprintf("id : %s ; walletId : %d", id, walletId), //fmt.Sprintf("%+v", newWallet),
	}).Infoln(ctx.Value(fiber.HeaderXRequestID).(string))

	/* call repo func */
	wallet, err := w.walletRepo.FindById(ctx, beginTx, id, walletId)
	if err != nil {
		return nil, err
	}

	/*logging start*/
	w.log.WithFields(logrus.Fields{
		"state": "START", "payload": fmt.Sprintf("%+v", wallet),
	}).Infoln(ctx.Value(fiber.HeaderXRequestID).(string))

	/*return object*/
	return wallet, nil
}

func (w *walletServiceImpl) AddWallet(ctx context.Context, newWallet *model.Wallet) (*model.Wallet, error) {

	/* validation data */
	if err := w.validate.Struct(newWallet); err != nil {
		return nil, err
	}

	/* create db transaction */
	conn, err := w.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.Close(err, beginTx, conn)
	}()
	if err != nil {
		return nil, err
	}

	/*set user_id and account_id*/
	id, _ := helper.GetUserIdAndAccountId(newWallet.UserId)
	newWallet.UserId = id

	wallet, err := w.walletRepo.Save(ctx, beginTx, newWallet)

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
		helper.Close(err, beginTx, conn)
	}()
	if err != nil {
		return err
	}

	/*get user_id and account_id*/
	id, _ := helper.GetUserIdAndAccountId(userid)

	err = w.walletRepo.DeleteById(ctx, beginTx, id, walletId)

	return err
}

func NewWalletService(validate *validator.Validate, database *sql.DB, walletRepo repository.WalletRepository, log *logrus.Logger) WalletService {
	return &walletServiceImpl{validate: validate, db: database, walletRepo: walletRepo, log: log}
}
