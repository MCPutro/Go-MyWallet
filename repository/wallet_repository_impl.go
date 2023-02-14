package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/query"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type walletRepositoryImpl struct {
	log *logrus.Logger
}

func (w *walletRepositoryImpl) AddAmount(ctx context.Context, tx *sql.Tx, walletId uint32, uid string, amount uint32, category string) (uint32, error) {
	var queryUpdate string
	querySQL := "UPDATE public.wallets SET amount = (amount %s $2) WHERE wallet_id = $1 and user_id = $3 returning amount;"
	if category == "EXP" {
		queryUpdate = fmt.Sprintf(querySQL, "-")
	} else {
		queryUpdate = fmt.Sprintf(querySQL, "+")
	}

	//_, err := tx.ExecContext(ctx, queryUpdate, walletId, amount, uid)
	var updateAmount uint32 = 0
	err := tx.QueryRowContext(ctx, queryUpdate, walletId, amount, uid).Scan(&updateAmount)
	if err != nil {
		return 0, err
	}

	return updateAmount, nil
}

func (w *walletRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, newWallet *model.Wallet) (*model.Wallet, error) {
	querySQL := "INSERT INTO public.wallets (user_id, wallet_name, type, amount) VALUES ($1, $2, $3, $4) RETURNING wallet_id;"

	/*logging start*/
	w.log.WithFields(logrus.Fields{"state": "START", "payload": fmt.Sprintf("%+v", newWallet), "query": querySQL}).Infoln(ctx.Value(fiber.HeaderXRequestID))

	var insertId uint32
	err := tx.QueryRowContext(ctx, querySQL, newWallet.UserId, newWallet.Name, newWallet.Type, newWallet.Amount).Scan(&insertId)
	if err != nil {
		/*logging error*/
		w.log.WithField("ERROR", err).Errorln(ctx.Value(fiber.HeaderXRequestID).(string))
		return nil, err
	}

	newWallet.WalletId = insertId

	return newWallet, nil
}

func (w *walletRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, newWallet *model.Wallet) (*model.Wallet, error) {
	querySQL := "UPDATE public.wallets SET user_id = $1, wallet_name = $2, type = $3, is_active = $4, amount = $5 WHERE wallet_id = $6 and user_id = $1;"

	/*logging start*/
	w.log.WithFields(logrus.Fields{"state": "START", "payload": fmt.Sprintf("%+v", newWallet), "query": querySQL}).Infoln(ctx.Value(fiber.HeaderXRequestID).(string))

	if newWallet.IsActive == "" {
		newWallet.IsActive = "Y"
	}

	_, err := tx.ExecContext(ctx, querySQL, newWallet.UserId, newWallet.Name, newWallet.Type, newWallet.IsActive, newWallet.Amount, newWallet.WalletId)
	if err != nil {
		/*logging error*/
		w.log.WithField("ERROR", err).Errorln(ctx.Value(fiber.HeaderXRequestID).(string))
		return nil, err
	}

	return newWallet, nil
}

func (w *walletRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, uid string) ([]*model.Wallet, error) {
	querySQL := fmt.Sprintf(query.GetWalletById, "w.user_id = $1")

	/*logging*/
	w.log.WithFields(logrus.Fields{"state": "START", "payload": fmt.Sprintf("userId:%s", uid), "query": querySQL}).Infoln(ctx.Value(fiber.HeaderXRequestID))

	rows, err := tx.QueryContext(ctx, querySQL, uid)
	defer rows.Close()
	if err != nil {
		/*logging error*/
		w.log.WithField("ERROR", err).Errorln(ctx.Value(fiber.HeaderXRequestID))
		return nil, err
	}

	var walletList []*model.Wallet = nil

	for rows.Next() {
		var tWallet model.Wallet
		err = rows.Scan(&tWallet.UserId, &tWallet.WalletId, &tWallet.Name, &tWallet.Type, &tWallet.Amount)
		if err != nil {
			/*logging error*/
			w.log.WithField("ERROR", err).Errorln(ctx.Value(fiber.HeaderXRequestID))
			return nil, err
		}

		walletList = append(walletList, &tWallet)
	}

	if len(walletList) > 0 {
		return walletList, nil
	}

	return nil, nil
}

func (w *walletRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userid string, walletId uint32) (*model.Wallet, error) {
	querySQL := fmt.Sprintf(query.GetWalletById, "w.wallet_id = $1 and w.user_id = $2 ")

	/*logging*/
	w.log.WithFields(logrus.Fields{
		"state": "START", "payload": fmt.Sprintf("userId:%s ; walletId:%d", userid, walletId), "query": querySQL,
	}).Infoln(ctx.Value(fiber.HeaderXRequestID))

	rows, err := tx.QueryContext(ctx, querySQL, walletId, userid)
	defer rows.Close()
	if err != nil {
		/*logging error*/
		w.log.WithField("ERROR", err).Errorln(ctx.Value(fiber.HeaderXRequestID))

		return nil, err
	}

	if rows.Next() {
		var tWallet model.Wallet
		err = rows.Scan(&tWallet.UserId, &tWallet.WalletId, &tWallet.Name, &tWallet.Type, &tWallet.Amount)
		if err != nil {
			/*logging error*/
			w.log.WithField("ERROR", err).Errorln(ctx.Value(fiber.HeaderXRequestID))

			return nil, err
		}

		return &tWallet, nil
	}

	return nil, nil
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
		err = rows.Scan(&itemCode, &itemName)
		if err != nil {
			return nil, err
		}

		resp[itemCode] = itemName

	}

	return resp, nil
}

func (w *walletRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, userid string, walletId uint32) error {

	querySQL := "delete from public.wallets where user_id = $1 and wallet_id = $2 ;"

	/*logging*/
	w.log.WithFields(logrus.Fields{"state": "START", "payload": fmt.Sprintf("userId:%s ; walletId:%d", userid, walletId), "query": querySQL}).Infoln(ctx.Value(fiber.HeaderXRequestID))

	_, err := tx.ExecContext(ctx, querySQL, userid, walletId)
	if err != nil {
		/*logging error*/
		w.log.WithField("ERROR", err).Errorln(ctx.Value(fiber.HeaderXRequestID))
		return err
	}

	return nil
}

func NewWalletRepository(log *logrus.Logger) WalletRepository {
	return &walletRepositoryImpl{log: log}
}
