package controller

import (
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/gofiber/fiber/v2"
)

type activityControllerImpl struct {
	activityService service.ActivityService
}

func (a *activityControllerImpl) DeleteActivity(c *fiber.Ctx) error {
	//userid := c.Get("userid")
	//param := c.Get("activityId")
	//activityId, err := strconv.ParseUint(param, 10, 32)
	//if err != nil {
	//	return c.SendString(err.Error())
	//}

	var body model.Activity

	//parse data
	if err := c.BodyParser(&body); err != nil {
		return c.JSON(web.Response{
			Status:  "ERROR",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err := a.activityService.DeleteActivity(c.Context(), body.ActivityId, body.UserId)

	return helper.PrintResponse(err, nil, c)
}

func (a *activityControllerImpl) GetAllActivity(c *fiber.Ctx) error {
	userid := c.Get("userId")

	list, err := a.activityService.GetActivityList(c.Context(), userid) //"7e65f8d1-bd30-4d2c-95bb-5bcbdb0e5561")

	return helper.PrintResponse(err, list, c)
}

func (a *activityControllerImpl) GetActivityTypes(c *fiber.Ctx) error {
	activityType, err := a.activityService.GetActivityCategory(c.Context())

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

	activity, err := a.activityService.AddActivity(c.Context(), body)
	return helper.PrintResponse(err, activity, c)
}

func NewActivityController(activityService service.ActivityService) ActivityController {
	return &activityControllerImpl{activityService: activityService}
}
