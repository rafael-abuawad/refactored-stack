package types

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignupDTO struct {
	LoginDTO
	Name            string `json:"name" validate:"required,min=3"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
}

type ForgotPasswordDTO struct {
	Email string `json:"email" validate:"required,email"`
}

type PasswordResetDTO struct {
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
}

type UserResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	ResetToken string `json:"-"`
}

type AccessResponse struct {
	Token string `json:"token"`
}

type AuthResponse struct {
	User *UserResponse   `json:"user"`
	Auth *AccessResponse `json:"auth"`
}
