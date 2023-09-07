package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
)

func FlashMiddleware(c *fiber.Ctx) error {
	c.Locals("Flash", flash.Get(c))
	return c.Next()
}
