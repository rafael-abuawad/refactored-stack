package utils

import "github.com/gofiber/fiber/v2"

type httpError struct {
	Statuscode int    `json:"statusCode"`
	Error      string `json:"error"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(&httpError{
		Statuscode: code,
		Error:      err.Error(),
	})
}
