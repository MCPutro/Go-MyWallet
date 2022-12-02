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

func (u *userServiceImpl) Login(ctx context.Context, param string, password string) (*model.Users, error) {
	conn, err := u.DB.Conn(ctx)
	defer helper.ConnClose(conn)

	beginTx, err := conn.BeginTx(ctx, nil)
	defer helper.CommitOrRollback(err, beginTx)
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
		err := bcrypt.CompareHashAndPassword([]byte(findUserByUsernameOrEmail.Authentication.Password), []byte(password))
		if err == nil {
			token := u.JwtService.GenerateToken(findUserByUsernameOrEmail.UserId)
			findUserByUsernameOrEmail.Authentication.Token = token
			findUserByUsernameOrEmail.Authentication.RefreshToken = strings.ReplaceAll(uuid.New().String(), "-", "") + ";" + base64.StdEncoding.EncodeToString([]byte(findUserByUsernameOrEmail.UserId))
		} else {
			return nil, errors.New("Your account and password not match. Please try again")
		}
	}

	return findUserByUsernameOrEmail, nil
}

func (u *userServiceImpl) Registration(ctx context.Context, userRegistration *web.UserRegistration) (*web.UserRegistrationResp, error) {

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

	conn, err := u.DB.Conn(ctx)
	defer helper.ConnClose(conn)

	beginTx, err := conn.BeginTx(ctx, nil)
	defer helper.CommitOrRollback(err, beginTx)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegistration.Password), bcrypt.DefaultCost)

	newUser := model.Users{
		Username:    userRegistration.Username,
		FirstName:   userRegistration.FirstName,
		LastName:    userRegistration.LastName,
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

	//save new user
	userSaved, err := u.UserRepo.Save(ctx, beginTx, newUser)
	if err == nil {
		return &web.UserRegistrationResp{
			UID:      userSaved.UserId,
			Username: userSaved.Username,
			Password: userSaved.Authentication.Password,
		}, nil
	} else {
		return nil, err
	}

}
