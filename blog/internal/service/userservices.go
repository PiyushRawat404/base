package service

import (
	"blog/internal/model"
	"blog/internal/repository"
	"blog/pkg/utils"
	"errors"
	"strings"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type UserService struct {
	Repo      repository.UserRepository
	JWTSecret string
}

func NewUserService(repo repository.UserRepository, jwtSecret string) *UserService {
	return &UserService{
		Repo:      repo,
		JWTSecret: jwtSecret,
	}
}

func (s *UserService) CreateUser(user model.User) (model.User, error) {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))

	if user.Role == "" {
		user.Role = model.UserRole
	}

	if err := user.ValidateForCreate(); err != nil {
		return model.User{}, err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return model.User{}, err
	}

	user.Password = hashedPassword

	createdUser, err := s.Repo.CreateUser(user)
	if err != nil {
		if repository.IsUniqueViolation(err) {
			return model.User{}, ErrEmailAlreadyExists
		}
		return model.User{}, err
	}

	createdUser.Password = ""
	return createdUser, nil
}

func (s *UserService) Login(email, password string) (string, model.User, error) {
	loginUser := model.User{
		Email:    strings.ToLower(strings.TrimSpace(email)),
		Password: password,
	}

	if err := loginUser.ValidateForLogin(); err != nil {
		return "", model.User{}, err
	}

	user, err := s.Repo.GetUserByEmail(loginUser.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return "", model.User{}, ErrInvalidCredentials
		}
		return "", model.User{}, err
	}

	if err := utils.ComparePassword(user.Password, loginUser.Password); err != nil {
		return "", model.User{}, ErrInvalidCredentials
	}

	token, err := utils.GenerateJWT(s.JWTSecret, user)
	if err != nil {
		return "", model.User{}, err
	}

	user.Password = ""
	return token, user, nil
}
