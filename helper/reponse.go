package helper

import (
	"github.com/MCPutro/Go-MyWallet/entity/web"
	"github.com/gofiber/fiber/v2"
)

func PrintResponse(err error, data interface{}, c *fiber.Ctx) error {
	if err != nil {
		//c.Status(fiber.StatusUnauthorized)
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
			Data:    data,
		})
	}
}
