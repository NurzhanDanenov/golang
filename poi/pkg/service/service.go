package service

import (
	"mime/multipart"
	"restapi/internal/entity"
	"restapi/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type UploadImage interface {
	Upload(userId int, image entity.Image, file multipart.File, filePath string) (int, error)
}

type Service struct {
	Authorization
	UploadImage
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
