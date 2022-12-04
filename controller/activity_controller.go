package controller

import "github.com/gofiber/fiber/v2"

type ActivityController interface {
	GetActivityTypes(c *fiber.Ctx) error
}
