package handler

import "github.com/gin-gonic/gin"

type Handler struct {
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
}
