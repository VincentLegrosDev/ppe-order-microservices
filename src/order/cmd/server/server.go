package server

import (
	"github.com/gin-gonic/gin"
	"ppe4peeps.com/services/topics"
	"ppe4peeps.com/services/order/internal/handlers"
)

func SetupRouter() *gin.Engine {
	api := gin.Default()
	api.GET("/ping", handlers.Health)

	api.POST("orderReceived", func(c *gin.Context) {
		handlers.PublishOrderEvent(c, topics.OrderReceived)
	})

	return api
}
