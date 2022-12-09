package controller

import "github.com/gofiber/fiber/v2"

type ActivityController interface {
	GetAllActivity(c *fiber.Ctx) error
	GetActivityTypes(c *fiber.Ctx) error
	AddActivity(c *fiber.Ctx) error
}
