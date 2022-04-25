package handlers

import (
	"github.com/gin-gonic/gin"
	"reddit/pck/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h Handler) InitRouters() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		posts := api.Group("/posts")
		{
			posts.GET("/:item_id")
		}
	}

	return router
}
