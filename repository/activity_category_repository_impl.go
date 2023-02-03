package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/query"
)

type activityCategoryRepositoryImpl struct {
}

func NewActivityCategoryRepositoryImpl() ActivityCategoryRepository {
	return &activityCategoryRepositoryImpl{}
}

func (a *activityCategoryRepositoryImpl) FindActivityCategory(ctx context.Context, tx *sql.Tx) ([]*model.ActivityCategory, error) {

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

func (a *activityCategoryRepositoryImpl) FindActivityCategoryById(ctx context.Context, tx *sql.Tx, categoryId uint) (*model.ActivityCategory, error) {

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
