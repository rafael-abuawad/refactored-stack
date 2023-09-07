package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rafael-abuawad/refactored-stack/app/services"
)

func Login(c *fiber.Ctx) error {
	switch c.Method() {
	case http.MethodPost:
		payload, err := services.Login(c)
		if err != nil {
			return err
		}

		services.GenereateAuthCookies(c, *payload)
		return c.Redirect("/")
	}

	return c.Render("login", fiber.Map{})
}

func Signup(c *fiber.Ctx) error {
	switch c.Method() {
	case http.MethodPost:
		payload, err := services.Signup(c)
		if err != nil {
			return err
		}

		services.GenereateAuthCookies(c, *payload)
		return c.Redirect("/")
	}

	return c.Render("signup", fiber.Map{})
}

func Logout(c *fiber.Ctx) error {
	services.Logout(c)
	return c.Redirect("/auth/login")
}

func ForgotPassword(c *fiber.Ctx) error {
	switch c.Method() {
	case http.MethodPost:
		if err := services.ForgotPassword(c); err != nil {
			return err
		}
		return c.Render("forgot-password-done", fiber.Map{})
	}

	return c.Render("forgot-password", fiber.Map{})
}

func ResetPassword(c *fiber.Ctx) error {
	switch c.Method() {
	case http.MethodPost:
		if err := services.ResetPassword(c); err != nil {
			return err
		}
		return c.Render("reset-password-done", fiber.Map{})
	}

	return c.Render("reset-password", fiber.Map{})
}

func AuthController(app fiber.Router) {
	r := app.Group("/auth")

	// auth/login
	r.Get("/login", Login)
	r.Post("/login", Login)

	// auth/signup
	r.Get("/signup", Signup)
	r.Post("/signup", Signup)

	// auth/logout
	r.Get("/logout", Logout)

	// auth/forgot-password
	r.Get("/forgot-password", ForgotPassword)
	r.Post("/forgot-password", ForgotPassword)

	// auth/reset-password/:token
	r.Get("/reset-password/:token", ResetPassword)
	r.Post("/reset-password/:token", ResetPassword)
}
