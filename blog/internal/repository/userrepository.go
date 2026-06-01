package repository

import (
	"blog/internal/model"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	GetUserByID(id int) (model.User, error)
	GetUsers() ([]model.User, error)
	CountUsers() (int, error)
}

type UserRepo struct {
	DB *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) UserRepository {
	return &UserRepo{DB: db}
}

func (r *UserRepo) CreateUser(user model.User) (model.User, error) {
	query := `
	INSERT INTO users(name, email, password_hash, role)
	VALUES($1, $2, $3, $4)
	RETURNING id, created_at
	`

	err := r.DB.QueryRow(
		context.Background(),
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
	).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepo) GetUserByEmail(email string) (model.User, error) {
	query := `
	SELECT id, name, email, password_hash, role, created_at
	FROM users
	WHERE email = $1
	`

	var user model.User
	err := r.DB.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepo) GetUserByID(id int) (model.User, error) {
	query := `
	SELECT id, name, email, password_hash, role, created_at
	FROM users
	WHERE id = $1
	`

	var user model.User
	err := r.DB.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepo) GetUsers() ([]model.User, error) {
	query := `
	SELECT id, name, email, password_hash, role, created_at
	FROM users
	ORDER BY created_at DESC
	`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) CountUsers() (int, error) {
	query := `SELECT COUNT(*) FROM users`

	var count int
	if err := r.DB.QueryRow(context.Background(), query).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
