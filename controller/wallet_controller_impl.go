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
}

func (w *walletControllerImpl) GetWalletByUID(c *fiber.Ctx) error {
	userid := c.Get("userid")
	walletsByUserId, err := w.walletService.GetWalletByUserId(c.Context(), userid)

	return helper.PrintResponse(err, walletsByUserId, c)
}

func (w *walletControllerImpl) GetWalletById(c *fiber.Ctx) error {
	fmt.Println(c.UserContext().Value(fiber.HeaderXRequestID).(string))

	userid := c.Get("UserId")
	//param := c.Get("wallet")
	paramWalletId := c.Params("WalletId")

	walletId, err := strconv.ParseUint(paramWalletId, 10, 32)
	if err != nil {
		return helper.PrintResponse(err, nil, c)
	}

	getWalletById, err := w.walletService.GetWalletById(c.UserContext(), userid, uint32(walletId))
	c.Set(fiber.HeaderXRequestID, c.UserContext().Value(fiber.HeaderXRequestID).(string))
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
	userid := c.Get("UserId")

	body := new(web.WalletReq)
	//body.UserId = userid

	if err := c.BodyParser(body); err != nil {
		return c.JSON(web.Response{
			Status:  "ERROR",
			Message: err.Error(),
			Data:    nil,
		})
	}

	wallet, err := w.walletService.AddWallet(c.Context(), &model.Wallet{
		UserId: userid,
		Name:   body.Name,
		Type:   body.Type,
		Amount: body.Amount,
	})

	return helper.PrintResponse(err, wallet, c)
}

func (w *walletControllerImpl) DeleteWallet(c *fiber.Ctx) error {
	//body := new(model.Wallet)
	////body.UserId = userid
	//
	//if err := c.BodyParser(body); err != nil {
	//	return c.JSON(web.Response{
	//		Status:  "ERROR",
	//		Message: err.Error(),
	//		Data:    nil,
	//	})
	//}

	userid := c.Get("UserId")
	paramWalletId := c.Params("WalletId")

	walletId, err := strconv.ParseUint(paramWalletId, 10, 32)

	err = w.walletService.DeleteWallet(c.Context(), userid, uint32(walletId))

	return helper.PrintResponse(err, nil, c)
}

func NewWalletController(walletService service.WalletService) WalletController {
	return &walletControllerImpl{walletService: walletService}
}
