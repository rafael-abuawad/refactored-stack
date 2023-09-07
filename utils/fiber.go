package utils

import "github.com/gofiber/fiber/v2"

func ParseBody(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ParseBody(ctx, body); err != nil {
		return err
	}

	return Validate(body)
}

func GetUser(c *fiber.Ctx) *uint {
	id, _ := c.Locals("USER").(uint)
	return &id
}
