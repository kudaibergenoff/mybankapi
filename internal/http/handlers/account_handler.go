package handlers

import "github.com/gofiber/fiber/v2"

func GetAccounts(c *fiber.Ctx) error {
	return c.SendString("Hello, Zhanuzak!")
}
