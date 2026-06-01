package repository

import (
	"blog/internal/model"
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
)

var ErrDuplicateLike = errors.New("like already exists")

type BlogRepository interface {
	CreateBlog(blog model.BlogPost) error
	GetBlogs() ([]model.BlogPost, error)
	UpdateBlog(blog model.BlogPost) (bool, error)
	DeleteBlog(id int) (bool, error)
	BlogExists(blogID int) (bool, error)
	CreateComment(comment model.Comment) (model.Comment, error)
	GetComments(blogID int) ([]model.Comment, error)
	AddLike(like model.Like) error
	RemoveLike(blogID, userID int) (bool, error)
	GetLikeCount(blogID int) (int, error)
}

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
	ORDER BY created_at DESC
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

func (r *BlogRepo) UpdateBlog(blog model.BlogPost) (bool, error) {
	query := `
	UPDATE blogs
	SET title=$1, content=$2, author=$3, is_published=$4
	WHERE id=$5
	`

	result, err := r.DB.Exec(
		context.Background(),
		query,
		blog.Title,
		blog.Content,
		blog.Author,
		blog.Is_Published,
		blog.Id,
	)
	if err != nil {
		return false, err
	}

	return result.RowsAffected() > 0, nil
}

func (r *BlogRepo) DeleteBlog(id int) (bool, error) {
	query := `
	DELETE FROM blogs
	WHERE id=$1
	`

	result, err := r.DB.Exec(
		context.Background(),
		query,
		id,
	)
	if err != nil {
		return false, err
	}

	return result.RowsAffected() > 0, nil
}

func (r *BlogRepo) BlogExists(blogID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM blogs WHERE id = $1)`

	var exists bool
	if err := r.DB.QueryRow(context.Background(), query, blogID).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *BlogRepo) CreateComment(comment model.Comment) (model.Comment, error) {
	query := `
	INSERT INTO comments(blog_id, user_id, content)
	VALUES($1, $2, $3)
	RETURNING id, created_at
	`

	err := r.DB.QueryRow(
		context.Background(),
		query,
		comment.BlogID,
		comment.UserID,
		comment.Content,
	).Scan(&comment.ID, &comment.CreatedAt)
	if err != nil {
		return model.Comment{}, err
	}

	return comment, nil
}

func (r *BlogRepo) GetComments(blogID int) ([]model.Comment, error) {
	query := `
	SELECT c.id, c.blog_id, c.user_id, u.name, c.content, c.created_at
	FROM comments c
	INNER JOIN users u ON u.id = c.user_id
	WHERE c.blog_id = $1
	ORDER BY c.created_at DESC
	`

	rows, err := r.DB.Query(context.Background(), query, blogID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []model.Comment
	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.BlogID,
			&comment.UserID,
			&comment.UserName,
			&comment.Content,
			&comment.CreatedAt,
		); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *BlogRepo) AddLike(like model.Like) error {
	query := `
	INSERT INTO likes(blog_id, user_id)
	VALUES($1, $2)
	`

	_, err := r.DB.Exec(context.Background(), query, like.BlogID, like.UserID)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate key") {
			return ErrDuplicateLike
		}
		return err
	}

	return nil
}

func (r *BlogRepo) RemoveLike(blogID, userID int) (bool, error) {
	query := `
	DELETE FROM likes
	WHERE blog_id = $1 AND user_id = $2
	`

	result, err := r.DB.Exec(context.Background(), query, blogID, userID)
	if err != nil {
		return false, err
	}

	return result.RowsAffected() > 0, nil
}

func (r *BlogRepo) GetLikeCount(blogID int) (int, error) {
	query := `
	SELECT COUNT(*)
	FROM likes
	WHERE blog_id = $1
	`

	var count int
	if err := r.DB.QueryRow(context.Background(), query, blogID).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
