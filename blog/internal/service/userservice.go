package service

import (
	"blog/internal/auth"
	"blog/internal/model"
	"blog/internal/repository"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAdminBootstrapOnly = errors.New("admin registration is only allowed for the first account")
)

type UserService struct {
	repo        repository.UserRepository
	tokenSecret string
	tokenTTL    time.Duration
}

func NewUserService(repo repository.UserRepository, tokenSecret string) *UserService {
	return &UserService{
		repo:        repo,
		tokenSecret: tokenSecret,
		tokenTTL:    24 * time.Hour,
	}
}

func (s *UserService) Register(request model.RegisterUserRequest) (model.AuthResponse, error) {
	if err := model.Validate(request); err != nil {
		return model.AuthResponse{}, err
	}

	existingUser, err := s.repo.GetUserByEmail(strings.ToLower(request.Email))
	if err == nil && existingUser.ID != 0 {
		return model.AuthResponse{}, ErrEmailAlreadyExists
	}
	if err != nil && !errors.Is(err, repository.ErrUserNotFound) {
		return model.AuthResponse{}, err
	}

	count, err := s.repo.CountUsers()
	if err != nil {
		return model.AuthResponse{}, err
	}

	role := model.UserRole
	if count == 0 {
		role = model.AdminRole
	} else if request.Role == model.AdminRole {
		return model.AuthResponse{}, ErrAdminBootstrapOnly
	} else if request.Role == model.UserRole {
		role = request.Role
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.AuthResponse{}, err
	}

	user, err := s.repo.CreateUser(model.User{
		Name:     request.Name,
		Email:    strings.ToLower(request.Email),
		Password: string(passwordHash),
		Role:     role,
	})
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate key") {
			return model.AuthResponse{}, ErrEmailAlreadyExists
		}
		return model.AuthResponse{}, err
	}

	token, err := auth.GenerateToken(user, s.tokenSecret, s.tokenTTL)
	if err != nil {
		return model.AuthResponse{}, err
	}

	user.Password = ""

	return model.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *UserService) Login(request model.LoginRequest) (model.AuthResponse, error) {
	if err := model.Validate(request); err != nil {
		return model.AuthResponse{}, err
	}

	user, err := s.repo.GetUserByEmail(strings.ToLower(request.Email))
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return model.AuthResponse{}, ErrInvalidCredentials
		}
		return model.AuthResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return model.AuthResponse{}, ErrInvalidCredentials
	}

	token, err := auth.GenerateToken(user, s.tokenSecret, s.tokenTTL)
	if err != nil {
		return model.AuthResponse{}, err
	}

	user.Password = ""

	return model.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *UserService) GetProfile(userID int) (model.User, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return model.User{}, err
	}

	user.Password = ""
	return user, nil
}

func (s *UserService) GetUsers() ([]model.User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	for index := range users {
		users[index].Password = ""
	}

	return users, nil
}
