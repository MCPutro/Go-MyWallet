package controller

import (
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type activityControllerImpl struct {
	activityService service.ActivityService
}

func (a *activityControllerImpl) DeleteActivity(c *fiber.Ctx) error {
	userId := c.Get("UserId")
	paramActivityId := c.Params("ActivityId")
	activityId, err := strconv.ParseUint(paramActivityId, 10, 32)

	err = a.activityService.DeleteActivity(c.UserContext(), uint32(activityId), userId)

	return helper.PrintResponse(err, nil, c)
}

func (a *activityControllerImpl) GetAllActivity(c *fiber.Ctx) error {
	userid := c.Get("userId")

	list, err := a.activityService.GetActivityList(c.UserContext(), userid) //"7e65f8d1-bd30-4d2c-95bb-5bcbdb0e5561")

	return helper.PrintResponse(err, list, c)
}

func (a *activityControllerImpl) GetActivityTypes(c *fiber.Ctx) error {
	activityType, err := a.activityService.GetActivityCategory(c.UserContext())

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

	activity, err := a.activityService.AddActivity(c.UserContext(), body)
	return helper.PrintResponse(err, activity, c)
}

func NewActivityController(activityService service.ActivityService) ActivityController {
	return &activityControllerImpl{activityService: activityService}
}
