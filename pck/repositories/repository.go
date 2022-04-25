package repositories

import (
	"github.com/jmoiron/sqlx"
	"reddit/models"
)

type Post interface {
	GetById(id string) (models.Post, error)
	GetList(page int, limit int) (models.OutputPostList, error)
	Create(post models.InputPost) (models.OutputPost, error)
	Update(post models.InputUpdatePost) error
	Delete(id string) error
}

type Repository struct {
	Post Post
}

func (r Repository) GetById(id string) (models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetList(page int, limit int) (models.OutputPostList, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Create(post models.InputPost) (models.OutputPost, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Update(post models.InputUpdatePost) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Post: NewPostPostgres(db),
	}
}
