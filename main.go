package main

import (
	"database/sql"
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

	logger := app.InitLog(app.LogApp)

	logger.Infoln("Starting Application...")

	validate := validator.New()
	db, err := app.InitDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	logger.Infoln("Initial Variable...")

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate, jwtService)
	userController := controller.NewUserController(userService)

	walletRepository := repository.NewWalletRepository(logger)
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

	logger.Infoln("Running in port", PORT)

	err = newRouter.Listen(":" + PORT)
	if err != nil {
		fmt.Println(err)
		logger.Errorln(err)
	}

}
