package service

import (
	"context"
	"errors"
	"product/internal/models"
	"product/internal/repository"
	"product/pkg/utils"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrValidation        = errors.New("validation failed")
	ErrWeakPassword      = errors.New("password must be at least 8 characters and include upper, lower, number, and special character")
	ErrDuplicateEmail    = errors.New("email already exists")
	ErrInvalidCredential = errors.New("invalid credentials")
	ErrProductNotFound   = errors.New("product not found")
)

type Service struct {
	repo      *repository.Repository
	jwtSecret []byte
}

func NewService(repo *repository.Repository, jwtSecret string) *Service {
	return &Service{
		repo:      repo,
		jwtSecret: []byte(jwtSecret),
	}
}

func (s *Service) Signup(ctx context.Context, req models.SignupRequest) (*models.AuthResponse, error) {
	email := strings.TrimSpace(strings.ToLower(req.Email))

	existingUser, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrDuplicateEmail
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.CreateUser(ctx, email, string(passwordHash))
	if err != nil {
		return nil, err
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{Token: token, User: sanitizeUser(*user)}, nil
}

func (s *Service) Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error) {
	email := strings.TrimSpace(strings.ToLower(req.Email))
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidCredential
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredential
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{Token: token, User: sanitizeUser(*user)}, nil
}

func (s *Service) CreateProduct(ctx context.Context, product models.Product) (*models.Product, error) {
	return s.repo.CreateProduct(ctx, product)
}

func (s *Service) GetProducts(ctx context.Context, pagination utils.PaginationParams, filter utils.ProductFilter) (*models.ProductListResponse, error) {
	products, total, err := s.repo.GetProducts(ctx, pagination, filter)
	if err != nil {
		return nil, err
	}

	return &models.ProductListResponse{
		Data:       products,
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		Total:      total,
		TotalPages: utils.CalculateTotalPages(total, pagination.Limit),
	}, nil
}

func (s *Service) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, ErrProductNotFound
	}
	return product, nil
}

func (s *Service) UpdateProduct(ctx context.Context, id int, product models.Product) (*models.Product, error) {
	updated, err := s.repo.UpdateProduct(ctx, id, product)
	if err != nil {
		return nil, err
	}
	if updated == nil {
		return nil, ErrProductNotFound
	}
	return updated, nil
}

func (s *Service) DeleteProduct(ctx context.Context, id int) error {
	deleted, err := s.repo.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}
	if !deleted {
		return ErrProductNotFound
	}
	return nil
}

func (s *Service) generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

func sanitizeUser(user models.User) models.User {
	user.PasswordHash = ""
	return user
}
