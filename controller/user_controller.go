package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	Login(c *fiber.Ctx) error
	Registration(c *fiber.Ctx) error
}
