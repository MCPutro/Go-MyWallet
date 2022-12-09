package controller

import (
	"fmt"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/MCPutro/Go-MyWallet/entity/web"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type walletControllerImpl struct {
	walletService service.WalletService
}

func (w *walletControllerImpl) GetWalletType(c *fiber.Ctx) error {
	walletType, err := w.walletService.GetWalletType(c.Context())

	return helper.PrintResponse(err, walletType, c)
	//if err != nil {
	//	//c.Status(fiber.StatusUnauthorized)
	//	//return c.SendString(fmt.Sprint("error :", err))
	//	return c.JSON(web.Response{
	//		Status:  "ERROR",
	//		Message: err.Error(),
	//		Data:    nil,
	//	})
	//} else {
	//	c.Status(fiber.StatusOK)
	//	return c.JSON(web.Response{
	//		Status:  "SUCCESS",
	//		Message: nil,
	//		Data:    walletType,
	//	})
	//}
}

func (w *walletControllerImpl) GetWalletByUID(c *fiber.Ctx) error {
	userid := c.Get("userid")
	walletsByUserId, err := w.walletService.GetWalletByUserId(c.Context(), userid)

	return helper.PrintResponse(err, walletsByUserId, c)
	//if err != nil {
	//	//c.Status(fiber.StatusUnauthorized)
	//	return c.JSON(web.Response{
	//		Status:  "ERROR",
	//		Message: err.Error(),
	//		Data:    nil,
	//	})
	//} else {
	//	c.Status(fiber.StatusOK)
	//	return c.JSON(web.Response{
	//		Status:  "SUCCESS",
	//		Message: nil,
	//		Data:    walletsByUserId,
	//	})
	//}
}

func (w *walletControllerImpl) GetWalletId(c *fiber.Ctx) error {
	userid := c.Get("userid")
	param := c.Get("wallet")

	walletId, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return helper.PrintResponse(err, nil, c)
	}

	fmt.Println(userid, walletId)

	getWalletById, err := w.walletService.GetWalletById(c.Context(), userid, uint32(walletId))
	return helper.PrintResponse(err, getWalletById, c)
}

//func (w *walletControllerImpl) UpdateWallet(c *fiber.Ctx) error {
//	body := new(model.Wallet)
//	//body.UserId = userid
//
//	if err := c.BodyParser(body); err != nil {
//		return c.JSON(web.Response{
//			Status:  "ERROR",
//			Message: err.Error(),
//			Data:    nil,
//		})
//	}
//
//	wallet, err := w.walletService.UpdateWallet(c.Context(), body)
//
//	return helper.PrintResponse(err, wallet, c)
//}

func (w *walletControllerImpl) AddWallet(c *fiber.Ctx) error {
	//userid := c.Get("userid")

	body := new(model.Wallet)
	//body.UserId = userid

	if err := c.BodyParser(body); err != nil {
		return c.JSON(web.Response{
			Status:  "ERROR",
			Message: err.Error(),
			Data:    nil,
		})
	}

	wallet, err := w.walletService.AddWallet(c.Context(), body)

	return helper.PrintResponse(err, wallet, c)
	//if err != nil {
	//	//c.Status(fiber.StatusUnauthorized)
	//	//return c.SendString(fmt.Sprint("error :", err))
	//	return c.JSON(web.Response{
	//		Status:  "ERROR",
	//		Message: err.Error(),
	//		Data:    nil,
	//	})
	//} else {
	//	c.Status(fiber.StatusOK)
	//	return c.JSON(web.Response{
	//		Status:  "SUCCESS",
	//		Message: nil,
	//		Data:    wallet,
	//	})
	//}
}

func (w *walletControllerImpl) DeleteWallet(c *fiber.Ctx) error {
	body := new(model.Wallet)
	//body.UserId = userid

	if err := c.BodyParser(body); err != nil {
		return c.JSON(web.Response{
			Status:  "ERROR",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err := w.walletService.DeleteWallet(c.Context(), body.UserId, body.WalletId)

	return helper.PrintResponse(err, nil, c)
}

func NewWalletController(walletService service.WalletService) WalletController {
	return &walletControllerImpl{walletService: walletService}
}
