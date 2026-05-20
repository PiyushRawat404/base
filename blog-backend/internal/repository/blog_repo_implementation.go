package repository

import (
	"blog-backend/internal/model"
	"context"

	"github.com/jackc/pgx/v5"
)

type BlogRepo struct {
	DB *pgx.Conn
}

func NewBlogRepository(db *pgx.Conn) BlogRepository {
	return &BlogRepo{
		DB: db,
	}
}

func (r *BlogRepo) CreateBlog(blog model.BlogPost) error {

	query := `
	INSERT INTO blogs(title, content)
	VALUES($1, $2)
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		blog.Title,
		blog.Content,
	)

	return err
}

func (r *BlogRepo) GetBlogs() ([]model.BlogPost, error) {

	query := `
	SELECT id, title, content
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
			&blog.ID,
			&blog.Title,
			&blog.Content,
		)

		if err != nil {
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (r *BlogRepo) UpdateBlog(blog model.BlogPost) error {

	query := `
	UPDATE blogs
	SET title=$1, content=$2
	WHERE id=$3
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		blog.Title,
		blog.Content,
		blog.ID,
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
