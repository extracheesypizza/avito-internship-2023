package handler

import (
	"avito-app/pkg/service"

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

	segment := router.Group("/segment")
	{
		segment.POST("/create", h.createSegment)
		segment.POST("/delete", h.deleteSegment)
	}

	user := router.Group("/user")
	{
		user.GET("/", h.getUserList)
	}

	return router
}
