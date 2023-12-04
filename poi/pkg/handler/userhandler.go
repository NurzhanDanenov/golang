package handler

import (
	"net/http"
	"restapi/internal/entity"

	"github.com/gin-gonic/gin"
)

type UserBuilder interface {
	WithName(name string) UserBuilder
	WithEmail(email string) UserBuilder
	WithAge(age int) UserBuilder
	WithGender(gender string) UserBuilder
	WithPassword(password string) UserBuilder
	Build() *entity.User
}

type UserBuilderImpl struct {
	user *entity.User
}

func NewUserBuilder() UserBuilder {
	return &UserBuilderImpl{user: &entity.User{}}
}

func (b *UserBuilderImpl) WithName(name string) UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilderImpl) WithEmail(email string) UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilderImpl) WithAge(age int) UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilderImpl) WithGender(gender string) UserBuilder {
	b.user.Gender = gender
	return b
}

func (b *UserBuilderImpl) WithPassword(password string) UserBuilder {
	b.user.Password = password
	return b
}

func (b *UserBuilderImpl) Build() *entity.User {
	return b.user
}

func (h *Handler) Register(ctx *gin.Context) {
	var input entity.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := NewUserBuilder().
		WithName(input.Name).
		WithEmail(input.Email).
		WithAge(input.Age).
		WithGender(input.Gender).
		WithPassword(input.Password).
		Build()

	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) Login(ctx *gin.Context) {
	var input entity.Token

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
