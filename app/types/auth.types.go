package types

type LoginDTO struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type SignupDTO struct {
	LoginDTO
	Name            string `validate:"required,min=3"`
	PasswordConfirm string `validate:"required,eqfield=Password"`
}

type ForgotPasswordDTO struct {
	Email string `validate:"required,email"`
}

type PasswordResetDTO struct {
	Password        string `validate:"required"`
	PasswordConfirm string `validate:"required,eqfield=Password"`
}

type UserResponse struct {
	ID         uint
	Name       string
	Email      string
	Password   string
	ResetToken string
}

type AccessResponse struct {
	Token string
}

type AuthResponse struct {
	User *UserResponse
	Auth *AccessResponse
}
