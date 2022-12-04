package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/Go-MyWallet/query"
)

type activityRepositoryImpl struct {
}

func NewActivityRepository() ActivityRepository {
	return &activityRepositoryImpl{}
}

func (a *activityRepositoryImpl) GetActivityTypes(ctx context.Context, tx *sql.Tx) (map[string]map[string]map[uint]string, error) {

	SQL := query.GetActivityTypes

	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	//resp := make(map[string]map[string][]string)
	resp := make(map[string]map[string]map[uint]string)
	var typeCode, typeName, category, subCategory string
	var typeId uint

	for rows.Next() {
		err := rows.Scan(&typeId, &typeCode, &typeName, &category, &subCategory)
		if err != nil {
			return nil, err
		}

		_, cek1 := resp[typeName]
		if !cek1 {
			//resp[typeName] = map[string][]string{
			//	category: {subCategory},
			//}
			resp[typeName] = map[string]map[uint]string{
				category: {typeId: subCategory},
			}
		} else {
			//m := resp[typeName]
			//m[category] = append(m[category], subCategory)

			_, cek2 := resp[typeName][category]
			if !cek2 {
				resp[typeName][category] = map[uint]string{
					typeId: subCategory,
				}
			} else {
				resp[typeName][category][typeId] = subCategory
			}

		}
	}

	//fmt.Println("hasil : ")
	//fmt.Println(resp)

	return resp, nil
}
