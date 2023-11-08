package repository

import (
	"restapi/meet"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user meet.User) (int, error)
	GetUser(email, password string) (meet.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
