package service

import (
	"context"
	"database/sql"
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

func (a *activityServiceImpl) GetActivityTypeById(ctx context.Context, categoryId uint) (*model.ActivityCategory, error) {
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

	activityCategory, err := a.activityRepo.GetActivityTypeById(ctx, beginTx, categoryId)
	if err != nil {
		return nil, err
	}
	return activityCategory, nil
}

func (a *activityServiceImpl) AddActivity(ctx context.Context, activity *model.Activity) (*web.ActivityResponse, error) {
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

	//save data activity
	activitySave, err := a.activityRepo.Save(ctx, beginTx, activity)
	if err != nil {
		return nil, err
	} else {
		//get detail category
		category, err2 := a.activityRepo.GetActivityTypeById(ctx, beginTx, activity.CategoryId)
		if err2 != nil {
			return nil, err2
		}

		if category.Multiplier == 1 || category.Multiplier == -1 { //income = 1 ; expense = -1
			updateAmount, err := a.walletRepository.AddAmount(ctx, beginTx, activity.WalletIdFrom, activity.UserId, activity.Amount*category.Multiplier)
			if err != nil {
				return nil, err
			}

			return &web.ActivityResponse{
				ActivityId:       activitySave.ActivityId,
				Type:             category.CategoryName,
				Category:         category.SubCategory[0].CategoryName,
				WalletIdFrom:     activitySave.WalletIdFrom,
				WalletIdTo:       activitySave.WalletIdTo,
				ActivityDate:     activitySave.ActivityDate,
				AmountActivity:   activity.Amount,
				AmountWalletFrom: updateAmount,
			}, nil
		} else {
			//transfer own wallet
			updateAmountFrom, err := a.walletRepository.AddAmount(ctx, beginTx, activity.WalletIdFrom, activity.UserId, activity.Amount*-1)
			if err != nil {
				return nil, err
			}
			updateAmountTo, err := a.walletRepository.AddAmount(ctx, beginTx, activity.WalletIdTo, activity.UserId, activity.Amount*1)
			if err != nil {
				return nil, err
			}

			return &web.ActivityResponse{
				ActivityId:       activitySave.ActivityId,
				Type:             category.CategoryName,
				Category:         category.SubCategory[0].CategoryName,
				WalletIdFrom:     activitySave.WalletIdFrom,
				WalletIdTo:       activitySave.WalletIdTo,
				ActivityDate:     activitySave.ActivityDate,
				AmountActivity:   activity.Amount,
				AmountWalletFrom: updateAmountFrom,
				AmountWalletTo:   updateAmountTo,
			}, nil
		}
	}

	//return activity with id
	return nil, err
}

func (a *activityServiceImpl) GetActivityType(ctx context.Context) (*web.ResponseActivityType, error) {
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

	activityTypes, err := a.activityRepo.GetActivityTypes(ctx, beginTx)
	if err != nil {
		return nil, err
	}

	var income []model.ActivityCategory = nil

	for k, v := range activityTypes["Income"]["Income"] {
		income = append(income, model.ActivityCategory{
			CategoryCode: k,
			CategoryName: v,
		})
	}

	var expense []model.ActivityCategory = nil
	var subCategory []model.ActivityCategory
	for k, v := range activityTypes["Expense"] {
		subCategory = nil
		for u, s := range v {
			subCategory = append(subCategory, model.ActivityCategory{
				CategoryCode: u,
				CategoryName: s,
			})
		}

		expense = append(expense, model.ActivityCategory{
			CategoryName: k,
			SubCategory:  subCategory,
		})
	}

	//fmt.Println(income)
	//fmt.Println(expense)

	var transfer []model.ActivityCategory = nil
	for k, v := range activityTypes["Transfer"]["Transfer"] {
		transfer = append(transfer, model.ActivityCategory{
			CategoryCode: k,
			CategoryName: v,
		})
	}

	return &web.ResponseActivityType{
		Status:   "SUCCESS",
		Message:  nil,
		Income:   income,
		Expense:  expense,
		Transfer: transfer,
	}, nil

}

func NewActivityService(activityRepo repository.ActivityRepository, walletRepository repository.WalletRepository, db *sql.DB) ActivityService {
	return &activityServiceImpl{activityRepo: activityRepo, walletRepository: walletRepository, db: db}
}
