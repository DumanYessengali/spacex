package handler

import (
	"garyshker/pkg/service"
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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.CreateUser)
		auth.POST("/sign-in", h.Login)
	}
	api := router.Group("api", h.TokenAuthMiddleware())
	{
		api.POST("/logout", h.LogOut)
	}
	return router
}
