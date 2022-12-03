package controller

import "github.com/gofiber/fiber/v2"

type WalletController interface {
	AddWallet(c *fiber.Ctx) error
	GetWalletByUid(c *fiber.Ctx) error
	GetWalletType(c *fiber.Ctx) error
}
