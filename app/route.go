package app

import (
	"github.com/MCPutro/Go-MyWallet/controller"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(UserController controller.UserController) *fiber.App {
	app := fiber.New()

	userAPI := app.Group("/user")

	userAPI.Post("/signup", UserController.Registration)
	userAPI.Post("/signin", UserController.Login)
	userAPI.Get("/all", UserController.ShowALl)

	return app
}
