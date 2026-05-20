package repository

import "blog-backend/internal/model"

type BlogRepository interface {
	CreateBlog(blog model.BlogPost) error
	GetBlogs() ([]model.BlogPost, error)
	UpdateBlog(blog model.BlogPost) error
	DeleteBlog(id int) error
}
