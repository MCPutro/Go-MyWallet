package router

import (
	"github.com/MCPutro/Go-MyWallet/controller"
	"github.com/MCPutro/Go-MyWallet/middleware"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(UserController controller.UserController, walletController controller.WalletController, activityController controller.ActivityController, jwtService service.JwtService) *fiber.App {
	app := fiber.New()

	customMiddleware := middleware.CustomMiddleware(jwtService)

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Pong")
	})

	userAPI := app.Group("/user")

	userAPI.Post("/signup", UserController.Registration)
	userAPI.Post("/signin", UserController.Login)
	userAPI.Get("/all", UserController.ShowALl)

	walletAPI := app.Group("/wallet", customMiddleware)

	walletAPI.Post("/", walletController.AddWallet)
	//walletAPI.Post("/update", walletController.UpdateWallet)
	walletAPI.Get("/uid", walletController.GetWalletByUID)
	walletAPI.Get("/:WalletId", walletController.GetWalletId)
	walletAPI.Get("/type", walletController.GetWalletType)
	walletAPI.Delete("/", walletController.DeleteWallet)

	activityGroup := app.Group("/activity", customMiddleware)

	activityGroup.Get("/category", activityController.GetActivityTypes)
	activityGroup.Post("/", activityController.AddActivity)
	activityGroup.Get("/", activityController.GetAllActivity)
	activityGroup.Delete("/", activityController.DeleteActivity)

	return app
}
