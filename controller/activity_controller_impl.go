package controller

import (
	"fmt"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/gofiber/fiber/v2"
)

type activityControllerImpl struct {
	activityService service.ActivityService
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

func (a *activityControllerImpl) AddActivity(c *fiber.Ctx) error {
	body := new(model.Activity)

	//parse data
	if err := c.BodyParser(body); err != nil {
		return c.JSON(web.Response{
			Status:  "ERROR",
			Message: err.Error(),
			Data:    nil,
		})
	}

	fmt.Println(">>>", body)

	activity, err := a.activityService.AddActivity(c.Context(), body)
	return helper.PrintResponse(err, activity, c)
}

func NewActivityController(activityService service.ActivityService) ActivityController {
	return &activityControllerImpl{activityService: activityService}
}
