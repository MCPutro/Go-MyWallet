package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/MCPutro/Go-MyWallet/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type userServiceImpl struct {
	UserRepo   repository.UserRepository
	DB         *sql.DB
	Validate   *validator.Validate
	JwtService JwtService
}

func NewUserService(userRepo repository.UserRepository, DB *sql.DB, validate *validator.Validate, jwtService JwtService) UserService {
	return &userServiceImpl{UserRepo: userRepo, DB: DB, Validate: validate, JwtService: jwtService}
}

func (u *userServiceImpl) FindAll(ctx context.Context) ([]*model.Users, error) {
	conn, err := u.DB.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer helper.Close(err, beginTx, conn)

	findAll, err := u.UserRepo.FindAll(ctx, beginTx)
	if err != nil {
		return nil, err
	}

	return findAll, nil
}

func (u *userServiceImpl) Login(ctx context.Context, param string, password string) (*model.Users, error) {
	conn, err := u.DB.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer helper.Close(err, beginTx, conn)
	if err != nil {
		return nil, err
	}

	//check username or email is existing or not
	findUserByUsernameOrEmail, err := u.UserRepo.FindByUsernameOrEmail(ctx, beginTx, param)
	if err != nil {
		return nil, err
	}

	//account is already exist
	if findUserByUsernameOrEmail != nil {
		if findUserByUsernameOrEmail.Status == "ACT" {
			err = bcrypt.CompareHashAndPassword([]byte(findUserByUsernameOrEmail.Authentication.Password), []byte(password))
			if err == nil {
				findUserByUsernameOrEmail.Authentication.Token = u.JwtService.GenerateToken(findUserByUsernameOrEmail.UserId, findUserByUsernameOrEmail.AccountId)
				findUserByUsernameOrEmail.Authentication.RefreshToken = u.GenerateRefreshToken(findUserByUsernameOrEmail.UserId, findUserByUsernameOrEmail.AccountId)
				/*hide pass from resp*/
				findUserByUsernameOrEmail.Authentication.Password = ""
			} else {
				//return nil, errors.New("Your account and password not match. Please try again")
				return nil, errors.New("your account and password not match. please try again")
			}
		} else {
			return nil, errors.New("your account is inactive")
		}
	}

	findUserByUsernameOrEmail.UserId += "-" + findUserByUsernameOrEmail.AccountId
	return findUserByUsernameOrEmail, nil
}

func (u *userServiceImpl) Registration(ctx context.Context, userRegistration *web.UserRegistration) (*model.Users, error) {

	//validate data struct from ui
	if err2 := u.Validate.Struct(userRegistration); err2 != nil {
		validationErrors := err2.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error : ", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
			fmt.Println(fieldError.Value())       //value dari variable yang di validasi
			fmt.Println(fieldError.Param())       //nilai dari requirement
			fmt.Println(fieldError.StructField()) //nama variable yang di validasi
		}
		return nil, err2
	}

	//open connection
	conn, err := u.DB.Conn(ctx)
	beginTx, err := conn.BeginTx(ctx, nil)
	defer helper.Close(err, beginTx, conn)
	if err != nil {
		return nil, err
	}

	//validation data is exists or not
	listAccount, err := u.UserRepo.GetListUsernameAndEmail(ctx, beginTx)
	if err != nil {
		return nil, err
	} else {
		//cek username is already exists in db
		var message []string
		if listAccount[userRegistration.Username] == true {
			message = append(message, "Username")
		}

		//cek email is already exists in db
		if listAccount[userRegistration.Email] == true {
			message = append(message, "Email")
		}

		//return message when username or email already use
		if len(message) > 0 {
			return nil, errors.New(strings.Join(message, " and ") + " already use")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegistration.Password), bcrypt.DefaultCost)

	newUser := model.Users{
		Username:    userRegistration.Username,
		FullName:    userRegistration.FullName,
		CreatedDate: time.Now(),
		Authentication: model.UserAuthentication{
			Password: string(hashedPassword),
		},
		Data: map[string]string{
			"EMAIL": userRegistration.Email,
		},
	}

	if userRegistration.Imei != "" {
		newUser.Data["IMEI"] = userRegistration.Imei
	}

	if userRegistration.DeviceId != "" {
		newUser.Data["DEVICE_ID"] = userRegistration.DeviceId
	}

	//save new user
	userSaved, err := u.UserRepo.Save(ctx, beginTx, newUser)
	if err == nil {
		//generate Token
		userSaved.Authentication.Token = u.JwtService.GenerateToken(userSaved.UserId, "zzz")
		userSaved.Authentication.RefreshToken = u.GenerateRefreshToken(userSaved.UserId, "zzz")
		//hide password
		userSaved.Authentication.Password = ""

		userSaved.UserId += "-zzz"
		return userSaved, nil
	} else {
		return nil, err
	}

}

func (u *userServiceImpl) GenerateRefreshToken(uid, accountId string) string {
	return strings.ReplaceAll(uuid.New().String(), "-", "") + ";" + base64.StdEncoding.EncodeToString([]byte(uid))
}
