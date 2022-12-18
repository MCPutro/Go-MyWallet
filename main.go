package main

import (
	"fmt"
	"github.com/MCPutro/Go-MyWallet/config"
	"github.com/MCPutro/Go-MyWallet/controller"
	"github.com/MCPutro/Go-MyWallet/repository"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/go-playground/validator/v10"
	"os"
)

func main() {
	jwtService := service.NewJwtService("Go-MyWallet")

	validate := validator.New()
	db, err := config.InitDatabase()
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
	activityService := service.NewActivityService(activityRepository, walletRepository, db)
	activityController := controller.NewActivityController(activityService)

	//customMiddleware := middleware.CustomMiddleware(jwtService)

	newRouter := config.NewRouter(userController, walletController, activityController, jwtService)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9999"
	}

	err = newRouter.Listen(":" + PORT)
	if err != nil {
		fmt.Println(err)
	}

}

//func main1() {
//	validate := validator.New()
//	walletRepository := repository.NewWalletRepository()
//
//	ctx := context.Background()
//
//	firebase, err := config.InitFirebase(ctx)
//	if err != nil {
//		return
//	}
//	initDb, err := firebase.Database(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	database := initDb.NewRef("blogs")
//
//	walletService := service.NewWalletService(validate, database, walletRepository)
//
//	wallet := model.Wallet{
//		UserId:   "123456789",
//		WalletId: 2,
//		Name:     "BNI",
//		Type:     "BANK",
//		IsActive: "Y",
//	}
//
//	addWallet, err := walletService.AddWallet(ctx, &wallet)
//	fmt.Println(err)
//	fmt.Println(addWallet)
//
//}
