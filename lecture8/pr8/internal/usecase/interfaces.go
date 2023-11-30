package usecase

import (
	"context"

	"github.com/evrone/go-clean-template/internal/controller/http/v1/dto"
	"github.com/evrone/go-clean-template/internal/entity"
)

type (

	// User
	UserUseCase interface {
		Users(ctx context.Context) ([]*entity.User, error)
		CreateUser(ctx context.Context, user *entity.User) (int, error)
		GetUserByEmail(ctx context.Context, id string) (*entity.User, error)
		GetUserByID(ctx context.Context, id int) (*entity.User, error)

		Register(ctx context.Context, name, email string, age int, password string) error
		Login(ctx context.Context, email, password string) (*dto.LoginResponse, error)
	}
)
