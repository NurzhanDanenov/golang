package handler

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
		imageHandler := apiHandler.Group("/images")
		{
			imageHandler.POST("/", h.createAva)
			imageHandler.GET("/", h.getAllImages)
			imageHandler.GET("/:id", h.getImageById)
			imageHandler.DELETE("/:id", h.deleteImage)
			imageHandler.GET("/metrics", gin.WrapH(promhttp.Handler()))
		}
	}
	return router
}
