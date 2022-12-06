package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/query"
)

type activityRepositoryImpl struct {
}

func (a *activityRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, act *model.Activity) (*model.Activity, error) {

	SQL := "INSERT INTO public.user_activity (user_id, category_id, wallet_id_from, wallet_id_to, period, activity_date, amount) \n" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING activity_id"

	var insertId uint8
	//result, err := tx.ExecContext(ctx, SQL, act.UserId, act.CategoryId, act.WalletIdFrom, act.WalletIdTo, act.Period, act.ActivityDate, act.Amount)
	err := tx.QueryRowContext(ctx, SQL, act.UserId, act.CategoryId, act.WalletIdFrom, act.WalletIdTo, act.Period, act.ActivityDate, act.Amount).Scan(&insertId)
	if err != nil {
		return nil, err
	}

	//insertId, err := result.LastInsertId()
	//fmt.Println(insertId, "<<<<")
	act.ActivityId = uint8(insertId)

	return act, nil
}

func (a *activityRepositoryImpl) GetActivityTypes(ctx context.Context, tx *sql.Tx) (map[string]map[string]map[uint]string, error) {

	SQL := query.GetActivityTypes + " where data.is_active = 'Y';"

	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	//resp := make(map[string]map[string][]string)
	resp := make(map[string]map[string]map[uint]string)
	var typeCode, typeName, category, subCategory string
	var subCategoryId uint
	var multip int

	for rows.Next() {
		err = rows.Scan(&typeCode, &typeName, &category, &subCategoryId, &subCategory, &multip)
		if err != nil {
			return nil, err
		}

		_, cek1 := resp[typeName]
		if !cek1 {
			//resp[typeName] = map[string][]string{
			//	category: {subCategory},
			//}
			resp[typeName] = map[string]map[uint]string{
				category: {subCategoryId: subCategory},
			}
		} else {
			//m := resp[typeName]
			//m[category] = append(m[category], subCategory)

			_, cek2 := resp[typeName][category]
			if !cek2 {
				resp[typeName][category] = map[uint]string{
					subCategoryId: subCategory,
				}
			} else {
				resp[typeName][category][subCategoryId] = subCategory
			}

		}
	}

	//fmt.Println("hasil : ")
	//fmt.Println(resp)

	return resp, nil
}

func (a *activityRepositoryImpl) GetActivityTypeById(ctx context.Context, tx *sql.Tx, categoryId uint) (*model.ActivityCategory, error) {

	SQL := query.GetActivityTypes + " where data.is_active = 'Y' and data.category_id = $1 ;"

	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		var category model.ActivityCategory
		var subCategory model.ActivityCategory
		var temp string
		//err = rows.Scan(&category.SubCategoryCode, &temp, &temp, &category.CategoryName, &category.CategoryName, &category.Multiplier)
		err = rows.Scan(&temp, &temp, &category.CategoryName, &subCategory.CategoryCode, &subCategory.CategoryName, &category.Multiplier)
		if err != nil {
			return nil, err
		}

		return &category, nil
	}

	return nil, nil
}

func NewActivityRepository() ActivityRepository {
	return &activityRepositoryImpl{}
}
