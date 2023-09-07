package middleware

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rafael-abuawad/refactored-stack/utils/jwt"
)

func isAuthenticated(c *fiber.Ctx) bool {
	userIDStr := c.Cookies("UserID")
	if userIDStr == "" {
		log.Error("no user id")
		return false
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Error(err)
		return false
	}

	token := c.Cookies("Authorization")
	if token == "" {
		log.Error("empty token")
		return false
	}

	payload, err := jwt.Verify(token)
	if err != nil {
		log.Error(err)
		return false
	}

	if payload.ID != uint(userID) {
		log.Error("user id and payload id do not match")
		return false
	}

	return true
}

// Auth is the authentication middleware
func AuthMiddleware(c *fiber.Ctx) error {
	currentRoute := c.OriginalURL()

	if isAuthenticated(c) {
		if strings.HasPrefix(currentRoute, "/auth/") {
			return c.Redirect("/")
		}
		return c.Next()
	}

	if strings.HasPrefix(currentRoute, "/auth/") {
		return c.Next()
	}

	return c.Redirect("/auth/login")
}
