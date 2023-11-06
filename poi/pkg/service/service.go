package service

import (
	"restapi/meet"
	"restapi/pkg/repository"
)

type Authorization interface {
	CreateUser(user meet.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
