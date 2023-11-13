package repository

import (
	"fmt"
	"restapi/meet"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user meet.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, age, password) values($1, $2, $3, $4) RETURNING id", UserTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Age, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (meet.User, error) {
	var user meet.User
	query := fmt.Sprintf("SELECT id FROM  %s WHERE email=$1 AND password=$2", UserTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}
