package controller

import "github.com/gofiber/fiber/v2"

type WalletController interface {
	AddWallet(c *fiber.Ctx) error
	GetWalletByUID(c *fiber.Ctx) error
	GetWalletById(c *fiber.Ctx) error
	GetWalletType(c *fiber.Ctx) error
	DeleteWallet(c *fiber.Ctx) error
}
