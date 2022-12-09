package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/MCPutro/Go-MyWallet/repository"
)

type activityServiceImpl struct {
	activityRepo     repository.ActivityRepository
	walletRepository repository.WalletRepository
	db               *sql.DB
}

func (a *activityServiceImpl) GetActivityList(ctx context.Context, UID string) ([]*web.Activity, error) {
	//open db trx
	conn, err := a.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.CommitOrRollback(err, beginTx)
		helper.ConnClose(conn)
	}()
	if err != nil {
		return nil, err
	}

	uid, err := a.activityRepo.FindCompleteActivityByUID(ctx, beginTx, UID)
	if err != nil {
		return nil, err
	}

	return uid, nil
}

func (a *activityServiceImpl) AddActivity(ctx context.Context, activity *model.Activity) (*web.NewActivityResponse, error) {
	//open db trx
	conn, err := a.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.CommitOrRollback(err, beginTx)
		helper.ConnClose(conn)
	}()
	if err != nil {
		return nil, err
	}

	//parse ActivityDate to period yyyy-mm
	activity.Period = activity.ActivityDate.Format("2006-01")

	//get detail category
	category, err := a.activityRepo.FindActivityCategoryById(ctx, beginTx, activity.CategoryId)
	if err != nil {
		return nil, err
	}

	//check current balance wallet from is greater than nominal activity
	walletFrom, err := a.walletRepository.FindById(ctx, beginTx, activity.UserId, activity.WalletIdFrom)
	if walletFrom.Amount < activity.Nominal && category.Type == "EXP" {
		return nil, errors.New("balance not enough")
	}

	//save data activity
	activitySave, err := a.activityRepo.Save(ctx, beginTx, activity)
	if err != nil {
		return nil, err
	} else {
		if category.Type == "EXP" || category.Type == "INC" { //income = 1 ; expense = -1
			updateAmount, err := a.walletRepository.AddAmount(ctx, beginTx, activity.WalletIdFrom, activity.UserId, activity.Nominal, category.Type)
			if err != nil {
				return nil, err
			}

			return &web.NewActivityResponse{
				ActivityId:         activitySave.ActivityId,
				Type:               category.Type, //category.CategoryName,
				Category:           category.SubCategoryName,
				WalletIdFrom:       activitySave.WalletIdFrom,
				WalletIdTo:         activitySave.WalletIdTo,
				ActivityDate:       activitySave.ActivityDate,
				Nominal:            activitySave.Nominal,
				AmountWalletIdFrom: updateAmount,
				Desc:               activitySave.Desc,
			}, nil
		} else {
			//transfer own wallet
			updateAmountFrom, err := a.walletRepository.AddAmount(ctx, beginTx, activity.WalletIdFrom, activity.UserId, activity.Nominal, category.Type)
			if err != nil {
				return nil, err
			}
			updateAmountTo, err2 := a.walletRepository.AddAmount(ctx, beginTx, activity.WalletIdTo, activity.UserId, activity.Nominal, category.Type)
			if err2 != nil {
				return nil, err2
			}

			return &web.NewActivityResponse{
				ActivityId:         activitySave.ActivityId,
				Type:               category.CategoryName,
				Category:           category.SubCategoryName,
				WalletIdFrom:       activitySave.WalletIdFrom,
				WalletIdTo:         activitySave.WalletIdTo,
				ActivityDate:       activitySave.ActivityDate,
				Nominal:            activitySave.Nominal,
				AmountWalletIdFrom: updateAmountFrom,
				AmountWalletIdTo:   updateAmountTo,
				Desc:               activitySave.Desc,
			}, nil
		}
	}

	//return activity with id
	//return nil, err
}

func (a *activityServiceImpl) GetActivityCategory(ctx context.Context) (*web.ResponseActivityType, error) {
	//begin db trx
	conn, err := a.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.CommitOrRollback(err, beginTx)
		helper.ConnClose(conn)
	}()
	if err != nil {
		return nil, err
	}

	categories, err := a.activityRepo.FindActivityCategory(ctx, beginTx)
	if err != nil {
		return nil, err
	}

	var income []*model.ActivityCategory = nil
	var expense []*model.ActivityCategory = nil
	var transfer []*model.ActivityCategory = nil

	for _, category := range categories {
		if category.Type == "INC" {
			income = append(income, category)
		} else if category.Type == "EXP" {
			expense = append(expense, category)
		} else {
			transfer = append(transfer, category)
		}
	}

	return &web.ResponseActivityType{
		Status:   "SUCCESS",
		Message:  nil,
		Income:   income,
		Expense:  expense,
		Transfer: transfer,
	}, nil

}

func (a *activityServiceImpl) GetActivityCategoryById(ctx context.Context, categoryId uint) (*model.ActivityCategory, error) {
	//open db trx
	conn, err := a.db.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer func() {
		helper.CommitOrRollback(err, beginTx)
		helper.ConnClose(conn)
	}()
	if err != nil {
		return nil, err
	}

	activityCategory, err := a.activityRepo.FindActivityCategoryById(ctx, beginTx, categoryId)
	if err != nil {
		return nil, err
	}
	return activityCategory, nil
}

func NewActivityService(activityRepo repository.ActivityRepository, walletRepository repository.WalletRepository, db *sql.DB) ActivityService {
	return &activityServiceImpl{activityRepo: activityRepo, walletRepository: walletRepository, db: db}
}
