package repository

import (
	"restapi/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(email, password string) (entity.User, error)
}

type UploadImage interface {
	UploadImage(id int, image entity.Image) (int, error)
	GetAll(userId int) ([]entity.Image, error)
	GetById(userId, imageId int) (entity.Image, error)
	Delete(userId, imageId int) error
}

type Repository struct {
	Authorization
	UploadImage
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UploadImage:   NewUploadImagePostgres(db),
	}
}
