package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetCookie(c *fiber.Ctx, name string, value string, duration string) {
	v, err := time.ParseDuration(duration)
	if err != nil {
		panic(err.Error())
	}

	expiration := time.Now().Add(v)
	c.Cookie(buildCookie(name, value, expiration))
}

func ClearCookie(c *fiber.Ctx, name string) {
	c.Cookie(buildCookie(name, "", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)))
}

func buildCookie(name string, value string, expires time.Time) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expires

	return cookie
}
