package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
	"github.com/MCPutro/Go-MyWallet/query"
	"log"
)

type activityRepositoryImpl struct {
}

func (a *activityRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, actId uint8, userId string) (*model.Activity, error) {
	querySQL := `select uc.activity_id, uc.user_id, uc.category_id, uc.wallet_id_from, uc.wallet_id_to, uc.period, uc.activity_date, uc.amount, uc."desc" as description
					from public.user_activity uc
					where uc.activity_id = $1 and uc.user_id = $2
					;`

	rows, err := tx.QueryContext(ctx, querySQL, actId, userId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		var activity model.Activity
		err = rows.Scan(&activity.ActivityId, &activity.UserId, &activity.CategoryId, &activity.WalletIdFrom, &activity.WalletIdTo, &activity.Period, &activity.ActivityDate, &activity.Nominal, &activity.Desc)
		if err != nil {
			log.Println("[ERROR] activity repo impl - FindById")
			return nil, err
		}
		return &activity, nil
	}

	return nil, errors.New("no data")
}

func (a *activityRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, actId uint8, userId string) error {
	querySQL := `DELETE
					FROM public.user_activity
					WHERE activity_id = $1
					  AND user_id = $2 
					;`

	_, err := tx.ExecContext(ctx, querySQL, actId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (a *activityRepositoryImpl) FindCompleteActivityByUID(ctx context.Context, tx *sql.Tx, userId string) ([]*web.Activity, error) {

	querySQL := `select ac.activity_id, a.type, a.sub_category_name as category, ac.wallet_id_from, w.wallet_name, ac.wallet_id_to, w2.wallet_name, ac.activity_date, ac.amount, ac."desc" as description
					from public.user_activity ac
					inner join activity_category a on a.category_id = ac.category_id
					inner join wallets w on w.wallet_id = ac.wallet_id_from
					inner join wallets w2 on w2.wallet_id = ac.wallet_id_to
					where ac.user_id = $1
					order by ac.activity_date DESC
					;`

	rows, err := tx.QueryContext(ctx, querySQL, userId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var list []*web.Activity

	for rows.Next() {
		var activity web.Activity
		err = rows.Scan(&activity.ActivityId, &activity.Type, &activity.Category, &activity.WalletIdFrom, &activity.WalletNameFrom, &activity.WalletIdTo, &activity.WalletNameTo, &activity.ActivityDate, &activity.Nominal, &activity.Desc)
		if err != nil {
			return nil, err
		}

		list = append(list, &activity)
	}

	return list, nil
}

//func (a *activityRepositoryImpl) FindByUserIdxxx(ctx context.Context, tx *sql.Tx, userId string) (*[]model.Activity, error) {
//
//	querySQL := fmt.Sprintf(query.GetActivityList, "t.user_id = $1")
//
//	rows, err := tx.QueryContext(ctx, querySQL, userId)
//	if err != nil {
//		return nil, err
//	}
//
//	var resp []model.Activity = nil
//	var item model.Activity
//
//	for rows.Next() {
//		err = rows.Scan(&item.ActivityId, &item.UserId, &item.WalletIdFrom, &item.WalletIdTo, &item.Period, &item.ActivityDate)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	return &resp, nil
//}

func (a *activityRepositoryImpl) FindActivityCategory(ctx context.Context, tx *sql.Tx) ([]*model.ActivityCategory, error) {

	SQL := fmt.Sprintf(query.GetActivityTypes, " where ac.is_active = 'Y' ;")

	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var categoryList []*model.ActivityCategory

	for rows.Next() {
		var category model.ActivityCategory
		err = rows.Scan(&category.CategoryId, &category.Type, &category.CategoryName, &category.SubCategoryName)
		if err != nil {
			return nil, err
		}

		categoryList = append(categoryList, &category)
	}

	if len(categoryList) > 0 {
		return categoryList, nil
	} else {
		return nil, errors.New("no data")
	}
}

func (a *activityRepositoryImpl) FindActivityCategoryById(ctx context.Context, tx *sql.Tx, categoryId uint) (*model.ActivityCategory, error) {

	SQL := fmt.Sprintf(query.GetActivityTypes, " where ac.is_active = 'Y' and ac.category_id = $1 ;")

	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		var category model.ActivityCategory
		err = rows.Scan(&category.CategoryId, &category.Type, &category.CategoryName, &category.SubCategoryName)
		if err != nil {
			return nil, err
		}

		return &category, nil
	}

	return nil, errors.New("no data")
}

func (a *activityRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, act *model.Activity) (*model.Activity, error) {

	SQL := `INSERT INTO public.user_activity (user_id, category_id, wallet_id_from, wallet_id_to, period, activity_date, amount, "desc") 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING activity_id`

	var insertId uint8
	_, err := tx.ExecContext(ctx, SQL, act.UserId, act.CategoryId, act.WalletIdFrom, act.WalletIdTo, act.Period, act.ActivityDate, act.Nominal, act.Desc)
	//err := tx.QueryRowContext(ctx, SQL, act.UserId, act.CategoryId, act.WalletIdFrom, act.WalletIdTo, act.Period, act.ActivityDate, act.Nominal, act.Desc).Scan(&insertId)
	if err != nil {
		return nil, err
	}

	//insertId, err := result.LastInsertId()
	//fmt.Println(insertId, "<<<<")
	act.ActivityId = insertId

	return act, nil
}

func NewActivityRepository() ActivityRepository {
	return &activityRepositoryImpl{}
}
