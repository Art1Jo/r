package services

import (
	"reddit/models"
	"reddit/pck/repositories"
)

type Post interface {
	GetById(id string) (models.Post, error)
	GetList(page int, limit int) (models.OutputPostList, error)
	Create(post models.InputPost) (models.OutputPost, error)
	Update(post models.InputUpdatePost) error
	Delete(id string) error
}

type Service struct {
	Post Post
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Post: NewPostService(repos.Post),
	}
}
