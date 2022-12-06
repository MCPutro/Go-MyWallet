package controller

import "github.com/gofiber/fiber/v2"

type WalletController interface {
	AddWallet(c *fiber.Ctx) error
	UpdateWallet(c *fiber.Ctx) error
	GetWalletByUID(c *fiber.Ctx) error
	GetWalletId(c *fiber.Ctx) error
	GetWalletType(c *fiber.Ctx) error
}
