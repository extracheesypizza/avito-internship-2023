package handler

import (
	"avito-app/pkg/service"

	_ "avito-app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	segment := router.Group("/segment")
	{
		segment.POST("/create/", h.createSegment)
		segment.POST("/delete/", h.deleteSegment)
	}

	user := router.Group("/user")
	{
		user.GET("/getSegments/:id", h.getUserSegments)
		user.POST("/getActions/", h.getUserActions)
		user.POST("/addToSegment/", h.addUserToSegment)
		user.POST("/removeFromSegment/", h.removeUserFromSegment)
	}

	return router
}
