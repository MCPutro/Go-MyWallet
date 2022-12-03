package main

import (
	"context"
	"fmt"
	"github.com/MCPutro/Go-MyWallet/app"
	"github.com/MCPutro/Go-MyWallet/controller"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/repository"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/go-playground/validator/v10"
	"log"
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

func main1() {
	validate := validator.New()
	walletRepository := repository.NewWalletRepository()

	ctx := context.Background()

	firebase, err := app.InitFirebase(ctx)
	if err != nil {
		return
	}
	initDb, err := firebase.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}
	database := initDb.NewRef("blogs")

	walletService := service.NewWalletService(validate, database, walletRepository)

	wallet := model.Wallet{
		UserId:   "123456789",
		WalletId: 2,
		Name:     "BNI",
		Type:     "BANK",
		IsActive: "Y",
	}

	addWallet, err := walletService.AddWallet(ctx, &wallet)
	fmt.Println(err)
	fmt.Println(addWallet)

}
