package handler

import (
	"restapi/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	adminHandler := router.Group("/admin/user")
	{

		adminHandler.GET("/all", h.GetUsers)
		adminHandler.POST("/", h.CreateUser)
		adminHandler.GET("/", h.GetUserByEmail)
	}

	userHandler := router.Group("/user")
	{
		userHandler.POST("/register", h.Register)
		userHandler.POST("/login", h.Login)
	}

	apiHandler := router.Group("/api", h.userIdentity)
	{
		apiHandler.POST("/", h.createAva)
	}
	return router
}
