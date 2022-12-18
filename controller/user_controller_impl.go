package controller

import (
	"github.com/MCPutro/Go-MyWallet/entity/web"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/gofiber/fiber/v2"
)

type userControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{UserService: userService}
}

func (u *userControllerImpl) ShowALl(c *fiber.Ctx) error {
	findAll, err := u.UserService.FindAll(c.Context())

	return helper.PrintResponse(err, findAll, c)
}

func (u *userControllerImpl) Login(c *fiber.Ctx) error {
	l := new(web.UserLogin)

	//parse raw body to variable l
	if err := c.BodyParser(l); err != nil {
		return err
	}

	//call func login in user service
	userLogin, err := u.UserService.Login(c.Context(), l.Account, l.Password)

	return helper.PrintResponse(err, userLogin, c)
}

func (u *userControllerImpl) Registration(c *fiber.Ctx) error {
	p := new(web.UserRegistration)

	//parse raw body to variable p
	if err := c.BodyParser(p); err != nil {
		return err
	}

	//call func login in user service
	userReg, err := u.UserService.Registration(c.Context(), p)

	return helper.PrintResponse(err, userReg, c)

}
