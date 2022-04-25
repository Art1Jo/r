package services

import (
	"reddit/models"
	"reddit/pck/repositories"
)

type PostService struct {
	repo repositories.Post
}

func NewPostService(repo repositories.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) GetById(id string) (models.Post, error) {

}

func (s *PostService) GetList(page int, limit int) (models.OutputPostList, error) {

}

func (s *PostService) Create(post models.InputPost) (models.OutputPost, error) {

}

func (s *PostService) Update(post models.InputUpdatePost) error {

}

func (s *PostService) Delete(id string) error {

}
