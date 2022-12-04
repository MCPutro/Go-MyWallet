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
	walletAPI.Get("/", walletController.GetWalletByUid)
	walletAPI.Get("/type", walletController.GetWalletType)

	app.Get("/activity", activityController.GetActivityTypes)

	return app
}
