package services

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rafael-abuawad/refactored-stack/app/models"
	"github.com/rafael-abuawad/refactored-stack/app/types"
	"github.com/rafael-abuawad/refactored-stack/config"
	"github.com/rafael-abuawad/refactored-stack/utils"
	"github.com/rafael-abuawad/refactored-stack/utils/jwt"
	"github.com/rafael-abuawad/refactored-stack/utils/password"
	"gorm.io/gorm"
)

func GenereateAuthCookies(c *fiber.Ctx, payload types.AuthResponse) {
	exp := config.TOKEN_EXPIRATION
	userId := strconv.FormatUint(uint64(payload.User.ID), 10)

	utils.SetCookie(c, "UserID", userId, exp)
	utils.SetCookie(c, "UserEmail", payload.User.Email, exp)
	utils.SetCookie(c, "UserName", payload.User.Name, exp)
	utils.SetCookie(c, "Authorization", payload.Auth.Token, exp)
}

func Login(c *fiber.Ctx) (*types.AuthResponse, error) {
	b := types.LoginDTO{}

	if err := utils.ParseBodyAndValidate(c, &b); err != nil {
		return nil, err
	}

	u := &models.User{}

	err := models.FindUserByEmail(u, b.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	if err := password.Verify(u.Password, b.Password); err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	t := jwt.Generate(&jwt.TokenPayload{
		ID: u.ID,
	}, config.TOKEN_EXPIRATION)

	return &types.AuthResponse{
		User: &types.UserResponse{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		},
		Auth: &types.AccessResponse{
			Token: t,
		},
	}, nil
}

func Signup(c *fiber.Ctx) (*types.AuthResponse, error) {
	b := types.SignupDTO{}

	if err := utils.ParseBodyAndValidate(c, &b); err != nil {
		return nil, err
	}

	err := models.FindUserByEmail(&struct{ ID string }{}, b.Email).Error

	// If email already exists, return
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	user := &models.User{
		Name:     b.Name,
		Password: password.Generate(b.Password),
		Email:    b.Email,
	}

	// Create a user, if error return
	if err := models.CreateUser(user); err.Error != nil {
		return nil, fiber.NewError(fiber.StatusConflict, err.Error.Error())
	}

	// generate access token
	t := jwt.Generate(&jwt.TokenPayload{
		ID: user.ID,
	}, config.TOKEN_EXPIRATION)

	return &types.AuthResponse{
		User: &types.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Auth: &types.AccessResponse{
			Token: t,
		},
	}, nil
}

func ForgotPassword(c *fiber.Ctx) error {
	b := types.ForgotPasswordDTO{}

	if err := utils.ParseBodyAndValidate(c, &b); err != nil {
		return err
	}

	u := &models.User{}

	err := models.FindUserByEmail(u, b.Email).Error

	// If email already exists, return
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Email does not exists")
	}

	// Generate token for password reset
	t := jwt.Generate(&jwt.TokenPayload{
		ID: u.ID,
	}, "15m")

	// Update the user
	u.ResetToken = t

	// Update a user, if error return
	if err := models.UpdateUser(u.ID, u); err.Error != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error.Error())
	}

	return nil
}

func ResetPassword(c *fiber.Ctx) error {
	token := c.Params("token")

	if token == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	// Parse body
	b := types.PasswordResetDTO{}

	if err := utils.ParseBodyAndValidate(c, &b); err != nil {
		return err
	}

	payload, err := jwt.Verify(token)
	if err != nil {
		return err
	}

	u := &models.User{}

	err = models.FindUserByToken(u, token).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	if u.ID != payload.ID {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	// Update the user
	u.Password = password.Generate(b.Password)

	// Update a user, if error return
	if err := models.UpdateUser(u.ID, u); err.Error != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error.Error())
	}

	return nil
}

func Logout(c *fiber.Ctx) {
	utils.ClearCookie(c, "UserID")
	utils.ClearCookie(c, "UserEmail")
	utils.ClearCookie(c, "UserName")
	utils.ClearCookie(c, "Authorization")
}
