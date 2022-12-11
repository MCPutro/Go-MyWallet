package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func CustomMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get(fiber.HeaderAuthorization, "xxx")
		if !strings.HasPrefix(auth, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).SendString("auth nya salah bro")
		}

		//fmt.Println(c.)
		return c.Next()
	}
}
