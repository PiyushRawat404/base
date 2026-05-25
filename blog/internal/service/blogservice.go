package service

import (
	"blog/internal/model"
	"blog/internal/repository"
)

type BlogService struct {
	Repo repository.BlogRepository
}

func NewBlogService(repo repository.BlogRepository) *BlogService {
	return &BlogService{
		Repo: repo,
	}
}

func (s *BlogService) CreateBlog(blog model.BlogPost) error {
	return s.Repo.CreateBlog(blog)
}

func (s *BlogService) GetBlogs() ([]model.BlogPost, error) {
	return s.Repo.GetBlogs()
}
func (s *BlogService) UpdateBlog(blog model.BlogPost) error {
	return s.Repo.UpdateBlog(blog)
}

func (s *BlogService) DeleteBlog(id int) error {
	return s.Repo.DeleteBlog(id)
}
