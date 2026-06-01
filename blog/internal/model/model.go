package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var validate = validator.New()

func Validate(value any) error {
	return validate.Struct(value)
}

type Role string

const (
	UserRole  Role = "USER"
	AdminRole Role = "ADMIN"
)

type BlogPost struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Author       string    `json:"author"`
	Is_Published bool      `json:"is_published"`
	Created_At   time.Time `json:"created_at"`
}

type CreateBlogRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=200"`
	Content     string `json:"content" validate:"required,min=10"`
	IsPublished bool   `json:"is_published"`
}

type UpdateBlogRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=200"`
	Content     string `json:"content" validate:"required,min=10"`
	IsPublished bool   `json:"is_published"`
}

type Comment struct {
	ID        int       `json:"id"`
	BlogID    int       `json:"blog_id"`
	UserID    int       `json:"user_id"`
	UserName  string    `json:"user_name,omitempty"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Like struct {
	ID        int       `json:"id"`
	BlogID    int       `json:"blog_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateCommentRequest struct {
	BlogID  int    `json:"blog_id" validate:"required,gt=0"`
	Content string `json:"content" validate:"required,min=1,max=1000"`
}

type CreateLikeRequest struct {
	BlogID int `json:"blog_id" validate:"required,gt=0"`
}

type LikeSummary struct {
	BlogID int `json:"blog_id"`
	Count  int `json:"count"`
}

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"-"`
	Role      Role      `json:"role" validate:"required,oneof=USER ADMIN"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     Role   `json:"role,omitempty" validate:"omitempty,oneof=USER ADMIN"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
