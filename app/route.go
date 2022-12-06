package app

import (
	"github.com/MCPutro/Go-MyWallet/controller"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(UserController controller.UserController, walletController controller.WalletController, activityController controller.ActivityController) *fiber.App {
	app := fiber.New()

	userAPI := app.Group("/user")

	userAPI.Post("/signup", UserController.Registration)
	userAPI.Post("/signin", UserController.Login)
	userAPI.Get("/all", UserController.ShowALl)

	walletAPI := app.Group("/wallet")

	walletAPI.Post("/", walletController.AddWallet)
	walletAPI.Post("/update", walletController.UpdateWallet)
	walletAPI.Get("/uid", walletController.GetWalletByUID)
	walletAPI.Get("/id", walletController.GetWalletId)
	walletAPI.Get("/type", walletController.GetWalletType)

	app.Get("/activityTypes", activityController.GetActivityTypes)
	app.Post("/activity", activityController.AddActivity)

	return app
}
