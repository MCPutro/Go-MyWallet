package controller

import (
	"fmt"
	"github.com/MCPutro/Go-MyWallet/entity/web"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/gofiber/fiber/v2"
)

type userControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{UserService: userService}
}

func (u *userControllerImpl) Login(c *fiber.Ctx) error {
	l := new(web.UserLogin)

	if err := c.BodyParser(l); err != nil {
		return err
	}

	userLogin, err := u.UserService.Login(c.Context(), l.Account, l.Password)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		//return c.SendString(fmt.Sprint("error :", err))
		return c.JSON(web.Response{
			Status:  "ERROR",
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		c.Status(fiber.StatusOK)
		return c.JSON(web.Response{
			Status:  "SUCCESS",
			Message: nil,
			Data:    userLogin,
		})
	}
}

func (u *userControllerImpl) Registration(c *fiber.Ctx) error {
	p := new(web.UserRegistration)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	resp, err := u.UserService.Registration(c.Context(), p)
	if err != nil {
		return c.SendString(fmt.Sprint("error :", err))
	} else {
		return c.JSON(resp)
	}
}
