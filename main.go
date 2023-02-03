package main

import (
	"fmt"
	"github.com/MCPutro/Go-MyWallet/app"
	"github.com/MCPutro/Go-MyWallet/app/router"
	"github.com/MCPutro/Go-MyWallet/controller"
	"github.com/MCPutro/Go-MyWallet/repository"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	jwtService := service.NewJwtService("Go-MyWallet")

	validate := validator.New()
	db, err := app.InitDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate, jwtService)
	userController := controller.NewUserController(userService)

	walletRepository := repository.NewWalletRepository()
	walletService := service.NewWalletService(validate, db, walletRepository)
	walletController := controller.NewWalletController(walletService)

	activityRepository := repository.NewActivityRepository()
	activityCategoryRepository := repository.NewActivityCategoryRepositoryImpl()
	activityService := service.NewActivityService(activityRepository, activityCategoryRepository, walletRepository, db)
	activityController := controller.NewActivityController(activityService)

	//customMiddleware := middleware.CustomMiddleware(jwtService)

	newRouter := router.NewRouter(userController, walletController, activityController, jwtService)

	PORT := app.AppPort
	if PORT == "" {
		PORT = "9999"
	}

	err = newRouter.Listen(":" + PORT)
	if err != nil {
		fmt.Println(err)
	}

}
