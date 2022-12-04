package service

import (
	"context"
	"database/sql"
	"github.com/MCPutro/Go-MyWallet/entity/web"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/MCPutro/Go-MyWallet/repository"
)

type activityServiceImpl struct {
	activityRepo repository.ActivityRepository
	db           *sql.DB
}

func NewActivityService(activityRepo repository.ActivityRepository, db *sql.DB) ActivityService {
	return &activityServiceImpl{activityRepo: activityRepo, db: db}
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

	var income []web.ActivityType

	for k, v := range activityTypes["Income"]["Income"] {
		income = append(income, web.ActivityType{
			Code: k,
			Name: v,
		})
	}

	var expense []web.ActivityType
	var temp []web.ActivityType
	for k, v := range activityTypes["Expense"] {
		temp = nil
		for u, s := range v {
			temp = append(temp, web.ActivityType{
				Code: u,
				Name: s,
			})
		}

		expense = append(expense, web.ActivityType{
			Name:        k,
			SubCategory: temp,
		})
	}

	//fmt.Println(income)
	//fmt.Println(expense)

	return &web.ResponseActivityType{
		Status:  "SUCCESS",
		Message: nil,
		Income:  income,
		Expense: expense,
	}, nil
}
