package service

import (
	"blog/internal/model"
	"blog/internal/repository"
	"errors"
)

var (
	ErrBlogNotFound = errors.New("blog not found")
	ErrLikeExists   = errors.New("blog already liked by user")
	ErrLikeNotFound = errors.New("like not found")
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
	updated, err := s.Repo.UpdateBlog(blog)
	if err != nil {
		return err
	}
	if !updated {
		return ErrBlogNotFound
	}
	return nil
}

func (s *BlogService) DeleteBlog(id int) error {
	deleted, err := s.Repo.DeleteBlog(id)
	if err != nil {
		return err
	}
	if !deleted {
		return ErrBlogNotFound
	}
	return nil
}

func (s *BlogService) CreateComment(comment model.Comment) (model.Comment, error) {
	exists, err := s.Repo.BlogExists(comment.BlogID)
	if err != nil {
		return model.Comment{}, err
	}
	if !exists {
		return model.Comment{}, ErrBlogNotFound
	}

	return s.Repo.CreateComment(comment)
}

func (s *BlogService) GetComments(blogID int) ([]model.Comment, error) {
	exists, err := s.Repo.BlogExists(blogID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrBlogNotFound
	}

	return s.Repo.GetComments(blogID)
}

func (s *BlogService) AddLike(blogID, userID int) error {
	exists, err := s.Repo.BlogExists(blogID)
	if err != nil {
		return err
	}
	if !exists {
		return ErrBlogNotFound
	}

	err = s.Repo.AddLike(model.Like{
		BlogID: blogID,
		UserID: userID,
	})
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateLike) {
			return ErrLikeExists
		}
		return err
	}

	return nil
}

func (s *BlogService) RemoveLike(blogID, userID int) error {
	exists, err := s.Repo.BlogExists(blogID)
	if err != nil {
		return err
	}
	if !exists {
		return ErrBlogNotFound
	}

	removed, err := s.Repo.RemoveLike(blogID, userID)
	if err != nil {
		return err
	}
	if !removed {
		return ErrLikeNotFound
	}

	return nil
}

func (s *BlogService) GetLikeSummary(blogID int) (model.LikeSummary, error) {
	exists, err := s.Repo.BlogExists(blogID)
	if err != nil {
		return model.LikeSummary{}, err
	}
	if !exists {
		return model.LikeSummary{}, ErrBlogNotFound
	}

	count, err := s.Repo.GetLikeCount(blogID)
	if err != nil {
		return model.LikeSummary{}, err
	}

	return model.LikeSummary{
		BlogID: blogID,
		Count:  count,
	}, nil
}
