package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"log"
	"strings"
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
	SQL1 := "INSERT INTO public.users (user_id, username, first_name, last_name) VALUES ($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, SQL1, newUser.UserId, newUser.Username, newUser.FirstName, newUser.LastName)
	if err != nil {
		//fmt.Println("[LOG] User_repository_impl - Save1 :", err)
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
	}

	//fmt.Println(result.LastInsertId())
	//fmt.Println(result.RowsAffected())

	//insert into user_data
	for key, value := range newUser.Data {
		SQL2 := "INSERT INTO public.user_data (user_id, data_key, data_value) VALUES ($1, $2, $3);"
		_, err = tx.ExecContext(ctx, SQL2, newUser.UserId, key, value)
		if err != nil {
			fmt.Println("[LOG] User_repository_impl - Save2 :", err)
			return nil, err
		}
	}

	//insert into user_authentication
	SQL3 := "INSERT INTO public.user_authentication (user_id, password) VALUES ($1, $2);"
	_, err = tx.ExecContext(ctx, SQL3, newUser.UserId, newUser.Authentication.Password)
	if err != nil {
		fmt.Println("[LOG] User_repository_impl - Save3 :", err)
		return nil, err
	}

	return &newUser, nil
}

func (u *userRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) (*[]model.Users, error) {
	SQL := "select u.user_id, u.username, u.first_name, u.last_name, u.status, ua.password, string_agg(concat(ud.data_key ,'|', ud.data_value),'##') as user_data " + //string_agg(ud.data_key,'|') as keys , string_agg(ud.data_value,'|') as values " +
		"from public.users u " +
		"left join user_authentication ua on u.user_id = ua.user_id " +
		"left join user_data ud on u.user_id = ud.user_id " +
		"group by (u.user_id, u.username, u.first_name, u.last_name, ua.password)"

	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	if err != nil {
		fmt.Println("[LOG] User_repository_impl - FindAll :", err)
		return nil, err
	}

	var users []model.Users

	m := model.Users{}
	var userData string

	for rows.Next() {
		err := rows.Scan(&m.UserId, &m.Username, &m.FirstName, &m.LastName, &m.Status, &m.Authentication.Password, &userData)
		if err != nil {
			return nil, err
		}
		tmap := make(map[string]string)
		for _, s := range strings.Split(userData, "##") {
			data := strings.Split(s, "|")
			tmap[data[0]] = data[1]
		}
		m.Data = tmap
		m.Authentication.UserId = m.UserId

		users = append(users, m)
	}

	return &users, nil
}

func (u *userRepositoryImpl) FindByUsernameOrEmail(ctx context.Context, tx *sql.Tx, param string) (*model.Users, error) {
	var SQL string
	if helper.IsEmail(param) {

		SQL = "select ud2.user_id, u.username, u.first_name, u.last_name, u.status, ua.password, u.created_date, string_agg(concat(ud.data_key ,'|', ud.data_value),'##') as user_data " +
			"from user_data ud2 " +
			"left join users u on u.user_id = ud2.user_id " +
			"left join user_data ud on u.user_id = ud.user_id and ud.user_id = ud2.user_id " +
			"left join user_authentication ua on ud2.user_id = ua.user_id " +
			"where ud2.data_key = 'EMAIL' and ud2.data_value = $1 " +
			"group by (ud2.user_id, u.username, u.first_name, u.last_name, u.status, ua.password, u.created_date);"
	} else {
		SQL = "select u.user_id, u.username, u.first_name, u.last_name, u.status, ua.password, u.created_date, string_agg(concat(ud.data_key ,'|', ud.data_value),'##') as user_data " +
			"from public.users u " +
			"left join user_authentication ua on u.user_id = ua.user_id " +
			"left join user_data ud on u.user_id = ud.user_id " +
			"where u.username = $1 " +
			"group by (u.user_id, u.username, u.first_name, u.last_name, u.status, ua.password, u.created_date);"
	}

	row, err := tx.QueryContext(ctx, SQL, param)
	defer row.Close()
	if err != nil {
		fmt.Println("[LOG] User_repository_impl - FindByUsernameOrEmail :", err)
		return nil, err
	}

	if row.Next() {
		user := model.Users{}
		var userData string
		err = row.Scan(&user.UserId, &user.Username, &user.FirstName, &user.LastName, &user.Status, &user.Authentication.Password, &user.CreatedDate, &userData)
		user.Authentication.UserId = user.UserId
		tmap := make(map[string]string)
		for _, s := range strings.Split(userData, "##") {
			data := strings.Split(s, "|")
			tmap[data[0]] = data[1]
		}
		user.Data = tmap
		return &user, nil
	}

	return nil, nil

}
