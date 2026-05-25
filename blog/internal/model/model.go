package model

import (
	"time"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Role string
const (
	UserRole Role ="USER";
	AdminRole Role="ADMIN"	
)
type BlogPost struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
	Is_Published bool `json:"is_published"`
	Created_At time.Time `json:"created_at"`
}

type User struct{
	ID int `json:"id" validate:"required,min=8"`
	Name string `json:"name" validate:"required,min=12"`
	Email    string  `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role Role `json:"role" validate:"required,oneof=USER ADMIN"`
}