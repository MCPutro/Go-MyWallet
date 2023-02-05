package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/MCPutro/Go-MyWallet/query"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"log"
)

type userRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (u *userRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, newUser model.Users) (*model.Users, error) {

	//generate uid
	uid := uuid.New().String()
	newUser.UserId = uid
	newUser.Authentication.UserId = uid

	//insert into table users
	SQL1 := "INSERT INTO public.users (user_id, username, full_name) VALUES ($1, $2, $3) RETURNING account_id"
	//_, err := tx.ExecContext(ctx, SQL1, newUser.UserId, newUser.Username, newUser.FullName)
	var accountId string
	err := tx.QueryRowContext(ctx, SQL1, newUser.UserId, newUser.Username, newUser.FullName).Scan(&accountId)
	if err != nil {
		log.Println("[LOG] User_repository_impl - Save1 :", err)
		if err2, ok := err.(*pq.Error); ok {
			// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
			//fmt.Println("pq error0:", err2.Code)
			//fmt.Println("pq error1:", err2.Code.Name())
			//fmt.Println("pq error2:", err2.Severity)
			//fmt.Println("pq error3:", err2.Message)
			fmt.Println("pq error4:", err2.Detail)
			//fmt.Println("pq error5:", err2.Hint)
			//fmt.Println("pq error6:", err2.Position)
			//fmt.Println("pq error7:", err2.InternalPosition)
			//fmt.Println("pq error8:", err2.InternalQuery)
			//fmt.Println("pq error9:", err2.Where)
			//fmt.Println("pq error10:", err2.Schema)
			//fmt.Println("pq error11:", err2.Table)
			//fmt.Println("pq error12:", err2.Column)
			//fmt.Println("pq error13:", err2.DataTypeName)
			//fmt.Println("pq error14:", err2.Constraint)
			//fmt.Println("pq error15:", err2.File)
			//fmt.Println("pq error16:", err2.Line)
			//fmt.Println("pq error17:", err2.Routine)
		}

		return nil, err
	} else {
		log.Println("[LOG] User_repository_impl - Save1 :", err)
		newUser.AccountId = accountId
	}

	//insert into user_data
	for key, value := range newUser.Data {
		//fmt.Println(key, " : ", value)

		SQL2 := "INSERT INTO public.user_data (user_id, data_key, data_value) VALUES ($1, $2, $3);"
		_, err = tx.ExecContext(ctx, SQL2, newUser.UserId, key, value)
		if err != nil {
			fmt.Println("[LOG] User_repository_impl - Save2 :", err)
			return nil, err
		}
	}

	//insert into user_authentication
	SQL3 := "INSERT INTO public.user_authentication (user_id, data_key, data_value) VALUES ($1, 'PASSWORD', $2);"
	_, err = tx.ExecContext(ctx, SQL3, newUser.UserId, newUser.Authentication.Password)
	if err != nil {
		fmt.Println("[LOG] User_repository_impl - Save3 :", err)
		return nil, err
	}

	return &newUser, nil
}

func (u *userRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Users, error) {
	SQL := query.GetUserAll + ";"

	//fmt.Println("FindAll user : ", SQL)

	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	if err != nil {
		fmt.Println("[LOG] User_repository_impl - FindAll :", err)
		return nil, err
	}

	var users []*model.Users
	var userData string

	for rows.Next() {
		var m model.Users
		err = rows.Scan(&m.UserId, &m.AccountId, &m.Username, &m.FullName, &m.Status, &m.CreatedDate, &m.Authentication.Password, &userData)
		if err != nil {
			fmt.Println("[LOG] User_repository_impl - FindAll - fetch row :", err)
			return nil, err
		}
		m.UserId += "-" + m.AccountId

		m.Authentication.UserId = m.UserId

		m.Authentication = model.UserAuthentication{}

		//fmt.Println(m.UserId, " -> ", userData, " -> ", len(userData))

		//for _, s := range strings.Split(userData, "##") {
		//	data := strings.Split(s, "|")
		//	tmap[data[0]] = data[1]
		//}
		if len(userData) > 0 {
			tmap := make(map[string]string)
			if err = json.Unmarshal([]byte(userData), &tmap); err != nil {
				return nil, err
			}
			m.Data = tmap
		} else {
			//m.Data = nil
		}

		users = append(users, &m)
		//fmt.Println(len(users))
	}

	return users, nil
}

func (u *userRepositoryImpl) FindByUsernameOrEmail(ctx context.Context, tx *sql.Tx, param string) (*model.Users, error) {
	var SQL string
	if helper.IsEmail(param) {
		SQL = query.GetUserByEmail + ";"
	} else {
		SQL = query.GetUserAll + "where data.username = $1 ;"
	}

	row, err := tx.QueryContext(ctx, SQL, param)
	defer row.Close()
	if err != nil {
		fmt.Println("[LOG] User_repository_impl - FindByUsernameOrEmail :", err)
		return nil, err
	}

	//user if else for only 1 row
	if row.Next() {
		var m model.Users
		var userData string
		err = row.Scan(&m.UserId, &m.AccountId, &m.Username, &m.FullName, &m.Status, &m.CreatedDate, &m.Authentication.Password, &userData)
		if err != nil {
			return nil, err
		}
		m.Authentication.UserId = m.UserId
		tmap := make(map[string]string)
		//for _, s := range strings.Split(userData, "##") {
		//	data := strings.Split(s, "|")
		//	tmap[data[0]] = data[1]
		//}
		if err = json.Unmarshal([]byte(userData), &tmap); err != nil {
			return nil, err
		}
		m.Data = tmap
		return &m, nil
	}

	return nil, err

}

func (u *userRepositoryImpl) GetListUsernameAndEmail(ctx context.Context, tx *sql.Tx) (map[string]bool, error) {
	SQL := query.GetListAccount

	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	resp := make(map[string]bool)
	var item string

	for rows.Next() {

		err = rows.Scan(&item)
		if err != nil {
			return nil, err
		}

		resp[item] = true
	}

	return resp, nil
}
