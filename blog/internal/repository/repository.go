package repository

import (
	"blog/internal/model"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type BlogRepository interface {
	CreateBlog(blog model.BlogPost) error
	GetBlogs() ([]model.BlogPost, error)
	UpdateBlog(blog model.BlogPost) error
	DeleteBlog(id int) error
}

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
}

var ErrUserNotFound = errors.New("user not found")

type BlogRepo struct {
	DB *pgx.Conn
}

type UserRepo struct {
	DB *pgx.Conn
}

func NewBlogRepository(db *pgx.Conn) BlogRepository {
	return &BlogRepo{
		DB: db,
	}
}

func NewUserRepository(db *pgx.Conn) UserRepository {
	return &UserRepo{
		DB: db,
	}
}

func (r *BlogRepo) CreateBlog(blog model.BlogPost) error {
	query := `
	INSERT INTO blogs(title, content, author, is_published)
	VALUES($1, $2, $3, $4)
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		blog.Title,
		blog.Content,
		blog.Author,
		blog.Is_Published,
	)

	return err
}

func (r *BlogRepo) GetBlogs() ([]model.BlogPost, error) {
	query := `
	SELECT id, title, content, author, is_published, created_at
	FROM blogs
	`

	rows, err := r.DB.Query(
		context.Background(),
		query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []model.BlogPost

	for rows.Next() {
		var blog model.BlogPost

		err := rows.Scan(
			&blog.Id,
			&blog.Title,
			&blog.Content,
			&blog.Author,
			&blog.Is_Published,
			&blog.Created_At,
		)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (r *BlogRepo) UpdateBlog(blog model.BlogPost) error {
	query := `
	UPDATE blogs
	SET title=$1, content=$2, author=$3, is_published=$4
	WHERE id=$5
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		blog.Title,
		blog.Content,
		blog.Author,
		blog.Is_Published,
		blog.Id,
	)

	return err
}

func (r *BlogRepo) DeleteBlog(id int) error {
	query := `
	DELETE FROM blogs
	WHERE id=$1
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}

func (r *UserRepo) CreateUser(user model.User) (model.User, error) {
	query := `
	INSERT INTO users(name, email, password, role)
	VALUES($1, $2, $3, $4)
	RETURNING id, name, email, role
	`

	var createdUser model.User
	err := r.DB.QueryRow(
		context.Background(),
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
	).Scan(
		&createdUser.ID,
		&createdUser.Name,
		&createdUser.Email,
		&createdUser.Role,
	)
	if err != nil {
		return model.User{}, err
	}

	return createdUser, nil
}

func (r *UserRepo) GetUserByEmail(email string) (model.User, error) {
	query := `
	SELECT id, name, email, password, role
	FROM users
	WHERE email=$1
	`

	var user model.User
	err := r.DB.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}

	return user, nil
}

func IsUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return false
	}

	return pgErr.Code == "23505"
}
