package models

import (
	"time"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func init() {
	_ = Validate.RegisterValidation("strong_password", func(fl validator.FieldLevel) bool {
		return IsStrongPassword(fl.Field().String())
	})
}

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type SignupRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,strong_password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"required,min=2,max=255"`
	Description string    `json:"description" validate:"max=1000"`
	Price       float64   `json:"price" validate:"required,gt=0"`
	Quantity    int       `json:"quantity" validate:"gte=0"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductListResponse struct {
	Data       []Product `json:"data"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
	Total      int       `json:"total"`
	TotalPages int       `json:"total_pages"`
}

func IsStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		default:
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}
