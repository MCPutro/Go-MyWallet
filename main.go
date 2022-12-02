package main

import (
	"context"
	"fmt"
	"github.com/MCPutro/Go-MyWallet/app"
	"github.com/MCPutro/Go-MyWallet/controller"
	"github.com/MCPutro/Go-MyWallet/repository"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/go-playground/validator/v10"
	"os"
)

func main() {
	jwtService := service.NewJwtService("goWallet", "emchepe")

	validate := validator.New()
	db, err := app.InitDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	userRepository := repository.NewUserRepository()

	userService := service.NewUserService(userRepository, db, validate, jwtService)

	userController := controller.NewUserController(userService)

	newRouter := app.NewRouter(userController)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9999"
	}

	err = newRouter.Listen(":" + PORT)
	if err != nil {
		fmt.Println(err)
	}

}

func main2() {

	db, err := app.InitDatabase()

	if err != nil {
		fmt.Println("error : ", err)
		return
	}

	//load all
	tx, err := db.Begin()

	userRepository := repository.NewUserRepository()

	//all := userRepository.FindAll(context.Background(), tx)

	userByUsernameOrEmail, err := userRepository.FindByUsernameOrEmail(context.Background(), tx, "akulup2a3")

	fmt.Println(userByUsernameOrEmail)
}

func main1() {

	jwtService := service.NewJwtService("goWallet", "emchepe")

	validate := validator.New()
	db, err := app.InitDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	userRepository := repository.NewUserRepository()

	service.NewUserService(userRepository, db, validate, jwtService)

	//dummyNewUser := web.UserRegistration{
	//	Username:  "akulupa4",
	//	FirstName: "aku",
	//	LastName:  "lupa",
	//	Password:  "123456789",
	//	Email:     "aku@lupa.oh",
	//}
	//
	//err = userService.Registration(context.Background(), dummyNewUser)
	//fmt.Println(err)

	//err2 := userService.Login(context.Background(), "akulupa4", "1234567891")
	//
	//fmt.Println(err2)
}
