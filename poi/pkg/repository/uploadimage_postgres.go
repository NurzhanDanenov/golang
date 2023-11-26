package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"restapi/internal/entity"
)

type UploadImagePostgres struct {
	db *sqlx.DB
}

func NewUploadImagePostgres(db *sqlx.DB) *UploadImagePostgres {
	return &UploadImagePostgres{db: db}
}

func (r *UploadImagePostgres) UploadImage(id int, image entity.Image) (int, error) {
	var imageId int
	query := fmt.Sprintf("INSERT INTO %s (image) values($1) RETURNING id", ImageTable)
	row := r.db.QueryRow(query, image.Image)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return imageId, nil
}
