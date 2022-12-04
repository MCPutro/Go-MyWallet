package controller

import (
	"github.com/MCPutro/Go-MyWallet/entity/web"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/gofiber/fiber/v2"
)

type activityControllerImpl struct {
	activityService service.ActivityService
}

func NewActivityController(activityService service.ActivityService) ActivityController {
	return &activityControllerImpl{activityService: activityService}
}

func (a *activityControllerImpl) GetActivityTypes(c *fiber.Ctx) error {
	activityType, err := a.activityService.GetActivityType(c.Context())

	if err != nil {
		return c.JSON(web.Response{
			Status:  "ERROR",
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		return c.JSON(activityType)
	}
}
